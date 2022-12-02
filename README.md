# Simple Priority Queue Service Over TCP

## Build
Just run make
```
    $ make 
```

## Usage 
After build command, there are two executable files in `bin/` directory

Runing the queue-server:
```
$ ./bin/queue-server -b 0.0.0.0:1080
```

Running the queue-cli:
```
$ ./bin/queue-cli -h <server_address>
>
```
It gives you an interactive command line.
Currently only `enqueue` and `dequeue` commands are supported.


