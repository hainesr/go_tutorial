package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (r13 *rot13Reader) Read(buffer []byte) (int, error) {
  n, err := r13.r.Read(buffer)

  for i := range buffer {
    if (buffer[i] >= 'A' && buffer[i] < 'N') || (buffer[i] >= 'a' && buffer[i] < 'n') {
      buffer[i] += 13
    } else if (buffer[i] >= 'N' && buffer[i] < 'Z') || (buffer[i] >= 'n' && buffer[i] < 'z') {
      buffer[i] -= 13
    }
  }

  return n, err
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
