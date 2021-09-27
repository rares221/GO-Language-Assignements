package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func ReadSetOfPoints(file *os.File, noOfPointsTop int) ([]float64, []float64) {

	var x []float64
	var y []float64

	for i := 0; i < noOfPointsTop; i++ {
		var currentX, currentY float64
		fmt.Fscan(file, &currentX)
		fmt.Fscan(file, &currentY)
		x = append(x, currentX)
		y = append(y, currentY)
	}

	return x, y
}

func PrintPoints(x, y []float64, noOfPoints int, where string) {
	fmt.Println("Points", where, ":")
	for i := 0; i < noOfPoints; i++ {
		fmt.Println(x[i], y[i])
	}
	fmt.Println("")
}

func DrawSimpleInterpolatedLine() {
	// Open file
	file, err := os.Open("input_simple.txt")
	if err != nil {
		panic(err)
	}

	var noOfPoints int
	fmt.Fscan(file, &noOfPoints)

	fmt.Println("Number of points:", noOfPoints)
	fmt.Println("")

	// Read points
	var x []float64
	var y []float64

	for i := 0; i < noOfPoints; i++ {
		var currentX, currentY float64
		fmt.Fscan(file, &currentX)
		fmt.Fscan(file, &currentY)
		x = append(x, currentX)
		y = append(y, currentY)
	}

	fmt.Println("Points:")
	for i := 0; i < noOfPoints; i++ {
		fmt.Println(x[i], y[i])
	}
	fmt.Println("")

	// Interpolare
	// CubicSecondDeriv = valorile derivatei de ordin 2 in capete
	s := newSpline(x, y, CubicSecondDeriv, 0, 0)
	fmt.Println("Valoarea spline in x=7:", s.At(7))

	p := 20                                       // numarul de puncte de pe grafic
	var h = (x[noOfPoints-1] - x[0]) / float64(p) // distanta dintre puncte

	output, err := os.Create("..\\Pointshow\\output_simple.obj")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(output)

	// incepand de la primul punct de dupa x[0]
	for i := 1; i < p; i++ {
		xTemp := x[0] + float64(i)*h
		yTemp := s.At(xTemp)
		zTemp := 100.0

		fmt.Fprintf(w, "v %f %f %f \n", xTemp, yTemp, zTemp)
	}
	w.Flush()
}

func DrawComplexInterpolatedLines() {

	// Open file
	file, err := os.Open("input_complex.txt")
	if err != nil {
		panic(err)
	}

	var noOfPointsTop int
	var noOfPointsBottom int
	fmt.Fscan(file, &noOfPointsTop)
	fmt.Fscan(file, &noOfPointsBottom)

	fmt.Println("Number of points:", noOfPointsTop, "and", noOfPointsBottom)
	fmt.Println("")

	// Read top set of points
	xTop, yTop := ReadSetOfPoints(file, noOfPointsTop)
	PrintPoints(xTop, yTop, noOfPointsTop, "top")

	// Read bottom set of points
	xBottom, yBottom := ReadSetOfPoints(file, noOfPointsBottom)
	PrintPoints(xBottom, yBottom, noOfPointsBottom, "bottom")

	// Interpolare
	sTop := newSpline(xTop, yTop, CubicSecondDeriv, 0, 0)
	fmt.Println("Valoarea spline (top) in x=3:", sTop.At(3))
	sBottom := newSpline(xBottom, yBottom, CubicSecondDeriv, 0, 0)
	fmt.Println("Valoarea spline (bottom) in x=3:", sBottom.At(3))

	p := 20                                                // numarul de puncte de pe grafic
	var h = (xTop[noOfPointsTop-1] - xTop[0]) / float64(p) // distanta dintre puncte

	output, err := os.Create("..\\Pointshow\\output_complex.obj")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(output)

	// incepand de la primul punct de dupa x[0]
	zTemp := 100.0
	for i := 1; i < p; i++ {
		xTempTop := xTop[0] + float64(i)*h
		yTempTop := sTop.At(xTempTop)
		fmt.Fprintf(w, "v %f %f %f \n", xTempTop, yTempTop, zTemp)
	}
	for i := 1; i < p; i++ {
		xTempBottom := xBottom[0] + float64(i)*h
		yTempBottom := sBottom.At(xTempBottom) + 0.75
		fmt.Fprintf(w, "v %f %f %f \n", xTempBottom, yTempBottom, zTemp)
	}
	w.Flush()
}

