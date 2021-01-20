package otu

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/devtools/cli/user"

	"github.com/Galdoba/utils"
)

type InfoRetriver interface {
	Sector() string
	SubSector() string
	Hex() string
	Name() string
	UWP() string
	Bases() []string
	Remarks() []string
	Zone() string
	PBG() string
	Allegiance() string
	Stellar() []string
	Iextention() string
	Eextention() string
	Cextention() string
	Nobility() string
	Worlds() string
	RU() string
}

func MapDataByHex(data []string) map[string]string {
	dataMap := make(map[string]string)
	for _, v := range data {
		data := strings.Split(v, "	")
		dataMap[data[2]] = v
	}
	return dataMap
}

func MapDataByName(data []string) map[string]string {
	dataMap := make(map[string]string)
	for _, v := range data {
		data := strings.Split(v, "	")
		dataMap[data[3]] = v
	}
	return dataMap
}

func MapDataByUWP(data []string) map[string]string {
	dataMap := make(map[string]string)
	for _, v := range data {
		data := strings.Split(v, "	")
		dataMap[data[4]] = v
	}
	return dataMap
}

func subSectorOffset(ss string) (int, int) {
	xOffset := 0
	yOffset := 0
	switch ss {
	case "A":
		xOffset = 0
		yOffset = 0
	case "B":
		xOffset = 8
		yOffset = 0
	case "C":
		xOffset = 16
		yOffset = 0
	case "D":
		xOffset = 24
		yOffset = 0
	case "E":
		xOffset = 0
		yOffset = 10
	case "F":
		xOffset = 8
		yOffset = 10
	case "G":
		xOffset = 16
		yOffset = 10
	case "H":
		xOffset = 24
		yOffset = 10
	case "I":
		xOffset = 0
		yOffset = 20
	case "J":
		xOffset = 8
		yOffset = 20
	case "K":
		xOffset = 16
		yOffset = 20
	case "L":
		xOffset = 24
		yOffset = 20
	case "M":
		xOffset = 0
		yOffset = 30
	case "N":
		xOffset = 8
		yOffset = 30
	case "O":
		xOffset = 16
		yOffset = 30
	case "P":
		xOffset = 24
		yOffset = 30
	}
	return xOffset, yOffset
}

func hex5ToHex4(hex5 string) string {
	if len(hex5) != 5 {
		return "Wrong Format"
	}
	hexParts := strings.Split(hex5, "")
	ss := hexParts[0]
	xOffset, yOffset := subSectorOffset(ss)
	x, _ := strconv.Atoi(hexParts[1] + hexParts[2])
	y, _ := strconv.Atoi(hexParts[3] + hexParts[4])
	x += xOffset
	y += yOffset
	res := ""
	if x < 10 {
		res += "0"
	}
	res += strconv.Itoa(x)
	if y < 10 {
		res += "0"
	}
	res += strconv.Itoa(y)
	return res
}

