# pprof-cpu-prof
Tells how cpu utilisation can be measured

$ go build -o binary
$ time ./binary > cpu.prof
$ go tool pprof cpu.prof
$ go tool pprof binary cpu.prof
pprof> top
pprof> top 5
pprof> list main.pixel   # This will give inside function info

pprof> png // Generates a pictorial profile in PNG format
pprof> pdf // Generates a pictorial profile in pdf format
pprof> svg // Generates a pictorial profile in svg format

Some technicals
================
Flat and Flat%  : This will tell the memory allocated in that particular function
Cum and Cum%    : This will tell the cumulative heap data which includes memory created in its stack as well as called func.
