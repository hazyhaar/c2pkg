// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2blueteam

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"
	"unsafe"
)

const (
	C2BT_SUB_PROC    = 1
	C2BT_SUB_FS      = 2
	C2BT_SUB_NET     = 3
	C2BT_SUB_MCP     = 4
	C2BT_SUB_HARNESS = 5
	C2BT_SUB_ENTROPY = 6

	C2BT_ACT_EXEC      = 1
	C2BT_ACT_OPEN      = 2
	C2BT_ACT_CONNECT   = 3
	C2BT_ACT_TOOL_CALL = 4
	C2BT_ACT_READ      = 5
	C2BT_ACT_WRITE     = 6

	C2BT_FLAG_NONE              = 0x0000
	C2BT_FLAG_VERDICT_OK        = 0x0001
	C2BT_FLAG_BLOCKED           = 0x0002
	C2BT_FLAG_ANOMALY           = 0x0004
	C2BT_FLAG_LOLBAS            = 0x0008
	C2BT_FLAG_BEACONING         = 0x0010
	C2BT_FLAG_DRIFT             = 0x0020
	C2BT_FLAG_HEX_PAYLOAD       = 0x0040
	C2BT_FLAG_BASE64_PAYLOAD    = 0x0080
	C2BT_FLAG_CRYPTO_PAYLOAD    = 0x0100
	C2BT_FLAG_CORRELATED_THREAT = 0x0200
)

func TestStructureInvariants(t *testing.T) {
	var ev Probe_event_t
	if sz := unsafe.Sizeof(ev); sz != 128 {
		t.Fatalf("sizeof Probe_event_t = %d, expected 128", sz)
	}
	if off := unsafe.Offsetof(ev.Ts_ns); off != 0 {
		t.Errorf("offsetof Ts_ns = %d, expected 0", off)
	}
	if off := unsafe.Offsetof(ev.Pid); off != 8 {
		t.Errorf("offsetof Pid = %d, expected 8", off)
	}
	if off := unsafe.Offsetof(ev.Tid); off != 12 {
		t.Errorf("offsetof Tid = %d, expected 12", off)
	}
	if off := unsafe.Offsetof(ev.Subsystem); off != 16 {
		t.Errorf("offsetof Subsystem = %d, expected 16", off)
	}
	if off := unsafe.Offsetof(ev.Action); off != 18 {
		t.Errorf("offsetof Action = %d, expected 18", off)
	}
	if off := unsafe.Offsetof(ev.Flags); off != 20 {
		t.Errorf("offsetof Flags = %d, expected 20", off)
	}
	if off := unsafe.Offsetof(ev.Src); off != 24 {
		t.Errorf("offsetof Src = %d, expected 24", off)
	}
	if off := unsafe.Offsetof(ev.Payload); off != 32 {
		t.Errorf("offsetof Payload = %d, expected 32", off)
	}

	var ch Probe_channel_t
	if sz := unsafe.Sizeof(ch); sz != 131200 {
		t.Fatalf("sizeof Probe_channel_t = %d, expected 131200", sz)
	}
	if off := unsafe.Offsetof(ch.Head); off != 131072 {
		t.Errorf("offsetof Head = %d, expected 131072", off)
	}
	if off := unsafe.Offsetof(ch.Drops); off != 131080 {
		t.Errorf("offsetof Drops = %d, expected 131080", off)
	}
	if off := unsafe.Offsetof(ch.Tail); off != 131136 {
		t.Errorf("offsetof Tail = %d, expected 131136", off)
	}
}