func DrawComplexInterpolatedLinesWithSlices() {

	// Open file
	file, err := os.Open("input_complex_slices.txt")
	if err != nil {
		panic(err)
	}

	var noOfSlices int
	fmt.Fscan(file, &noOfSlices)

	output, err := os.Create("..\\Pointshow\\output_complex_slices.obj")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(output)

	for iteration := 0; iteration < noOfSlices; iteration++ {

		var noOfPointsTop int
		var noOfPointsBottom int
		fmt.Fscan(file, &noOfPointsTop)
		fmt.Fscan(file, &noOfPointsBottom)

		fmt.Println("Number of slices:", noOfSlices)
		fmt.Println("Number of points:", noOfPointsTop, "and", noOfPointsBottom)
		fmt.Println("")

		// Read top set of points
		xTop, yTop := ReadSetOfPoints(file, noOfPointsTop)
		PrintPoints(xTop, yTop, noOfPointsTop, "top")

		// Read bottom set of points
		xBottom, yBottom := ReadSetOfPoints(file, noOfPointsBottom)
		PrintPoints(xBottom, yBottom, noOfPointsBottom, "bottom")

		// Interpolare
		sTop := newSpline(xTop, yTop, CubicSecondDeriv, 0, 0)
		fmt.Println("Valoarea spline (top) in x=3:", sTop.At(3))
		sBottom := newSpline(xBottom, yBottom, CubicSecondDeriv, 0, 0)
		fmt.Println("Valoarea spline (bottom) in x=3:", sBottom.At(3))

		p := 20
		var h = (xTop[noOfPointsTop-1] - xTop[0]) / float64(p)

		zTemp := 4
		for i := 1; i < p; i++ {
			xTempTop := xTop[0] + float64(i)*h
			yTempTop := sTop.At(xTempTop)
			fmt.Fprintf(w, "v %f %f %d \n", xTempTop, yTempTop, iteration*zTemp)
		}
		for i := 1; i < p; i++ {
			xTempBottom := xBottom[0] + float64(i)*h
			yTempBottom := sBottom.At(xTempBottom)
			fmt.Fprintf(w, "v %f %f %d \n", xTempBottom, yTempBottom, iteration*zTemp)
		}
	}
	w.Flush()
}

