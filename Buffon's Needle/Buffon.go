package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var L float64
	fmt.Println("Introduceti lungimea dintre linii:")
	fmt.Scan(&L)


		var linii float64
		fmt.Println("Introduceti cate linii doriti sa trasati:")
		fmt.Scan(&linii)



	var n int
	fmt.Println("Introduceti numarul de aruncari:")
	fmt.Scan(&n)

	minY := 0.0
	maxY := L
	minAlpha := 0.0
	maxAlpha := math.Pi / 2.0
	m := 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		y := (rand.Float64() * (maxY - minY)) + minY
		alpha := (rand.Float64() * (maxAlpha - minAlpha)) + minAlpha
		y1 := y + L*math.Sin(alpha)
		if int(y/L) != int(y1/L) {
			m++
		}
		fmt.Println("y=", y, "  alpha=", alpha, "  y1=", y1)

			for j := 0; j <= int(linii); j++ {
				if y <= float64(j)*L && y1 >= float64(j)*L {
					m++
				}
			}

	}

	fmt.Println("Numarul de ace care au cazut pe linie:", m)

	pi := 2.0 * float64(n) / float64(m)
	fmt.Println("Pi=", pi)

}