func TestRingBufferSPSC(t *testing.T) {
	var ch Probe_channel_t
	C2bt_channel_init(&ch)
	if drops := C2bt_channel_get_drops(&ch); drops != 0 {
		t.Fatalf("initial drops = %d, expected 0", drops)
	}

	var evIn Probe_event_t
	evIn.Ts_ns = 1000000000
	evIn.Pid = 4242
	evIn.Subsystem = C2BT_SUB_PROC
	evIn.Action = C2BT_ACT_EXEC
	copy(evIn.Payload[:], "python")

	w := C2bt_channel_write(&ch, &evIn)
	if w != 0 {
		t.Fatalf("C2bt_channel_write failed, code %d", w)
	}

	var evOut Probe_event_t
	r := C2bt_channel_read(&ch, &evOut)
	if r != 1 {
		t.Fatalf("C2bt_channel_read expected 1, got %d", r)
	}
	if evOut.Pid != 4242 {
		t.Errorf("read pid = %d, expected 4242", evOut.Pid)
	}
	if !bytes.HasPrefix(evOut.Payload[:], []byte("python")) {
		t.Errorf("payload mismatch: %s", string(evOut.Payload[:]))
	}

	r = C2bt_channel_read(&ch, &evOut)
	if r != 0 {
		t.Fatalf("expected empty channel read 0, got %d", r)
	}

	// Test saturation et compteur de rejets
	C2bt_channel_init(&ch)
	for i := 0; i < 1024; i++ {
		evIn.Pid = uint32(1000 + i)
		if res := C2bt_channel_write(&ch, &evIn); res != 0 {
			t.Fatalf("write #%d failed with %d", i, res)
		}
	}
	if drops := C2bt_channel_get_drops(&ch); drops != 0 {
		t.Fatalf("drops before overflow = %d, expected 0", drops)
	}

	for i := 0; i < 15; i++ {
		if res := C2bt_channel_write(&ch, &evIn); res != -2 {
			t.Fatalf("overflow write #%d expected -2, got %d", i, res)
		}
	}
	if drops := C2bt_channel_get_drops(&ch); drops != 15 {
		t.Fatalf("drops after overflow = %d, expected 15", drops)
	}

	// Libérer 1 créneau
	if res := C2bt_channel_read(&ch, &evOut); res != 1 {
		t.Fatalf("read expected 1, got %d", res)
	}
	if res := C2bt_channel_write(&ch, &evIn); res != 0 {
		t.Fatalf("write on freed slot expected 0, got %d", res)
	}
	if drops := C2bt_channel_get_drops(&ch); drops != 15 {
		t.Fatalf("drops after successful write = %d, expected 15", drops)
	}
}

func TestEntropyShannonKAT(t *testing.T) {
	// Zero entropy buffer (all same byte 'A')
	zeroBuf := make([]byte, 256)
	for i := range zeroBuf {
		zeroBuf[i] = 'A'
	}
	entZero := C2bt_calc_entropy_8_8(zeroBuf, uint64(len(zeroBuf)))
	if entZero != 0 {
		t.Errorf("entropy for uniform buffer = %d, expected 0", entZero)
	}

	// Max entropy buffer (all 256 distinct byte values 0..255)
	maxBuf := make([]byte, 256)
	for i := range maxBuf {
		maxBuf[i] = byte(i)
	}
	entMax := C2bt_calc_entropy_8_8(maxBuf, uint64(len(maxBuf)))
	if entMax != 2048 {
		t.Errorf("entropy for max distinct buffer = %d, expected 2048", entMax)
	}
}

func TestRulesEngineBatch(t *testing.T) {
	inBatch := make([]Probe_event_t, 3)
	outBatch := make([]Probe_event_t, 3)

	inBatch[0].Subsystem = C2BT_SUB_PROC
	inBatch[0].Action = C2BT_ACT_EXEC
	copy(inBatch[0].Payload[:], "curl")

	inBatch[1].Subsystem = C2BT_SUB_PROC
	inBatch[1].Action = C2BT_ACT_EXEC
	copy(inBatch[1].Payload[:], "ls")

	inBatch[2].Subsystem = C2BT_SUB_MCP
	inBatch[2].Action = C2BT_ACT_TOOL_CALL
	copy(inBatch[2].Payload[:], "rm -rf /tmp/data")

	evalCount := C2bt_eval_rules_batch(inBatch, outBatch, 3)
	if evalCount != 3 {
		t.Fatalf("C2bt_eval_rules_batch count = %d, expected 3", evalCount)
	}

	// Check LOLBAS detection on curl
	if (outBatch[0].Flags & C2BT_FLAG_LOLBAS) == 0 {
		t.Errorf("expected LOLBAS flag on curl, got flags=0x%x", outBatch[0].Flags)
	}
	if (outBatch[0].Flags & C2BT_FLAG_ANOMALY) == 0 {
		t.Errorf("expected ANOMALY flag on curl, got flags=0x%x", outBatch[0].Flags)
	}

	// Check clean execution on ls
	if (outBatch[1].Flags & C2BT_FLAG_VERDICT_OK) == 0 {
		t.Errorf("expected VERDICT_OK on ls, got flags=0x%x", outBatch[1].Flags)
	}
	if (outBatch[1].Flags & C2BT_FLAG_ANOMALY) != 0 {
		t.Errorf("unexpected ANOMALY on ls, got flags=0x%x", outBatch[1].Flags)
	}

	// Check MCP injection block on rm -rf
	if (outBatch[2].Flags & C2BT_FLAG_BLOCKED) == 0 {
		t.Errorf("expected BLOCKED on rm -rf, got flags=0x%x", outBatch[2].Flags)
	}
	if (outBatch[2].Flags & C2BT_FLAG_ANOMALY) == 0 {
		t.Errorf("expected ANOMALY on rm -rf, got flags=0x%x", outBatch[2].Flags)
	}
}

