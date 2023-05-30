package main

import (
	"math"
	"fmt"
)

func main() {
	var a, c, x, xAnt, cortes, i int64
	a = 1
	c = 1

	for {
		if (a == 0) && (c == 0) {
			break
		}
		fmt.Scanf("%d %d", &a, &c)

		if (1 <= a) && (float64(a) <= (math.Pow(10, 4))) && (1 <= c) && (float64(c) <= (math.Pow(10, 4))) {

			for i = 0; i < c; i++ {
				fmt.Scanf("%d", &x)

				if i == 0 {
					cortes = cortes + a - x
				} else if x < xAnt {
					cortes = cortes + xAnt - x
				}
				xAnt = x

			}
			fmt.Printf("%d\n", cortes)
			cortes = 0

		}
	}
}
