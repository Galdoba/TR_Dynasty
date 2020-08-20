package Astrogation

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/Galdoba/utils"
)

const (
	directionN  = 0
	directionNE = 1
	directionSE = 2
	directionS  = 3
	directionSW = 4
	directionNW = 5
)

type stellarHex struct {
	hexStr string
	hex    hexCoords
	cube   cubeCoords
}

func Hex(hexStr string) stellarHex {
	h := stellarHex{}
	h.hexStr = hexStr
	bts := []byte(hexStr)
	col, err := strconv.Atoi(string(bts[0]) + string(bts[1]))
	if err != nil {
		fmt.Println(err)
	}
	row, err := strconv.Atoi(string(bts[2]) + string(bts[3]))
	if err != nil {
		fmt.Println(err)
	}
	h.hex = setHexCoords(col, row)
	h.cube = evenQToCube(h.hex)
	return h
}

type cubeCoords struct {
	x int
	y int
	z int
}

func cubeCoordsStr(cube cubeCoords) string {
	fmt.Println(cube)
	xStr := coordNumToStr("X", cube.x)
	yStr := coordNumToStr("Y", cube.y)
	zStr := coordNumToStr("Z", cube.z)
	output := xStr + " " + yStr + " " + zStr
	return output
}

func coordNumToStr(coordName string, x int) string {
	xStr := coordName
	if x < 0 {
		xStr += "-"
		x = x * -1
	} else {
		xStr += " "
	}
	fmt.Println("1:", xStr)
	if x < 10 && x > -10 {
		xStr += "0"
		xStr += strconv.Itoa(x)
	} else {
		xStr += strconv.Itoa(x)
	}
	return xStr
}

func setCubeCoords(x, y, z int) cubeCoords {
	cube := cubeCoords{}
	cube.x = x
	cube.y = y
	cube.z = z
	return cube
}

func oddQToCube(hex hexCoords) cubeCoords {
	x := hex.col
	z := hex.row - (hex.col-(hex.col&1))/2
	y := -x - z
	return setCubeCoords(x, y, z)
}

func evenQToCube(hex hexCoords) cubeCoords {
	var x = hex.col
	var z = hex.row - (hex.col+(hex.col&1))/2
	var y = -x - z
	return setCubeCoords(x, y, z)
}

func cubeToEvenq(cube cubeCoords) hexCoords {
	var col = cube.x
	var row = cube.z + (cube.x+(cube.x&1))/2
	return setHexCoords(col, row)
}

type hexCoords struct {
	col int
	row int
}

func setHexCoords(c, r int) hexCoords {
	offCrds := hexCoords{}
	offCrds.col = c
	offCrds.row = r
	return offCrds
}

func cubeToHex(cube cubeCoords) hexCoords {
	col := cube.x
	row := cube.z + (cube.x-(cube.x&1))/2
	return setHexCoords(col, row)
}

var hexDirections [][]hexCoords

func init() {
	// hexDirections = [][]hexCoords{
	// 	{hexCoords{1, 0}, hexCoords{1, -1}, hexCoords{0, -1}, hexCoords{-1, -1}, hexCoords{-1, 0}, hexCoords{0, 1}},
	// 	{hexCoords{1, 1}, hexCoords{1, 0}, hexCoords{0, -1}, hexCoords{-1, 0}, hexCoords{-1, 1}, hexCoords{0, 1}},
	// }
	hexDirections = [][]hexCoords{
		{hexCoords{0, -1}, hexCoords{1, -1}, hexCoords{1, 0}, hexCoords{0, 1}, hexCoords{-1, 0}, hexCoords{-1, -1}},
		{hexCoords{0, -1}, hexCoords{1, 0}, hexCoords{1, 1}, hexCoords{0, 1}, hexCoords{-1, 1}, hexCoords{-1, 0}},
	}

}

func hexNeighbor(hex hexCoords, direction int) hexCoords {
	parity := hex.col & 1
	dir := hexDirections[parity][direction]
	return setHexCoords(hex.col+dir.col, hex.row+dir.row)
}

func cubeNeighbor(cube cubeCoords, direction int) cubeCoords {
	hex := cubeToHex(cube)
	parity := hex.col & 1
	dir := hexDirections[parity][direction]
	hexN := setHexCoords(hex.col+dir.col, hex.row+dir.row)
	return oddQToCube(hexN)
}

func cubeDistance(cubeA, cubeB cubeCoords) int {
	return int((math.Abs(float64(cubeA.x-cubeB.x)) + math.Abs(float64(cubeA.y-cubeB.y)) + math.Abs(float64(cubeA.z-cubeB.z))) / 2)
}

func hexDistance(hexA, hexB hexCoords) int {
	cubeA := evenQToCube(hexA)
	cubeB := evenQToCube(hexB)
	return cubeDistance(cubeA, cubeB)
}

func JumpDistance(h1, h2 stellarHex) int {
	return cubeDistance(h1.cube, h2.cube)
}

func evenQToStr(hx hexCoords) string {
	hexStr := ""
	if hx.col < 10 {
		hexStr += "0"
	}
	hexStr += strconv.Itoa(hx.col)
	if hx.row < 10 {
		hexStr += "0"
	}
	hexStr += strconv.Itoa(hx.row)
	return hexStr
}

func addCube(ac, bc cubeCoords) cubeCoords {
	return setCubeCoords(ac.x+bc.x, ac.y+bc.y, ac.z+bc.z)
}

func JumpCoordinatesFrom(initHex string, j int) []string {
	var coords []string
	start := Hex(initHex)
	for x := -j; x <= j; x++ {
		for y := utils.Max(-j, -x-j); y <= utils.Min(j, -x+j); y++ {
			z := -x - y
			cb := addCube(setCubeCoords(x, y, z), start.cube)
			hx := cubeToEvenq(cb)
			coords = append(coords, evenQToStr(hx))
		}
	}
	//fmt.Println(coords)
	sort.Strings(coords)
	return coords
}
