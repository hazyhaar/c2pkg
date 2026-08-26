# libc2blueteam — C99 Standalone Zero-Allocation Defense Library

High-throughput security supervision, eBPF telemetry processing, and ARCHTIME Shannon entropy engine for Linux / Unix environments.

---

## 1. Key Features

- **Pure C99 & Zero Allocation :** No `malloc`, no FPU instructions, 100% L1D-friendly static tables.
- **Shannon Entropy in Q8.8 Fixed-Point :** $3.42\text{ GB/s}$ throughput on a single CPU core, continuous resolution from 1 byte to 4 GB.
- **Single-Pass Character Classification :** Joint extraction of `PROSE`, `HEX`, `BASE64`, `JWT`, `CRYPTO_COMPRESSED`.
- **Lock-Free SPSC Ring Buffer :** Cache-line isolated producer/consumer indices (`131,200 bytes` constant footprint) with atomic drop accounting.
- **32 KiB Multi-Stream Correlator :** Sub-12 ns sliding-window correlation across MCP, Process, Filesystem, and Network events.

---

## 2. Build & Test

```bash
# Build static library (libc2blueteam.a) and run test suite
make all

# Run throughput benchmark
make bench
```

---

## 3. Library Usage

```c
#include "c2blueteam.h"
#include "correlator.h"

int main(void) {
    // 1. Initialize context in-place
    c2bt_ctx_t ctx;
    c2bt_config_t cfg = {
        .enable_proc = 1,
        .enable_mcp = 1,
        .enforce_mode = C2BT_MODE_ACTIVE
    };
    c2bt_init_inplace(&ctx, &cfg);
    c2bt_start(&ctx);

    // 2. Poll and evaluate events with fair-share round-robin & correlation
    probe_event_t batch[16];
    int count = c2bt_poll_batch(&ctx, batch, 16);

    c2bt_stop(&ctx);
    return 0;
}
```
