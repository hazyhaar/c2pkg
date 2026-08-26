#include <stdint.h>
#include <stddef.h>
#include <string.h>

int c2_myers_ses(const int *a, int n, const int *b, int m, int *v, int vlen) {
    if (n == 0 && m == 0) {
        return 0;
    }
    if (n < 0 || m < 0) {
        return -1;
    }
    int max_d = n + m;
    int need = max_d + max_d + 1;
    if (need > vlen) {
        return -1;
    }
    int offset = max_d;
    int i = 0;
    for (i = 0; i < need; i++) {
        v[i] = 0;
    }
    int d;
    for (d = 0; d <= max_d; d++) {
        int k;
        for (k = 0 - d; k <= d; k = k + 2) {
            int x;
            if (k == (0 - d) || (k != d && v[(k - 1) + offset] < v[(k + 1) + offset])) {
                x = v[(k + 1) + offset];
            } else {
                x = v[(k - 1) + offset] + 1;
            }
            int y = x - k;
            while (x < n && y < m && a[x] == b[y]) {
                x = x + 1;
                y = y + 1;
            }
            v[k + offset] = x;
            if (x >= n && y >= m) {
                return d;
            }
        }
    }
    return -1;
}

int c2_validate_patch_mode(const char *mode_str) {
    if (mode_str == NULL) return 1;
    if (strcmp(mode_str, "120000") == 0) return 0;
    if (strcmp(mode_str, "160000") == 0) return 0;
    return 1;
}
