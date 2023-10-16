package ch5

import (
	"errors"
	"io"
	"os"
)

func CatCancel(args []string) error {
	if len(args) < 2 {
		return errors.New("no file specified")
	}
	f, closer, err := getFile(args[1])
	if err != nil {
		return err
	}
	defer closer()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}
	return nil
}

func getFile(name string) (*os.File, func(), error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return f, func() {
		f.Close()
	}, err
}
