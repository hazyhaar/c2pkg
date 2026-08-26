# x/crypto v0.54.0 — 2026-08-26T13:17Z — go version go1.27.0 linux/amd64 — count=3 benchtime=100ms pin='taskset -c 0-15'

## asm (référence) vs archsimd
goos: linux
goarch: amd64
pkg: golang.org/x/crypto/chacha20poly1305
cpu: Intel(R) Core(TM) i9-14900K
                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                      sec/op                       │            sec/op              vs base                 │
Chacha20Poly1305/Open-64-16                                            110.9n ± ∞ ¹                    224.7n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-64-16                                            107.0n ± ∞ ¹                    230.2n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-64-X-16                                          174.9n ± ∞ ¹                    304.1n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-64-X-16                                          172.8n ± ∞ ¹                    278.7n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-16                                          522.4n ± ∞ ¹                   1060.0n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-16                                          555.2n ± ∞ ¹                   1130.0n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-X-16                                        628.7n ± ∞ ¹                   1229.0n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-X-16                                        596.1n ± ∞ ¹                   1277.0n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-16                                          2.646µ ± ∞ ¹                    3.802µ ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-16                                          2.652µ ± ∞ ¹                    4.203µ ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-X-16                                        2.625µ ± ∞ ¹                    4.333µ ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-X-16                                        2.682µ ± ∞ ¹                    4.420µ ± ∞ ¹        ~ (p=0.100 n=3) ²
geomean                                                                593.9n                          1.080µ        +81.87%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                        B/s                        │              B/s               vs base                 │
Chacha20Poly1305/Open-64-16                                           550.4Mi ± ∞ ¹                   271.6Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-64-16                                           570.3Mi ± ∞ ¹                   265.1Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-64-X-16                                         348.9Mi ± ∞ ¹                   200.7Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-64-X-16                                         353.3Mi ± ∞ ¹                   219.0Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-16                                         2.407Gi ± ∞ ¹                   1.187Gi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-16                                         2.265Gi ± ∞ ¹                   1.112Gi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-X-16                                       2.000Gi ± ∞ ¹                   1.023Gi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-X-16                                      2159.6Mi ± ∞ ¹                  1007.9Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-16                                         2.884Gi ± ∞ ¹                   2.007Gi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-16                                         2.877Gi ± ∞ ¹                   1.815Gi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-X-16                                       2.907Gi ± ∞ ¹                   1.761Gi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-X-16                                       2.844Gi ± ∞ ¹                   1.726Gi ± ∞ ¹        ~ (p=0.100 n=3) ²
geomean                                                               1.397Gi                         786.8Mi        -45.02%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                       B/op                        │              B/op               vs base                │
Chacha20Poly1305/Open-64-16                                             0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-64-16                                             0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-64-X-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-64-X-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-1350-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-1350-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-1350-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-1350-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-8192-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-8192-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-8192-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-8192-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
geomean                                                                           ³                                   +0.00%               ³
¹ need >= 6 samples for confidence interval at level 0.95
² all samples are equal
³ summaries must be >0 to compute geomean

                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                     allocs/op                     │           allocs/op             vs base                │
Chacha20Poly1305/Open-64-16                                             0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-64-16                                             0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-64-X-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-64-X-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-1350-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-1350-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-1350-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-1350-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-8192-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-8192-16                                           0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Open-8192-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-8192-X-16                                         0.000 ± ∞ ¹                      0.000 ± ∞ ¹       ~ (p=1.000 n=3) ²
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
                                │                        sec/op                        │            sec/op              vs base                 │
Chacha20Poly1305/Open-64-16                                               232.3n ± ∞ ¹                    224.7n ± ∞ ¹        ~ (p=0.700 n=3) ²
Chacha20Poly1305/Seal-64-16                                               218.9n ± ∞ ¹                    230.2n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-64-X-16                                             296.7n ± ∞ ¹                    304.1n ± ∞ ¹        ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-64-X-16                                             286.7n ± ∞ ¹                    278.7n ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-16                                             1.879µ ± ∞ ¹                    1.060µ ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-16                                             1.870µ ± ∞ ¹                    1.130µ ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-X-16                                           1.960µ ± ∞ ¹                    1.229µ ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-X-16                                           1.961µ ± ∞ ¹                    1.277µ ± ∞ ¹        ~ (p=0.700 n=3) ²
Chacha20Poly1305/Open-8192-16                                            10.774µ ± ∞ ¹                    3.802µ ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-16                                            10.598µ ± ∞ ¹                    4.203µ ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-X-16                                          10.409µ ± ∞ ¹                    4.333µ ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-X-16                                          10.465µ ± ∞ ¹                    4.420µ ± ∞ ¹        ~ (p=0.100 n=3) ²
geomean                                                                   1.732µ                          1.080µ        -37.62%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                         B/s                          │              B/s               vs base                 │
Chacha20Poly1305/Open-64-16                                              262.7Mi ± ∞ ¹                   271.6Mi ± ∞ ¹        ~ (p=0.700 n=3) ²
Chacha20Poly1305/Seal-64-16                                              278.8Mi ± ∞ ¹                   265.1Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-64-X-16                                            205.7Mi ± ∞ ¹                   200.7Mi ± ∞ ¹        ~ (p=1.000 n=3) ²
Chacha20Poly1305/Seal-64-X-16                                            212.9Mi ± ∞ ¹                   219.0Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-16                                            685.3Mi ± ∞ ¹                  1215.1Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-16                                            688.6Mi ± ∞ ¹                  1139.1Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-X-16                                          657.0Mi ± ∞ ¹                  1047.3Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-X-16                                          656.5Mi ± ∞ ¹                  1007.9Mi ± ∞ ¹        ~ (p=0.700 n=3) ²
Chacha20Poly1305/Open-8192-16                                            725.1Mi ± ∞ ¹                  2055.0Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-16                                            737.2Mi ± ∞ ¹                  1858.9Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-X-16                                          750.5Mi ± ∞ ¹                  1803.1Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-X-16                                          746.5Mi ± ∞ ¹                  1767.5Mi ± ∞ ¹        ~ (p=0.100 n=3) ²
geomean                                                                  490.8Mi                         786.8Mi        +60.31%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                         B/op                         │               B/op                 vs base             │
Chacha20Poly1305/Open-64-16                                                32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-64-16                                                32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-64-X-16                                              32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-64-X-16                                              32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-16                                              32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-16                                              32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-X-16                                            32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-X-16                                            32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-16                                              32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-16                                              32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-X-16                                            32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-X-16                                            32.00 ± ∞ ¹                          0.00 ± ∞ ¹  ~ (p=0.100 n=3) ²
geomean                                                                    32.00                                            ?               ³ ⁴
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
³ summaries must be >0 to compute geomean
⁴ ratios must be >0 to compute geomean

                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                      allocs/op                       │             allocs/op              vs base             │
Chacha20Poly1305/Open-64-16                                                1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-64-16                                                1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-64-X-16                                              1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-64-X-16                                              1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-16                                              1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-16                                              1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-1350-X-16                                            1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-1350-X-16                                            1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-16                                              1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-16                                              1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Open-8192-X-16                                            1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
Chacha20Poly1305/Seal-8192-X-16                                            1.000 ± ∞ ¹                         0.000 ± ∞ ¹  ~ (p=0.100 n=3) ²
geomean                                                                    1.000                                            ?               ³ ⁴
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
³ summaries must be >0 to compute geomean
⁴ ratios must be >0 to compute geomean
