package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dot struct {
	X, Y int
}

type fold struct {
	ToLeft bool //false: toUp
	Num    int
}

func readInput(path string) ([]dot, []fold) {
	dots := []dot{}
	folds := []fold{}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, ",")
		if len(arr) == 2 {
			x, _ := strconv.Atoi(arr[0])
			y, _ := strconv.Atoi(arr[1])
			dots = append(dots, dot{X: x, Y: y})
		}
		if strings.Contains(arr[0], "=") {
			fs := strings.Split(arr[0], "=")
			num, _ := strconv.Atoi(fs[1])
			folds = append(folds, fold{ToLeft: strings.Contains(fs[0], "x"), Num: num})
		}
	}
	return dots, folds
}

func getSize(dots map[dot]bool) (int, int, int, int) {
	xMin, yMin, xMax, yMax := 100, 100, 0, 0
	for d := range dots {
		if d.X > xMax {
			xMax = d.X
		}
		if d.X < xMin {
			xMin = d.X
		}
		if d.Y > yMax {
			yMax = d.Y
		}
		if d.Y < yMin {
			yMin = d.Y
		}
	}
	return xMin, yMin, xMax, yMax
}

func part01(dots []dot, folds []fold) int {
	for _, f := range folds {
		newDots := map[dot]bool{}
		if f.ToLeft {
			for _, d := range dots {
				if d.X < f.Num {
					newDots[dot{X: d.X + (2 * (f.Num - d.X)), Y: d.Y}] = true
				} else if d.X > f.Num {
					newDots[dot{X: d.X, Y: d.Y}] = true
				}
			}
		} else {
			for _, d := range dots {
				if d.Y < f.Num {
					newDots[dot{X: d.X, Y: d.Y}] = true
				} else if d.Y > f.Num {
					newDots[dot{X: d.X, Y: d.Y - (2 * (d.Y - f.Num))}] = true
				}
			}
		}

		fmt.Println(f)
		fmt.Println(len(newDots))
		// for v := range newDots {
		// 	fmt.Println(v)
		// }
		// xMin, yMin, xMax, yMax := getSize(newDots)
		// fmt.Println(xMin, yMin, xMax, yMax)
		break
	}
	return 0
}

func main() {
	part01(readInput("example.txt"))
	part01(readInput("input.txt"))
}
