package main

import (
	"fmt"
)

func main() {

	// if, else if & else

	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// if-else

	var point1 = 8840.0

	if percent := point1 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// switch - case
	var point2 = 6

	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
		fmt.Println("you can be better!")
	}

	// switch case dengan gaya if-else

	var point3 = 6

	switch {
	case point3 == 8:
		fmt.Println("perfect")
	case (point3 < 8) && (point3 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// fallthrough dalam switch
	var point4 = 6

	switch {
	case point4 == 8:
		fmt.Println("perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awesome")
		fallthrough
	case point4 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// if bersarang
	var point5 = 10

	if point5 > 7 {
		switch point5 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point5 == 5 {
			fmt.Println("not bad")
		} else if point5 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point5 == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
