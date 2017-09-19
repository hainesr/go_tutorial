package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
  z := 1.0

	if x < 0.0 {
		return 0.0, ErrNegativeSqrt(x)
	}

  for i := 0; i < 10; i++ {
    z = z - (((z * z) - x) / (2 * z))
    i += 1
  }

  return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
