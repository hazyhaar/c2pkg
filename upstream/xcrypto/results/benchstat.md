# x/crypto v0.54.0 — 2026-08-26T09:52Z — go version go1.27.0 linux/amd64 — count=10 benchtime=300ms pin='taskset -c 0-15'

## asm (référence) vs archsimd
goos: linux
goarch: amd64
pkg: golang.org/x/crypto/chacha20poly1305
cpu: Intel(R) Core(TM) i9-14900K
                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                      sec/op                       │            sec/op              vs base                 │
Chacha20Poly1305/Open-64-16                                             102.2n ± 1%                    204.7n ±  1%  +100.24% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-16                                             101.1n ± 1%                    205.8n ±  2%  +103.61% (p=0.000 n=10)
Chacha20Poly1305/Open-64-X-16                                           162.6n ± 1%                    260.6n ±  2%   +60.35% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-X-16                                           158.5n ± 2%                    267.2n ± 11%   +68.63% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-16                                           468.8n ± 1%                    999.0n ±  2%  +113.10% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-16                                           489.5n ± 2%                   1054.0n ±  2%  +115.32% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-X-16                                         528.1n ± 1%                   1115.0n ±  1%  +111.13% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-X-16                                         545.8n ± 1%                   1109.5n ±  2%  +103.28% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-16                                           2.345µ ± 1%                    3.592µ ±  1%   +53.21% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-16                                           2.364µ ± 1%                    3.766µ ±  2%   +59.31% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-X-16                                         2.409µ ± 1%                    3.721µ ±  2%   +54.44% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-X-16                                         2.417µ ± 3%                    3.756µ ±  2%   +55.43% (p=0.000 n=10)
geomean                                                                 536.4n                         973.3n         +81.43%

                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                        B/s                        │              B/s                vs base                │
Chacha20Poly1305/Open-64-16                                            597.3Mi ± 1%                    298.2Mi ±  1%  -50.07% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-16                                            603.4Mi ± 1%                    296.5Mi ±  2%  -50.86% (p=0.000 n=10)
Chacha20Poly1305/Open-64-X-16                                          375.5Mi ± 1%                    234.2Mi ±  2%  -37.63% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-X-16                                          385.3Mi ± 2%                    228.4Mi ± 10%  -40.71% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-16                                          2.682Gi ± 1%                    1.258Gi ±  2%  -53.08% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-16                                          2.569Gi ± 2%                    1.193Gi ±  2%  -53.56% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-X-16                                        2.381Gi ± 1%                    1.128Gi ±  1%  -52.63% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-X-16                                        2.304Gi ± 1%                    1.133Gi ±  2%  -50.81% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-16                                          3.254Gi ± 1%                    2.124Gi ±  1%  -34.73% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-16                                          3.227Gi ± 1%                    2.026Gi ±  2%  -37.22% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-X-16                                        3.167Gi ± 1%                    2.051Gi ±  2%  -35.25% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-X-16                                        3.158Gi ± 3%                    2.031Gi ±  2%  -35.68% (p=0.000 n=10)
geomean                                                                1.547Gi                         873.3Mi        -44.88%

                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                       B/op                        │             B/op               vs base                 │
Chacha20Poly1305/Open-64-16                                            0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-64-16                                            0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-64-X-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-64-X-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-1350-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-1350-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-1350-X-16                                        0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-1350-X-16                                        0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-8192-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-8192-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-8192-X-16                                        0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-8192-X-16                                        0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
geomean                                                                           ²                                  +0.00%                ²
¹ all samples are equal
² summaries must be >0 to compute geomean

                                │ /devhoros/c2simd/upstream/xcrypto/results/asm.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                     allocs/op                     │           allocs/op            vs base                 │
Chacha20Poly1305/Open-64-16                                            0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-64-16                                            0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-64-X-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-64-X-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-1350-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-1350-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-1350-X-16                                        0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-1350-X-16                                        0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-8192-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-8192-16                                          0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Open-8192-X-16                                        0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
Chacha20Poly1305/Seal-8192-X-16                                        0.000 ± 0%                        0.000 ± 0%       ~ (p=1.000 n=10) ¹
geomean                                                                           ²                                  +0.00%                ²
¹ all samples are equal
² summaries must be >0 to compute geomean