func TestContextLifecycle(t *testing.T) {
	var ctx C2bt_ctx_t
	var cfg C2bt_config_t

	if C2bt_init_inplace(&ctx, &cfg) != 0 {
		t.Fatalf("C2bt_init_inplace failed")
	}
	if C2bt_start(&ctx) != 0 {
		t.Fatalf("C2bt_start failed")
	}
	if ctx.Running != 1 {
		t.Errorf("expected ctx.Running == 1, got %d", ctx.Running)
	}
	if C2bt_stop(&ctx) != 0 {
		t.Fatalf("C2bt_stop failed")
	}
	if ctx.Running != 0 {
		t.Errorf("expected ctx.Running == 0, got %d", ctx.Running)
	}
}

// TestVsCOracleBitExact exécute l'oracle binaire C compilé avec gcc -O2
// et vérifie la stricte parité bit-exacte requise par la doctrine c2simd.
func TestVsCOracleBitExact(t *testing.T) {
	tmpBin := filepath.Join(t.TempDir(), "c_oracle")
	srcDir := filepath.Join("..", "..", "sources", "c2blueteam")

	cmdBuild := exec.Command("gcc", "-O2", "-Wall", "-Wextra", "-std=gnu99",
		"-I", srcDir,
		filepath.Join(srcDir, "test_oracle.c"),
		filepath.Join(srcDir, "ring_buffer.c"),
		filepath.Join(srcDir, "probe_entropy.c"),
		filepath.Join(srcDir, "rules_simd.c"),
		filepath.Join(srcDir, "probe_proc.c"),
		filepath.Join(srcDir, "probe_mcp.c"),
		filepath.Join(srcDir, "probe_fs.c"),
		filepath.Join(srcDir, "c2blueteam.c"),
		filepath.Join(srcDir, "correlator.c"),
		"-o", tmpBin,
	)
	if out, err := cmdBuild.CombinedOutput(); err != nil {
		t.Fatalf("échec de compilation de l'oracle C: %v\nSortie: %s", err, string(out))
	}

	cmdRun := exec.Command(tmpBin)
	out, err := cmdRun.CombinedOutput()
	if err != nil {
		t.Fatalf("échec d'exécution de l'oracle C: %v\nSortie: %s", err, string(out))
	}
	t.Logf("Sortie Oracle C (gcc -O2):\n%s", string(out))
}

func BenchmarkRingBuffer_WriteRead_Nominal(b *testing.B) {
	var ch Probe_channel_t
	C2bt_channel_init(&ch)
	var evIn Probe_event_t
	evIn.Pid = 42
	var evOut Probe_event_t

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		C2bt_channel_write(&ch, &evIn)
		C2bt_channel_read(&ch, &evOut)
	}
}

func BenchmarkRingBuffer_Write_SaturatedDrops(b *testing.B) {
	var ch Probe_channel_t
	C2bt_channel_init(&ch)
	var ev Probe_event_t
	ev.Pid = 100

	// Remplissage complet préalable
	for i := 0; i < 1024; i++ {
		C2bt_channel_write(&ch, &ev)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		C2bt_channel_write(&ch, &ev)
	}
}

func BenchmarkRingBuffer_BatchRead_8(b *testing.B) {
	var ch Probe_channel_t
	C2bt_channel_init(&ch)
	var ev Probe_event_t
	ev.Pid = 100
	batch := make([]Probe_event_t, 8)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 8; j++ {
			C2bt_channel_write(&ch, &ev)
		}
		b.StartTimer()
		C2bt_channel_read_batch(&ch, batch, 8)
	}
}
