#include <stdint.h>
#include <stddef.h>
/* Minimal strspn-like for probe triangle (oracle C). accept = "hel" */
size_t strlenspn_lab(const uint8_t *s, size_t n) {
    const uint8_t *accept = (const uint8_t *)"hel";
    size_t i;
    for (i = 0; i < n; i++) {
        const uint8_t *a;
        int ok;
        ok = 0;
        for (a = accept; *a != 0; a++) {
            if (*a == s[i]) {
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