## purego (générique) vs archsimd
goos: linux
goarch: amd64
pkg: golang.org/x/crypto/chacha20poly1305
cpu: Intel(R) Core(TM) i9-14900K
                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                        sec/op                        │             sec/op              vs base                │
Chacha20Poly1305/Open-64-16                                                210.4n ± 1%                     204.7n ±  1%   -2.73% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-16                                                198.6n ± 0%                     205.8n ±  2%   +3.65% (p=0.000 n=10)
Chacha20Poly1305/Open-64-X-16                                              275.4n ± 1%                     260.6n ±  2%   -5.36% (p=0.002 n=10)
Chacha20Poly1305/Seal-64-X-16                                              256.4n ± 2%                     267.2n ± 11%   +4.21% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-16                                             1762.5n ± 1%                     999.0n ±  2%  -43.32% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-16                                              1.736µ ± 1%                     1.054µ ±  2%  -39.27% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-X-16                                            1.805µ ± 1%                     1.115µ ±  1%  -38.21% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-X-16                                            1.794µ ± 1%                     1.110µ ±  2%  -38.15% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-16                                              9.409µ ± 1%                     3.592µ ±  1%  -61.82% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-16                                              9.374µ ± 1%                     3.766µ ±  2%  -59.82% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-X-16                                            9.484µ ± 2%                     3.721µ ±  2%  -60.77% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-X-16                                            9.480µ ± 2%                     3.756µ ±  2%  -60.38% (p=0.000 n=10)
geomean                                                                    1.574µ                          973.3n        -38.18%

                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                         B/s                          │              B/s               vs base                 │
Chacha20Poly1305/Open-64-16                                               290.1Mi ± 1%                   298.2Mi ±  1%    +2.80% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-16                                               307.3Mi ± 0%                   296.5Mi ±  2%    -3.51% (p=0.000 n=10)
Chacha20Poly1305/Open-64-X-16                                             221.6Mi ± 1%                   234.2Mi ±  2%    +5.68% (p=0.002 n=10)
Chacha20Poly1305/Seal-64-X-16                                             238.0Mi ± 2%                   228.4Mi ± 10%    -4.03% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-16                                             730.4Mi ± 1%                  1288.5Mi ±  2%   +76.41% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-16                                             741.9Mi ± 1%                  1221.5Mi ±  2%   +64.64% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-X-16                                           713.6Mi ± 1%                  1154.9Mi ±  1%   +61.85% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-X-16                                           717.6Mi ± 1%                  1160.4Mi ±  2%   +61.69% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-16                                             830.3Mi ± 1%                  2175.2Mi ±  1%  +161.98% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-16                                             833.5Mi ± 1%                  2074.4Mi ±  2%  +148.89% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-X-16                                           823.8Mi ± 2%                  2099.7Mi ±  2%  +154.88% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-X-16                                           824.1Mi ± 2%                  2079.8Mi ±  2%  +152.37% (p=0.000 n=10)
geomean                                                                   539.9Mi                        873.3Mi         +61.75%

                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                         B/op                         │           B/op             vs base                     │
Chacha20Poly1305/Open-64-16                                                 32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-16                                                 32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-64-X-16                                               32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-X-16                                               32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-16                                               32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-16                                               32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-X-16                                             32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-X-16                                             32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-16                                               32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-16                                               32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-X-16                                             32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-X-16                                             32.00 ± 0%                   0.00 ± 0%  -100.00% (p=0.000 n=10)
geomean                                                                     32.00                                   ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

                                │ /devhoros/c2simd/upstream/xcrypto/results/purego.txt │ /devhoros/c2simd/upstream/xcrypto/results/archsimd.txt │
                                │                      allocs/op                       │         allocs/op          vs base                     │
Chacha20Poly1305/Open-64-16                                                 1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-16                                                 1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-64-X-16                                               1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-64-X-16                                               1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-16                                               1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-16                                               1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-1350-X-16                                             1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-1350-X-16                                             1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-16                                               1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-16                                               1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Open-8192-X-16                                             1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
Chacha20Poly1305/Seal-8192-X-16                                             1.000 ± 0%                  0.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                                                                     1.000                                   ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean
