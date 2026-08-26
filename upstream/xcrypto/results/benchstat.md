# x/crypto v0.54.0 — 2026-08-26T16:20Z — go version go1.27.0 linux/amd64 — count=5 benchtime=300ms pin='taskset -c 0-15'

## asm (référence) vs archsimd
goos: linux
goarch: amd64
pkg: golang.org/x/crypto/chacha20poly1305
cpu: Intel(R) Core(TM) i9-14900K
                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                      sec/op                       │             sec/op              vs base                │
Chacha20Poly1305/Open-64-16                                            112.1n ± ∞ ¹                     244.2n ± ∞ ¹  +117.84% (p=0.008 n=5)
Chacha20Poly1305/Seal-64-16                                            113.4n ± ∞ ¹                     255.7n ± ∞ ¹  +125.49% (p=0.008 n=5)
Chacha20Poly1305/Open-64-X-16                                          179.8n ± ∞ ¹                     322.9n ± ∞ ¹   +79.59% (p=0.008 n=5)
Chacha20Poly1305/Seal-64-X-16                                          184.1n ± ∞ ¹                     339.6n ± ∞ ¹   +84.46% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-16                                          548.1n ± ∞ ¹                    1105.0n ± ∞ ¹  +101.61% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-16                                          621.8n ± ∞ ¹                    1191.0n ± ∞ ¹   +91.54% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-X-16                                        642.3n ± ∞ ¹                    1355.0n ± ∞ ¹  +110.96% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-X-16                                        630.8n ± ∞ ¹                    1273.0n ± ∞ ¹  +101.81% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-16                                          2.854µ ± ∞ ¹                     4.111µ ± ∞ ¹   +44.04% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-16                                          2.757µ ± ∞ ¹                     4.039µ ± ∞ ¹   +46.50% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-X-16                                        2.972µ ± ∞ ¹                     4.100µ ± ∞ ¹   +37.95% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-X-16                                        2.846µ ± ∞ ¹                     4.231µ ± ∞ ¹   +48.66% (p=0.008 n=5)
geomean                                                                629.4n                           1.133µ         +80.01%
¹ need >= 6 samples for confidence interval at level 0.95

                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                        B/s                        │               B/s                vs base               │
Chacha20Poly1305/Open-64-16                                           544.6Mi ± ∞ ¹                     250.0Mi ± ∞ ¹  -54.10% (p=0.008 n=5)
Chacha20Poly1305/Seal-64-16                                           538.1Mi ± ∞ ¹                     238.7Mi ± ∞ ¹  -55.64% (p=0.008 n=5)
Chacha20Poly1305/Open-64-X-16                                         339.4Mi ± ∞ ¹                     189.0Mi ± ∞ ¹  -44.30% (p=0.008 n=5)
Chacha20Poly1305/Seal-64-X-16                                         331.5Mi ± ∞ ¹                     179.7Mi ± ∞ ¹  -45.79% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-16                                         2.294Gi ± ∞ ¹                     1.138Gi ± ∞ ¹  -50.39% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-16                                         2.022Gi ± ∞ ¹                     1.055Gi ± ∞ ¹  -47.80% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-X-16                                      2004.6Mi ± ∞ ¹                     950.1Mi ± ∞ ¹  -52.61% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-X-16                                      2041.1Mi ± ∞ ¹                    1011.1Mi ± ∞ ¹  -50.46% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-16                                         2.674Gi ± ∞ ¹                     1.856Gi ± ∞ ¹  -30.59% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-16                                         2.767Gi ± ∞ ¹                     1.889Gi ± ∞ ¹  -31.74% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-X-16                                       2.567Gi ± ∞ ¹                     1.861Gi ± ∞ ¹  -27.52% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-X-16                                       2.681Gi ± ∞ ¹                     1.803Gi ± ∞ ¹  -32.74% (p=0.008 n=5)
geomean                                                               1.319Gi                           750.1Mi        -44.45%
¹ need >= 6 samples for confidence interval at level 0.95

                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                       B/op                        │              B/op               vs base                │
Chacha20Poly1305/Open-64-16                                             0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-64-16                                             0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-64-X-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-64-X-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-1350-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-1350-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-1350-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-1350-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-8192-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-8192-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-8192-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-8192-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
geomean                                                                           ³                                   +0.00%               ³
¹ need >= 6 samples for confidence interval at level 0.95
² all samples are equal
³ summaries must be >0 to compute geomean

                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                     allocs/op                     │           allocs/op             vs base                │
Chacha20Poly1305/Open-64-16                                             0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-64-16                                             0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-64-X-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-64-X-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-1350-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-1350-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-1350-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-1350-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-8192-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-8192-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Open-8192-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
Chacha20Poly1305/Seal-8192-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=5) ²
geomean                                                                           ³                                   +0.00%               ³
¹ need >= 6 samples for confidence interval at level 0.95
² all samples are equal
³ summaries must be >0 to compute geomean

