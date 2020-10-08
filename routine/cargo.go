package routine

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"

	trade "github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

const (
	cargoFile = "mgt2_traffic.config"
)

type cargoManifest struct {
	entry        []cargoLot
	maximumTonns int
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
	id               int
}

var portCargo []cargoLot

func newCargoLot() cargoLot {
	cl := cargoLot{}
	cl.id = int(time.Now().UnixNano())
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
	cl.eta = data[9]
	// if isDate(data[9]) {
	// 	cl.eta = data[9]
	// }
	cl.insurance, err = strconv.Atoi(data[10])
	cl.suplierType = data[11]
	cl.comment = data[12]
	cl.id, err = strconv.Atoi(data[13])
	if cl.id == 0 {
		cl.id = int(time.Now().UnixNano())
	}
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

func (cl *cargoLot) SetID(newval int) {
	cl.id = newval
}

func (cl *cargoLot) GetID() int {
	return cl.id
}

func loadCargoManifest() cargoManifest {
	cm := cargoManifest{}
	cm.maximumTonns = getShipData("SHIP_CARGO_VOLUME")
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
	cm := loadCargoManifest()
	for i := range cm.entry {
		fmt.Println(i, cm.entry[i])
	}
	cl := newCargoLot()
	cl.SetTGCode("164")
	cl.SetDescr(trade.GetDescription("164"))
	cl.SetVolume(326)
	cl.SetComment("Test Entry")
	fmt.Println(cl)
	cm.entry = append(cm.entry, cl)
	fmt.Println(cm)
	saveCargoManifest(cm)
}

func (cm *cargoManifest) addCargo(cl cargoLot) {
	cm.entry = append(cm.entry, cl)
}

func saveCargoManifest(cm cargoManifest) {
	lines := utils.LinesFromTXT(exPath + cargoFile)
entry:
	for i := range cm.entry {
		for l, val := range lines {
			if strings.Contains(val, strconv.Itoa(cm.entry[i].GetID())) {
				utils.EditLineInFile(exPath+cargoFile, l, cm.entry[i].SeedData())
				continue entry
			}
		}
		utils.AddLineToFile(exPath+cargoFile, cm.entry[i].SeedData())
	}

}

func (cl *cargoLot) SeedData() string {
	str := "CARGOENTRY:"
	//CARGOENTRY:148_Workable Alloys_16_Drinax_4600_-6_0_TRUE_Asim_118-1106_50_1_Freight
	str += cl.GetTGCode() + "_" + cl.GetDescr() + "_" + strconv.Itoa(cl.GetVolume()) + "_" + cl.GetOrigin() + "_" + strconv.Itoa(cl.GetCost()) +
		"_" + strconv.Itoa(cl.GetDangerDM()) + "_" + strconv.Itoa(cl.GetRiskDM()) + "_" + strconv.FormatBool(cl.GetLegality()) + "_" + cl.GetDestination() + "_" + cl.GetETA() + "_" +
		strconv.Itoa(cl.GetInsurance()) + "_" + cl.GetSupplierType() + "_" + cl.GetComment() + "_" + strconv.Itoa(cl.GetID())
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
		mn = 1
		if i == 2 {
			mn = 30
		}
		if i == 3 {
			mn = 30 * 30
		}
		if i == 4 {
			mn = 30 * 30 * 30
		}
		if i == 5 {
			mn = 30 * 30 * 30 * 30
		}
		if i == 6 {
			mn = 30 * 30 * 30 * 30 * 30
		}
		if i == 7 {
			mn = 30 * 30 * 30 * 30 * 30 * 30
		}
		if i == 8 {
			mn = 30 * 30 * 30 * 30 * 30 * 30 * 30
		}
		if i == 9 {
			mn = 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30
		}
		if i == 10 {
			mn = 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30
		}
		if i == 11 {
			mn = 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30
		}
		if i == 12 {
			mn = 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30 * 30
		}
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

func (cl *cargoLot) detailsFreight(code string, volume int, fee int) {
	cl.SetTGCode(code)
	cl.SetDestination(targetWorld.Hex())
	cl.SetLegality(true)
	cl.SetOrigin(sourceWorld.Hex())
	cl.SetSupplierType(constant.MerchantTypeTrade)
	cl.SetETA(integerToEhexCode(eta))
	cl.SetDescr(trade.GetDescription(code))
	cl.SetInsurance(dice.Roll("1d10").Sum() * 10)
	cl.SetComment("Freight fee " + strconv.Itoa(fee) + " Cr")
	cl.SetVolume(volume)
	cl.SetID(int(time.Now().UnixNano()))
	cl.SetCost(fee)

}

func deleteFromPortCargo(i int) {
	portCargo = append(portCargo[:i], portCargo[i+1:]...)
}

func freeCargoVolume() int {
	overall := getShipData("SHIP_CARGO_VOLUME")
	cm := loadCargoManifest()
	filled := 0
	for _, val := range cm.entry {

		filled += val.GetVolume()
	}
	return overall - filled
}

func unloadCargo() {
	cm := loadCargoManifest()
	totalProfit := 0
	for _, val := range cm.entry {
		if val.GetDestination() == sourceWorld.Hex() {
			fmt.Print("Unloaded ", val.GetVolume(), " tons of "+val.GetDescr(), "\n")
			profit := val.GetCost()
			missedETA := rawDay - ehexCodeToInteger(val.GetETA())
			//fmt.Println("missedETA", missedETA, ehexCodeToInteger(val.GetETA()), rawDay, ehexCodeToInteger(val.GetETA()), val.GetETA())
			penny := 0
			if missedETA > 7 {
				penny = (dice.Roll("1d6").Sum() + 4)
			}
			if missedETA > 14 {
				penny = (dice.Roll("1d6").Sum() + 4) * 2
			}
			if missedETA > 21 {
				penny = (dice.Roll("1d6").Sum() + 4) * 4
			}
			if missedETA > 28 {
				penny = (dice.Roll("1d6").Sum() + 4) * 8
			}
			if missedETA > 35 {
				penny = (dice.Roll("1d6").Sum() + 4) * 10
			}
			//fmt.Println("penny", penny)
			profit = profit - (int((float64(profit) / 100)) * penny)
			//fmt.Println("profit", profit)
			fmt.Print(profit, " Cr was received")
			if penny > 0 {
				fmt.Print(" (", penny, "% was taxed due to ETA was missed by ", missedETA, " days)")
			}
			fmt.Print("\n")
			totalProfit += profit
			deleteFromCargoManifest(val.GetID())
		}

	}
	fmt.Print("------------------------------\n")
	fmt.Print("Total Profit: ", totalProfit, " Cr\n")
}

func loadCargo() {
	done := false
	cm := loadCargoManifest()
	for !done {
		clrScrn()
		fmt.Println("Free Volume: ", freeCargoVolume())
		allLots := []string{}
		allLots = append(allLots, "Return")
		for i := range portCargo {
			if freeCargoVolume() >= portCargo[i].GetVolume() {
				allLots = append(allLots, lotInfo(portCargo[i]))
			}
		}
		if len(allLots) == 1 {
			break
		}
		selected, lot := menu("Load Cargo:", allLots...)
		if selected == 0 {
			done = true
			continue
		}
		cm.entry = append(cm.entry, portCargo[selected-1])
		fmt.Println(lot, "was loaded to ship")
		deleteFromPortCargo(selected - 1)
		saveCargoManifest(cm)

	}
	clrScrn()
}

func cargoDesignatedTo(w wrld.World) int {
	cm := loadCargoManifest()
	d := 0
	for _, val := range cm.entry {
		if val.destinationWorld == w.Hex() {
			d++
		}
	}
	return d
}

func reserveCargoSpace() {
	clrScrn()
	cm := loadCargoManifest()
	tons := -1
	descr := ""
	err := errors.New("No value for tons")
	for err != nil {
		fmt.Print("Enter volume in tons to designate: ")
		tons, err = user.InputInt()
		if tons == 0 {
			err = errors.New("WARNING: Zero value entered")
			reportErr(err)
			if userConfirm("Abort operation?") {
				return
			}
			continue
		}
		if tons < 0 {
			err = errors.New("Error: Tons can't be negative")
		}
		if tons > freeCargoVolume() {
			err = errors.New("Error: Tons can't be more than free volume")
		}
		reportErr(err)
	}
	err = errors.New("No value for descr")
	for err != nil {
		fmt.Print("Enter volume reserve description: ")
		descr, err = user.InputStr()
		reportErr(err)
	}
	cl := newCargoLot()
	cl.SetVolume(tons)
	cl.SetDescr(descr)
	cm.entry = append(cm.entry, cl)
	saveCargoManifest(cm)
}

func editCargoEntryVolume() {
	clrScrn()
	cm := loadCargoManifest()
	allLots := []string{}
	allLots = append(allLots, "Return")
	for _, val := range cm.entry {
		allLots = append(allLots, strconv.Itoa(val.GetVolume())+" tons of "+val.GetDescr()+"	(ID:"+strconv.Itoa(val.GetID())+")")
	}
	i, _ := menu("Select Cargo to edit:", allLots...)
	if i == 0 {
		return
	}
	err := errors.New("No value for 'newVal'")
	newVal := 0
	for err != nil {
		fmt.Print("Enter new volume: ")
		newVal, err = user.InputInt()
		if newVal == 0 {
			err = errors.New("Zero value entered") //delete?
			reportErr(err)
			if userConfirm("Delete entry: " + allLots[i] + "?") {
				deleteFromCargoManifest(cm.entry[i-1].GetID())
				return
			}
			err = nil
		}
		if newVal < 0 {
			err = errors.New("Error: New volume can't be negative")
		}
		if newVal > cm.entry[i-1].GetVolume()+freeCargoVolume() {
			err = errors.New("Error: Not enough Cargo Volume")
		}
		cm.entry[i-1].SetVolume(newVal)
		reportErr(err)
	}
	saveCargoManifest(cm)
	menuPosition = "HANGAR"
}

func deleteFromCargoManifest(idValue int) cargoManifest {
	id := strconv.Itoa(idValue)
	n := utils.InFileContains(exPath+cargoFile, id)
	utils.DeleteLineFromFileN(exPath+cargoFile, n)
	return loadCargoManifest()
}
