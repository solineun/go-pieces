package ch5

import (
	"errors"
	"io"
	"os"
)

func Cat(args []string) error{
	if len(args) < 2 {
		return errors.New("no file specified")
	}
	f, err := os.Open(args[1])
	if err != nil {
		return err
	}
	defer f.Close()
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