## purego (générique) vs archsimd
goos: linux
goarch: amd64
pkg: golang.org/x/crypto/chacha20poly1305
cpu: Intel(R) Core(TM) i9-14900K
                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                        sec/op                        │             sec/op               vs base               │
Chacha20Poly1305/Open-64-16                                               252.5n ± ∞ ¹                      244.2n ± ∞ ¹        ~ (p=0.222 n=5)
Chacha20Poly1305/Seal-64-16                                               241.0n ± ∞ ¹                      255.7n ± ∞ ¹        ~ (p=0.690 n=5)
Chacha20Poly1305/Open-64-X-16                                             321.7n ± ∞ ¹                      322.9n ± ∞ ¹        ~ (p=0.310 n=5)
Chacha20Poly1305/Seal-64-X-16                                             319.8n ± ∞ ¹                      339.6n ± ∞ ¹   +6.19% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-16                                             2.050µ ± ∞ ¹                      1.105µ ± ∞ ¹        ~ (p=0.151 n=5)
Chacha20Poly1305/Seal-1350-16                                             2.056µ ± ∞ ¹                      1.191µ ± ∞ ¹  -42.07% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-X-16                                           2.137µ ± ∞ ¹                      1.355µ ± ∞ ¹  -36.59% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-X-16                                           2.111µ ± ∞ ¹                      1.273µ ± ∞ ¹  -39.70% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-16                                            10.858µ ± ∞ ¹                      4.111µ ± ∞ ¹  -62.14% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-16                                            10.754µ ± ∞ ¹                      4.039µ ± ∞ ¹  -62.44% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-X-16                                          11.086µ ± ∞ ¹                      4.100µ ± ∞ ¹  -63.02% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-X-16                                          10.953µ ± ∞ ¹                      4.231µ ± ∞ ¹  -61.37% (p=0.008 n=5)
geomean                                                                   1.858µ                            1.133µ        -39.00%
¹ need >= 6 samples for confidence interval at level 0.95

                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                         B/s                          │              B/s                vs base                │
Chacha20Poly1305/Open-64-16                                              241.7Mi ± ∞ ¹                    250.0Mi ± ∞ ¹         ~ (p=0.222 n=5)
Chacha20Poly1305/Seal-64-16                                              253.2Mi ± ∞ ¹                    238.7Mi ± ∞ ¹         ~ (p=0.690 n=5)
Chacha20Poly1305/Open-64-X-16                                            189.7Mi ± ∞ ¹                    189.0Mi ± ∞ ¹         ~ (p=0.310 n=5)
Chacha20Poly1305/Seal-64-X-16                                            190.8Mi ± ∞ ¹                    179.7Mi ± ∞ ¹    -5.84% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-16                                            628.0Mi ± ∞ ¹                   1165.3Mi ± ∞ ¹         ~ (p=0.151 n=5)
Chacha20Poly1305/Seal-1350-16                                            626.2Mi ± ∞ ¹                   1080.8Mi ± ∞ ¹   +72.59% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-X-16                                          602.4Mi ± ∞ ¹                    950.1Mi ± ∞ ¹   +57.72% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-X-16                                          610.0Mi ± ∞ ¹                   1011.1Mi ± ∞ ¹   +65.77% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-16                                            719.5Mi ± ∞ ¹                   1900.2Mi ± ∞ ¹  +164.10% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-16                                            726.5Mi ± ∞ ¹                   1934.2Mi ± ∞ ¹  +166.25% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-X-16                                          704.7Mi ± ∞ ¹                   1905.4Mi ± ∞ ¹  +170.38% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-X-16                                          713.3Mi ± ∞ ¹                   1846.3Mi ± ∞ ¹  +158.85% (p=0.008 n=5)
geomean                                                                  457.5Mi                          750.1Mi         +63.94%
¹ need >= 6 samples for confidence interval at level 0.95

                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                         B/op                         │            B/op             vs base                    │
Chacha20Poly1305/Open-64-16                                                32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-64-16                                                32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-64-X-16                                              32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-64-X-16                                              32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-16                                              32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-16                                              32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-X-16                                            32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-X-16                                            32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-16                                              32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-16                                              32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-X-16                                            32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-X-16                                            32.00 ± ∞ ¹                   0.00 ± ∞ ¹  -100.00% (p=0.008 n=5)
geomean                                                                    32.00                                     ?                      ² ³
¹ need >= 6 samples for confidence interval at level 0.95
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean

                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                      allocs/op                       │         allocs/op           vs base                    │
Chacha20Poly1305/Open-64-16                                                1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-64-16                                                1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-64-X-16                                              1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-64-X-16                                              1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-16                                              1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-16                                              1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-1350-X-16                                            1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-1350-X-16                                            1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-16                                              1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-16                                              1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Open-8192-X-16                                            1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
Chacha20Poly1305/Seal-8192-X-16                                            1.000 ± ∞ ¹                  0.000 ± ∞ ¹  -100.00% (p=0.008 n=5)
geomean                                                                    1.000                                     ?                      ² ³
¹ need >= 6 samples for confidence interval at level 0.95
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean
