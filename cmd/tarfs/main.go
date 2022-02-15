package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/drew-richardson/fsdemo/internal/demo"
	"github.com/nlepage/go-tarfs"
)

func run() error {
	fmt.Println("*** tarfs demo ***")
	f, err := os.Open("static.tar")
	if err != nil {
		return err
	}
	var fsys fs.FS
	if fsys, err = tarfs.New(f); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return demo.Demo(fsys)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("OK")
	os.Exit(0)
}
