package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func Map(Size_x int, Size_y int) [][]string {
	Map := make([][]string, Size_x)
	for i := 0; i < Size_x; i++ {
		Map[i] = make([]string, Size_y)
	}
	return Map
}
func DefinMap(Difficulty int, Map_With_Size [][]string, Size_x int) ([][]string, int) {
	Count_Bomber := 0
	for i := 0; i < Size_x; i++ {
		for j := range Map_With_Size[i] {
			Rdn_Int := rand.Intn(Difficulty) // <-- Difficult
			if Rdn_Int == 1 || Rdn_Int == 2 || Rdn_Int == 3 || Rdn_Int == 4 {
				Map_With_Size[i][j] = "0"
			} else {
				Map_With_Size[i][j] = "x"
			}
		}
	}
	//fmt.Println(Map_With_Size)
	for i := range Map_With_Size {
		for j := range Map_With_Size[i] {
			if Map_With_Size[i][j] == "x" {
				Count_Bomber++
			}
		}
	}
	return Map_With_Size, Count_Bomber
}

func StartNewgame(Size_x int, Size_y int, Difficulty int) ([][]string, int) {
	Map_With_Size := Map(Size_x, Size_y)
	return DefinMap(Difficulty, Map_With_Size, Size_x)
}

func initMap(Map [][]string, Bombe int) [][]string {
	for i := 0; i < len(Map)-1; i++ {
		for j := range Map[i] {
			POS := Map[i][j]
			if POS == "x" {
				//Conditions array
				if i == 0 && j == 0 { //[0][0] Exemple : [X . . . .]
					if Map[0][1] != "x" {
						a_a, _ := strconv.Atoi(Map[0][1])
						Map[0][1] = strconv.Itoa(a_a + 1)
						a_a = 0
					}
					if Map[1][0] != "x" {
						c, _ := strconv.Atoi(Map[1][0])
						Map[1][0] = strconv.Itoa(c + 1)
						c = 0
					}
					if Map[1][1] != "x" {
						a_c, _ := strconv.Atoi(Map[1][1])
						Map[1][1] = strconv.Itoa(a_c + 1)
						a_c = 0
					}
				} else if i == 0 && j == len(Map[0])-1 { //[0][max] Exemple : [. . . . X]

					if Map[0][len(Map[0])-2] != "x" {
						a_a, _ := strconv.Atoi(Map[0][len(Map[0])-2])
						Map[0][len(Map[0])-2] = strconv.Itoa(a_a + 1)
						a_a = 0
					}
					if Map[1][len(Map[0])-1] != "x" {
						a_b, _ := strconv.Atoi(Map[1][len(Map[0])-1])
						Map[1][len(Map[0])-1] = strconv.Itoa(a_b + 1)
						a_b = 0
					}
					if Map[1][len(Map[0])-2] != "x" {
						a_c, _ := strconv.Atoi(Map[1][len(Map[0])-2])
						Map[1][len(Map[0])-2] = strconv.Itoa(a_c + 1)
						a_c = 0

					}

				} else if i == len(Map)-1 && j == 0 { //[max][0] Exemple : [X . . . .]

					if Map[len(Map)-1][1] != "x" {
						a_a, _ := strconv.Atoi(Map[len(Map)-1][1])
						Map[len(Map)-1][1] = strconv.Itoa(a_a + 1)
						a_a = 0
					}

					if Map[len(Map)-2][0] != "x" {
						a_b, _ := strconv.Atoi(Map[len(Map)-2][0])
						Map[len(Map)-2][0] = strconv.Itoa(a_b + 1)
						a_b = 0
					}

					if Map[len(Map)-2][1] != "x" {

						a_c, _ := strconv.Atoi(Map[len(Map)-2][1])
						Map[len(Map)-2][1] = strconv.Itoa(a_c + 1)
						a_c = 0

					}
				} else if i == len(Map)-1 && j == len(Map[len(Map)-1])-1 { //[max][max] Exemple : [. . . . X]

					if Map[len(Map)-1][len(Map[0])-2] != "x" {
						a_a, _ := strconv.Atoi(Map[len(Map)-1][len(Map[0])-2])
						Map[len(Map)-1][len(Map[0])-2] = strconv.Itoa(a_a + 1)
						a_a = 0
					}
					if Map[len(Map)-2][len(Map[0])-2] != "x" {
						a_b, _ := strconv.Atoi(Map[len(Map)-2][len(Map[0])-2])
						Map[len(Map)-2][len(Map[0])-2] = strconv.Itoa(a_b + 1)
						a_b = 0
					}

					if Map[len(Map)-2][len(Map[0])-1] != "x" {

						a_c, _ := strconv.Atoi(Map[len(Map)-2][len(Map[0])-1])
						Map[len(Map)-2][len(Map[0])-1] = strconv.Itoa(a_c + 1)
						a_c = 0
					}
				} else if i == 0 {

					if j != 0 && j != len(Map[0])-1 { //In [0][...] Exemple : [. X X X .]
						if Map[0][j-1] != "x" {
							a_a, _ := strconv.Atoi(Map[0][j-1])
							Map[0][j-1] = strconv.Itoa(a_a + 1)
							a_a = 0
						}
						if Map[0][j+1] != "x" {
							a_b, _ := strconv.Atoi(Map[0][j+1])
							Map[0][j+1] = strconv.Itoa(a_b + 1)
							a_b = 0
						}

						if Map[1][j-1] != "x" {
							a_c, _ := strconv.Atoi(Map[1][j-1])
							Map[1][j-1] = strconv.Itoa(a_c + 1)
							a_c = 0
						}
						if Map[1][j] != "x" {
							a_d, _ := strconv.Atoi(Map[1][j])
							Map[1][j] = strconv.Itoa(a_d + 1)
							a_d = 0
						}
						if Map[1][j+1] != "x" {
							a_e, _ := strconv.Atoi(Map[1][j+1])
							Map[1][j+1] = strconv.Itoa(a_e + 1)
							a_e = 0
						}
					}
				} else if i == len(Map)-1 {

					if j != 0 && j != len(Map[0])-1 { //In [Max][...] Exemple : [. X X X .]

						if Map[len(Map)-1][j-1] != "x" {
							a_a, _ := strconv.Atoi(Map[len(Map)-1][j-1])
							Map[len(Map)-1][j-1] = strconv.Itoa(a_a + 1)
							a_a = 0
						}
						if Map[len(Map)-1][j+1] != "x" {
							a_b, _ := strconv.Atoi(Map[len(Map)-1][j+1])
							Map[len(Map)-1][j+1] = strconv.Itoa(a_b + 1)
							a_b = 0
						}

						if Map[len(Map)-2][j-1] != "x" {
							a_c, _ := strconv.Atoi(Map[len(Map)-2][j-1])
							Map[len(Map)-2][j-1] = strconv.Itoa(a_c + 1)
							a_c = 0
						}
						if Map[len(Map)-2][j] != "x" {
							a_d, _ := strconv.Atoi(Map[len(Map)-2][j])
							Map[len(Map)-2][j] = strconv.Itoa(a_d + 1)
							a_d = 0
						}
						if Map[len(Map)-2][j+1] != "x" {
							a_e, _ := strconv.Atoi(Map[len(Map)-2][j+1])
							Map[len(Map)-2][j+1] = strconv.Itoa(a_e + 1)
							a_e = 0
						}
					}
				} else if i != 0 && i != len(Map)-1 { //border left
					if j == 0 {
						if Map[i-1][0] != "x" {
							a_a, _ := strconv.Atoi(Map[i-1][0])
							Map[i-1][0] = strconv.Itoa(a_a + 1)
							a_a = 0
						}
						if Map[i-1][1] != "x" {
							a_b, _ := strconv.Atoi(Map[i-1][1])
							Map[i-1][1] = strconv.Itoa(a_b + 1)
							a_b = 0
						}
						if Map[i][1] != "x" {
							a_c, _ := strconv.Atoi(Map[i][1])
							Map[i][1] = strconv.Itoa(a_c + 1)
							a_c = 0
						}
						if Map[i+1][0] != "x" {
							a_d, _ := strconv.Atoi(Map[i+1][0])
							Map[i+1][0] = strconv.Itoa(a_d + 1)
							a_d = 0
						}
						if Map[i+1][1] != "x" {
							a_e, _ := strconv.Atoi(Map[i+1][1])
							Map[i+1][1] = strconv.Itoa(a_e + 1)
							a_e = 0
						}
					} else if j == len(Map[0])-1 { //Border left
						if Map[i-1][len(Map[0])-1] != "x" {
							a_a, _ := strconv.Atoi(Map[i-1][len(Map[0])-1])
							Map[i-1][len(Map[0])-1] = strconv.Itoa(a_a + 1)
							a_a = 0

						}
						if Map[i-1][len(Map[0])-2] != "x" {
							a_b, _ := strconv.Atoi(Map[i-1][len(Map[0])-2])
							Map[i-1][len(Map[0])-2] = strconv.Itoa(a_b + 1)
							a_b = 0
						}

						if Map[i][len(Map[0])-2] != "x" {
							a_c, _ := strconv.Atoi(Map[i][len(Map[0])-2])
							Map[i][len(Map[0])-2] = strconv.Itoa(a_c + 1)
							a_c = 0
						}
						if Map[i+1][len(Map[0])-1] != "x" {
							a_d, _ := strconv.Atoi(Map[i+1][len(Map[0])-1])
							Map[i+1][len(Map[0])-1] = strconv.Itoa(a_d + 1)
							a_d = 0
						}
						if Map[i+1][len(Map[0])-2] != "x" {
							a_e, _ := strconv.Atoi(Map[i+1][len(Map[0])-2])
							Map[i+1][len(Map[0])-2] = strconv.Itoa(a_e + 1)
							a_e = 0
						}
					} else { //the rest
						if Map[i-1][j-1] != "x" {
							a_a, _ := strconv.Atoi(Map[i-1][j-1])
							Map[i-1][j-1] = strconv.Itoa(a_a + 1)
							a_a = 0
						}
						if Map[i-1][j] != "x" {
							a_b, _ := strconv.Atoi(Map[i-1][j])
							Map[i-1][j] = strconv.Itoa(a_b + 1)
							a_b = 0
						}
						if Map[i-1][j+1] != "x" {
							a_c, _ := strconv.Atoi(Map[i-1][j+1])
							Map[i-1][j+1] = strconv.Itoa(a_c + 1)
							a_c = 0
						}
						if Map[i][j-1] != "x" {
							a_d, _ := strconv.Atoi(Map[i][j-1])
							Map[i][j-1] = strconv.Itoa(a_d + 1)
							a_d = 0
						}
						if Map[i][j+1] != "x" {
							a_e, _ := strconv.Atoi(Map[i][j+1])
							Map[i][j+1] = strconv.Itoa(a_e + 1)
							a_e = 0
						}
						if Map[i+1][j-1] != "x" {
							a_e, _ := strconv.Atoi(Map[i+1][j-1])
							Map[i+1][j-1] = strconv.Itoa(a_e + 1)
							a_e = 0
						}
						if Map[i+1][j] != "x" {
							a_e, _ := strconv.Atoi(Map[i+1][j])
							Map[i+1][j] = strconv.Itoa(a_e + 1)
							a_e = 0
						}
						if Map[i+1][j+1] != "x" {
							a_e, _ := strconv.Atoi(Map[i+1][j+1])
							Map[i+1][j+1] = strconv.Itoa(a_e + 1)
							a_e = 0
						}
					}
				}
			}
		}
	}
	return Map
}

func main() {
	Size_x := 15
	Size_y := 16
	Map, Bombes := StartNewgame(Size_x, Size_y, 2)
	Map_End := initMap(Map, Bombes)
	for i := range Map_End {
		fmt.Println(Map_End[i])
	}
}
