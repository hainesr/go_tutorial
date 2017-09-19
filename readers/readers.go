package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(buffer []byte) (int, error) {
  for i := range buffer {
    buffer[i] = byte('A')
  }

  return len(buffer), nil
}

func main() {
	reader.Validate(MyReader{})
}
