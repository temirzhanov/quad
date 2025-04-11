package piscine

import "fmt"

func QuadA(x, y int) {
	quad(x, y, "o", "o", "o", "o", "-", "-", "|", "|")
}

func QuadB(x, y int) {
	quad(x, y, "/", "\\", "\\", "/", "*", "*", "*", "*")
}

func QuadC(x, y int) {
	quad(x, y, "A", "A", "C", "C", "B", "B", "B", "B")
}

func QuadD(x, y int) {
	quad(x, y, "A", "C", "A", "C", "B", "B", "B", "B")
}

func QuadE(x, y int) {
	quad(x, y, "A", "C", "C", "A", "B", "B", "B", "B")
}

func quad(x, y int, topLeft, topRight, bottomLeft, bottomRight, topBorder, bottomBorder, leftBorder, rightBorder string) {
	if x < 1 || y < 1 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// top left
			if i == 0 && j == 0 {
				fmt.Print(topLeft)
				continue
			}
			// top right
			if i == 0 && j == x-1 {
				fmt.Print(topRight)
				continue
			}
			// bottom left
			if i == y-1 && j == 0 {
				fmt.Print(bottomLeft)
				continue
			}
			// bottom right
			if i == y-1 && j == x-1 {
				fmt.Print(bottomRight)
				continue
			}
			// top border
			if i == 0 && j < x-1 {
				fmt.Print(topBorder)
				continue
			}
			// bottom border
			if i == y-1 && j < x-1 {
				fmt.Print(bottomBorder)
				continue
			}
			// left border
			if i < y-1 && j == 0 {
				fmt.Print(leftBorder)
				continue
			}
			// right border
			if i < y-1 && j == x-1 {
				fmt.Print(rightBorder)
				continue
			}

			fmt.Print(" ")
		}
		fmt.Println()
	}
}
