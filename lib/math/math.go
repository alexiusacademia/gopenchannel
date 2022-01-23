package math

// Function for linear interpolation
func Interpolate(x1 float64, x3 float64, y1 float64, y2 float64, y3 float64) (x2 float64) {
	x2 = (x1-x3)*(y2-y3)/(y1-y3) + x3

	return
}

// Coordinate type
type Node struct {
	X float64
	Y float64
}
