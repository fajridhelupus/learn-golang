package main

import "fmt"

func main() {

	//penggunaan if_else
	point := 0
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	//penggunaan switch case

	point2 := 8
	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("good")
	default:
		fmt.Println("lumayan")
	}

	//penggunaan switch case dengan banyak kondisi

	point3 := 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awosome")
	default:
		fmt.Println("lumayan")
	}

	//switch case gaya if_else
	// var point := 6
	// switch {
	// case point == 8:
	// 	fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 	fmt.Println("awesome")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("you need to learn more")
	// 	}
	// }

	point4 := 1
	switch {
	case point4 == 8:
		fmt.Println("perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awsome")
	default:
		{
			fmt.Println("lumayan")
			fmt.Println("butuh belajar lagi")
		}
	}

	//seleksi kondisi bersarang
	var point5 = 0
	if point > 7 {
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
