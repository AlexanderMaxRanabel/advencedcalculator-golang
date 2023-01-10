//go 1.19.4

package main

import "fmt"

func main() {
	var x, y int
	var z string
	var control string
	var recmode int
	var recshort int
	var reclong int

	fmt.Print("Please enter the first number: ")
	fmt.Scan(&x)

	fmt.Print("Please enter the second numbers: ")
	fmt.Scan(&y)
	fmt.Println("Please Enter A Number. Symbols +,-,*,/ are for first four main equations. +!,-!,*!,/! are what if these numbers place is switched. Write AD To enable advenced Calculations:")
	fmt.Scan(&z)

	var d, d1 int
	if x == 0 || y == 0 {
		fmt.Println("Cannot divide by 0")
	} else {
		d = x / y
		d1 = y / x
	}

	var sqcontrol int
	var sq int
	var trimode int
	var triheight int
	var tribase int
	var triareax int
	var triareay int

	var tria, trib, tric int

	a := x + y
	b := x - y
	c := x * y

	a1 := y + x
	b1 := y - x
	c1 := y * x

	switch z {
	case "+":
		fmt.Println(a)
		fmt.Scan(&control)
	case "-":
		fmt.Println(b)
		fmt.Scan(&control)
	case "*":
		fmt.Println(c)
		fmt.Scan(&control)
	case "/":
		fmt.Println(d)
		fmt.Scan(&control)
	case "+!":
		fmt.Println(a1)
		fmt.Scan(&control)
	case "-!":
		fmt.Println(b1)
		fmt.Scan(&control)
	case "*!":
		fmt.Println(c1)
		fmt.Scan(&control)
	case "/!":
		fmt.Println(d1)
		fmt.Scan(&control)
	case "AD":
		fmt.Println("Please enter a mode. 1 to calculate the area of a square. 2 to calculate the perimeter of square. 3 to start rectangle mode. 4 to start triangle mode ")
		fmt.Scan(&sqcontrol)
	default:
		fmt.Println("İnvalid")
		fmt.Scan(&control)
	}

	switch sqcontrol {
	case 1:
		fmt.Println("Please Enter how long a side of the square")
		fmt.Scan(&sq)
		fmt.Println(sq * sq)
		fmt.Scan(&control)
	case 2:
		fmt.Println("Please Enter How Long a side of the square")
		fmt.Scan(&sq)
		fmt.Println(sq + sq + sq + sq)
		fmt.Scan(&control)
	case 3:
		fmt.Println("Please Enter A mode. 1 to calculate the area of the rectangle. 2 to calculate the perimeter of a rectangle")
		fmt.Scan(&recmode)
	case 4:
		fmt.Println("Please Enter A mode. 1 to calculate the area of the triangle. 2 to calculate the perimeter of a triangle")
		fmt.Scan(&trimode)
	default:
		fmt.Println("İnvalid")
	}

	switch recmode {
	case 1:
		fmt.Println("Please enter how long the short side of the rectangle")
		fmt.Scan(&recshort)
		fmt.Println("Please enter how long the long side of the rectangle")
		fmt.Scan(&reclong)

		fmt.Println(recshort * reclong)
		fmt.Scan(&control)
	case 2:
		fmt.Println("Please enter how long the short side of the rectangle")
		fmt.Scan(&recshort)
		fmt.Println("Please enter how long the long side of the rectangle")
		fmt.Scan(&reclong)

		fmt.Println(recshort + recshort + reclong + reclong)
		fmt.Scan(&control)
	}

	switch trimode {
	case 1:
		fmt.Println("Please Enter the height of the triangle")
		fmt.Scan(&triheight)
		fmt.Println("Please Enter how long is the base of the triangle")
		fmt.Scan(&tribase)

		triareax = tribase * triheight
		triareay = triareax / 2

		fmt.Println(triareay)
		fmt.Scan(&control)
	case 2:
		fmt.Println("Please Enter How Long a side of the triangle")
		fmt.Scan(&tria)
		fmt.Println("Please Enter How Long a side of the triangle")
		fmt.Scan(&trib)
		fmt.Println("Please Enter How Long a side of the triangle")
		fmt.Scan(&tric)

		fmt.Println(tria + trib + tric)
		fmt.Scan(&control)
	}
}