func DrawTriangulation(closeFigure bool) {

	// Open file
	file, err := os.Open("input_complex_slices.txt")
	if err != nil {
		panic(err)
	}

	var noOfSlices int
	fmt.Fscan(file, &noOfSlices)

	output, err := os.Create("..\\Pointshow\\output_triangulation.obj")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(output)

	p := 20
	for iteration := 0; iteration < noOfSlices; iteration++ {

		var noOfPointsTop int
		var noOfPointsBottom int
		fmt.Fscan(file, &noOfPointsTop)
		fmt.Fscan(file, &noOfPointsBottom)

		fmt.Println("Number of slices:", noOfSlices)
		fmt.Println("Number of points:", noOfPointsTop, "and", noOfPointsBottom)
		fmt.Println("")

		// Read top set of points
		xTop, yTop := ReadSetOfPoints(file, noOfPointsTop)
		PrintPoints(xTop, yTop, noOfPointsTop, "top")

		// Read bottom set of points
		xBottom, yBottom := ReadSetOfPoints(file, noOfPointsBottom)
		PrintPoints(xBottom, yBottom, noOfPointsBottom, "bottom")

		// Interpolare
		var xCoord = 3.0
		sTop := newSpline(xTop, yTop, CubicSecondDeriv, 0, 0)
		fmt.Println("Valoarea spline (top) in x=3:", sTop.At(xCoord))
		sBottom := newSpline(xBottom, yBottom, CubicSecondDeriv, 0, 0)
		fmt.Println("Valoarea spline (bottom) in x=3:", sBottom.At(xCoord))

		var h = (xTop[noOfPointsTop-1] - xTop[0]) / float64(p)

		zTemp := 4
		for i := 1; i < p; i++ {
			xTempTop := xTop[0] + float64(i)*h
			yTempTop := sTop.At(xTempTop)
			fmt.Fprintf(w, "v %f %f %d \n", xTempTop, yTempTop, iteration*zTemp)
		}
		for i := 1; i < p; i++ {
			xTempBottom := xBottom[0] + float64(i)*h
			yTempBottom := sBottom.At(xTempBottom)
			fmt.Fprintf(w, "v %f %f %d \n", xTempBottom, yTempBottom, iteration*zTemp)
		}
	}

	// draw faces and close the figure => no open-ended faces
	if closeFigure {

		p--
		for i := 1; i < p; i++ {
			fmt.Fprintf(w, "f %d %d %d \n", i, 2*p+i, 2*p+i+1)
			fmt.Fprintf(w, "f %d %d %d \n", i, i+1, 2*p+i+1)
			fmt.Fprintf(w, "f %d %d %d \n", p+i, 3*p+i, 3*p+i+1)
			fmt.Fprintf(w, "f %d %d %d \n", p+i, p+i+1, 3*p+i+1)
		}

		p++
		fmt.Fprintf(w, "f %d %d %d \n", 1, p, 3*p-2)
		fmt.Fprintf(w, "f %d %d %d \n", 1, 2*p-1, 3*p-2)
		fmt.Fprintf(w, "f %d %d %d \n", p-1, 3*p-3, 4*p-4)
		fmt.Fprintf(w, "f %d %d %d \n", p-1, 2*p-2, 4*p-4)

	} else {
		for i := 1; i < p-4; i++ {
			fmt.Fprintf(w, "f %d %d %d \n", i, 2*p+i, 2*p+i+1)
			fmt.Fprintf(w, "f %d %d %d \n", i, i+1, 2*p+i+1)
			fmt.Fprintf(w, "f %d %d %d \n", p+i, 3*p+i, 3*p+i+1)
			fmt.Fprintf(w, "f %d %d %d \n", p+i, p+i+1, 3*p+i+1)
		}
	}

	w.Flush()
}

type Point struct {
	x, y, z float64
}

func CalculateDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2) + math.Pow(p1.z-p2.z, 2))
}