//GetData -
func GetDataUser() SecondSurveyReportT5SS {
	fmt.Print("Enter Search Key: ")
	key := "No Key"
	err := errors.New("")
	for err != nil {
		fmt.Print(err.Error())
		key, err = user.InputStr()
	}
	responce, err := searchByName(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	return responce
}

func GetDataOn(key string) (SecondSurveyReportT5SS, error) {
	err := errors.New("")
	responce, err := searchByName(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	return responce, err
}

func searchByName(key string) (SecondSurveyReportT5SS, error) {
	err := errors.New("No Data Found")
	keyUP := strings.ToUpper(key)
	search := make(map[int][]SecondSurveyReportT5SS)
	for _, line := range RawData() {
		lineUP := strings.ToUpper(line)
		//keyParts := strings.Split(keyUP, " ")
		if containsAny(lineUP, strings.Split(keyUP, " ")...) {
			ssr := SecondSurveyReportT5SS{line}
			if strings.ToUpper(ssr.Name()) == keyUP {
				search[1] = append(search[1], ssr)
			}
			if strings.ToUpper(ssr.UWP()) == keyUP {
				search[4] = append(search[4], ssr)
			}
			if strings.Contains(strings.ToUpper(ssr.Name()), keyUP) && len(key) > 3 {
				search[2] = append(search[2], ssr)
			}
			if strings.Contains(keyUP, strings.ToUpper(ssr.Sector())) && strings.Contains(keyUP, strings.ToUpper(ssr.Hex())) {
				search[3] = append(search[3], ssr)
			}

			err = nil
		}
	}
	if len(search[1]) == 1 {
		return search[1][0], nil
	}
	if len(search[4]) == 1 {
		return search[4][0], nil
	}
	if len(search[2]) == 1 {
		return search[2][0], nil
	}
	if len(search[3]) == 1 {
		return search[3][0], nil
	}
	if len(search[1]) > 5 {
		err = errors.New("Search Key returns " + strconv.Itoa(len(search[1])) + " results")
	}
	if len(search[1]) > 0 {
		return selectFromFound(search[1])
	}
	if len(search[4]) > 0 {
		return selectFromFound(search[4])
	}
	if len(search[2]) > 0 {
		return selectFromFound(search[2])
	}
	if len(search[3]) > 0 {
		return selectFromFound(search[3])
	}
	err = errors.New("No Data Found")
	return SecondSurveyReportT5SS{}, err
}

func selectFromFound(found []SecondSurveyReportT5SS) (SecondSurveyReportT5SS, error) {
	searchRes := []string{}
	for _, val := range found {
		searchRes = append(searchRes, val.String())
	}
	i, _ := utils.TakeOptions("Select Entry:", searchRes...)
	fmt.Println("\033[A                          \r")
	return found[i-1], nil
}

func containsAny(str string, subStr ...string) bool {
	for _, v := range subStr {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

func handleSearchError(err error) {
	if err == nil {
		return
	}
	switch err.Error() {
	default:
		msg := "Warning! " + err.Error()
		if ok, _ := user.Confirm(msg + "\ncontinue?"); ok {
			return
		} else {
			panic(err)
		}

	}
	return
}

//SecondSurveyReportT5SS - содержит строку с данными с https://travellermap.com
//данные хранятся одной строкой и разделены пробелами.
//Карта данных сейчас:
//6_2_4_28_9_5 44 4 3 10 29 7 7 6 8 2 5  - Line 66191 - max len(line):179
//TODO: сделать конструктор маски для вывода нескольких репортов
type SecondSurveyReportT5SS struct {
	data string
}

//Data - возвращает не изменненную строку из общей базы
func (ssr *SecondSurveyReportT5SS) Data() string {
	return ssr.data
}

//Sector - возвращает абревиатуру сектора
func (ssr *SecondSurveyReportT5SS) Sector() string {
	return string(ssr.data[:4])
}

//SubSector - возвращает абревиатуру субсектора
func (ssr *SecondSurveyReportT5SS) SubSector() string {
	return string(ssr.data[7:8])
}

//Hex - возвращает координаты хекса внутри сектора
func (ssr *SecondSurveyReportT5SS) Hex() string {
	return string(ssr.data[10:14])
}

//Name - возвращает имя главной планеты
func (ssr *SecondSurveyReportT5SS) Name() string {
	name := string(ssr.data[15:43])
	name = trimAllSpaces(name)
	return name
}

//UWP - возвращает UWP главной планеты
func (ssr *SecondSurveyReportT5SS) UWP() string {
	return string(ssr.data[44:53])
}

//Bases - возвращает базы находящиеся в системе.
func (ssr *SecondSurveyReportT5SS) Bases() []string {
	basStr := string(ssr.data[54:59])
	trStr := trimAllSpaces(basStr)
	bases := strings.Split(trStr, "")
	return bases
}

//BasesStr - возвращает базы находящиеся в системе.
func (ssr *SecondSurveyReportT5SS) BasesStr() string {
	basStr := string(ssr.data[54:59])
	return trimAllSpaces(basStr)
}

//Remarks - возвращает ремарки в виде слайса.
func (ssr *SecondSurveyReportT5SS) Remarks() []string {
	basStr := string(ssr.data[60:104])
	trStr := trimAllSpaces(basStr)
	remarks := strings.Split(trStr, " ")
	return remarks
}

//Zone - возвращает Travel Zone главной планеты
func (ssr *SecondSurveyReportT5SS) Zone() string {
	return string(ssr.data[105:106])
}

//PBG - возвращает PBG главной планеты
func (ssr *SecondSurveyReportT5SS) PBG() string {
	return string(ssr.data[110:113])
}

//Allegiance - возвращает Принадлежность главной планеты
func (ssr *SecondSurveyReportT5SS) Allegiance() string {
	return trimAllSpaces(string(ssr.data[114:125]))
}

//Stellar - возвращает перечень звезд системы
func (ssr *SecondSurveyReportT5SS) Stellar() string {
	return trimAllSpaces(string(ssr.data[125:154]))
}

//Iextention - возвращает Importance Extention
func (ssr *SecondSurveyReportT5SS) Iextention() string {
	return trimAllSpaces(string(ssr.data[155:162]))
}

//IextentionVal - возвращает Параметр Importance в числовом значении
func (ssr *SecondSurveyReportT5SS) IextentionVal() int {
	iExStr := ssr.Iextention()
	iExStr = strings.TrimPrefix(iExStr, "{ ")
	iExStr = strings.TrimSuffix(iExStr, " }")
	iEx, err := strconv.Atoi(iExStr)
	resolveErr(err)
	return iEx
}

//Eextention - возвращает Econoimic Extention
func (ssr *SecondSurveyReportT5SS) Eextention() string {
	return trimAllSpaces(string(ssr.data[163:170]))
}

//Cextention - возвращает Cultural Extention
func (ssr *SecondSurveyReportT5SS) Cextention() string {
	return trimAllSpaces(string(ssr.data[171:177]))
}

//Nobility - возвращает перечень титулов возможных на главном мире.
func (ssr *SecondSurveyReportT5SS) Nobility() string {
	return trimAllSpaces(string(ssr.data[178:186]))
}

//Worlds - возвращает количество миров помимо главного
func (ssr *SecondSurveyReportT5SS) Worlds() string {
	return trimAllSpaces(string(ssr.data[187:189]))
}

//RU - возвращает экономический показатель мира (годовой ВВП)
//1 RU примерно равен 10MCr
//TODO: обдумать возможность/необходимость расчетов бюджетов по модулю Pocket Empire
func (ssr *SecondSurveyReportT5SS) RU() string {
	return trimAllSpaces(string(ssr.data[190:195]))
}

//RUint - возвращает экономический показатель мира (годовой ВВП) в виде Int
//1 RU примерно равен 10MCr
func (ssr *SecondSurveyReportT5SS) RUint() int {
	ru, err := strconv.Atoi(ssr.RU())
	resolveErr(err)
	return ru
}

func (ssr *SecondSurveyReportT5SS) String() string {
	if len(ssr.data) == 0 {
		return "--NO DATA--"
	}
	str := ""
	str += ssr.Sector() + "  "
	str += ssr.SubSector() + "  "
	str += ssr.Hex() + "  "
	str += ssr.Name() + "  "
	str += ssr.UWP() + "  "
	for _, v := range ssr.Remarks() {
		str += v + " "
	}
	str += " "
	str += ssr.Iextention() + "  "
	str += ssr.Eextention() + "  "
	str += ssr.Cextention() + "  "
	str += ssr.Nobility() + "  "
	str += ssr.BasesStr() + "  "
	str += ssr.Zone() + "  "
	str += ssr.PBG() + "  "
	str += ssr.Worlds() + "  "
	str += ssr.Allegiance() + "  "
	str += ssr.Stellar() + "  "
	str += ssr.RU()
	return str
}

func trimAllSpaces(s string) string {
	l := len(s) + 1
	for len(s) != l {
		l = len(s)
		s = strings.TrimSuffix(s, " ")
	}
	return s
}

func resolveErr(err error) {
	if err != nil {
		switch err.Error() {
		default:
			panic(err)
		}
	}
}

/*
//2 5  - Line 66191 - max len(line):179
type InfoRetriver interface {
	Worlds() string
	RU() string
}
*/
