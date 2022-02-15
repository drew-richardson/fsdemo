# fsdemo

Compare and contrast the behavior of different [fs.FS](https://pkg.go.dev/io/fs#FS) implementations.

Call [`run.sh`](run.sh) to start the comparison.

    $ ./run.sh
    *** embed demo ***
    hello world
    open static/not-found: file does not exist
    open /intentinally/malformed: file does not exist
    OK
    *** dirfs demo ***
    hello world
    open internal/data/static/not-found: no such file or directory
    open /intentinally/malformed: invalid argument
    OK
    *** tarfs demo ***
    hello world
    readfile static/not-found: file does not exist
    readfile /intentinally/malformed: invalid argument
    OK
    $
