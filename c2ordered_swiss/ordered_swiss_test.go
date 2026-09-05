// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2ordered_swiss

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const (
	elemInt = 1
	elemStr = 2
)

func newTable(capHash, capData uint32) *Swiss_table_t {
	var ht Swiss_table_t
	ctrl := make([]byte, capHash+16)
	slots := make([]uint32, capHash)
	entries := make([]Swiss_element_t, capData)
	Swiss_init(&ht, ctrl, slots, entries, capHash, capData)
	return &ht
}

func TestSwissInitPacked(t *testing.T) {
	ht := newTable(32, 32)
	if ht.N_elements != 0 || ht.N_used != 0 {
		t.Fatalf("table non vide apres init: elems=%d used=%d", ht.N_elements, ht.N_used)
	}
	for i := int64(0); i < 10; i++ {
		if !Swiss_insert_packed_int(ht, i, elemInt, i*100) {
			t.Fatalf("echec insert packed pour i=%d", i)
		}
	}
	if ht.N_elements != 10 || ht.N_used != 10 {
		t.Fatalf("apres inserts: elems=%d used=%d", ht.N_elements, ht.N_used)
	}
	for i := int64(0); i < 10; i++ {
		idx := Swiss_find_int(ht, i)
		if idx != int(i) {
			t.Fatalf("find_int(%d): got %d, want %d", i, idx, i)
		}
		if ht.Entries[idx].Val.I64 != i*100 {
			t.Fatalf("val.i64 pour %d: got %d, want %d", i, ht.Entries[idx].Val.I64, i*100)
		}
	}
	if idx := Swiss_find_int(ht, 999); idx != -1 {
		t.Fatalf("find_int(999) attendu -1, got %d", idx)
	}
}

// Avant correction, (uint32_t)key tronquait 2^32+3 en 3 et retournait l'indice 3.
func TestSwissFindIntNoTruncation(t *testing.T) {
	ht := newTable(32, 32)
	for i := int64(0); i < 10; i++ {
		Swiss_insert_packed_int(ht, i, elemInt, i)
	}
	for _, key := range []int64{1<<32 + 3, 1 << 32, -1, -1 << 40, 1<<63 - 1} {
		if idx := Swiss_find_int(ht, key); idx != -1 {
			t.Fatalf("find_int(%d) en mode packed: got %d, want -1", key, idx)
		}
	}
}

func TestSwissInsertIntSwitchesToHashed(t *testing.T) {
	ht := newTable(64, 64)
	for i := int64(0); i < 8; i++ {
		if !Swiss_insert_int(ht, i, elemInt, i*3) {
			t.Fatalf("insert_int(%d) packed refuse", i)
		}
	}
	if ht.Flags&2 == 0 {
		t.Fatal("la table devrait encore etre packed")
	}
	if !Swiss_insert_int(ht, 1<<32+3, elemInt, 42) {
		t.Fatal("insert_int(2^32+3) refuse")
	}
	if ht.Flags&2 != 0 {
		t.Fatal("la table devrait etre hashed")
	}
	for i := int64(0); i < 8; i++ {
		idx := Swiss_find_int(ht, i)
		if idx != int(i) || ht.Entries[idx].Val.I64 != i*3 {
			t.Fatalf("find_int(%d) apres bascule: idx=%d", i, idx)
		}
	}
	if idx := Swiss_find_int(ht, 3); idx != 3 {
		t.Fatalf("find_int(3) devrait donner 3, got %d", idx)
	}
	idx := Swiss_find_int(ht, 1<<32+3)
	if idx != 8 || ht.Entries[idx].Val.I64 != 42 {
		t.Fatalf("find_int(2^32+3): idx=%d", idx)
	}
	if !Swiss_insert_int(ht, 3, elemInt, 777) || ht.Entries[3].Val.I64 != 777 {
		t.Fatal("mise a jour de la cle 3 echouee")
	}
	if ht.N_elements != 9 {
		t.Fatalf("n_elements=%d, want 9", ht.N_elements)
	}
}

