package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  y := make([][]uint8, dy)

  for i := range y {
    y[i] = make([]uint8, dx)

    for j := range y[i] {
      y[i][j] = uint8(i + j)
    }
  }

  return y
}

func main() {
  pic.Show(Pic)
}
