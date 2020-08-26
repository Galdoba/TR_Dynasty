package trademgt2

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
	trade "github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/devtools/cli/prettytable"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

//var sectorMapByHex map[string]string
//var sectorMapByName map[string]string
//var sectorMapByUWP map[string]string
var sourceWorld world.World
var targetWorld world.World
var ftv int
var dist int
var freigtTable prettytable.PrettyTable
var freightLots []*freightLot

func init() {
	//sectorData := otu.TrojanReachData()
	//sectorMapByHex = otu.MapDataByHex(sectorData)
	//sectorMapByName = otu.MapDataByName(sectorData)
	//sectorMapByUWP = otu.MapDataByUWP(sectorData)

}

func printHead() {
	fmt.Println("Curent World: " + sourceWorld.Hex() + " " + sourceWorld.Name() + " (" + sourceWorld.UWP() + ") " + sliceToStr(sourceWorld.TradeCodes()))
	fmt.Println("Target World: " + targetWorld.Hex() + " " + targetWorld.Name() + " (" + targetWorld.UWP() + ") " + sliceToStr(targetWorld.TradeCodes()))
	fmt.Println("Direct Jump Distance:", dist)
	fmt.Println("Freight Trafic Value:", ftv)
	fmt.Println("Passenger Trafic Value:", ptv)
	fmt.Println("Transit Hazard: [Not Implemented]")
	fmt.Println("------------------------------------------------------")
}

func sliceToStr(sl []string) string {
	str := ""
	for i := range sl {
		str = str + sl[i] + " "
	}
	str = strings.TrimSuffix(str, " ")
	return str
}

//RunMerchantPrince - точка входа
func RunMerchantPrince() {
	clrScrn()
	sourceWorld = LoadWorld("Current World (Name, Hex or UWP): ")
	clrScrn()
	targetWorld = LoadWorld("Target World (Name, Hex or UWP): ")
	dist = Astrogation.JumpDistance(sourceWorld.Hex(), targetWorld.Hex())
	ftv = freightTrafficValue(sourceWorld, targetWorld)
	clrScrn()
	fmt.Println("Searching Freight Contracts...")
	//playerEffect := userInputInt("Enter Diplomat(8), Investigate(8) or Streetwise(8) check effect: ")
	playerEffect := dice.Roll("2d6").DM(-8).Sum()
	fmt.Println("Auto Roll:", playerEffect)

	inLot, mnLot, mjLot := AvailableFreightLots(ftv + playerEffect)
	freightLots = LotList(inLot, mnLot, mjLot)
	clrScrn()
	FreightProcedure()

	MailProcedure()
	PassengerProcedure()
}

func renegotiateFreightLots() {
	for {
		clrScrn()
		informAboutLots()

		fmt.Println("------------------------------------------------------")

		if len(freightLots) == 0 {
			break
		}
		uchoise := userInputInt("Pick lot (-1 if none): ")
		if uchoise < 0 || uchoise > len(freightLots)-1 {
			if uchoise == -1 {
				break
			}
			continue
		}
		freightLots[uchoise].Negotiate()
	}
}

func informAboutLots() {
	defer fmt.Println("------------------------------------------------------")
	if len(freightLots) < 1 {
		fmt.Println("No Freight Lots available")
		return
	}
	tb := prettytable.New()
	tb.AddRow([]string{"Lot #", "Tons volume", "Freight Offer", "Cargo Manifest", "Renegotiated", "Risk Factor"})

	for i, lot := range freightLots {
		neg := "FALSE"
		if lot.Negotiated() {
			neg = "TRUE"
		}

		tb.AddRow([]string{"Lot " + strconv.Itoa(i), strconv.Itoa(lot.Tonns()) + " tons", strconv.Itoa(lot.Price()) + " Cr", lot.Descr(), neg, lot.Risk()})
	}

	tb.PTPrint()
}

