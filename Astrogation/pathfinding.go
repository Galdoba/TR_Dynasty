package Astrogation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/otu"

	"github.com/Galdoba/utils"
)

/*
Определения:
Start - точка в которой путешественник находится в момент 0
End - точка в которой путешественник должен оказаться не познее чем через remainingT итераций
MaxJumps - количество доступных итераций
remainingT - (MaxJumps - T) количество оставшихся итераций
Т - количество завершенных итераций

Логика алгоритма:
поиск подходящих путей расчитывается на N итераций.
в каждый конкретный момент времени путешественник должен находиться не дальше чем
(ramainingT * Jumpdrive) полей от точки End (множество E)
в каждый конкретный момент времени путешественник должен находиться не дальше чем
(T * Jumpdrive) полей от точки Start (множество S)
пересечение множеств S и E - являются возможными точками нахождения путешественника в момент T
*/

func Test() {
	fmt.Println("Run Test")
	path, err := PlotCourse("Hilfer", "Arunisiir", 2)
	fmt.Println("Path", path, err)
	if err != nil {
		return
	}
	points := strings.Split(path, " ")
	way := ""
	for _, val := range points {
		name, err := otu.GetDataOn(val)
		if err != nil {
			return
		}
		way += name.Name() + " -> "
	}
	way = strings.TrimSuffix(way, " -> ")
	fmt.Println(way)
}

func PlotCourse(start, end string, drive int) (string, error) {
	data1, err := otu.GetDataOn(start)
	plot := jumpPlot{}
	if err != nil {
		return "", err
	}
	data2, err := otu.GetDataOn(end)
	if err != nil {
		return "", err
	}
	startHex := data1.Hex()
	endHex := data2.Hex()
	for maxJumps := 0; maxJumps < 99; maxJumps++ {
		plot, err = newPlot(startHex, endHex, drive, maxJumps)
		if plot.err != nil {
			return "", plot.err
		}
		break
	}
	return chooseBest(plot.jumpMap, plot.drives), plot.err
}

func chooseBest(crsmap map[int]string, drives int) string {
	minSum := 1000000
	candidate := 0
mapCycle:
	for k, val := range crsmap {
		//fmt.Println("Test:", k, "of", len(crsmap))
		points := strings.Split(val, " ")
		//fmt.Println("POINTS", points)
		sum := 0
		minJump := 0
		for r := range points {
			last := r - 1
			if r == 0 {
				last = 0
			}
			if JumpDistance(points[last], points[r]) > drives {
				//	fmt.Println("Exclude:", crsmap[k])
				continue mapCycle
			}
			sum += JumpDistance(points[last], points[r])
			if minJump < JumpDistance(points[last], points[r]) {
				minJump = JumpDistance(points[last], points[r])
			}
		}
		//fmt.Println("Check:", crsmap[k], minSum, sum, minJump)
		if sum < minSum {
			minSum = sum
			candidate = k
			//fmt.Println("SET CANDIDATE:", k)
		}
	}
	//fmt.Println("Chosen:", crsmap[candidate], drives)
	return crsmap[candidate]
}

type jumpPlot struct {
	start    string
	end      string
	drives   int
	maxJumps int
	pMap     map[int][]string
	jumpMap  map[int]string
	err      error
}

func newPlot(startHex, endHex string, drive, maxJumps int) (jumpPlot, error) {
	jp := jumpPlot{}
	jp.start = startHex
	jp.end = endHex
	jp.drives = drive
	jp.maxJumps = maxJumps
	jp.pMap = make(map[int][]string)
	jp.jumpMap = make(map[int]string)
	jp.calcJumpWaves()
	if jp.err != nil {
		return jp, jp.err
	}
	jp.connectDots()
	//jp.removeImpossibleRoads()

	return jp, jp.err
}

func commonInSlices(sl1, sl2 []string) []string {
	sl3 := []string{}
	for i := range sl1 {
		for j := range sl2 {
			if sl1[i] == sl2[j] {
				sl3 = utils.AppendUniqueStr(sl3, sl1[i])
			}
		}
	}
	return sl3
}

// func commonInSlices2(sl1, sl2 []string) []string {
// 	for j := range sl2 {
// 		sl1 = utils.AppendUniqueStr(sl1, sl2[j])
// 	}
// 	return sl1
// }

func appendSlice(sl1, sl2 []string) []string {
	for j := range sl2 {
		sl1 = utils.AppendUniqueStr(sl1, sl2[j])
	}
	return sl1
}

func removeEmpty(sl []string) []string {
	newSl := []string{}
	for i := range sl {
		data, err := otu.GetDataOn(sl[i])
		if err != nil {
			continue
		}
		newSl = append(newSl, data.Hex())
	}
	return newSl
}

func (jp jumpPlot) connectDots() {
	routeMapMax := []int{}
	currentRoad := []int{}
	for k := 0; k < len(jp.pMap); k++ {
		toadd := len(jp.pMap[k]) - 1
		if toadd < 0 {
			toadd = 0
		}
		routeMapMax = append(routeMapMax, toadd)
		currentRoad = append(currentRoad, 0)
	}
	v := 1
	for i := range routeMapMax {
		v = v * (routeMapMax[i] + 1)
	}
	route := 0
	jp.jumpMap[route] = projectCourse(jp.pMap, currentRoad, jp.end)
	//fmt.Println("j", jp.jumpMap[route])
	for !roadIsEqual(routeMapMax, currentRoad) {
		//	fmt.Println("c", currentRoad)
		route++
		currentRoad = pushCloser(currentRoad, routeMapMax)
		jp.jumpMap[route] = projectCourse(jp.pMap, currentRoad, jp.end)
		//	fmt.Println("j", jp.jumpMap[route])
	}

	//return roadMap
}

