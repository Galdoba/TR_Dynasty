package Astrogation

import (
	"errors"
	"fmt"
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
	fmt.Println(PlotCourse("2714", "3117", 3))
}

func PlotCourse(startHex, endHex string, drive int) (string, error) {
	courseMap := make(map[int]string)
	for maxJumps := 0; maxJumps < 7; maxJumps++ {
		plot, err := newPlot(startHex, endHex, drive, maxJumps)
		if err != nil {
			//fmt.Println(err.Error() + strconv.Itoa(maxJumps) + " jumps")
			continue
		}
		courseMap = plot.jumpMap
		if maxJumps == 6 {
			return "", errors.New("DEBUG: It will take to long to calculate :(")
		}
		break
		//return plot, nil
	}

	return chooseBest(courseMap), nil
}

func chooseBest(crsmap map[int]string) string {
	minSum := 1000000
	candidate := 0
	for k, val := range crsmap {
		points := strings.Split(val, " ")
		sum := 0
		for r := range points {
			last := r
			if r == 0 {
				last = 0
			}
			sum += JumpDistance(points[last], points[r])
		}
		if sum < minSum {
			minSum = sum
			candidate = k
		}
	}
	return crsmap[candidate]
}

type jumpPlot struct {
	start    string
	end      string
	drives   int
	maxJumps int
	sMap     map[int][]string
	eMap     map[int][]string
	pMap     map[int][]string
	jumpMap  map[int]string
}

func newPlot(startHex, endHex string, drive, maxJumps int) (jumpPlot, error) {
	jp := jumpPlot{}
	jp.start = startHex
	jp.end = endHex
	jp.drives = drive
	jp.maxJumps = maxJumps
	if !jp.possible() {
		return jumpPlot{}, errors.New("Jump Plot not possible with ")
	}
	//jp.testNextJump()
	//panic(0)
	jp.sMap = make(map[int][]string)
	jp.eMap = make(map[int][]string)
	jp.pMap = make(map[int][]string)
	//jp.possiblePointsMap[0] = JumpCoordinatesFrom(jp.start, drive)
	//jp.testNextJump()
search:
	for i := 0; i <= maxJumps; i++ {
		jp.sMap[i] = JumpCoordinatesFrom(jp.start, (i)*drive) //sMap - все что находится не дальше чем (i)*drive
		//	jp.eMap[i] = JumpCoordinatesFrom(jp.end, (maxJumps-i)*drive)
		endCoords := JumpCoordinatesFrom(jp.end, (maxJumps-i)*drive)
		endCoords = removeEmpty(endCoords)
		for e := range endCoords {
			if JumpDistance(endCoords[e], jp.end) <= (maxJumps-i)*drive {
				jp.eMap[i] = append(jp.eMap[i], endCoords[e])
				//	fmt.Println("ADD", JumpDistance(endCoords[e], jp.end), (maxJumps-i)*drive, endCoords[e])
			} else {
				//	fmt.Println("DONT ADD")
			}
		}
		jp.eMap[i] = removeEmpty(jp.eMap[i])
		jp.pMap[i] = commonInSlices(jp.sMap[i], jp.eMap[i])
		jp.pMap[i] = removeEmpty(jp.pMap[i])
		//		fmt.Println("--------------------")
		//		fmt.Println(jp.pMap)
		for _, val := range jp.pMap {
			for p := range val {
				if val[p] == jp.end {
					break search // типа надо отсечь все что дальше необходимого
					//Прекращаем поиск если найден порядок точек с N прыжков меньше максимального
				}
			}
		}
		//88panic(0)
	}
	// jp.testNextJump()
	// if len(jp.pMap) == 0 {
	// 	return jumpPlot{}, errors.New("Jump Plot not possible with ")
	// }

	//fmt.Println("Connect Dots", len(jp.pMap))
	jp.jumpMap = jp.connectDots(jp.pMap)
	// for k, val := range jp.jumpMap {
	// 	fmt.Println(k, val)
	// }
	jp.removeImpossibleRoads()
	if len(jp.jumpMap) == 0 {
		err := errors.New("1 Jump Plot not possible with ")
		return jp, err
	}
	//findShortest(jp.jumpMap, jp.end)
	return jp, nil
}