//LoadWorld - кандидад на вынос (загружает данные из OTU таблицы)
func LoadWorld(msg string) world.World {
	sectorData := otu.TrojanReachData()
	sectorMapByHex := otu.MapDataByHex(sectorData)
	sectorMapByName := otu.MapDataByName(sectorData)
	sectorMapByUWP := otu.MapDataByUWP(sectorData)
	done := false
	key := ""
	otuData := ""
	for !done {
		key = userInputStr(msg)
		if val, ok := sectorMapByHex[key]; ok {
			otuData = val
			done = true
			continue
		}
		if val, ok := sectorMapByName[key]; ok {
			otuData = val
			done = true
			continue
		}
		key2 := strings.ToUpper(key)
		if val, ok := sectorMapByUWP[key2]; ok {
			otuData = val
			done = true
			continue
		}
		fmt.Println("No data by key '" + key + "'")
	}
	fmt.Println("Loading world data:")
	fmt.Println("Trojan Reach "+otu.Info{otuData}.Hex(), "-", otu.Info{otuData}.Name(), "("+otu.Info{otuData}.UWP()+")")
	w, err := world.FromOTUdata(otuData)
	if err != nil {
		panic(err)
	}
	return w
}

func userInputStr(msg string) string {
	done := false
	fmt.Print(msg)
	for !done {
		uwp, err := user.InputStr()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return uwp
	}
	return "Must not happen !!"
}

func userInputInt(msg string) int {
	done := false
	fmt.Print(msg)
	for !done {
		i, err := user.InputInt()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return i
	}
	return -999
}

func freightTrafficValue(sourceWorld, targetWorld world.World) int {
	dm := TrvCore.EhexToDigit(sourceWorld.PlanetaryData("Pops"))
	for _, val := range sourceWorld.TradeCodes() {
		switch val {
		default:
		case constant.TradeCodeAgricultural:
			dm += 2
		case constant.TradeCodeAsteroid:
			dm += -3
		case constant.TradeCodeBarren:
			return 0
		case constant.TradeCodeDesert:
			dm += -3
		case constant.TradeCodeFluidOceans:
			dm += -3
		case constant.TradeCodeGarden:
			dm += 2
		case constant.TradeCodeHighPopulation:
			dm += 2
		case constant.TradeCodeIceCapped:
			dm += -3
		case constant.TradeCodeIndustrial:
			dm += 3
		case constant.TradeCodeLowPopulation:
			dm += -5
		case constant.TradeCodeNonAgricultural:
			dm += -3
		case constant.TradeCodeNonIndustrial:
			dm += -3
		case constant.TradeCodePoor:
			dm += -3
		case constant.TradeCodeRich:
			dm += 2
		case constant.TradeCodeWaterWorld:
			dm += -3
		case constant.TravelCodeAmber:
			dm += 5
		case constant.TravelCodeRed:
			dm += -5
		}
	}
	//dm += TrvCore.EhexToDigit(targetWorld.PlanetaryData("Pops")) // - да хер пойми надо оно или нет (конфликт правил MgT1:MP p.66)
	for _, val := range targetWorld.TradeCodes() {
		switch val {
		default:
		case constant.TradeCodeAgricultural:
			dm++
		case constant.TradeCodeAsteroid:
			dm++
		case constant.TradeCodeBarren:
			dm += -5
		case constant.TradeCodeDesert:
			dm += 0
		case constant.TradeCodeFluidOceans:
			dm += 0
		case constant.TradeCodeGarden:
			dm++
		case constant.TradeCodeHighPopulation:
			dm += 0
		case constant.TradeCodeIceCapped:
			dm += 0
		case constant.TradeCodeIndustrial:
			dm += 2
		case constant.TradeCodeLowPopulation:
			dm += 0
		case constant.TradeCodeNonAgricultural:
			dm++
		case constant.TradeCodeNonIndustrial:
			dm++
		case constant.TradeCodePoor:
			dm += -3
		case constant.TradeCodeRich:
			dm += 2
		case constant.TradeCodeWaterWorld:
			dm += 0
		case constant.TravelCodeAmber:
			dm += -5
		case constant.TravelCodeRed:
			return 0
		}
	}

	tl1 := TrvCore.EhexToDigit(sourceWorld.PlanetaryData(constant.PrTL))
	tl2 := TrvCore.EhexToDigit(targetWorld.PlanetaryData(constant.PrTL))
	tlDiff := utils.Max(tl1, tl2) - utils.Min(tl1, tl2)
	if tlDiff > 5 {
		tlDiff = 5
	}
	dm -= tlDiff
	return dm
}

