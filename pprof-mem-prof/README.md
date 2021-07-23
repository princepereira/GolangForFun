# pprof-mem-prof
Tells how heap profile can be measured

$ go build -o binary
$ time ./binary // Running the binary with time // Run this in terminal 1
$ go tool pprof http://localhost:6060/debug/pprof/heap // Generating a heap profile. Run this in terminal 2 when the above program is under execution
pprof> top // List all high heap consuming functions
pprof> top 5 // If your function list is too big, this command will show only top 5 functions
pprof> list main.go_ticker // This will show the internals of function where exectly the heap is formed
pprof> png // Generates a pictorial profile in PNG format
pprof> pdf // Generates a pictorial profile in pdf format
pprof> svg // Generates a pictorial profile in svg format

Some technicals
================
Flat and Flat%  : This will tell the memory allocated in that particular function
Cum and Cum%    : This will tell the cumulative heap data which includes memory created in its stack as well as called func.
        : In above example main.main calls main.iter() which inturn calls main.norm_ticker. If norm_ticker makes 100MB data and main.iter() makes another 20MB, the cum of main.main() is 120MB.