func (jp *jumpPlot) testNextJump() {
	//fmt.Println("Start testNextJump()")
	jump := make(map[int][]string)
	allCoords := []string{jp.start}
	lastLen := len(allCoords)
	for i := 0; i < 1000; i++ {
		if i >= jp.maxJumps {
			return
		}
		//fmt.Println("Go i =", i, lastLen)
		for r := range allCoords {
			//	fmt.Println("Go r =", r, lastLen)
			coords := JumpCoordinatesFrom(allCoords[r], jp.drives)
			coords = removeEmpty(coords)
			//fmt.Println(coords)
			allCoords = commonInSlices2(allCoords, coords)
			jump[i] = allCoords
		}
		for k := range allCoords {
			if allCoords[k] == jp.end {
				//			fmt.Println("Stop!!")
				//			fmt.Println("End testNextJump() 3")

				jp.pMap = jump
				//fmt.Println(jump)
				//fmt.Println(jp.pMap)
				return
			}
		}
		if lastLen < len(allCoords) {
			lastLen = len(allCoords)
		} else {
			//		fmt.Println("Stop Here")
			//	fmt.Println(jump)
			jp.pMap = jump
			//		fmt.Println("End testNextJump() 1")
			return
		}
		jump[i] = allCoords
	}
	//fmt.Println("End testNextJump() 2")
	// for k, v := range jump {
	// 	fmt.Println("1111JUMP", k, "=", v)
	// 	jump[k] = removeEmpty(jump[k])
	// 	fmt.Println("2222JUMP", k, "=", jump[k])
	// }

	//fmt.Println("JUMP:", jump)
	//fmt.Println("allCoords:", allCoords)
}

func (jp jumpPlot) totalDistance() int {
	return JumpDistance(jp.start, jp.end)
}

func (jp jumpPlot) possible() bool {
	if jp.totalDistance() <= jp.drives*jp.maxJumps {
		return true
	}
	return false
}

func commonInSlices(sl1, sl2 []string) []string {
	sl3 := []string{}
	//fmt.Println(sl1)
	//fmt.Println(sl2)
	for i := range sl1 {
		for j := range sl2 {
			if sl1[i] == sl2[j] {
				sl3 = utils.AppendUniqueStr(sl3, sl1[i])
				//			fmt.Println("+++")
			}
		}
	}
	return sl3
}

func commonInSlices2(sl1, sl2 []string) []string {
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

func (jp jumpPlot) connectDots(jmap map[int][]string) map[int]string {
	routeMapMax := []int{}
	currentRoad := []int{}
	for k := 0; k < len(jmap); k++ {
		toadd := len(jmap[k]) - 1
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
	//fmt.Println(v)
	route := 0
	roadMap := make(map[int]string)
	roadMap[route] = projectCourse(jmap, currentRoad, jp.end)
	for !roadIsEqual(routeMapMax, currentRoad) {

		//	fmt.Println(routeMapMax, currentRoad)
		route++
		currentRoad = pushCloser(currentRoad, routeMapMax)
		roadMap[route] = projectCourse(jmap, currentRoad, jp.end)

		//fmt.Print("Processing Route: ", route, " of ", v, "\r")
	}

	return roadMap
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
	//fmt.Println(shortJumpsf, 1)
	shortJumpsS := []string{}
	for i := range shortJumpsf {
		jumpPoints := strings.Split(shortJumpsf[i], " ")
		if jumpPoints[leastJumps] != end {
			continue
		}
		shortJumpsS = append(shortJumpsS, shortJumpsf[i])
	}

	//fmt.Println(shortJumpsS, "TEst")
	//removeImpossibleRoads(shortJumpsS)
	return shortJumpsS
}

func (jp *jumpPlot) removeImpossibleRoads() {
	//	newMap := make(map[int]string)
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
			//fmt.Println("p2 =", p2)
			if p2 == "" {
				p1 = jp.start
				p2 = jp.end
			}
			if JumpDistance(p1, p2) > jp.drives {
				invalid++
				delete(jp.jumpMap, key)
				break
			}
		}
	}
	//fmt.Println(invalid, len(jp.jumpMap), shortestLen)
	newMap := make(map[int]string)
	validKey := 0
	for _, road := range jp.jumpMap {
		points := strings.Split(road, " ")
		//verdict := "Long"
		if points[shortestLen] == jp.end {
			//	verdict = "short"
			//	fmt.Println(key, points, verdict)
			newMap[validKey] = road
		}
	}
	jp.jumpMap = newMap
}