func roadIsEqual(sli1, sli2 []int) bool {
	if len(sli1) != len(sli2) {
		return false
	}
	for i := range sli1 {
		if sli1[i] != sli2[i] {
			return false
		}
	}
	return true
}

func pushCloser(sli1, sli2 []int) []int {
	for i := range sli1 {
		if sli1[i] < sli2[i] {
			sli1[i]++
			return sli1
		}
		sli1[i] = 0
	}
	return sli1
}

func projectCourse(jmap map[int][]string, cource []int, end string) string {
	path := ""
	for i, val := range cource {
		if len(jmap[i]) == 0 {
			return path
		}
		path = path + jmap[i][val] + " "
		if jmap[i][val] == end {
			break
		}
	}
	path = strings.TrimSuffix(path, " ")
	return path
}

func findShortest(jmap map[int]string, end string) []string {
	shortJumpsf := []string{}
	leastJumps := 1000
	for _, v := range jmap {
		jumpPoints := strings.Split(v, " ")
		for i := range jumpPoints {
			if jumpPoints[i] == end && leastJumps > i {
				leastJumps = i
				shortJumpsf = append(shortJumpsf, v)
			}
		}
	}
	shortJumpsS := []string{}
	for i := range shortJumpsf {
		jumpPoints := strings.Split(shortJumpsf[i], " ")
		if jumpPoints[leastJumps] != end {
			continue
		}
		shortJumpsS = append(shortJumpsS, shortJumpsf[i])
	}
	return shortJumpsS
}

func (jp *jumpPlot) removeImpossibleRoads() {
	p1 := ""
	p2 := ""
	invalid := 0
	shortestLen := 10000000
	for key, road := range jp.jumpMap {
		points := strings.Split(road, " ")
		for i := range points {
			if points[i] == jp.end && shortestLen > i {
				shortestLen = i
			}
			if i > shortestLen {
				delete(jp.jumpMap, key) // слишком длинный маршрут и его можно даже не расчитывать
			}
			if i == 0 {
				p1 = jp.start
			} else {
				p1 = points[i-1]
			}
			p2 = points[i]
			if p2 == "" {
				p1 = jp.start
				p2 = jp.end
			}
			if JumpDistance(p1, p2) > jp.drives {
				fmt.Println("WTF?", p1, p2)
				invalid++
				delete(jp.jumpMap, key)
				break
			}
		}
	}
	newMap := make(map[int]string)
	validKey := 0
	for _, road := range jp.jumpMap {
		points := strings.Split(road, " ")
		if points[shortestLen] == jp.end {
			newMap[validKey] = road
		}
	}
	jp.jumpMap = newMap
}

func (jp *jumpPlot) calcJumpWaves() {
	pointsPool := []string{jp.start}
	workingPool := pointsPool
	workingPoolR := []string{jp.end}
	currentWaveLen := 0
	currentWaveLenR := 0
	forwardSearch := make(map[int][]string)
	reverseSearch := make(map[int][]string)
	fmt.Print("Constructing Jump Plot.")
fSearch:
	for i := 0; i < 99; i++ {
		fmt.Print(".")
		forwardSearch[i] = workingPool
		workingPool = addCoordsInRange(workingPool, jp.drives)
		if currentWaveLen == len(workingPool) {
			fmt.Print("\r")
			jp.err = errors.New("Jump Plot Impossible with drives " + strconv.Itoa(jp.drives))
			return
		}
		if currentWaveLen < len(workingPool) {
			currentWaveLen = len(workingPool)
		}
		for _, val := range forwardSearch[i] {
			if val == jp.end {
				break fSearch
			}
		}

	}
	fmt.Println("ok")
	fmt.Print("Calculating Path.")
rSearch:
	for i := 0; i < 99; i++ {
		fmt.Print(".")
		reverseSearch[i] = workingPoolR
		workingPoolR = addCoordsInRange(workingPoolR, jp.drives)
		if currentWaveLenR == len(workingPoolR) {
			break
		}
		if currentWaveLenR < len(workingPoolR) {
			currentWaveLenR = len(workingPoolR)
		}
		for _, val := range reverseSearch[i] {
			if val == jp.start {
				break rSearch
			}
		}
	}
	fmt.Println("ok")
	max := utils.Max(len(forwardSearch), len(reverseSearch))
	for i := 0; i < max; i++ {
		jp.pMap[i] = commonInSlices(forwardSearch[i], reverseSearch[len(reverseSearch)-i-1])
	}
}

func addCoordsInRange(cPool []string, jRange int) []string {
	var newCoords []string
	for p := range cPool {
		pick := cPool[p]
		for _, suggest := range JumpCoordinatesFrom(pick, jRange) {
			if havePlanet(suggest) {
				newCoords = utils.AppendUniqueStr(newCoords, suggest)
			}
		}
	}
	cPool = appendSlice(cPool, newCoords)
	return cPool
}

func havePlanet(hex string) bool {
	_, err := otu.GetDataOn(hex)
	if err != nil {
		return false
	}
	return true
}
