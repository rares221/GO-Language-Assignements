package main

import (
	"fmt"
	"math"
	"os"
)

type Point struct {
	x float64
	y float64
}

func LeastSquaresMethod(points *[]Point) (a float64, b float64) {

	n := float64(len(*points))

	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumXX := 0.0

	for _, p := range *points {
		sumX += p.x
		sumY += p.y
		sumXY += p.x * p.y
		sumXX += p.x * p.x
	}

	base := n*sumXX - sumX*sumX
	a = (n*sumXY - sumX*sumY) / base
	b = (sumXX*sumY - sumXY*sumX) / base

	return a, b
}

func Rotate(p Point, alpha float64) Point {
	newPoint := Point{
		x: p.x*math.Cos(alpha) - p.y*math.Sin(alpha),
		y: p.x*math.Sin(alpha) + p.y*math.Cos(alpha),
	}
	return newPoint
}

func main() {
	points := make([]Point, 0)
	var numberOfPoints int
	alpha := math.Pi / 2

	file, err := os.Open("points.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fscan(file, &numberOfPoints)

	for i := 0; i < numberOfPoints; i++ {
		var currentPoint Point
		fmt.Fscan(file, &currentPoint.x)
		fmt.Fscan(file, &currentPoint.y)
		fmt.Print(currentPoint, " ")
		points = append(points, currentPoint)
	}

	fmt.Print("\n")
	a, b := LeastSquaresMethod(&points)
	fmt.Print("Before rotation", "\n")
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)

	for i := 0; i < numberOfPoints; i++ {
		points[i] = Rotate(points[i], alpha)
	}

	a, b = LeastSquaresMethod(&points)
	fmt.Print("After rotation", "\n")
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)
}
