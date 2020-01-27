package main

import ("fmt"; "math")

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}
func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

func (c *Circle) perimeter() float64{
	return math.Pi * c.r*2
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) perimeter() float64{
    l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return (l+w)*2
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func main(){
	c := Circle{10, 20, 5}
	fmt.Println(c.area())
	fmt.Println(c.perimeter())
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())	
}