func AvailableFreightLots(tfv int) (int, int, int) {
	inLots := utils.Max(0, dice.Roll("1d6").DM(tfv-13).Sum())
	mnLots := utils.Max(0, dice.Roll("1d6").DM(tfv-8).Sum())
	mjLots := utils.Max(0, dice.Roll("1d6").DM(tfv-6).Sum())
	return inLots, mnLots, mjLots
}

type freightLot struct {
	tonns      int
	basePrice  int
	source     string
	target     string
	dist       int
	negotiated bool
	cat        string
}

// type lot interface {
// 	Price() int
// 	Tonns() int
// 	Negotiated() bool
// 	Descr() string
// 	Risk() string
// 	Negotiate(int)
// }

func NewFreightLot(sourceHex, targetHex string, tons int) *freightLot {
	frL := freightLot{}
	frL.tonns = tons
	frL.source = sourceHex
	frL.target = targetHex
	frL.dist = Astrogation.JumpDistance(frL.source, frL.target)
	frL.basePrice = 500
	frL.negotiated = false
	frL.cat = trade.RandomTGCategory(sourceWorld) + dice.Roll("2d6").SumStr()
	return &frL
}

func (l *freightLot) Price() int {
	price := l.basePrice * l.tonns * l.dist
	bonus := (price / 5) * l.dist
	return price + bonus
}

func (l *freightLot) Tonns() int {
	return l.tonns
}

func (l *freightLot) Descr() string {
	return trade.GetDescription(l.cat)
}

func (l *freightLot) Risk() string {
	risk := trade.GetDangerousGoodsDM(l.cat)
	return strconv.Itoa(risk)
}

func (l *freightLot) Negotiated() bool {
	return l.negotiated
}

func LotList(inLot, mnLot, mjLot int) []*freightLot {
	var tons []int
	fmt.Print("Searching available Freight lots")
	for i := 0; i < mjLot; i++ {
		tons = append(tons, dice.Roll("1d6").Sum()*10)
	}
	for i := 0; i < mnLot; i++ {
		tons = append(tons, dice.Roll("1d6").Sum()*5)
	}
	for i := 0; i < inLot; i++ {
		tons = append(tons, dice.Roll("1d6").Sum()*1)
	}
	sort.Ints(tons)
	for i := range tons {
		freightLots = append(freightLots, NewFreightLot(sourceWorld.Hex(), targetWorld.Hex(), tons[i]))
	}
	return freightLots
}

func (l *freightLot) Negotiate() {
	if l.negotiated {
		fmt.Println("Price on this lot was already renegotiated")
		return
	}
	tn := brokerPersuadeDiff(ftv)
	eff := userInputInt("Enter Broker(" + strconv.Itoa(tn) + ") or Persuade(" + strconv.Itoa(tn) + ") check effect: ")
	l.basePrice = negotiateEffect(eff)
	fmt.Println("After Negotiatios price for this lot was set as:", l.Price())
	l.negotiated = true
}

func negotiateEffect(eff int) int {
	base := 0
	switch eff {
	default:
		if eff < -4 {
			base = 200
		}
		if eff > 4 {
			base = 1000
		}
	case -3, -4:
		base = 300
	case -2, -1:
		base = 400
	case 0:
		base = 500
	case 1, 2:
		base = 600
	case 3, 4:
		base = 750
	}
	return base
}

func brokerPersuadeDiff(ftv int) int {
	switch ftv {
	default:
		if ftv < 5 {
			return 6
		}
		if ftv > 14 {
			return 10
		}
	case 5, 6, 7:
		return 7
	case 8, 9, 10, 11:
		return 8
	case 12, 13, 14:
		return 9
	}
	return 0
}

func clrScrn() {
	var clear map[string]func()
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
	printHead()
}

func FreightProcedure() {

	informAboutLots()
}

/*
1 собираем инфу по планетам
выводим шапку
2 собираем инфу по перевозкам
3 собираем инфу по почте
4 собираем инфу по пассажирам





*/
