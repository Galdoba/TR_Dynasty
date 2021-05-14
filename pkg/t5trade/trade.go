package t5trade

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/utils"
)

const (
	Consumables    = "Consumables"
	Data           = "Data"
	Entertainments = "Entertainments"
	Imbalances     = "Imbalances"
	Manufactureds  = "Manufactureds"
	Novelties      = "Novelties"
	Pharma         = "Pharma"
	Rares          = "Rares"
	Raws           = "Raws"
	RedTape        = "Red Tape"
	Samples        = "Samples"
	ScrapWaste     = "Scrap/Waste"
	Uniques        = "Uniques"
	Valuta         = "Valuta"
)

func Test() {
	fmt.Println("Select Source World:")
	sourceworld := wrld.PickWorld()
	fmt.Println(sourceworld.SecondSurvey())
	fmt.Println("Select Market World:")
	marketworld := wrld.PickWorld()
	fmt.Println(marketworld.SecondSurvey())
	fmt.Println("=============")
	cargo := NewCargo(sourceworld.GetСharacteristic(constant.PrTL).Value(), sourceworld.TradeCodes())
	sell := SellPrice(cargo, marketworld.GetСharacteristic(constant.PrTL).Value(), sourceworld.TradeCodes())
	fmt.Println(cargo.FullName())
	fmt.Println(cargo.ID())
	fmt.Println("Sell on", marketworld.Name(), "for", sell, "Cr")
	fmt.Println("=============")
	//cargo.determineTradeGoodDetails()
	//fmt.Println(cargo)

}

/*

Lot.
Freight.
Cargo.
Speculator.
Market World.
Cost.
Price.
Selling Price.
Delivery. - two common forms of shipment are
OTO. Orbit To Orbit.
STS. Surface To Surface.
*/

type Cargo struct {
	CargoID    string //A cargo is identified by stating its source world’s Tech Level, Trade Classifications, and Cost. Tech Level is derived directly from the source world UWP. All trade classifications possible are determined and then listed together
	Cost       int    //Purchase Cost at Sourceworld
	Imbalances int    //supply lines thing bonus to sell price
	ImbSource  string //supply lines thing bonus to sell price
	Name       string
	Type       string
	Prefix     string
}

func NewCargo(swTechLvl int, swTradeCodes []string) Cargo {
	c := Cargo{}
	c.CargoID = determineID(swTechLvl, swTradeCodes)
	c.Cost = determineCost(c.CargoID)
	c.CargoID = c.CargoID + "Cr" + strconv.Itoa(c.Cost)
	c.determineTradeGoodDetails()
	c.determinePrefix()
	return c
}

func (c *Cargo) FullName() string {
	nm := ""
	if c.Prefix != "" {
		nm += c.Prefix + " "
	}
	nm += c.Name
	return nm
}

func (c *Cargo) ID() string {
	return c.CargoID
}

func SellPrice(c Cargo, mwTL int, mwTC []string) int {
	tc := strings.Split(c.CargoID, " ")
	sell := 5000
	//sell = sell + (c.Imbalances * 1000)
	for _, cargoCode := range tc {
		switch cargoCode {
		case "Ag":
			if slicesOverlap([]string{"Ag", "As", "De", "Hi", "In", "Ri", "Va"}, mwTC) {
				sell += 1000
			}
		case "As":
			if slicesOverlap([]string{"As", "In", "Ri", "Va"}, mwTC) {
				sell += 1000
			}
		case "Ba":
			if slicesOverlap([]string{"In"}, mwTC) {
				sell += 1000
			}
		case "Fl":
			if slicesOverlap([]string{"Fl", "In"}, mwTC) {
				sell += 1000
			}
		case "Hi":
			if slicesOverlap([]string{"Hi"}, mwTC) {
				sell += 1000
			}
		case "In":
			if slicesOverlap([]string{"Ag", "As", "De", "Fl", "Hi", "In", "Ri", "Va"}, mwTC) {
				sell += 1000
			}
		case "Na":
			if slicesOverlap([]string{"As", "De", "Va"}, mwTC) {
				sell += 1000
			}
		case "Po":
			if slicesOverlap([]string{"Ag", "Hi", "In", "Ri"}, mwTC) {
				sell += 1000
			}
		case "Ri":
			if slicesOverlap([]string{"Ag", "De", "Hi", "In", "Ri"}, mwTC) {
				sell += 1000
			}
		case "Va":
			if slicesOverlap([]string{"As", "In", "Va"}, mwTC) {
				sell += 1000
			}
		}
	}
	tlEffect := sell * ((ehex.New(tc[0]).Value() - ehex.New(mwTL).Value()) * 10) / 100
	sell += tlEffect
	if utils.ListContains(mwTC, c.ImbSource) {
		sell += 1000
		fmt.Println("IT WORKS!!")
		fmt.Println(c)
		panic(sell)
	}
	return sell
}

func (c *Cargo) determinePrefix() {
	tc := strings.Split(c.CargoID, " ")
	options := []string{}
	if utils.ListContains(tc, "As") {
		options = append(options, "Strange")
	}
	if utils.ListContains(tc, "Ba") {
		options = append(options, "Gathered")
	}
	if utils.ListContains(tc, "De") {
		options = append(options, "Mineral")
	}
	if utils.ListContains(tc, "Di") {
		options = append(options, "Artifact")
	}
	if utils.ListContains(tc, "Fl") {
		options = append(options, "Unusual")
	}
	if utils.ListContains(tc, "Ga") {
		options = append(options, "Premium")
	}
	if utils.ListContains(tc, "He") {
		options = append(options, "Strange")
	}
	if utils.ListContains(tc, "Hi") {
		if !utils.ListContains(tc, "In") {
			options = append(options, "Processed")
		}
	}
	if utils.ListContains(tc, "Ic") {
		options = append(options, "Cryo")
	}
	if utils.ListContains(tc, "Lo") {
		options = append(options, "")
	}
	if utils.ListContains(tc, "Ni") {
		options = append(options, "Unprocessed")
	}
	if utils.ListContains(tc, "Oc") {
		options = append(options, "")
	}
	if utils.ListContains(tc, "Po") {
		options = append(options, "Obscure")
	}
	if utils.ListContains(tc, "Ri") {
		options = append(options, "Quality")
	}
	if utils.ListContains(tc, "Va") {
		if !utils.ListContains(tc, "As") {
			options = append(options, "Exotic")
		}
	}
	if utils.ListContains(tc, "Wa") {
		options = append(options, "Infused")
	}
	if len(options) == 0 {
		options = append(options, "")
	}
	c.Prefix = dice.New().RollFromList(options)
}

