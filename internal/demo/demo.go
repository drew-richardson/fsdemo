package demo

import (
	"fmt"
	"io/fs"
)

func Demo(fsys fs.FS) error {
	data, err := fs.ReadFile(fsys, "static/contents")
	if err != nil {
		return err
	}
	fmt.Print(string(data))
	_, err = fs.ReadFile(fsys, "static/not-found")
	fmt.Println(err)
	_, err = fs.ReadFile(fsys, "/intentionally/malformed")
	fmt.Println(err)
	return nil
}