func TestSwissInsertStr(t *testing.T) {
	ht := newTable(64, 64)
	keys := []string{"alpha", "beta", "gamma", "", "a", "alphA"}
	for i, k := range keys {
		if !Swiss_insert_str(ht, []byte(k), uint32(len(k)), elemStr, int64(i)) {
			t.Fatalf("insert_str(%q) refuse", k)
		}
	}
	for i, k := range keys {
		h := Swiss_hash_str([]byte(k), uint32(len(k)))
		idx := Swiss_find_str(ht, []byte(k), uint32(len(k)), h)
		if idx != i {
			t.Fatalf("find_str(%q): idx=%d want %d", k, idx, i)
		}
	}
	h := Swiss_hash_str([]byte("delta"), 5)
	if idx := Swiss_find_str(ht, []byte("delta"), 5, h); idx != -1 {
		t.Fatalf("find_str(delta) attendu -1, got %d", idx)
	}
	if !Swiss_insert_str(ht, []byte("beta"), 4, elemStr, 99) || ht.Entries[1].Val.I64 != 99 || ht.N_elements != 6 {
		t.Fatal("mise a jour de beta echouee")
	}
	// Une cle chaine ne se confond pas avec une cle entiere
	if idx := Swiss_find_int(ht, 0); idx != -1 {
		t.Fatalf("find_int(0) sur table de chaines: got %d", idx)
	}
}

// Seuil : (n_used + n_tombstones + 1) * 8 < capacity_hash * 7.
func TestSwissLoadFactor(t *testing.T) {
	ht := newTable(64, 64)
	inserted := 0
	for k := int64(1000); k < 1000+64; k++ {
		if !Swiss_insert_int(ht, k, elemInt, k) {
			break
		}
		inserted++
	}
	if inserted != 55 {
		t.Fatalf("insertions acceptees=%d, want 55 (< 87.5%% de 64)", inserted)
	}
	if float64(ht.N_used)/float64(ht.Capacity_hash) >= 0.875 {
		t.Fatalf("charge %d/%d atteint 87.5%%", ht.N_used, ht.Capacity_hash)
	}
	// Pop en mode hashed : le slot de controle devient un tombstone, l'occupation ne baisse pas
	if !Swiss_pop(ht) || ht.N_tombstones != 1 || ht.N_used != 54 {
		t.Fatalf("pop hashed: tomb=%d used=%d", ht.N_tombstones, ht.N_used)
	}
	if Swiss_find_int(ht, 1054) != -1 {
		t.Fatal("la cle poppee reste trouvable")
	}
	// Reinsertion apres pop : acceptee au seuil, car la sequence de sondage de la cle poppee
	// retrouve son propre slot DELETED (occupation constante)
	if !Swiss_insert_int(ht, 1054, elemInt, 7) || ht.N_tombstones != 0 || ht.N_used != 55 {
		t.Fatalf("reinsertion apres pop: tomb=%d used=%d", ht.N_tombstones, ht.N_used)
	}
	if idx := Swiss_find_int(ht, 1054); idx != 54 || ht.Entries[idx].Val.I64 != 7 {
		t.Fatalf("find_int(1054) apres reinsertion: idx=%d", idx)
	}
	if Swiss_insert_int(ht, 5000, elemInt, 1) {
		t.Fatal("insertion au-dela du seuil acceptee")
	}
	// Trou dense + compaction puis insertion acceptee
	ht.Entries[10].Val.Type_ = 0
	ht.N_tombstones = 1
	ht.N_elements--
	Swiss_compact(ht)
	if ht.N_tombstones != 0 || ht.N_used != 54 {
		t.Fatalf("apres compact: tomb=%d used=%d", ht.N_tombstones, ht.N_used)
	}
	if Swiss_find_int(ht, 1010) != -1 {
		t.Fatal("cle 1010 supprimee encore trouvable apres compact")
	}
	for k := int64(1000); k < 1055; k++ {
		if k == 1010 {
			continue
		}
		idx := Swiss_find_int(ht, k)
		if idx < 0 || ht.Entries[idx].Int_key != k {
			t.Fatalf("cle %d perdue apres compact (idx=%d)", k, idx)
		}
	}
	if !Swiss_insert_int(ht, 5000, elemInt, 1) {
		t.Fatal("insertion apres compaction refusee")
	}
}

func TestSwissPopAndCompaction(t *testing.T) {
	ht := newTable(32, 32)
	for i := int64(0); i < 5; i++ {
		Swiss_insert_packed_int(ht, i, elemInt, i+1)
	}
	ok := Swiss_pop(ht)
	if !ok || ht.N_elements != 4 || ht.N_used != 4 {
		t.Fatalf("pop failed: ok=%v elems=%d used=%d", ok, ht.N_elements, ht.N_used)
	}
	ht.N_tombstones = 1
	ht.Entries[1].Val.Type_ = 0
	Swiss_compact(ht)
	if ht.N_tombstones != 0 {
		t.Fatalf("tombstones apres compact: %d", ht.N_tombstones)
	}
}

