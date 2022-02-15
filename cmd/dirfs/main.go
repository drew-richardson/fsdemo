package main

import (
	"fmt"
	"os"

	"github.com/drew-richardson/fsdemo/internal/demo"
)

func run() error {
	fmt.Println("*** dirfs demo ***")
	return demo.Demo(os.DirFS("internal/data"))
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("OK")
	os.Exit(0)
}