func CalculateArea(a, b, c float64) float64 {
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func DrawTriangulationWithHeron() {

	// Open file
	file, err := os.Open("input_complex_slices_heron.txt")
	if err != nil {
		panic(err)
	}

	var noOfSlices int
	fmt.Fscan(file, &noOfSlices)

	output, err := os.Create("..\\Pointshow\\output_triangulation_heron.obj")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(output)

	var vertices []Point

	p := 20
	for iteration := 0; iteration < noOfSlices; iteration++ {

		var noOfPointsTop int
		var noOfPointsBottom int
		fmt.Fscan(file, &noOfPointsTop)
		fmt.Fscan(file, &noOfPointsBottom)

		fmt.Println("Number of slices:", noOfSlices)
		fmt.Println("Number of points:", noOfPointsTop, "and", noOfPointsBottom)
		fmt.Println("")

		// Read top set of points
		xTop, yTop := ReadSetOfPoints(file, noOfPointsTop)
		PrintPoints(xTop, yTop, noOfPointsTop, "top")

		// Read bottom set of points
		xBottom, yBottom := ReadSetOfPoints(file, noOfPointsBottom)
		PrintPoints(xBottom, yBottom, noOfPointsBottom, "bottom")

		// Interpolare
		var xCoord = 3.0
		sTop := newSpline(xTop, yTop, CubicSecondDeriv, 0, 0)
		fmt.Println("Valoarea spline (top) in x=3:", sTop.At(xCoord))
		sBottom := newSpline(xBottom, yBottom, CubicSecondDeriv, 0, 0)
		fmt.Println("Valoarea spline (bottom) in x=3:", sBottom.At(xCoord))

		var h = (xTop[noOfPointsTop-1] - xTop[0]) / float64(p)

		zTemp := 4
		for i := 1; i < p; i++ {
			xTempTop := xTop[0] + float64(i)*h
			yTempTop := sTop.At(xTempTop)
			fmt.Fprintf(w, "v %f %f %d \n", xTempTop, yTempTop, iteration*zTemp)
			vertices = append(vertices, Point{xTempTop, yTempTop, float64(iteration * zTemp)})
		}
		for i := 1; i < p; i++ {
			xTempBottom := xBottom[0] + float64(i)*h
			yTempBottom := sBottom.At(xTempBottom)
			fmt.Fprintf(w, "v %f %f %d \n", xTempBottom, yTempBottom, iteration*zTemp)
			vertices = append(vertices, Point{xTempBottom, yTempBottom, float64(iteration * zTemp)})
		}
	}

	area := 0.0
	var a, b, c float64
	p--
	for i := 1; i < p; i++ {
		fmt.Fprintf(w, "f %d %d %d \n", i, 2*p+i, 2*p+i+1)
		a = CalculateDistance(vertices[i-1], vertices[2*p+i-1])
		b = CalculateDistance(vertices[2*p+i-1], vertices[2*p+i])
		c = CalculateDistance(vertices[2*p+i], vertices[i-1])
		area += CalculateArea(a, b, c)

		fmt.Fprintf(w, "f %d %d %d \n", i, i+1, 2*p+i+1)
		a = CalculateDistance(vertices[i-1], vertices[i])
		b = CalculateDistance(vertices[i], vertices[2*p+i])
		area += CalculateArea(a, b, c)

		fmt.Fprintf(w, "f %d %d %d \n", p+i, 3*p+i, 3*p+i+1)
		a = CalculateDistance(vertices[p+i-1], vertices[3*p+i-1])
		b = CalculateDistance(vertices[3*p+i-1], vertices[3*p+i])
		c = CalculateDistance(vertices[3*p+i], vertices[p+i-1])
		area += CalculateArea(a, b, c)

		fmt.Fprintf(w, "f %d %d %d \n", p+i, p+i+1, 3*p+i+1)
		a = CalculateDistance(vertices[p+i-1], vertices[p+i])
		b = CalculateDistance(vertices[p+i], vertices[3*p+i])
		area += CalculateArea(a, b, c)
	}

	p++

	fmt.Fprintf(w, "f %d %d %d \n", 1, p, 3*p-2)
	a = CalculateDistance(vertices[0], vertices[p-1])
	b = CalculateDistance(vertices[p-1], vertices[3*p-3])
	c = CalculateDistance(vertices[3*p-3], vertices[0])
	area += CalculateArea(a, b, c)

	fmt.Fprintf(w, "f %d %d %d \n", 1, 2*p-1, 3*p-2)
	a = CalculateDistance(vertices[0], vertices[2*p-2])
	b = CalculateDistance(vertices[2*p-2], vertices[3*p-3])
	area += CalculateArea(a, b, c)

	fmt.Fprintf(w, "f %d %d %d \n", p-1, 3*p-3, 4*p-4)
	a = CalculateDistance(vertices[p-2], vertices[3*p-4])
	b = CalculateDistance(vertices[3*p-4], vertices[4*p-5])
	c = CalculateDistance(vertices[4*p-5], vertices[p-2])
	area += CalculateArea(a, b, c)

	fmt.Fprintf(w, "f %d %d %d \n", p-1, 2*p-2, 4*p-4)
	a = CalculateDistance(vertices[p-2], vertices[2*p-3])
	b = CalculateDistance(vertices[2*p-3], vertices[4*p-5])
	area += CalculateArea(a, b, c)

	fmt.Println("")
	fmt.Println("Calculated area:", area)

	w.Flush()
}

func main() {
	fmt.Println("")
	fmt.Println("Spline.go")
	fmt.Println("")

	//DrawSimpleInterpolatedLine()
	//DrawComplexInterpolatedLines()
	//DrawComplexInterpolatedLinesWithSlices()
	//DrawTriangulation(false)
	DrawTriangulationWithHeron()
}