func slicesOverlap(sl1, sl2 []string) bool {
	for _, v1 := range sl1 {
		for _, v2 := range sl2 {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func determineID(tl int, tc []string) string {
	id := ehex.New(tl).String() + " - "
	for _, val := range tc {
		for _, tcx := range applicableTradeClassifications() {
			if val == tcx {
				id += val + " "
			}
		}
	}
	return id
}

func applicableTradeClassifications() []string {
	return []string{"Ag", "As", "Ba", "De", "Fl", "Hi", "In", "Lo", "Ni", "Po", "Ri", "Va"}
}

func determineCost(cargoID string) int {
	cost := 3000
	tc := strings.Split(cargoID, " ")
	for _, code := range tc {
		switch code {
		case "Ag", "As", "Hi", "In", "Po":
			cost = cost - 1000
		case "Ba", "De", "Fl", "Lo", "Ni", "Ri", "Va":
			cost = cost + 1000
		}
	}
	cost = cost + (ehex.New(tc[0]).Value() * 100)
	return cost
}

func (c *Cargo) determineTradeGoodDetails() string {
	descr := ""
	varityTags := strings.Split(c.CargoID, " ")
	tgIndex := newIndex(varityTags)
	tgdata := mapIndex(tgIndex)
	for tgdata.tgType == Imbalances {
		c.Imbalances++
		c.ImbSource = tgdata.tgName
		varityTags = []string{tgdata.tgName}
		tgIndex = newIndex(varityTags)
		tgdata = mapIndex(tgIndex)
	}
	c.Name = tgdata.tgName
	c.Type = tgdata.tgType
	return descr
}

func newIndex(varityTags []string) tgIndex {
	chart := varityChart(varityTags)
	dp := dice.New()
	l := strconv.Itoa(len(chart))
	r := dp.RollNext("1d" + l).DM(-1).Sum()
	chrt := chart[r]
	gType := dp.RollNext("1d6").Sum()
	gIndex := dp.RollNext("1d6").Sum()
	return tgIndex{chrt, gType, gIndex}
}

func varityChart(varityTags []string) []int {
	chart := []int{}
	dp := dice.New()
	if utils.ListContains(varityTags, "Ag") {
		ag := dp.RollNext("1d2").Sum()
		chart = append(chart, ag)
	}
	if utils.ListContains(varityTags, "As") {
		chart = append(chart, 3)
	}
	if utils.ListContains(varityTags, "De") {
		chart = append(chart, 4)
	}
	if utils.ListContains(varityTags, "Fl") {
		chart = append(chart, 5)
	}
	if utils.ListContains(varityTags, "Ic") {
		chart = append(chart, 6)
	}
	if utils.ListContains(varityTags, "Na") {
		chart = append(chart, 7)
	}
	if utils.ListContains(varityTags, "In") {
		chart = append(chart, 8)
	}
	if utils.ListContains(varityTags, "Po") {
		chart = append(chart, 9)
	}
	if utils.ListContains(varityTags, "Ri") {
		chart = append(chart, 10)
	}
	if utils.ListContains(varityTags, "Va") {
		chart = append(chart, 11)
	}
	if utils.ListContains(varityTags, "Cp") || utils.ListContains(varityTags, "Cs") || utils.ListContains(varityTags, "Cx") {
		chart = append(chart, 12)
	}
	if len(chart) == 0 {
		chart = append(chart, 7)
	}
	return chart
}

type tgIndex struct { //tradeGoodIndex
	chart  int
	gType  int
	gIndex int
}

type tgData struct {
	tgName string
	tgType string
}

func mapIndex(tgi tgIndex) tgData {
	tgDataMap := make(map[tgIndex]tgData)
	//Ag-1
	tgDataMap[tgIndex{1, 1, 1}] = tgData{"Bulk Protein", Raws}
	tgDataMap[tgIndex{2, 1, 1}] = tgData{"Bulk Woods", Raws}
	tgDataMap[tgIndex{3, 1, 1}] = tgData{"Bulk Nitrates", Raws}
	tgDataMap[tgIndex{4, 1, 1}] = tgData{"Bulk Nitrates", Raws}
	tgDataMap[tgIndex{5, 1, 1}] = tgData{"Bulk Carbon", Raws}
	tgDataMap[tgIndex{6, 1, 1}] = tgData{"Bulk Ices", Raws}
	tgDataMap[tgIndex{1, 1, 2}] = tgData{"Bulk Carbs", Raws}
	tgDataMap[tgIndex{2, 1, 2}] = tgData{"Bulk Pelts", Raws}
	tgDataMap[tgIndex{3, 1, 2}] = tgData{"Bulk Carbon", Raws}
	tgDataMap[tgIndex{4, 1, 2}] = tgData{"Bulk Minerals", Raws}
	tgDataMap[tgIndex{5, 1, 2}] = tgData{"Bulk Petros", Raws}
	tgDataMap[tgIndex{6, 1, 2}] = tgData{"Bulk Precipitates", Raws}
	tgDataMap[tgIndex{1, 1, 3}] = tgData{"Bulk Fats", Raws}
	tgDataMap[tgIndex{2, 1, 3}] = tgData{"Bulk Herbs", Raws}
	tgDataMap[tgIndex{3, 1, 3}] = tgData{"Bulk Iron", Raws}
	tgDataMap[tgIndex{4, 1, 3}] = tgData{"Bulk Abrasives", Raws}
	tgDataMap[tgIndex{5, 1, 3}] = tgData{"Bulk Precipitates", Raws}
	tgDataMap[tgIndex{6, 1, 3}] = tgData{"Bulk Ephemerals", Raws}
	tgDataMap[tgIndex{1, 1, 4}] = tgData{"Bulk Pharma", Raws}
	tgDataMap[tgIndex{2, 1, 4}] = tgData{"Bulk Spices", Raws}
	tgDataMap[tgIndex{3, 1, 4}] = tgData{"Bulk Copper", Raws}
	tgDataMap[tgIndex{4, 1, 4}] = tgData{"Bulk Particulates", Raws}
	tgDataMap[tgIndex{5, 1, 4}] = tgData{"Exotic Fluids", Raws}
	tgDataMap[tgIndex{6, 1, 4}] = tgData{"Exotic Flora", Raws}
	tgDataMap[tgIndex{1, 1, 5}] = tgData{"Livestock", Raws}
	tgDataMap[tgIndex{2, 1, 5}] = tgData{"Bulk Nitrates", Raws}
	tgDataMap[tgIndex{3, 1, 5}] = tgData{"Radioactive Ores", Raws}
	tgDataMap[tgIndex{4, 1, 5}] = tgData{"Exotic Fauna", Raws}
	tgDataMap[tgIndex{5, 1, 5}] = tgData{"Organic Polymers", Raws}
	tgDataMap[tgIndex{6, 1, 5}] = tgData{"Bulk Gases", Raws}
	tgDataMap[tgIndex{1, 1, 6}] = tgData{"Seedstock", Raws}
	tgDataMap[tgIndex{2, 1, 6}] = tgData{"Foodstuffs", Raws}
	tgDataMap[tgIndex{3, 1, 6}] = tgData{"Bulk Ices", Raws}
	tgDataMap[tgIndex{4, 1, 6}] = tgData{"Exotic Flora", Raws}
	tgDataMap[tgIndex{5, 1, 6}] = tgData{"Bulk Synthetics", Raws}
	tgDataMap[tgIndex{6, 1, 6}] = tgData{"Bulk Oxygen", Raws}
	tgDataMap[tgIndex{1, 2, 1}] = tgData{"Flavored Waters", Consumables}
	tgDataMap[tgIndex{2, 2, 1}] = tgData{"Flowers", Consumables}
	tgDataMap[tgIndex{3, 2, 1}] = tgData{"Ores", Samples}
	tgDataMap[tgIndex{4, 2, 1}] = tgData{"Archeologicals", Samples}
	tgDataMap[tgIndex{5, 2, 1}] = tgData{"Archeologicals", Samples}
	tgDataMap[tgIndex{6, 2, 1}] = tgData{"Archeologicals", Samples}
	tgDataMap[tgIndex{1, 2, 2}] = tgData{"Wines", Consumables}
	tgDataMap[tgIndex{2, 2, 2}] = tgData{"Aromatics", Consumables}
	tgDataMap[tgIndex{3, 2, 2}] = tgData{"Ices", Samples}
	tgDataMap[tgIndex{4, 2, 2}] = tgData{"Fauna", Samples}
	tgDataMap[tgIndex{5, 2, 2}] = tgData{"Fauna", Samples}
	tgDataMap[tgIndex{6, 2, 2}] = tgData{"Fauna", Samples}
	tgDataMap[tgIndex{1, 2, 3}] = tgData{"Juices", Consumables}
	tgDataMap[tgIndex{2, 2, 3}] = tgData{"Pheromones", Consumables}
	tgDataMap[tgIndex{3, 2, 3}] = tgData{"Carbons", Samples}
	tgDataMap[tgIndex{4, 2, 3}] = tgData{"Flora", Samples}
	tgDataMap[tgIndex{5, 2, 3}] = tgData{"Flora", Samples}
	tgDataMap[tgIndex{6, 2, 3}] = tgData{"Flora", Samples}
	tgDataMap[tgIndex{1, 2, 4}] = tgData{"Nectars", Consumables}
	tgDataMap[tgIndex{2, 2, 4}] = tgData{"Secretions", Consumables}
	tgDataMap[tgIndex{3, 2, 4}] = tgData{"Metals", Samples}
	tgDataMap[tgIndex{4, 2, 4}] = tgData{"Minerals", Samples}
	tgDataMap[tgIndex{5, 2, 4}] = tgData{"Germanes", Samples}
	tgDataMap[tgIndex{6, 2, 4}] = tgData{"Minerals", Samples}
	tgDataMap[tgIndex{1, 2, 5}] = tgData{"Decoctions", Consumables}
	tgDataMap[tgIndex{2, 2, 5}] = tgData{"Adhesives", Consumables}
	tgDataMap[tgIndex{3, 2, 5}] = tgData{"Uranium", Samples}
	tgDataMap[tgIndex{4, 2, 5}] = tgData{"Ephemerals", Samples}
	tgDataMap[tgIndex{5, 2, 5}] = tgData{"Flill", Samples}
	tgDataMap[tgIndex{6, 2, 5}] = tgData{"Luminescents", Samples}
	tgDataMap[tgIndex{1, 2, 6}] = tgData{"Drinkable Lymphs", Consumables}
	tgDataMap[tgIndex{2, 2, 6}] = tgData{"Novel Flavorings", Consumables}
	tgDataMap[tgIndex{3, 2, 6}] = tgData{"Chelates", Samples}
	tgDataMap[tgIndex{4, 2, 6}] = tgData{"Polymers", Samples}
	tgDataMap[tgIndex{5, 2, 6}] = tgData{"Chelates", Samples}
	tgDataMap[tgIndex{6, 2, 6}] = tgData{"Polymers", Samples}
	tgDataMap[tgIndex{1, 3, 1}] = tgData{"Health Foods", Pharma}
	tgDataMap[tgIndex{2, 3, 1}] = tgData{"Antifungals", Pharma}
	tgDataMap[tgIndex{3, 3, 1}] = tgData{"Platinum", Valuta}
	tgDataMap[tgIndex{4, 3, 1}] = tgData{"Stimulants", Pharma}
	tgDataMap[tgIndex{5, 3, 1}] = tgData{"Antifungals", Pharma}
	tgDataMap[tgIndex{6, 3, 1}] = tgData{"Antifungals", Pharma}
	tgDataMap[tgIndex{1, 3, 2}] = tgData{"Nutraceuticals", Pharma}
	tgDataMap[tgIndex{2, 3, 2}] = tgData{"Antivirals", Pharma}
	tgDataMap[tgIndex{3, 3, 2}] = tgData{"Gold", Valuta}
	tgDataMap[tgIndex{4, 3, 2}] = tgData{"Bulk Herbs", Pharma}
	tgDataMap[tgIndex{5, 3, 2}] = tgData{"Antivirals", Pharma}
	tgDataMap[tgIndex{6, 3, 2}] = tgData{"Antivirals", Pharma}
	tgDataMap[tgIndex{1, 3, 3}] = tgData{"Fast Drug", Pharma}
	tgDataMap[tgIndex{2, 3, 3}] = tgData{"Panacea", Pharma}
	tgDataMap[tgIndex{3, 3, 3}] = tgData{"Gallium", Valuta}
	tgDataMap[tgIndex{4, 3, 3}] = tgData{"Palliatives", Pharma}
	tgDataMap[tgIndex{5, 3, 3}] = tgData{"Palliatives", Pharma}
	tgDataMap[tgIndex{6, 3, 3}] = tgData{"Palliatives", Pharma}
	tgDataMap[tgIndex{1, 3, 4}] = tgData{"Painkillers", Pharma}
	tgDataMap[tgIndex{2, 3, 4}] = tgData{"Pseudomones", Pharma}
	tgDataMap[tgIndex{3, 3, 4}] = tgData{"Silver", Valuta}
	tgDataMap[tgIndex{4, 3, 4}] = tgData{"Pheromones", Pharma}
	tgDataMap[tgIndex{5, 3, 4}] = tgData{"Counter-prions", Pharma}
	tgDataMap[tgIndex{6, 3, 4}] = tgData{"Restoratives", Pharma}
	tgDataMap[tgIndex{1, 3, 5}] = tgData{"Antiseptic", Pharma}
	tgDataMap[tgIndex{2, 3, 5}] = tgData{"Anagathics", Pharma}
	tgDataMap[tgIndex{3, 3, 5}] = tgData{"Thorium", Valuta}
	tgDataMap[tgIndex{4, 3, 5}] = tgData{"Antibiotics", Pharma}
	tgDataMap[tgIndex{5, 3, 5}] = tgData{"Antibiotics", Pharma}
	tgDataMap[tgIndex{6, 3, 5}] = tgData{"Antibiotics", Pharma}
	tgDataMap[tgIndex{1, 3, 6}] = tgData{"Antibiotics", Pharma}
	tgDataMap[tgIndex{2, 3, 6}] = tgData{"Slow Drug", Pharma}
	tgDataMap[tgIndex{3, 3, 6}] = tgData{"Radium", Valuta}
	tgDataMap[tgIndex{4, 3, 6}] = tgData{"Combat Drug", Pharma}
	tgDataMap[tgIndex{5, 3, 6}] = tgData{"Cold Sleep Pills", Pharma}
	tgDataMap[tgIndex{6, 3, 6}] = tgData{"Antiseptics", Pharma}
	tgDataMap[tgIndex{1, 4, 1}] = tgData{"Incenses", Novelties}
	tgDataMap[tgIndex{2, 4, 1}] = tgData{"Strange Seeds", Novelties}
	tgDataMap[tgIndex{3, 4, 1}] = tgData{"Unusual Rocks", Novelties}
	tgDataMap[tgIndex{4, 4, 1}] = tgData{"Envirosuits", Novelties}
	tgDataMap[tgIndex{5, 4, 1}] = tgData{"Silanes", Novelties}
	tgDataMap[tgIndex{6, 4, 1}] = tgData{"Heat Pumps", Novelties}
	tgDataMap[tgIndex{1, 4, 2}] = tgData{"Iridescents", Novelties}
	tgDataMap[tgIndex{2, 4, 2}] = tgData{"Motile Plants", Novelties}
	tgDataMap[tgIndex{3, 4, 2}] = tgData{"Fused Metals", Novelties}
	tgDataMap[tgIndex{4, 4, 2}] = tgData{"Reclamation Suits", Novelties}
	tgDataMap[tgIndex{5, 4, 2}] = tgData{"Lek Emitters", Novelties}
	tgDataMap[tgIndex{6, 4, 2}] = tgData{"Mag Emitters", Novelties}
	tgDataMap[tgIndex{1, 4, 3}] = tgData{"Photonics", Novelties}
	tgDataMap[tgIndex{2, 4, 3}] = tgData{"Reactive Plants", Novelties}
	tgDataMap[tgIndex{3, 4, 3}] = tgData{"Strange Crystals", Novelties}
	tgDataMap[tgIndex{4, 4, 3}] = tgData{"Navigators", Novelties}
	tgDataMap[tgIndex{5, 4, 3}] = tgData{"Aware Blockers", Novelties}
	tgDataMap[tgIndex{6, 4, 3}] = tgData{"Percept Blockers", Novelties}
	tgDataMap[tgIndex{1, 4, 4}] = tgData{"Pigments", Novelties}
	tgDataMap[tgIndex{2, 4, 4}] = tgData{"Reactive Woods", Novelties}
	tgDataMap[tgIndex{3, 4, 4}] = tgData{"Fine Dusts", Novelties}
	tgDataMap[tgIndex{4, 4, 4}] = tgData{"Dupe Masterpieces", Novelties}
	tgDataMap[tgIndex{5, 4, 4}] = tgData{"Soothants", Novelties}
	tgDataMap[tgIndex{6, 4, 4}] = tgData{"Silanes", Novelties}
	tgDataMap[tgIndex{1, 4, 5}] = tgData{"Noisemakers", Novelties}
	tgDataMap[tgIndex{2, 4, 5}] = tgData{"IR Emitters", Novelties}
	tgDataMap[tgIndex{3, 4, 5}] = tgData{"Magnetics", Novelties}
	tgDataMap[tgIndex{4, 4, 5}] = tgData{"ShimmerCloth", Novelties}
	tgDataMap[tgIndex{5, 4, 5}] = tgData{"Self-Solving Puzzles", Novelties}
	tgDataMap[tgIndex{6, 4, 5}] = tgData{"Cold Light Blocks", Novelties}
	tgDataMap[tgIndex{1, 4, 6}] = tgData{"Soundmakers", Novelties}
	tgDataMap[tgIndex{2, 4, 6}] = tgData{"Lek Emitters", Novelties}
	tgDataMap[tgIndex{3, 4, 6}] = tgData{"Light-Sensitives", Novelties}
	tgDataMap[tgIndex{4, 4, 6}] = tgData{"ANIFX Blocker", Novelties}
	tgDataMap[tgIndex{5, 4, 6}] = tgData{"Fluidic Timepieces", Novelties}
	tgDataMap[tgIndex{6, 4, 6}] = tgData{"VHDUS Blocker", Novelties}
	tgDataMap[tgIndex{1, 5, 1}] = tgData{"Fine Furs", Rares}
	tgDataMap[tgIndex{2, 5, 1}] = tgData{"Spices", Rares}
	tgDataMap[tgIndex{3, 5, 1}] = tgData{"Gemstones", Rares}
	tgDataMap[tgIndex{4, 5, 1}] = tgData{"Excretions", Rares}
	tgDataMap[tgIndex{5, 5, 1}] = tgData{"Flavorings", Rares}
	tgDataMap[tgIndex{6, 5, 1}] = tgData{"Unusual Ices", Rares}
	tgDataMap[tgIndex{1, 5, 2}] = tgData{"Meat Delicacies", Rares}
	tgDataMap[tgIndex{2, 5, 2}] = tgData{"Organic Gems", Rares}
	tgDataMap[tgIndex{3, 5, 2}] = tgData{"Alloys", Rares}
	tgDataMap[tgIndex{4, 5, 2}] = tgData{"Flavorings", Rares}
	tgDataMap[tgIndex{5, 5, 2}] = tgData{"Unusual Fluids", Rares}
	tgDataMap[tgIndex{6, 5, 2}] = tgData{"Cryo Alloys", Rares}
	tgDataMap[tgIndex{1, 5, 3}] = tgData{"Fruit Delicacies", Rares}
	tgDataMap[tgIndex{2, 5, 3}] = tgData{"Flavorings", Rares}
	tgDataMap[tgIndex{3, 5, 3}] = tgData{"Iridium Sponge", Rares}
	tgDataMap[tgIndex{4, 5, 3}] = tgData{"Nectars", Rares}
	tgDataMap[tgIndex{5, 5, 3}] = tgData{"Encapsulants", Rares}
	tgDataMap[tgIndex{6, 5, 3}] = tgData{"Rare Minerals", Rares}
	tgDataMap[tgIndex{1, 5, 4}] = tgData{"Candies", Rares}
	tgDataMap[tgIndex{2, 5, 4}] = tgData{"Aged Meats", Rares}
	tgDataMap[tgIndex{3, 5, 4}] = tgData{"Lanthanum", Rares}
	tgDataMap[tgIndex{4, 5, 4}] = tgData{"Pelts", Rares}
	tgDataMap[tgIndex{5, 5, 4}] = tgData{"Insidiants", Rares}
	tgDataMap[tgIndex{6, 5, 4}] = tgData{"Unusual Fluids", Rares}
	tgDataMap[tgIndex{1, 5, 5}] = tgData{"Textiles", Rares}
	tgDataMap[tgIndex{2, 5, 5}] = tgData{"Fermented Fluids", Rares}
	tgDataMap[tgIndex{3, 5, 5}] = tgData{"Isotopes", Rares}
	tgDataMap[tgIndex{4, 5, 5}] = tgData{"ANIFX Dyes", Rares}
	tgDataMap[tgIndex{5, 5, 5}] = tgData{"Corrosives", Rares}
	tgDataMap[tgIndex{6, 5, 5}] = tgData{"Cryogems", Rares}
	tgDataMap[tgIndex{1, 5, 6}] = tgData{"Exotic Sauces", Rares}
	tgDataMap[tgIndex{2, 5, 6}] = tgData{"Fine Aromatics", Rares}
	tgDataMap[tgIndex{3, 5, 6}] = tgData{"Anti-Matter", Rares}
	tgDataMap[tgIndex{4, 5, 6}] = tgData{"Seedstock", Rares}
	tgDataMap[tgIndex{5, 5, 6}] = tgData{"Exotic Aromatics", Rares}
	tgDataMap[tgIndex{6, 5, 6}] = tgData{"VHDUS Dyes", Rares}
	tgDataMap[tgIndex{1, 6, 1}] = tgData{"As", Imbalances}
	tgDataMap[tgIndex{2, 6, 1}] = tgData{"Po", Imbalances}
	tgDataMap[tgIndex{3, 6, 1}] = tgData{"Ag", Imbalances}
	tgDataMap[tgIndex{4, 6, 1}] = tgData{"Pheromones", Uniques}
	tgDataMap[tgIndex{5, 6, 1}] = tgData{"In", Imbalances}
	tgDataMap[tgIndex{6, 6, 1}] = tgData{"Fossils", Uniques}
	tgDataMap[tgIndex{1, 6, 2}] = tgData{"De", Imbalances}
	tgDataMap[tgIndex{2, 6, 2}] = tgData{"Ri", Imbalances}
	tgDataMap[tgIndex{3, 6, 2}] = tgData{"De", Imbalances}
	tgDataMap[tgIndex{4, 6, 2}] = tgData{"Artifacts", Uniques}
	tgDataMap[tgIndex{5, 6, 2}] = tgData{"Ri", Imbalances}
	tgDataMap[tgIndex{6, 6, 2}] = tgData{"Cryogems", Uniques}
	tgDataMap[tgIndex{1, 6, 3}] = tgData{"Fl", Imbalances}
	tgDataMap[tgIndex{2, 6, 3}] = tgData{"Va", Imbalances}
	tgDataMap[tgIndex{3, 6, 3}] = tgData{"Na", Imbalances}
	tgDataMap[tgIndex{4, 6, 3}] = tgData{"Sparx", Uniques}
	tgDataMap[tgIndex{5, 6, 3}] = tgData{"Ic", Imbalances}
	tgDataMap[tgIndex{6, 6, 3}] = tgData{"Vision Suppressant", Uniques}
	tgDataMap[tgIndex{1, 6, 4}] = tgData{"Ic", Imbalances}
	tgDataMap[tgIndex{2, 6, 4}] = tgData{"Ic", Imbalances}
	tgDataMap[tgIndex{3, 6, 4}] = tgData{"Po", Imbalances}
	tgDataMap[tgIndex{4, 6, 4}] = tgData{"Repulsant", Uniques}
	tgDataMap[tgIndex{5, 6, 4}] = tgData{"Na", Imbalances}
	tgDataMap[tgIndex{6, 6, 4}] = tgData{"Fission Suppressant", Uniques}
	tgDataMap[tgIndex{1, 6, 5}] = tgData{"Na", Imbalances}
	tgDataMap[tgIndex{2, 6, 5}] = tgData{"Na", Imbalances}
	tgDataMap[tgIndex{3, 6, 5}] = tgData{"Ri", Imbalances}
	tgDataMap[tgIndex{4, 6, 5}] = tgData{"Dominants", Uniques}
	tgDataMap[tgIndex{5, 6, 5}] = tgData{"Ag", Imbalances}
	tgDataMap[tgIndex{6, 6, 5}] = tgData{"Wafers", Uniques}
	tgDataMap[tgIndex{1, 6, 6}] = tgData{"In", Imbalances}
	tgDataMap[tgIndex{2, 6, 6}] = tgData{"In", Imbalances}
	tgDataMap[tgIndex{3, 6, 6}] = tgData{"Ic", Imbalances}
	tgDataMap[tgIndex{4, 6, 6}] = tgData{"Fossils", Uniques}
	tgDataMap[tgIndex{5, 6, 6}] = tgData{"Po", Imbalances}
	tgDataMap[tgIndex{6, 6, 6}] = tgData{"Cold Sleep Pills", Uniques}
	tgDataMap[tgIndex{7, 1, 1}] = tgData{"Bulk Abrasives", Raws}
	tgDataMap[tgIndex{8, 1, 1}] = tgData{"Electronics", Manufactureds}
	tgDataMap[tgIndex{9, 1, 1}] = tgData{"Bulk Nutrients", Raws}
	tgDataMap[tgIndex{10, 1, 1}] = tgData{"Bulk Foodstuffs", Raws}
	tgDataMap[tgIndex{11, 1, 1}] = tgData{"Bulk Dusts", Raws}
	tgDataMap[tgIndex{12, 1, 1}] = tgData{"Software", Data}
	tgDataMap[tgIndex{7, 1, 2}] = tgData{"Bulk Gases", Raws}
	tgDataMap[tgIndex{8, 1, 2}] = tgData{"Photonics", Manufactureds}
	tgDataMap[tgIndex{9, 1, 2}] = tgData{"Bulk Fibers", Raws}
	tgDataMap[tgIndex{10, 1, 2}] = tgData{"Bulk Protein", Raws}
	tgDataMap[tgIndex{11, 1, 2}] = tgData{"Bulk Minerals", Raws}
	tgDataMap[tgIndex{12, 1, 2}] = tgData{"Expert Systems", Data}
	tgDataMap[tgIndex{7, 1, 3}] = tgData{"Bulk Minerals", Raws}
	tgDataMap[tgIndex{8, 1, 3}] = tgData{"Magnetics", Manufactureds}
	tgDataMap[tgIndex{9, 1, 3}] = tgData{"Bulk Organics", Raws}
	tgDataMap[tgIndex{10, 1, 3}] = tgData{"Bulk Carbs", Raws}
	tgDataMap[tgIndex{11, 1, 3}] = tgData{"Bulk Metals", Raws}
	tgDataMap[tgIndex{12, 1, 3}] = tgData{"Databases", Data}
	tgDataMap[tgIndex{7, 1, 4}] = tgData{"Bulk Precipitates", Raws}
	tgDataMap[tgIndex{8, 1, 4}] = tgData{"Fluidics", Manufactureds}
	tgDataMap[tgIndex{9, 1, 4}] = tgData{"Bulk Minerals", Raws}
	tgDataMap[tgIndex{10, 1, 4}] = tgData{"Bulk Fats", Raws}
	tgDataMap[tgIndex{11, 1, 4}] = tgData{"Radioactive Ores", Raws}
	tgDataMap[tgIndex{12, 1, 4}] = tgData{"Upgrades", Data}
	tgDataMap[tgIndex{7, 1, 5}] = tgData{"Exotic Fauna", Raws}
	tgDataMap[tgIndex{8, 1, 5}] = tgData{"Polymers", Manufactureds}
	tgDataMap[tgIndex{9, 1, 5}] = tgData{"Bulk Textiles", Raws}
	tgDataMap[tgIndex{10, 1, 5}] = tgData{"Exotic Flora", Raws}
	tgDataMap[tgIndex{11, 1, 5}] = tgData{"Bulk Particulates", Raws}
	tgDataMap[tgIndex{12, 1, 5}] = tgData{"Backups", Data}
	tgDataMap[tgIndex{7, 1, 6}] = tgData{"Exotic Flora", Raws}
	tgDataMap[tgIndex{8, 1, 6}] = tgData{"Gravitics", Manufactureds}
	tgDataMap[tgIndex{9, 1, 6}] = tgData{"Exotic Flora", Raws}
	tgDataMap[tgIndex{10, 1, 6}] = tgData{"Exotic Fauna", Raws}
	tgDataMap[tgIndex{11, 1, 6}] = tgData{"Ephemerals", Raws}
	tgDataMap[tgIndex{12, 1, 6}] = tgData{"Raw Sensings", Data}
	tgDataMap[tgIndex{7, 2, 1}] = tgData{"Archeologicals", Samples}
	tgDataMap[tgIndex{8, 2, 1}] = tgData{"Obsoletes", ScrapWaste}
	tgDataMap[tgIndex{9, 2, 1}] = tgData{"Art", Entertainments}
	tgDataMap[tgIndex{10, 2, 1}] = tgData{"Echostones", Novelties}
	tgDataMap[tgIndex{11, 2, 1}] = tgData{"Branded Vacc Suits", Novelties}
	tgDataMap[tgIndex{12, 2, 1}] = tgData{"Incenses", Novelties}
	tgDataMap[tgIndex{7, 2, 2}] = tgData{"Fauna", Samples}
	tgDataMap[tgIndex{8, 2, 2}] = tgData{"Used Goods", ScrapWaste}
	tgDataMap[tgIndex{9, 2, 2}] = tgData{"Recordings", Entertainments}
	tgDataMap[tgIndex{10, 2, 2}] = tgData{"Self-Defenders", Novelties}
	tgDataMap[tgIndex{11, 2, 2}] = tgData{"Awareness Pinger", Novelties}
	tgDataMap[tgIndex{12, 2, 2}] = tgData{"Contemplatives", Novelties}
	tgDataMap[tgIndex{7, 2, 3}] = tgData{"Flora", Samples}
	tgDataMap[tgIndex{8, 2, 3}] = tgData{"Reparables", ScrapWaste}
	tgDataMap[tgIndex{9, 2, 3}] = tgData{"Writings", Entertainments}
	tgDataMap[tgIndex{10, 2, 3}] = tgData{"Attractants", Novelties}
	tgDataMap[tgIndex{11, 2, 3}] = tgData{"Strange Seeds", Novelties}
	tgDataMap[tgIndex{12, 2, 3}] = tgData{"Cold Welders", Novelties}
	tgDataMap[tgIndex{7, 2, 4}] = tgData{"Minerals", Samples}
	tgDataMap[tgIndex{8, 2, 4}] = tgData{"Radioactives", ScrapWaste}
	tgDataMap[tgIndex{9, 2, 4}] = tgData{"Tactiles", Entertainments}
	tgDataMap[tgIndex{10, 2, 4}] = tgData{"Sophont Cuisine", Novelties}
	tgDataMap[tgIndex{11, 2, 4}] = tgData{"Pigments", Novelties}
	tgDataMap[tgIndex{12, 2, 4}] = tgData{"Polymer Sheets", Novelties}
	tgDataMap[tgIndex{7, 2, 5}] = tgData{"Ephemerals", Samples}
	tgDataMap[tgIndex{8, 2, 5}] = tgData{"Metals", ScrapWaste}
	tgDataMap[tgIndex{9, 2, 5}] = tgData{"Osmancies", Entertainments}
	tgDataMap[tgIndex{10, 2, 5}] = tgData{"Sophont Hats", Novelties}
	tgDataMap[tgIndex{11, 2, 5}] = tgData{"Unusual Minerals", Novelties}
	tgDataMap[tgIndex{12, 2, 5}] = tgData{"Hats", Novelties}
	tgDataMap[tgIndex{7, 2, 6}] = tgData{"Polymers", Samples}
	tgDataMap[tgIndex{8, 2, 6}] = tgData{"Sludges", ScrapWaste}
	tgDataMap[tgIndex{9, 2, 6}] = tgData{"Wafers", Entertainments}
	tgDataMap[tgIndex{10, 2, 6}] = tgData{"Variable Tattoos", Novelties}
	tgDataMap[tgIndex{11, 2, 6}] = tgData{"Exotic Crystals", Novelties}
	tgDataMap[tgIndex{12, 2, 6}] = tgData{"Skin Tones", Novelties}
	tgDataMap[tgIndex{7, 3, 1}] = tgData{"Branded Tools", Novelties}
	tgDataMap[tgIndex{8, 3, 1}] = tgData{"Biologics", Manufactureds}
	tgDataMap[tgIndex{9, 3, 1}] = tgData{"Strange Crystals", Novelties}
	tgDataMap[tgIndex{10, 3, 1}] = tgData{"Branded Foods", Consumables}
	tgDataMap[tgIndex{11, 3, 1}] = tgData{"Branded Oxygen", Consumables}
	tgDataMap[tgIndex{12, 3, 1}] = tgData{"Branded Clothes", Consumables}
	tgDataMap[tgIndex{7, 3, 2}] = tgData{"Drinkable Lymphs", Novelties}
	tgDataMap[tgIndex{8, 3, 2}] = tgData{"Mechanicals", Manufactureds}
	tgDataMap[tgIndex{9, 3, 2}] = tgData{"Strange Seeds", Novelties}
	tgDataMap[tgIndex{10, 3, 2}] = tgData{"Branded Drinks", Consumables}
	tgDataMap[tgIndex{11, 3, 2}] = tgData{"Vacc Suit Scents", Consumables}
	tgDataMap[tgIndex{12, 3, 2}] = tgData{"Branded Devices", Consumables}
	tgDataMap[tgIndex{7, 3, 3}] = tgData{"Strange Seeds", Novelties}
	tgDataMap[tgIndex{8, 3, 3}] = tgData{"Textiles", Manufactureds}
	tgDataMap[tgIndex{9, 3, 3}] = tgData{"Pigments", Novelties}
	tgDataMap[tgIndex{10, 3, 3}] = tgData{"Branded Clothes", Consumables}
	tgDataMap[tgIndex{11, 3, 3}] = tgData{"Vacc Suit Patches", Consumables}
	tgDataMap[tgIndex{12, 3, 3}] = tgData{"Flavored Drinks", Consumables}
	tgDataMap[tgIndex{7, 3, 4}] = tgData{"Pattern Creators", Novelties}
	tgDataMap[tgIndex{8, 3, 4}] = tgData{"Weapons", Manufactureds}
	tgDataMap[tgIndex{9, 3, 4}] = tgData{"Emotion Lighting", Novelties}
	tgDataMap[tgIndex{10, 3, 4}] = tgData{"Flavored Drinks", Consumables}
	tgDataMap[tgIndex{11, 3, 4}] = tgData{"Branded Tools", Consumables}
	tgDataMap[tgIndex{12, 3, 4}] = tgData{"Flavorings", Consumables}
	tgDataMap[tgIndex{7, 3, 5}] = tgData{"Pigments", Novelties}
	tgDataMap[tgIndex{8, 3, 5}] = tgData{"Armor", Manufactureds}
	tgDataMap[tgIndex{9, 3, 5}] = tgData{"Silanes", Novelties}
	tgDataMap[tgIndex{10, 3, 5}] = tgData{"Flowers", Consumables}
	tgDataMap[tgIndex{11, 3, 5}] = tgData{"Holo-Companions", Consumables}
	tgDataMap[tgIndex{12, 3, 5}] = tgData{"Decorations", Consumables}
	tgDataMap[tgIndex{7, 3, 6}] = tgData{"Warm Leather", Novelties}
	tgDataMap[tgIndex{8, 3, 6}] = tgData{"Robots", Manufactureds}
	tgDataMap[tgIndex{9, 3, 6}] = tgData{"Flora", Novelties}
	tgDataMap[tgIndex{10, 3, 6}] = tgData{"Music", Consumables}
	tgDataMap[tgIndex{11, 3, 6}] = tgData{"Flavored Air", Consumables}
	tgDataMap[tgIndex{12, 3, 6}] = tgData{"Group Symbols", Consumables}
	tgDataMap[tgIndex{7, 4, 1}] = tgData{"Hummingsand", Rares}
	tgDataMap[tgIndex{8, 4, 1}] = tgData{"Nostrums", Pharma}
	tgDataMap[tgIndex{9, 4, 1}] = tgData{"Gemstones", Rares}
	tgDataMap[tgIndex{10, 4, 1}] = tgData{"Delicacies", Rares}
	tgDataMap[tgIndex{11, 4, 1}] = tgData{"Vacc Gems", Rares}
	tgDataMap[tgIndex{12, 4, 1}] = tgData{"Monumental Art", Rares}
	tgDataMap[tgIndex{7, 4, 2}] = tgData{"Masterpieces", Rares}
	tgDataMap[tgIndex{8, 4, 2}] = tgData{"Restoratives", Pharma}
	tgDataMap[tgIndex{9, 4, 2}] = tgData{"Antiques", Rares}
	tgDataMap[tgIndex{10, 4, 2}] = tgData{"Spices", Rares}
	tgDataMap[tgIndex{11, 4, 2}] = tgData{"Unusual Dusts", Rares}
	tgDataMap[tgIndex{12, 4, 2}] = tgData{"Holo Sculpture", Rares}
	tgDataMap[tgIndex{7, 4, 3}] = tgData{"Fine Carpets", Rares}
	tgDataMap[tgIndex{8, 4, 3}] = tgData{"Palliatives", Pharma}
	tgDataMap[tgIndex{9, 4, 3}] = tgData{"Collectibles", Rares}
	tgDataMap[tgIndex{10, 4, 3}] = tgData{"Tisanes", Rares}
	tgDataMap[tgIndex{11, 4, 3}] = tgData{"Insulants", Rares}
	tgDataMap[tgIndex{12, 4, 3}] = tgData{"Collectible Books", Rares}
	tgDataMap[tgIndex{7, 4, 4}] = tgData{"Isotopes", Rares}
	tgDataMap[tgIndex{8, 4, 4}] = tgData{"Chelates", Pharma}
	tgDataMap[tgIndex{9, 4, 4}] = tgData{"Allotropes", Rares}
	tgDataMap[tgIndex{10, 4, 4}] = tgData{"Nectars", Rares}
	tgDataMap[tgIndex{11, 4, 4}] = tgData{"Crafted Devices", Rares}
	tgDataMap[tgIndex{12, 4, 4}] = tgData{"Jewelry", Rares}
	tgDataMap[tgIndex{7, 4, 5}] = tgData{"Pelts", Rares}
	tgDataMap[tgIndex{8, 4, 5}] = tgData{"Antidotes", Pharma}
	tgDataMap[tgIndex{9, 4, 5}] = tgData{"Spices", Rares}
	tgDataMap[tgIndex{10, 4, 5}] = tgData{"Pelts", Rares}
	tgDataMap[tgIndex{11, 4, 5}] = tgData{"Rare Minerals", Rares}
	tgDataMap[tgIndex{12, 4, 5}] = tgData{"Museum Items", Rares}
	tgDataMap[tgIndex{7, 4, 6}] = tgData{"Seedstock", Rares}
	tgDataMap[tgIndex{8, 4, 6}] = tgData{"Antitoxins", Pharma}
	tgDataMap[tgIndex{9, 4, 6}] = tgData{"Seedstock", Rares}
	tgDataMap[tgIndex{10, 4, 6}] = tgData{"Variable Tattoos", Rares}
	tgDataMap[tgIndex{11, 4, 6}] = tgData{"Catalysts", Rares}
	tgDataMap[tgIndex{12, 4, 6}] = tgData{"Monumental Art", Rares}
	tgDataMap[tgIndex{7, 5, 1}] = tgData{"Masterpieces", Uniques}
	tgDataMap[tgIndex{8, 5, 1}] = tgData{"Software", Data}
	tgDataMap[tgIndex{9, 5, 1}] = tgData{"Masterpieces", Uniques}
	tgDataMap[tgIndex{10, 5, 1}] = tgData{"Antique Art", Uniques}
	tgDataMap[tgIndex{11, 5, 1}] = tgData{"Archeologicals", Samples}
	tgDataMap[tgIndex{12, 5, 1}] = tgData{"Coinage", Valuta}
	tgDataMap[tgIndex{7, 5, 2}] = tgData{"Unusual Rocks", Uniques}
	tgDataMap[tgIndex{8, 5, 2}] = tgData{"Databases", Data}
	tgDataMap[tgIndex{9, 5, 2}] = tgData{"Exotic Flora", Uniques}
	tgDataMap[tgIndex{10, 5, 2}] = tgData{"Masterpieces", Uniques}
	tgDataMap[tgIndex{11, 5, 2}] = tgData{"Fauna", Samples}
	tgDataMap[tgIndex{12, 5, 2}] = tgData{"Currency", Valuta}
	tgDataMap[tgIndex{7, 5, 3}] = tgData{"Artifacts", Uniques}
	tgDataMap[tgIndex{8, 5, 3}] = tgData{"Expert Systems", Data}
	tgDataMap[tgIndex{9, 5, 3}] = tgData{"Antiques", Uniques}
	tgDataMap[tgIndex{10, 5, 3}] = tgData{"Artifacts", Uniques}
	tgDataMap[tgIndex{11, 5, 3}] = tgData{"Flora", Samples}
	tgDataMap[tgIndex{12, 5, 3}] = tgData{"Money Cards", Valuta}
	tgDataMap[tgIndex{7, 5, 4}] = tgData{"Non-Fossil Carca", Uniques}
	tgDataMap[tgIndex{8, 5, 4}] = tgData{"Upgrades", Data}
	tgDataMap[tgIndex{9, 5, 4}] = tgData{"Incomprehensibles", Uniques}
	tgDataMap[tgIndex{10, 5, 4}] = tgData{"Fine Art", Uniques}
	tgDataMap[tgIndex{11, 5, 4}] = tgData{"Minerals", Samples}
	tgDataMap[tgIndex{12, 5, 4}] = tgData{"Gold", Valuta}
	tgDataMap[tgIndex{7, 5, 5}] = tgData{"Replicating Clays", Uniques}
	tgDataMap[tgIndex{8, 5, 5}] = tgData{"Backups", Data}
	tgDataMap[tgIndex{9, 5, 5}] = tgData{"Fossils", Uniques}
	tgDataMap[tgIndex{10, 5, 5}] = tgData{"Meson Barriers", Uniques}
	tgDataMap[tgIndex{11, 5, 5}] = tgData{"Ephemerals", Samples}
	tgDataMap[tgIndex{12, 5, 5}] = tgData{"Silver", Valuta}
	tgDataMap[tgIndex{7, 5, 6}] = tgData{"ANIFX Emitter", Uniques}
	tgDataMap[tgIndex{8, 5, 6}] = tgData{"Raw Sensings", Data}
	tgDataMap[tgIndex{9, 5, 6}] = tgData{"VHDUS Emitter", Uniques}
	tgDataMap[tgIndex{10, 5, 6}] = tgData{"Famous Wafers", Uniques}
	tgDataMap[tgIndex{11, 5, 6}] = tgData{"Polymers", Samples}
	tgDataMap[tgIndex{12, 5, 6}] = tgData{"Platinum", Valuta}
	tgDataMap[tgIndex{7, 6, 1}] = tgData{"Ag", Imbalances}
	tgDataMap[tgIndex{8, 6, 1}] = tgData{"Disposables", Consumables}
	tgDataMap[tgIndex{9, 6, 1}] = tgData{"In", Imbalances}
	tgDataMap[tgIndex{10, 6, 1}] = tgData{"Edutainments", Entertainments}
	tgDataMap[tgIndex{11, 6, 1}] = tgData{"Obsoletes", ScrapWaste}
	tgDataMap[tgIndex{12, 6, 1}] = tgData{"Regulations", RedTape}
	tgDataMap[tgIndex{7, 6, 2}] = tgData{"Ri", Imbalances}
	tgDataMap[tgIndex{8, 6, 2}] = tgData{"Respirators", Consumables}
	tgDataMap[tgIndex{9, 6, 2}] = tgData{"Ri", Imbalances}
	tgDataMap[tgIndex{10, 6, 2}] = tgData{"Recordings", Entertainments}
	tgDataMap[tgIndex{11, 6, 2}] = tgData{"Used Goods", ScrapWaste}
	tgDataMap[tgIndex{12, 6, 2}] = tgData{"Synchronizations", RedTape}
	tgDataMap[tgIndex{7, 6, 3}] = tgData{"In", Imbalances}
	tgDataMap[tgIndex{8, 6, 3}] = tgData{"Filter Masks", Consumables}
	tgDataMap[tgIndex{9, 6, 3}] = tgData{"Fl", Imbalances}
	tgDataMap[tgIndex{10, 6, 3}] = tgData{"Writings", Entertainments}
	tgDataMap[tgIndex{11, 6, 3}] = tgData{"Reparables", ScrapWaste}
	tgDataMap[tgIndex{12, 6, 3}] = tgData{"Expert Systems", RedTape}
	tgDataMap[tgIndex{7, 6, 4}] = tgData{"Ic", Imbalances}
	tgDataMap[tgIndex{8, 6, 4}] = tgData{"Combination", Consumables}
	tgDataMap[tgIndex{9, 6, 4}] = tgData{"Ic", Imbalances}
	tgDataMap[tgIndex{10, 6, 4}] = tgData{"Tactiles", Entertainments}
	tgDataMap[tgIndex{11, 6, 4}] = tgData{"Plutonium", ScrapWaste}
	tgDataMap[tgIndex{12, 6, 4}] = tgData{"Educationals", RedTape}
	tgDataMap[tgIndex{7, 6, 5}] = tgData{"De", Imbalances}
	tgDataMap[tgIndex{8, 6, 5}] = tgData{"Parts", Consumables}
	tgDataMap[tgIndex{9, 6, 5}] = tgData{"Ag", Imbalances}
	tgDataMap[tgIndex{10, 6, 5}] = tgData{"Osmancies", Entertainments}
	tgDataMap[tgIndex{11, 6, 5}] = tgData{"Metals", ScrapWaste}
	tgDataMap[tgIndex{12, 6, 5}] = tgData{"Mandates", RedTape}
	tgDataMap[tgIndex{7, 6, 6}] = tgData{"Fl", Imbalances}
	tgDataMap[tgIndex{8, 6, 6}] = tgData{"Improvements", Consumables}
	tgDataMap[tgIndex{9, 6, 6}] = tgData{"Va", Imbalances}
	tgDataMap[tgIndex{10, 6, 6}] = tgData{"Wafers", Entertainments}
	tgDataMap[tgIndex{11, 6, 6}] = tgData{"Sludges", ScrapWaste}
	tgDataMap[tgIndex{12, 6, 6}] = tgData{"Accountings", RedTape}
	return tgDataMap[tgi]
}
