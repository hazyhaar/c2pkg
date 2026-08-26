#include <stdint.h>
#include <stddef.h>
/* Minimal strspn-like — no string literals (front subset). accept = hel\0 */
static const uint8_t accept_hel[4] = {104, 101, 108, 0};

size_t strlenspn_lab(const uint8_t *s, size_t n) {
    size_t i;
    for (i = 0; i < n; i++) {
        size_t j;
        int ok;
        ok = 0;
        for (j = 0; j < 4; j++) {
            if (accept_hel[j] == 0) {
                break;
            }
            if (accept_hel[j] == s[i]) {
                ok = 1;
                break;
            }
        }
        if (ok == 0) {
            break;
        }
    }
    return i;
}
