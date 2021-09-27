package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Introduceti raza:")
	var r float64
	fmt.Scan(&r)

	fmt.Println("Introduceti numarul de trageri:")
	var n float64
	fmt.Scan(&n)

	min := 0.0
	max := 2.0 * r
	m := 0.0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int(n); i++ {
		x := (rand.Float64() * (max - min)) + min
		y := (rand.Float64() * (max - min)) + min
		fmt.Println("(",x,",",y,")")
		if (x-r)*(x-r)+(y-r)*(y-r) <= r*r {
			m++
		}
	}
	fmt.Println("Trageri in cerc:", m)
	pi := 4.0 * m / n
	fmt.Printf("%f", pi)
}
