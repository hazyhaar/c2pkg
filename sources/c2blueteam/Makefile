# Makefile — libc2blueteam (C99 Standalone Zero-Allocation Defense Library)
CC ?= gcc
CFLAGS ?= -O3 -Wall -Wextra -Werror -std=c99 -march=native -fPIC
AR ?= ar
ARFLAGS ?= rcs

LIB = libc2blueteam.a
OBJS = c2blueteam.o ring_buffer.o probe_entropy.o correlator.o rules_simd.o probe_proc.o probe_fs.o probe_mcp.o

.PHONY: all clean test bench lib

all: lib test

lib: $(LIB)

$(LIB): $(OBJS)
	$(AR) $(ARFLAGS) $@ $^

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

test: $(LIB)
	$(CC) $(CFLAGS) test_oracle.c $(LIB) -lm -o bin_test_oracle
	$(CC) $(CFLAGS) test_correlator_oracle.c $(LIB) -lm -o bin_test_correlator
	$(CC) $(CFLAGS) test_payload_classifier.c $(LIB) -lm -o bin_test_classifier
	./bin_test_oracle
	./bin_test_correlator
	./bin_test_classifier

bench: $(LIB)
	$(CC) $(CFLAGS) bench_entropy_c.c $(LIB) -lm -o bin_bench_entropy
	./bin_bench_entropy

clean:
	rm -f *.o $(LIB) bin_test_oracle bin_test_correlator bin_test_classifier bin_bench_entropy