// Scénario partagé C / Go : chaque étape émet une ligne, les deux traces doivent être identiques.
const oracleHarness = `
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <string.h>

typedef struct swiss_val_s { uint32_t type; uint32_t flags; int64_t i64; } swiss_val_t;
typedef struct swiss_element_s { swiss_val_t val; uint64_t hash; char *str_key; uint32_t str_len; int64_t int_key; } swiss_element_t;
typedef struct swiss_table_s {
    uint8_t *ctrl; uint32_t *ctrl_to_dense; swiss_element_t *entries;
    uint32_t capacity_hash, capacity_data, n_elements, n_used, n_tombstones, flags;
} swiss_table_t;

uint64_t swiss_hash_str(const char *str, uint32_t len);
void swiss_init(swiss_table_t *ht, uint8_t *ctrl_buf, uint32_t *slots_buf, swiss_element_t *entries_buf, uint32_t cap_hash, uint32_t cap_data);
bool swiss_insert_packed_int(swiss_table_t *ht, int64_t key, uint32_t val_type, int64_t val_i64);
bool swiss_insert_int(swiss_table_t *ht, int64_t key, uint32_t val_type, int64_t val_i64);
bool swiss_insert_str(swiss_table_t *ht, char *key, uint32_t key_len, uint32_t val_type, int64_t val_i64);
int32_t swiss_find_int(const swiss_table_t *ht, int64_t key);
int32_t swiss_find_str(const swiss_table_t *ht, const char *key, uint32_t key_len, uint64_t hash);
bool swiss_pop(swiss_table_t *ht);
void swiss_compact(swiss_table_t *ht);

static swiss_table_t ht;
static void state(const char *tag) {
    printf("%s used=%u elems=%u tomb=%u flags=%u\n", tag, ht.n_used, ht.n_elements, ht.n_tombstones, ht.flags);
}
static void find_i(int64_t key) {
    int32_t idx = swiss_find_int(&ht, key);
    long long v = idx >= 0 ? (long long)ht.entries[idx].val.i64 : 0;
    printf("FI %lld idx=%d val=%lld\n", (long long)key, idx, v);
}
static void find_s(const char *key) {
    uint32_t len = (uint32_t)strlen(key);
    int32_t idx = swiss_find_str(&ht, key, len, swiss_hash_str(key, len));
    long long v = idx >= 0 ? (long long)ht.entries[idx].val.i64 : 0;
    printf("FS %s idx=%d val=%lld\n", key, idx, v);
}
static void ins_i(int64_t key, int64_t val) {
    printf("II %lld ok=%d\n", (long long)key, (int)swiss_insert_int(&ht, key, 1, val));
}
static void ins_s(const char *key, int64_t val) {
    printf("IS %s ok=%d\n", key, (int)swiss_insert_str(&ht, (char *)key, (uint32_t)strlen(key), 2, val));
}
static char names[8][8] = {"alpha", "beta", "gamma", "delta", "eps", "zeta", "eta", "theta"};

int main(void) {
    static uint8_t ctrl[64 + 16];
    static uint32_t slots[64];
    static swiss_element_t entries[64];
    swiss_init(&ht, ctrl, slots, entries, 64, 64);
    for (int64_t i = 0; i < 20; ++i) swiss_insert_packed_int(&ht, i, 1, i * 7 + 3);
    state("S0");
    for (int64_t i = 0; i < 20; ++i) find_i(i);
    find_i(4294967299LL); find_i(4294967296LL); find_i(-1); find_i(20); find_i(9223372036854775807LL);
    ins_i(1000, 11); state("S1");
    for (int64_t i = 0; i < 20; ++i) find_i(i);
    find_i(1000); find_i(4294967299LL);
    ins_i(-7, 22); ins_i(1LL << 40, 33); ins_i(4294967299LL, 44); ins_i(5, 555); ins_i(1000, 1111);
    find_i(-7); find_i(1LL << 40); find_i(4294967299LL); find_i(5); find_i(1000); find_i(3);
    state("S2");
    for (int i = 0; i < 6; ++i) ins_s(names[i], 100 + i);
    ins_s("beta", 999);
    for (int i = 0; i < 8; ++i) find_s(names[i]);
    find_i(0); find_i(100);
    state("S3");
    int accepted = 0;
    for (int64_t k = 2000; k < 2100; ++k) { if (!swiss_insert_int(&ht, k, 1, k)) break; accepted++; }
    printf("FILL accepted=%d\n", accepted);
    state("S4");
    for (int i = 0; i < 3; ++i) { printf("POP ok=%d\n", (int)swiss_pop(&ht)); state("P"); }
    ins_i(9000, 1); find_i(9000);
    swiss_compact(&ht); state("S5");
    ins_i(9000, 1); ins_i(9001, 2); ins_i(9002, 3); ins_i(9003, 4); state("S6");
    for (int64_t k = 2000; k < 2000 + accepted; ++k) find_i(k);
    find_i(9000); find_i(9003);
    for (int i = 0; i < 8; ++i) find_s(names[i]);
    ht.entries[3].val.type = 0; ht.entries[25].val.type = 0; ht.n_tombstones += 2; ht.n_elements -= 2;
    swiss_compact(&ht); state("S7");
    find_i(3); find_i(2); find_i(4);
    for (int64_t k = 2000; k < 2000 + accepted; ++k) find_i(k);
    for (int i = 0; i < 8; ++i) find_s(names[i]);
    ins_i(3, 303); find_i(3); state("S8");
    while (swiss_pop(&ht)) {}
    state("S9");
    find_i(0); find_s("alpha");
    return 0;
}
`

func goOracleTrace() []string {
	var out []string
	ht := newTable(64, 64)
	state := func(tag string) {
		out = append(out, fmt.Sprintf("%s used=%d elems=%d tomb=%d flags=%d", tag, ht.N_used, ht.N_elements, ht.N_tombstones, ht.Flags))
	}
	findI := func(key int64) {
		idx := Swiss_find_int(ht, key)
		var v int64
		if idx >= 0 {
			v = ht.Entries[idx].Val.I64
		}
		out = append(out, fmt.Sprintf("FI %d idx=%d val=%d", key, idx, v))
	}
	findS := func(key string) {
		b := []byte(key)
		idx := Swiss_find_str(ht, b, uint32(len(b)), Swiss_hash_str(b, uint32(len(b))))
		var v int64
		if idx >= 0 {
			v = ht.Entries[idx].Val.I64
		}
		out = append(out, fmt.Sprintf("FS %s idx=%d val=%d", key, idx, v))
	}
	b2i := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}
	insI := func(key, val int64) {
		out = append(out, fmt.Sprintf("II %d ok=%d", key, b2i(Swiss_insert_int(ht, key, elemInt, val))))
	}
	insS := func(key string, val int64) {
		b := []byte(key)
		out = append(out, fmt.Sprintf("IS %s ok=%d", key, b2i(Swiss_insert_str(ht, b, uint32(len(b)), elemStr, val))))
	}
	names := []string{"alpha", "beta", "gamma", "delta", "eps", "zeta", "eta", "theta"}

	for i := int64(0); i < 20; i++ {
		Swiss_insert_packed_int(ht, i, elemInt, i*7+3)
	}
	state("S0")
	for i := int64(0); i < 20; i++ {
		findI(i)
	}
	findI(4294967299)
	findI(4294967296)
	findI(-1)
	findI(20)
	findI(9223372036854775807)
	insI(1000, 11)
	state("S1")
	for i := int64(0); i < 20; i++ {
		findI(i)
	}
	findI(1000)
	findI(4294967299)
	insI(-7, 22)
	insI(1<<40, 33)
	insI(4294967299, 44)
	insI(5, 555)
	insI(1000, 1111)
	findI(-7)
	findI(1 << 40)
	findI(4294967299)
	findI(5)
	findI(1000)
	findI(3)
	state("S2")
	for i := 0; i < 6; i++ {
		insS(names[i], int64(100+i))
	}
	insS("beta", 999)
	for i := 0; i < 8; i++ {
		findS(names[i])
	}
	findI(0)
	findI(100)
	state("S3")
	accepted := 0
	for k := int64(2000); k < 2100; k++ {
		if !Swiss_insert_int(ht, k, elemInt, k) {
			break
		}
		accepted++
	}
	out = append(out, fmt.Sprintf("FILL accepted=%d", accepted))
	state("S4")
	for i := 0; i < 3; i++ {
		out = append(out, fmt.Sprintf("POP ok=%d", b2i(Swiss_pop(ht))))
		state("P")
	}
	insI(9000, 1)
	findI(9000)
	Swiss_compact(ht)
	state("S5")
	insI(9000, 1)
	insI(9001, 2)
	insI(9002, 3)
	insI(9003, 4)
	state("S6")
	for k := int64(2000); k < int64(2000+accepted); k++ {
		findI(k)
	}
	findI(9000)
	findI(9003)
	for i := 0; i < 8; i++ {
		findS(names[i])
	}
	ht.Entries[3].Val.Type_ = 0
	ht.Entries[25].Val.Type_ = 0
	ht.N_tombstones += 2
	ht.N_elements -= 2
	Swiss_compact(ht)
	state("S7")
	findI(3)
	findI(2)
	findI(4)
	for k := int64(2000); k < int64(2000+accepted); k++ {
		findI(k)
	}
	for i := 0; i < 8; i++ {
		findS(names[i])
	}
	insI(3, 303)
	findI(3)
	state("S8")
	for Swiss_pop(ht) {
	}
	state("S9")
	findI(0)
	findS("alpha")
	return out
}

