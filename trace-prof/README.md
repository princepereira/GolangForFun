# trace-prof$
Tells how trace can be used

$ go build -o binary
$ time ./binary > file.trace
$ go tool trace file.trace  # This will open a browser
$ go tool trace binary file.trace

Press 'w' : To expand 
Press 's' : To shrink
Click '?' at the top right corner to get more options.
