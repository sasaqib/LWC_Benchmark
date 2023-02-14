tinygo build -o program.uf2 -target=pico main.go

eg: extract a file named timing.go to a uf2 that benchmarks 4-byte TinyJAMBU

tinygo build -o golang4byte.uf2 -target=pico timing.go