func TestSwissVsCOracle(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc non disponible")
	}
	candidates := []string{
		filepath.Join("..", "..", "sources", "c2ordered_swiss", "ordered_swiss.c"),
		filepath.Join("..", "sources", "c2ordered_swiss", "ordered_swiss.c"),
		filepath.Join("sources", "c2ordered_swiss", "ordered_swiss.c"),
	}
	var cSrcPath string
	for _, cand := range candidates {
		if abs, err := filepath.Abs(cand); err == nil {
			if _, err := os.Stat(abs); err == nil {
				cSrcPath = abs
				break
			}
		}
	}
	if cSrcPath == "" {
		t.Skip("source C introuvable")
	}

	tmpDir := t.TempDir()
	harnessFile := filepath.Join(tmpDir, "harness.c")
	binFile := filepath.Join(tmpDir, "oracle_swiss")
	if err := os.WriteFile(harnessFile, []byte(oracleHarness), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("gcc", "-O2", "-Wall", "-o", binFile, harnessFile, cSrcPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc compile failed: %v\n%s", err, string(out))
	}
	cOut, err := exec.Command(binFile).CombinedOutput()
	if err != nil {
		t.Fatalf("oracle execution failed: %v\n%s", err, string(cOut))
	}

	cLines := strings.Split(strings.TrimSpace(string(cOut)), "\n")
	goLines := goOracleTrace()
	if len(cLines) < 150 {
		t.Fatalf("trace C trop courte (%d lignes), scenario non joue", len(cLines))
	}
	n := len(cLines)
	if len(goLines) < n {
		n = len(goLines)
	}
	for i := 0; i < n; i++ {
		if strings.TrimSpace(cLines[i]) != goLines[i] {
			t.Fatalf("Divergence C vs Go a la ligne %d:\n  C : %q\n  Go: %q", i+1, cLines[i], goLines[i])
		}
	}
	if len(cLines) != len(goLines) {
		t.Fatalf("longueur des traces: C=%d Go=%d", len(cLines), len(goLines))
	}
	// Garde : le scenario doit avoir exerce le seuil de charge et la bascule hashed
	joined := strings.Join(goLines, "\n")
	if !strings.Contains(joined, "FILL accepted=") || !strings.Contains(joined, "flags=1") {
		t.Fatal("le scenario n'a pas exerce le remplissage ou la bascule hashed")
	}
	t.Logf("parite C/Go sur %d lignes", len(goLines))
}

func TestSwissZeroAlloc(t *testing.T) {
	ht := newTable(64, 64)
	for i := int64(0); i < 20; i++ {
		Swiss_insert_int(ht, i, elemInt, i*10)
	}
	Swiss_insert_int(ht, 1<<33, elemInt, 1)
	key := []byte("alpha")
	Swiss_insert_str(ht, key, 5, elemStr, 2)
	h := Swiss_hash_str(key, 5)
	allocs := testing.AllocsPerRun(1000, func() {
		_ = Swiss_find_int(ht, 5)
		_ = Swiss_find_int(ht, 1<<33)
		_ = Swiss_find_str(ht, key, 5, h)
		_ = Swiss_insert_int(ht, 5, elemInt, 6)
	})
	if allocs != 0 {
		t.Fatalf("Allocs/op = %.2f, attendu 0", allocs)
	}
}
