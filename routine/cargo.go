package routine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/utils"

	trade "github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/otu"
)

const (
	cargoFile = "mgt2_traffic.config"
)

type cargoManifest struct {
	entry []cargoLot
}

type cargoLot struct {
	tgCode           string
	descr            string
	volume           int
	originWorld      string //HEX or 0000
	cost             int
	dangerDM         int
	riskDM           int
	legal            bool
	destinationWorld string
	eta              string
	insurance        int
	suplierType      string
	comment          string
}

func newCargoLot() cargoLot {
	cl := cargoLot{}

	return cl
}

func (cl *cargoLot) LeachData(rawData string) error {
	var err error
	data := strings.Split(rawData, "_")
	cl.tgCode = data[0]
	cl.descr = data[1]
	cl.volume, err = strconv.Atoi(data[2])
	origin, errO := otu.GetDataOn(data[3])
	cl.originWorld = "[NO DATA]"
	if errO == nil {
		cl.originWorld = origin.Hex()
	}
	cl.cost, err = strconv.Atoi(data[4])
	cl.dangerDM, err = strconv.Atoi(data[5])
	cl.riskDM, err = strconv.Atoi(data[6])
	cl.legal, err = strconv.ParseBool(data[7])
	destination, errD := otu.GetDataOn(data[8])
	cl.destinationWorld = "[NO DATA]"
	if errD == nil {
		cl.destinationWorld = destination.Hex()
	}
	cl.eta = "[NO DATA]"
	if isDate(data[9]) {
		cl.eta = data[9]
	}
	cl.insurance, err = strconv.Atoi(data[10])
	cl.suplierType = data[11]
	cl.comment = data[12]
	return err
}

func (cl *cargoLot) SetTGCode(newVal string) {
	cl.tgCode = newVal
}

func (cl *cargoLot) SetDescr(newVal string) {
	cl.descr = newVal
}

func (cl *cargoLot) SetVolume(newVal int) {
	cl.volume = newVal
}

func (cl *cargoLot) SetOrigin(newVal string) {
	cl.originWorld = newVal
}

func (cl *cargoLot) SetCost(newVal int) {
	cl.cost = newVal
}

func (cl *cargoLot) SetDangerDM(newVal int) {
	cl.dangerDM = newVal
}

func (cl *cargoLot) SetRiskDM(newVal int) {
	cl.riskDM = newVal
}

func (cl *cargoLot) SetLegality(newVal bool) {
	cl.legal = newVal
}

func (cl *cargoLot) SetDestination(newVal string) {
	cl.destinationWorld = newVal
}

func (cl *cargoLot) SetETA(newVal string) {
	cl.eta = newVal
}

func (cl *cargoLot) SetInsurance(newVal int) {
	cl.insurance = newVal
}

func (cl *cargoLot) SetSupplierType(newVal string) {
	cl.suplierType = newVal
}

func (cl *cargoLot) SetComment(newVal string) {
	cl.comment = newVal
}

func (cl *cargoLot) GetTGCode() string {
	return cl.tgCode
}

func (cl *cargoLot) GetDescr() string {
	return cl.descr
}

func (cl *cargoLot) GetVolume() int {
	return cl.volume
}

func (cl *cargoLot) GetOrigin() string {
	return cl.originWorld
}

func (cl *cargoLot) GetCost() int {
	return cl.cost
}

func (cl *cargoLot) GetDangerDM() int {
	return cl.dangerDM
}

func (cl *cargoLot) GetRiskDM() int {
	return cl.riskDM
}

func (cl *cargoLot) GetLegality() bool {
	return cl.legal
}

func (cl *cargoLot) GetDestination() string {
	return cl.destinationWorld
}

func (cl *cargoLot) GetETA() string {
	return cl.eta
}

func (cl *cargoLot) GetInsurance() int {
	return cl.insurance
}

func (cl *cargoLot) GetSupplierType() string {
	return cl.suplierType
}

func (cl *cargoLot) GetComment() string {
	return cl.comment
}

func newCargoManifest() cargoManifest {
	cm := cargoManifest{}
	rawData := getCargo()
	for i := range rawData {
		lot := newCargoLot()
		if lot.LeachData(rawData[i]) != nil {
			continue
		}
		cm.entry = append(cm.entry, lot)
	}
	return cm
}

func TestCargo() {
	cm := newCargoManifest()
	for i := range cm.entry {
		fmt.Println(i, cm.entry[i])
	}
	cl := newCargoLot()
	cl.SetTGCode("164")
	cl.SetDescr(trade.GetDescription("164"))
	cl.SetVolume(325)
	cl.SetComment("Test Entry")
	fmt.Println(cl)
	cm.entry = append(cm.entry, cl)
	saveCargo(cm)
}

func (cm *cargoManifest) addCargo(cl cargoLot) {
	cm.entry = append(cm.entry, cl)
}

func saveCargo(cm cargoManifest) {
	lines := utils.LinesFromTXT(exPath + cargoFile)
entry:
	for i := range cm.entry {
		fmt.Println("Go Entry", i)
		for l, val := range lines {
			if strings.Contains(val, cm.entry[i].GetComment()) {
				fmt.Println("Edit line", l, "entry", cm.entry[i])
				utils.EditLineInFile(exPath+cargoFile, l, cm.entry[i].SeedData())

				continue entry
			}
			fmt.Println(l)
		}
		utils.AddLineToFile(exPath+cargoFile, cm.entry[i].SeedData())

	}

}

func (cl *cargoLot) SeedData() string {
	str := "CARGOENTRY:"
	//CARGOENTRY:148_Workable Alloys_16_Drinax_4600_-6_0_TRUE_Asim_118-1106_50_1_Freight
	str += cl.GetTGCode() + "_" + cl.GetDescr() + "_" + strconv.Itoa(cl.GetVolume()) + "_" + cl.GetOrigin() + "_" + strconv.Itoa(cl.GetCost()) +
		"_" + strconv.Itoa(cl.GetDangerDM()) + "_" + strconv.Itoa(cl.GetRiskDM()) + "_" + strconv.FormatBool(cl.GetLegality()) + "_" + cl.GetDestination() + "_" + cl.GetETA() + "_" +
		strconv.Itoa(cl.GetInsurance()) + "_" + cl.GetSupplierType() + "_" + cl.GetComment()
	return str
}

func integerToEhexCode(i int) string {
	// fmt.Println(30*30*30*30*30*30*30*30*30*30*30*30, 12) - максимум
	str := "#"
	neg := false
	if i < 0 {
		neg = true
		i = i * -1
	}
	for i/30 >= 0 {
		str += TrvCore.DigitToEhex(i % 30)
		if i/30 == 0 {
			break
		}
		i = i / 30
	}
	if neg {
		str += "^"
	}
	//fmt.Print("\r" + str + "   " + strconv.Itoa(d))
	return str
}

func ehexCodeToInteger(s string) int {
	bts := []byte(s)
	neg := false
	d := 0
	mn := 0
	for i, b := range bts {
		if i == 0 && string(b) != "#" {
			return 0
		}
		if i == 0 {
			continue
		}
		if string(b) == "^" {
			neg = true
			continue
		}
		mn = 30 ^ i // - 1
		fmt.Println()
		d += TrvCore.EhexToDigit(string(b)) * mn
	}
	if neg {
		d = d * -1
	}
	return d
}

func isDate(str string) bool {
	data := strings.Split(str, "-")
	for i := range data {
		if i > 1 {
			return false
		}
		_, err := strconv.Atoi(data[i])
		if err != nil {
			return false
		}
	}
	return true
}
