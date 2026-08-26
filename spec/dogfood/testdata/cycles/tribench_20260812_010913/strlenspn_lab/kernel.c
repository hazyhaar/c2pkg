#include <stdint.h>
#include <stddef.h>
/* Minimal strspn-like for probe triangle (oracle C). accept = "hel\0" fixed. */
size_t strlenspn_lab(const uint8_t *s, size_t n) {
    static const uint8_t accept[] = {'h', 'e', 'l', 0};
    size_t i = 0;
    for (; i < n; i++) {
        const uint8_t *a = accept;
        int ok = 0;
        for (; *a; a++) {
            if (*a == s[i]) { ok = 1; break; }
        }
        if (!ok) break;
    }
    return i;
}
