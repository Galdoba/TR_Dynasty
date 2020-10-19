package trade

//CODE||TYPE            ||EXAMPLE                    ||PRODUCTION TAGS	||TON INCREMENT	||BASE PRICE||PURCHSE DM   							||SALE DM                                ||RISK||DANGER
//112||Basic Electronics||Calculators/Adding Machines||All				||1d6 x 12		||6000		||In +2, Ht +3, Ri +1	||Ni +2, Lt +1, Po +1||+0	||-6

import (
	"strconv"
	"strings"
)

const (
	//tradeTagAll             = "All"
	tradeTagAg             = "Ag"
	tradeTagAs             = "As"
	tradeTagBa             = "Ba"
	tradeTagDe             = "De"
	tradeTagFluidOceans    = "Fl"
	tradeTagGa             = "Ga"
	tradeTagHighPopulation = "Hi"
	tradeTagHighTechnology = "Ht"
	tradeTagIceCapped      = "Ic"
	tradeTagIn             = "In"
	tradeTagLowPopulation  = "Lo"
	tradeTagLowTechnology  = "Lt"
	tradeTagNonAg          = "Na"
	tradeTagNonIn          = "Ni"
	tradeTagPo             = "Po"
	tradeTagRi             = "Ri"
	tradeTagVa             = "Va"
	tradeTagWaterWorld     = "Wa"
	travelTagAmber         = "AZ"
	travelTagRed           = "RZ"
)

func tradeTagLIST() []string {
	return []string{
		//tradeTagAll,
		tradeTagAg,
		tradeTagAs,
		tradeTagBa,
		tradeTagDe,
		tradeTagFluidOceans,
		tradeTagGa,
		tradeTagHighPopulation,
		tradeTagHighTechnology,
		tradeTagIceCapped,
		tradeTagIn,
		tradeTagLowPopulation,
		tradeTagLowTechnology,
		tradeTagNonAg,
		tradeTagNonIn,
		tradeTagPo,
		tradeTagRi,
		tradeTagVa,
		tradeTagWaterWorld,
	}
}

func tradeCodeValid(tCode string) bool {
	switch tCode {
	default:
		return false
	case "Ag":
		return true
	case "As":
		return true
	case "Ba":
		return true
	case "De":
		return true
	case "Fl":
		return true
	case "Ga":
		return true
	case "Hi":
		return true
	case "Ht":
		return true
	case "Ic":
		return true
	case "In":
		return true
	case "Lo":
		return true
	case "Lt":
		return true
	case "Na":
		return true
	case "Ni":
		return true
	case "Po":
		return true
	case "Ri":
		return true
	case "Va":
		return true
	case "Wa":
		return true
	case "AZ":
		return true
	case "RZ":
		return true
	}
}

func tradeCodeFullName(tCode string) string {
	switch tCode {
	default:
		return "Error"
	case "Ag":
		return "Ag"
	case "As":
		return "As"
	case "Ba":
		return "Ba"
	case "De":
		return "De"
	case "Fl":
		return "Fl"
	case "Ga":
		return "Ga"
	case "Hi":
		return "Hi"
	case "Ht":
		return "Ht"
	case "Ic":
		return "Ic"
	case "In":
		return "In"
	case "Lo":
		return "Lo"
	case "Lt":
		return "Lt"
	case "Na":
		return "Na"
	case "Ni":
		return "Ni"
	case "Po":
		return "Po"
	case "Ri":
		return "Ri"
	case "Va":
		return "Va"
	case "Wa":
		return "Wa"
	case "AZ":
		return "AZ"
	case "RZ":
		return "RZ"
	}
}

func inListCustom(position string, list []string) bool {
	for i := range list {
		if position == list[i] {
			return true
		}
	}
	return false
}

//TGoodsData -
func TGoodsData() map[string][]string {
	tgrMap := make(map[string][]string)
	//fill map
	tgrMap["112"] = []string{"Basic Electronics", "Calculators/Adding Machines", "All", "1d6 x 12", "6000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["113"] = []string{"Basic Electronics", "Video Game and Entertainment Systems", "All", "1d6 x 10", "8000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["114"] = []string{"Basic Electronics", "Video Game and Entertainment Systems", "All", "1d6 x 10", "8000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["115"] = []string{"Basic Electronics", "Video Game and Entertainment Systems", "All", "1d6 x 10", "8000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["116"] = []string{"Basic Electronics", "Personal and Commercial Computers", "All", "1d6 x 10", "10000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["117"] = []string{"Basic Electronics", "Personal and Commercial Computers", "All", "1d6 x 10", "10000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["118"] = []string{"Basic Electronics", "Personal and Commercial Computers", "All", "1d6 x 10", "10000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["119"] = []string{"Basic Electronics", "Banking Machines and Security Systems", "All", "1d6 x 4", "12000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["1110"] = []string{"Basic Electronics", "Banking Machines and Security Systems", "All", "1d6 x 4", "12000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["1111"] = []string{"Basic Electronics", "Banking Machines and Security Systems", "All", "1d6 x 4", "12000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}
	tgrMap["1112"] = []string{"Basic Electronics", "Microprocessor Assemblies", "All", "1d6 x 2", "14000", "In +2, Ht +3, Ri +1", "Ni +2, Lt +1, Po +1", "+0", "-6", "1d6 x 10"}

	tgrMap["122"] = []string{"Basic Machine Parts", "Stamped/Poured Cogs and Sprockets", "All", "1d6 x 12", "8000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["123"] = []string{"Basic Machine Parts", "Piping and Attachment Pieces", "All", "1d6 x 10", "9000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["124"] = []string{"Basic Machine Parts", "Piping and Attachment Pieces", "All", "1d6 x 10", "9000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["125"] = []string{"Basic Machine Parts", "Piping and Attachment Pieces", "All", "1d6 x 10", "9000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["126"] = []string{"Basic Machine Parts", "Engine Components", "All", "1d6 x 10", "10000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["127"] = []string{"Basic Machine Parts", "Engine Components", "All", "1d6 x 10", "10000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["128"] = []string{"Basic Machine Parts", "Engine Components", "All", "1d6 x 10", "10000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["129"] = []string{"Basic Machine Parts", "Pneumatics and Hydraulics", "All", "1d6 x 6", "11000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["1210"] = []string{"Basic Machine Parts", "Pneumatics and Hydraulics", "All", "1d6 x 6", "11000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["1211"] = []string{"Basic Machine Parts", "Pneumatics and Hydraulics", "All", "1d6 x 6", "11000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}
	tgrMap["1212"] = []string{"Basic Machine Parts", "Starship-Quality Components", "All", "1d6 x 4", "12000", "Na +2, In +5", "Ni +3, Ag +2", "+0", "-6", "1d6 x 10"}

	tgrMap["132"] = []string{"Basic Manufactured Goods", "Second Stage Components", "All", "1d6 x 12", "8000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["133"] = []string{"Basic Manufactured Goods", "Uniforms/Clothing Products", "All", "1d6 x 10", "9000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["134"] = []string{"Basic Manufactured Goods", "Uniforms/Clothing Products", "All", "1d6 x 10", "9000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["135"] = []string{"Basic Manufactured Goods", "Uniforms/Clothing Products", "All", "1d6 x 10", "9000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["136"] = []string{"Basic Manufactured Goods", "Residential Appliances", "All", "1d6 x 10", "10000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["137"] = []string{"Basic Manufactured Goods", "Residential Appliances", "All", "1d6 x 10", "10000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["138"] = []string{"Basic Manufactured Goods", "Residential Appliances", "All", "1d6 x 10", "10000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["139"] = []string{"Basic Manufactured Goods", "Furniture/Storage Systems/Tools", "All", "1d6 x 5", "11000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["1310"] = []string{"Basic Manufactured Goods", "Furniture/Storage Systems/Tools", "All", "1d6 x 5", "11000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["1311"] = []string{"Basic Manufactured Goods", "Furniture/Storage Systems/Tools", "All", "1d6 x 5", "11000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}
	tgrMap["1312"] = []string{"Basic Manufactured Goods", "Vehicle/Survival Accessories", "All", "1d6 x 3", "12000", "Na +2, In +5", "Ni +3, Hi +2", "+0", "-6", "1d6 x 10"}

	tgrMap["142"] = []string{"Basic Raw Materials", "Foundation Stones and Base Elements", "All", "1d6 x 14", "1000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["143"] = []string{"Basic Raw Materials", "Workable Metals", "All", "1d6 x 12", "3000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["144"] = []string{"Basic Raw Materials", "Workable Metals", "All", "1d6 x 12", "3000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["145"] = []string{"Basic Raw Materials", "Workable Metals", "All", "1d6 x 12", "3000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["146"] = []string{"Basic Raw Materials", "Workable Alloys", "All", "1d6 x 10", "5000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["147"] = []string{"Basic Raw Materials", "Workable Alloys", "All", "1d6 x 10", "5000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["148"] = []string{"Basic Raw Materials", "Workable Alloys", "All", "1d6 x 10", "5000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["149"] = []string{"Basic Raw Materials", "Fabricated Plastics", "All", "1d6 x 5", "7000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["1410"] = []string{"Basic Raw Materials", "Fabricated Plastics", "All", "1d6 x 5", "7000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["1411"] = []string{"Basic Raw Materials", "Fabricated Plastics", "All", "1d6 x 5", "7000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}
	tgrMap["1412"] = []string{"Basic Raw Materials", "Chemical Solutions or Compounds", "All", "1d6 x 3", "9000", "Ag +3, Ga +2", "In +2, Po +2", "+0", "-6", "1d6 x 10"}

	tgrMap["152"] = []string{"Basic Consumables", "Feed-grade Vegetation", "All", "1d6 x 12", "500", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["153"] = []string{"Basic Consumables", "Food-grade Vegetation", "All", "1d6 x 10", "1000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["154"] = []string{"Basic Consumables", "Food-grade Vegetation", "All", "1d6 x 10", "1000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["155"] = []string{"Basic Consumables", "Food-grade Vegetation", "All", "1d6 x 10", "1000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["156"] = []string{"Basic Consumables", "Pre-packaged Food and Drink", "All", "1d6 x 10", "2000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["157"] = []string{"Basic Consumables", "Pre-packaged Food and Drink", "All", "1d6 x 10", "2000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["158"] = []string{"Basic Consumables", "Pre-packaged Food and Drink", "All", "1d6 x 10", "2000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["159"] = []string{"Basic Consumables", "Survival Rations and Storage-packed Liquids", "All", "1d6 x 8", "3000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["1510"] = []string{"Basic Consumables", "Survival Rations and Storage-packed Liquids", "All", "1d6 x 8", "3000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["1511"] = []string{"Basic Consumables", "Survival Rations and Storage-packed Liquids", "All", "1d6 x 8", "3000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}
	tgrMap["1512"] = []string{"Basic Consumables", "Junk Food/Soda/Beer", "All", "1d6 x 4", "5000", "Ag +3, Wa +2, Ga +1, As -4", "As +1, Fl +1, Ic +1, Hi +1", "+0", "-6", "1d6 x 10"}

	tgrMap["162"] = []string{"Basic Ore", "Bornite or Galena or Sedimentary Stone", "All", "1d6 x 14", "250", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["163"] = []string{"Basic Ore", "Chalcocite or Talc", "All", "1d6 x 12", "500", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["164"] = []string{"Basic Ore", "Chalcocite or Talc", "All", "1d6 x 12", "500", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["165"] = []string{"Basic Ore", "Chalcocite or Talc", "All", "1d6 x 12", "500", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["166"] = []string{"Basic Ore", "Bauxite, Coltan and Wolframite", "All", "1d6 x 10", "1000", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["167"] = []string{"Basic Ore", "Bauxite, Coltan and Wolframite", "All", "1d6 x 10", "1000", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["168"] = []string{"Basic Ore", "Bauxite, Coltan and Wolframite", "All", "1d6 x 10", "1000", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["169"] = []string{"Basic Ore", "Acanthite, Cobaltite or Magnetite", "All", "1d6 x 8", "1500", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["1610"] = []string{"Basic Ore", "Acanthite, Cobaltite or Magnetite", "All", "1d6 x 8", "1500", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["1611"] = []string{"Basic Ore", "Acanthite, Cobaltite or Magnetite", "All", "1d6 x 8", "1500", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}
	tgrMap["1612"] = []string{"Basic Ore", "Chromite or Cinnabar", "All", "1d6 x 4", "2000", "As +4, Ic +0", "In +3, Ni +1", "+0", "-6", "1d6 x 10"}

	tgrMap["212"] = []string{"Advanced Electronics", "Circuitry Bundles", "In, Ht", "1d6 x 6", "25000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["213"] = []string{"Advanced Electronics", "Fibre-optic Components", "In, Ht", "1d6 x 5", "50000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["214"] = []string{"Advanced Electronics", "Fibre-optic Components", "In, Ht", "1d6 x 5", "50000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["215"] = []string{"Advanced Electronics", "Fibre-optic Components", "In, Ht", "1d6 x 5", "50000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["216"] = []string{"Advanced Electronics", "VR Computer and Sensor Packages", "In, Ht", "1d6 x 5", "100000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["217"] = []string{"Advanced Electronics", "VR Computer and Sensor Packages", "In, Ht", "1d6 x 5", "100000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["218"] = []string{"Advanced Electronics", "VR Computer and Sensor Packages", "In, Ht", "1d6 x 5", "100000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["219"] = []string{"Advanced Electronics", "Weapon Components", "In, Ht", "1d6 x 2", "125000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["2110"] = []string{"Advanced Electronics", "Weapon Components", "In, Ht", "1d6 x 2", "125000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["2111"] = []string{"Advanced Electronics", "Weapon Components", "In, Ht", "1d6 x 2", "125000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}
	tgrMap["2112"] = []string{"Advanced Electronics", "Starship Bridge Components", "In, Ht", "1d6 x 1", "150000", "In +2, Ht +3", "Ni +1, Ri +2, As +3", "+2", "-2", "1d6 x 5"}

	tgrMap["222"] = []string{"Advanced Machine Parts", "Alloy and Plastic Tool Kits", "In, Ht", "1d6 x 6", "25000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["223"] = []string{"Advanced Machine Parts", "Starship Deckplate/Atmospheric Filters", "In, Ht", "1d6 x 5", "50000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["224"] = []string{"Advanced Machine Parts", "Starship Deckplate/Atmospheric Filters", "In, Ht", "1d6 x 5", "50000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["225"] = []string{"Advanced Machine Parts", "Starship Deckplate/Atmospheric Filters", "In, Ht", "1d6 x 5", "50000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["226"] = []string{"Advanced Machine Parts", "Fusion Conduits/Power Plant Shells", "In, Ht", "1d6 x 5", "75000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["227"] = []string{"Advanced Machine Parts", "Fusion Conduits/Power Plant Shells", "In, Ht", "1d6 x 5", "75000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["228"] = []string{"Advanced Machine Parts", "Fusion Conduits/Power Plant Shells", "In, Ht", "1d6 x 5", "75000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["229"] = []string{"Advanced Machine Parts", "Weapon Cores/Starship Hull", "In, Ht", "1d6 x 3", "90000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["2210"] = []string{"Advanced Machine Parts", "Weapon Cores/Starship Hull", "In, Ht", "1d6 x 3", "90000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["2211"] = []string{"Advanced Machine Parts", "Weapon Cores/Starship Hull", "In, Ht", "1d6 x 3", "90000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}
	tgrMap["2212"] = []string{"Advanced Machine Parts", "Gravitic Gyros, Navigation Magnetics", "In, Ht", "1d6 x 1", "100000", "In +2, Ht +1", "As +2, Ni +1", "+2", "-2", "1d6 x 5"}

	tgrMap["232"] = []string{"Advanced Manufactured Goods", "High-Pressure or Temperature-Resistant Components", "In, Ht", "1d6 x 6", "25000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["233"] = []string{"Advanced Manufactured Goods", "Protective or Specialised Clothing", "In, Ht", "1d6 x 5", "50000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["234"] = []string{"Advanced Manufactured Goods", "Protective or Specialised Clothing", "In, Ht", "1d6 x 5", "50000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["235"] = []string{"Advanced Manufactured Goods", "Protective or Specialised Clothing", "In, Ht", "1d6 x 5", "50000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["236"] = []string{"Advanced Manufactured Goods", "Survival Equipment/Colonisation Kits", "In, Ht", "1d6 x 5", "100000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["237"] = []string{"Advanced Manufactured Goods", "Survival Equipment/Colonisation Kits", "In, Ht", "1d6 x 5", "100000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["238"] = []string{"Advanced Manufactured Goods", "Survival Equipment/Colonisation Kits", "In, Ht", "1d6 x 5", "100000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["239"] = []string{"Advanced Manufactured Goods", "Computerised Job-related Gear", "In, Ht", "1d6 x 2", "125000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["2310"] = []string{"Advanced Manufactured Goods", "Computerised Job-related Gear", "In, Ht", "1d6 x 2", "125000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["2311"] = []string{"Advanced Manufactured Goods", "Computerised Job-related Gear", "In, Ht", "1d6 x 2", "125000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}
	tgrMap["2312"] = []string{"Advanced Manufactured Goods", "Starship Add-Ons/Powered Armour Components", "In, Ht", "1d6 x 1", "150000", "In +1, Ht +0", "Hi +1, Ri +2", "+2", "-2", "1d6 x 5"}

	tgrMap["242"] = []string{"Advanced Weapons", "(TL7 or less) Slug Weapons", "In, Ht", "1d6 x 7", "50000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["243"] = []string{"Advanced Weapons", "(TL10 or less) Slug Weapons", "In, Ht", "1d6 x 6", "100000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["244"] = []string{"Advanced Weapons", "(TL10 or less) Slug Weapons", "In, Ht", "1d6 x 6", "100000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["245"] = []string{"Advanced Weapons", "(TL10 or less) Slug Weapons", "In, Ht", "1d6 x 6", "100000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["246"] = []string{"Advanced Weapons", "(TL12 or less) Slug or Energy Weapons/Heavy Slug Weapons", "In, Ht", "1d6 x 5", "150000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["247"] = []string{"Advanced Weapons", "(TL12 or less) Slug or Energy Weapons/Heavy Slug Weapons", "In, Ht", "1d6 x 5", "150000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["248"] = []string{"Advanced Weapons", "(TL12 or less) Slug or Energy Weapons/Heavy Slug Weapons", "In, Ht", "1d6 x 5", "150000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["249"] = []string{"Advanced Weapons", "(TL15 or less) Slug or Energy Weapons/Explosives", "In, Ht", "1d6 x 3", "200000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["2410"] = []string{"Advanced Weapons", "(TL15 or less) Slug or Energy Weapons/Explosives", "In, Ht", "1d6 x 3", "200000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["2411"] = []string{"Advanced Weapons", "(TL15 or less) Slug or Energy Weapons/Explosives", "In, Ht", "1d6 x 3", "200000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}
	tgrMap["2412"] = []string{"Advanced Weapons", "Artillery, Heavy Energy Weapons", "In, Ht", "1d6 x 1", "250000", "In +0, Ht +2", "Po +1, AZ +2, RZ +4", "+3", "+0", "1d6 x 5"}

	tgrMap["252"] = []string{"Advanced Vehicles", "Engine Components or Packages", "In, Ht", "1d6 x 5", "100000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["253"] = []string{"Advanced Vehicles", "Seafaring or Mole Vehicle Components or Packages", "In, Ht", "1d6 x 5", "140000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["254"] = []string{"Advanced Vehicles", "Seafaring or Mole Vehicle Components or Packages", "In, Ht", "1d6 x 5", "140000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["255"] = []string{"Advanced Vehicles", "Seafaring or Mole Vehicle Components or Packages", "In, Ht", "1d6 x 5", "140000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["256"] = []string{"Advanced Vehicles", "Air Raft Components or Packages", "In, Ht", "1d6 x 5", "180000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["257"] = []string{"Advanced Vehicles", "Air Raft Components or Packages", "In, Ht", "1d6 x 5", "180000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["258"] = []string{"Advanced Vehicles", "Air Raft Components or Packages", "In, Ht", "1d6 x 5", "180000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["259"] = []string{"Advanced Vehicles", "Grav-Vehicle Components", "In, Ht", "1d6 x 2", "200000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["2510"] = []string{"Advanced Vehicles", "Grav-Vehicle Components", "In, Ht", "1d6 x 2", "200000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["2511"] = []string{"Advanced Vehicles", "Grav-Vehicle Components", "In, Ht", "1d6 x 2", "200000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}
	tgrMap["2512"] = []string{"Advanced Vehicles", "Spacecraft Components", "In, Ht", "1d6 x 1", "250000", "In +0, Ht +2", "As +2, Ri +2", "+3", "+0", "1d6 x 5"}

	tgrMap["262"] = []string{"Biochemicals", "Organic Glues, Acids or Bases/Vegetable Oil", "Ag, Wa", "1d6 x 6", "10000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["263"] = []string{"Biochemicals", "Ethanol/Fructose Syrup", "Ag, Wa", "1d6 x 5", "25000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["264"] = []string{"Biochemicals", "Ethanol/Fructose Syrup", "Ag, Wa", "1d6 x 5", "25000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["265"] = []string{"Biochemicals", "Ethanol/Fructose Syrup", "Ag, Wa", "1d6 x 5", "25000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["266"] = []string{"Biochemicals", "Biodiesel/Cooking Compounds", "Ag, Wa", "1d6 x 5", "50000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["267"] = []string{"Biochemicals", "Biodiesel/Cooking Compounds", "Ag, Wa", "1d6 x 5", "50000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["268"] = []string{"Biochemicals", "Biodiesel/Cooking Compounds", "Ag, Wa", "1d6 x 5", "50000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["269"] = []string{"Biochemicals", "Oxygenated Cleaner/Biodegradable Concentrates", "Ag, Wa", "1d6 x 3", "60000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["2610"] = []string{"Biochemicals", "Oxygenated Cleaner/Biodegradable Concentrates", "Ag, Wa", "1d6 x 3", "60000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["2611"] = []string{"Biochemicals", "Oxygenated Cleaner/Biodegradable Concentrates", "Ag, Wa", "1d6 x 3", "60000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}
	tgrMap["2612"] = []string{"Biochemicals", "Gelid Oxygen-Substitutes/Bio-fusion Cell Fuel", "Ag, Wa", "1d6 x 1", "80000", "Ag +1, Wa +2", "In +2", "+2", "+2", "1d6 x 5"}

	tgrMap["312"] = []string{"Crystals & Gem", "Rock Salt/Compressed Coal", "As, De, Ic", "1d6 x 7", "5000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["313"] = []string{"Crystals & Gem", "Graphite/Quartz", "As, De, Ic", "1d6 x 6", "10000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["314"] = []string{"Crystals & Gem", "Graphite/Quartz", "As, De, Ic", "1d6 x 6", "10000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["315"] = []string{"Crystals & Gem", "Graphite/Quartz", "As, De, Ic", "1d6 x 6", "10000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["316"] = []string{"Crystals & Gem", "Silica/Focuser-Quality Gems", "As, De, Ic", "1d6 x 5", "20000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["317"] = []string{"Crystals & Gem", "Silica/Focuser-Quality Gems", "As, De, Ic", "1d6 x 5", "20000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["318"] = []string{"Crystals & Gem", "Silica/Focuser-Quality Gems", "As, De, Ic", "1d6 x 5", "20000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["319"] = []string{"Crystals & Gem", "Photonics/Synthetic Gems", "As, De, Ic", "1d6 x 3", "30000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["3110"] = []string{"Crystals & Gem", "Photonics/Synthetic Gems", "As, De, Ic", "1d6 x 3", "30000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["3111"] = []string{"Crystals & Gem", "Photonics/Synthetic Gems", "As, De, Ic", "1d6 x 3", "30000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}
	tgrMap["3112"] = []string{"Crystals & Gem", "In Diamond/Jewellery-Quality Gems", "As, De, Ic", "1d6 x 2", "45000", "As +2, De +1, Ic +1", "In +3, Ri +2", "+2", "-1", "1d6 x 5"}

	tgrMap["322"] = []string{"Cybernetics", "Cybernetic Lubricants", "Ht", "1d6 x 1 + 2", "100000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["323"] = []string{"Cybernetics", "Cybernetic Components/Physical Augments", "Ht", "1d6 x 1 + 1", "200000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["324"] = []string{"Cybernetics", "Cybernetic Components/Physical Augments", "Ht", "1d6 x 1 + 1", "200000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["325"] = []string{"Cybernetics", "Cybernetic Components/Physical Augments", "Ht", "1d6 x 1 + 1", "200000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["326"] = []string{"Cybernetics", "Cyber-Prosthetics", "Ht", "1d6 x 1", "250000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["327"] = []string{"Cybernetics", "Cyber-Prosthetics", "Ht", "1d6 x 1", "250000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["328"] = []string{"Cybernetics", "Cyber-Prosthetics", "Ht", "1d6 x 1", "250000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["329"] = []string{"Cybernetics", "Cosmetic Prosthetics", "Ht", "1d6 x 1", "350000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["3210"] = []string{"Cybernetics", "Cosmetic Prosthetics", "Ht", "1d6 x 1", "350000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["3211"] = []string{"Cybernetics", "Cosmetic Prosthetics", "Ht", "1d6 x 1", "350000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}
	tgrMap["3212"] = []string{"Cybernetics", "Real-Life Replacements and Augments", "Ht", "0d6 x 1 + 1", "500000", "Ht +0", "As +1, Ic +1, Ri +2", "+3", "+1", "1d6 x 1"}

	tgrMap["332"] = []string{"Live Animals", "Beasts of Burden", "Ga", "1d6 x 12", "2500", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["333"] = []string{"Live Animals", "Untrained Riding Animals", "Ga", "1d6 x 10", "5000", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["334"] = []string{"Live Animals", "Untrained Riding Animals", "Ga", "1d6 x 10", "5000", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["335"] = []string{"Live Animals", "Untrained Riding Animals", "Ga", "1d6 x 10", "5000", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["336"] = []string{"Live Animals", "Trained Riding Animals/Common Pets", "Ga", "1d6 x 10", "10000", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["337"] = []string{"Live Animals", "Trained Riding Animals/Common Pets", "Ga", "1d6 x 10", "10000", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["338"] = []string{"Live Animals", "Trained Riding Animals/Common Pets", "Ga", "1d6 x 10", "10000", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["339"] = []string{"Live Animals", "Untrained Guard Animals", "Ga", "1d6 x 6", "12500", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["3310"] = []string{"Live Animals", "Untrained Guard Animals", "Ga", "1d6 x 6", "12500", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["3311"] = []string{"Live Animals", "Untrained Guard Animals", "Ga", "1d6 x 6", "12500", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}
	tgrMap["3312"] = []string{"Live Animals", "Trained Guard Animals/Exotic Pets", "Ga", "1d6 x 3", "15000", "Ag +2, Ga +0", "Lo +3", "+2", "+2", "1d6 x 10"}

	tgrMap["342"] = []string{"Luxury Consumables", "Common Desserts/Rare Food Additives", "Ag, Ga, Wa", "1d6 x 14", "5000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["343"] = []string{"Luxury Consumables", "Common Desserts/Common Wine", "Ag, Ga, Wa", "1d6 x 12", "10000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["344"] = []string{"Luxury Consumables", "Common Desserts/Common Wine", "Ag, Ga, Wa", "1d6 x 12", "10000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["345"] = []string{"Luxury Consumables", "Common Desserts/Common Wine", "Ag, Ga, Wa", "1d6 x 12", "10000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["346"] = []string{"Luxury Consumables", "Rare Foods/Common Liquor", "Ag, Ga, Wa", "1d6 x 10", "20000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["347"] = []string{"Luxury Consumables", "Rare Foods/Common Liquor", "Ag, Ga, Wa", "1d6 x 10", "20000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["348"] = []string{"Luxury Consumables", "Rare Foods/Common Liquor", "Ag, Ga, Wa", "1d6 x 10", "20000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["349"] = []string{"Luxury Consumables", "Exotic Foods/Rare Desserts/Rare Liquor", "Ag, Ga, Wa", "1d6 x 5", "30000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["3410"] = []string{"Luxury Consumables", "Exotic Foods/Rare Desserts/Rare Liquor", "Ag, Ga, Wa", "1d6 x 5", "30000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["3411"] = []string{"Luxury Consumables", "Exotic Foods/Rare Desserts/Rare Liquor", "Ag, Ga, Wa", "1d6 x 5", "30000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}
	tgrMap["3412"] = []string{"Luxury Consumables", "Exotic Desserts/Exotic Liquor", "Ag, Ga, Wa", "1d6 x 2", "50000", "Ag +2, Ga +0, Wa +1", "Ri +2, Hi +2", "+3", "+2", "1d6 x 10"}

	tgrMap["352"] = []string{"Luxury Goods", "Rare Literature/Art", "Hi", "1d6 x 1 + 2", "50000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["353"] = []string{"Luxury Goods", "Jewellery/Alien Textiles", "Hi", "1d6 x 1 + 1", "100000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["354"] = []string{"Luxury Goods", "Jewellery/Alien Textiles", "Hi", "1d6 x 1 + 1", "100000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["355"] = []string{"Luxury Goods", "Jewellery/Alien Textiles", "Hi", "1d6 x 1 + 1", "100000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["356"] = []string{"Luxury Goods", "Rare Clothing/Home Decorations", "Hi", "1d6 x 1", "200000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["357"] = []string{"Luxury Goods", "Rare Clothing/Home Decorations", "Hi", "1d6 x 1", "200000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["358"] = []string{"Luxury Goods", "Rare Clothing/Home Decorations", "Hi", "1d6 x 1", "200000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["359"] = []string{"Luxury Goods", "VR Electronic Entertainment Devices", "Hi", "1d6 x 1", "250000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["3510"] = []string{"Luxury Goods", "VR Electronic Entertainment Devices", "Hi", "1d6 x 1", "250000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["3511"] = []string{"Luxury Goods", "VR Electronic Entertainment Devices", "Hi", "1d6 x 1", "250000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}
	tgrMap["3512"] = []string{"Luxury Goods", "Exotic Furnishings/Exquisite Jewellery", "Hi", "1d6 x 0 + 1", "500000", "Hi +0", "Ri +4", "+3", "+2", "1d6 x 1"}

	tgrMap["362"] = []string{"Medical Supplies", "Medical Uniforms/Disposable Tools", "Ht, Hi", "1d6 x 6", "10000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["363"] = []string{"Medical Supplies", "Cosmetic Chemicals/Practitioner's Tools", "Ht, Hi", "1d6 x 5", "30000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["364"] = []string{"Medical Supplies", "Cosmetic Chemicals/Practitioner's Tools", "Ht, Hi", "1d6 x 5", "30000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["365"] = []string{"Medical Supplies", "Cosmetic Chemicals/Practitioner's Tools", "Ht, Hi", "1d6 x 5", "30000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["366"] = []string{"Medical Supplies", "General Medical Equipment or Supplies", "Ht, Hi", "1d6 x 5", "50000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["367"] = []string{"Medical Supplies", "General Medical Equipment or Supplies", "Ht, Hi", "1d6 x 5", "50000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["368"] = []string{"Medical Supplies", "General Medical Equipment or Supplies", "Ht, Hi", "1d6 x 5", "50000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["369"] = []string{"Medical Supplies", "Specialist Equipment or Supplies", "Ht, Hi", "1d6 x 2", "75000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["3610"] = []string{"Medical Supplies", "Specialist Equipment or Supplies", "Ht, Hi", "1d6 x 2", "75000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["3611"] = []string{"Medical Supplies", "Specialist Equipment or Supplies", "Ht, Hi", "1d6 x 2", "75000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}
	tgrMap["3612"] = []string{"Medical Supplies", "Micro-surgical Equipment or Supplies", "Ht, Hi", "1d6 x 1", "100000", "Ht +2, Hi +0", "In +2, Po +1, Ri +1", "+2", "+2", "1d6 x 5"}

	tgrMap["412"] = []string{"Petrochemicals", "Crude Oil/Diesel", "De, Fl, Ic, Wa", "1d6 x 12", "2500", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["413"] = []string{"Petrochemicals", "Refined Kerosene/Purified Oil", "De, Fl, Ic, Wa", "1d6 x 10", "5000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["414"] = []string{"Petrochemicals", "Refined Kerosene/Purified Oil", "De, Fl, Ic, Wa", "1d6 x 10", "5000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["415"] = []string{"Petrochemicals", "Refined Kerosene/Purified Oil", "De, Fl, Ic, Wa", "1d6 x 10", "5000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["416"] = []string{"Petrochemicals", "Gasoline/Machine Lubricants", "De, Fl, Ic, Wa", "1d6 x 10", "10000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["417"] = []string{"Petrochemicals", "Gasoline/Machine Lubricants", "De, Fl, Ic, Wa", "1d6 x 10", "10000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["418"] = []string{"Petrochemicals", "Gasoline/Machine Lubricants", "De, Fl, Ic, Wa", "1d6 x 10", "10000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["419"] = []string{"Petrochemicals", "Jet Fuel/Gelid Adhesives", "De, Fl, Ic, Wa", "1d6 x 8", "20000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["4110"] = []string{"Petrochemicals", "Jet Fuel/Gelid Adhesives", "De, Fl, Ic, Wa", "1d6 x 8", "20000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["4111"] = []string{"Petrochemicals", "Jet Fuel/Gelid Adhesives", "De, Fl, Ic, Wa", "1d6 x 8", "20000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}
	tgrMap["4112"] = []string{"Petrochemicals", "Rocket Fuel/Power Plant Starter Charges", "De, Fl, Ic, Wa", "1d6 x 4", "30000", "De +2, Fl +0, Ic +0, Wa +0", "In +2, Ag +1, Lt +2", "+2", "+2", "1d6 x 10"}

	tgrMap["422"] = []string{"Pharmaceuticals", "OTC Drugs/Antibiotics", "As, De, Hi, Wa", "1d6 x 1 + 3", "25000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["423"] = []string{"Pharmaceuticals", "Antivenin/Prescription Medications", "As, De, Hi, Wa", "1d6 x 1 + 2", "50000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["424"] = []string{"Pharmaceuticals", "Antivenin/Prescription Medications", "As, De, Hi, Wa", "1d6 x 1 + 2", "50000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["425"] = []string{"Pharmaceuticals", "Antivenin/Prescription Medications", "As, De, Hi, Wa", "1d6 x 1 + 2", "50000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["426"] = []string{"Pharmaceuticals", "Prescription Medications/Surgical", "As, De, Hi, Wa", "1d6 x 1", "100000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["427"] = []string{"Pharmaceuticals", "Prescription Medications/Surgical", "As, De, Hi, Wa", "1d6 x 1", "100000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["428"] = []string{"Pharmaceuticals", "Prescription Medications/Surgical", "As, De, Hi, Wa", "1d6 x 1", "100000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["429"] = []string{"Pharmaceuticals", "Anagathics", "As, De, Hi, Wa", "1d6 x 0 + 2", "200000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["4210"] = []string{"Pharmaceuticals", "Anagathics", "As, De, Hi, Wa", "1d6 x 0 + 2", "200000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["4211"] = []string{"Pharmaceuticals", "Anagathics", "As, De, Hi, Wa", "1d6 x 0 + 2", "200000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}
	tgrMap["4212"] = []string{"Pharmaceuticals", "Psi-Related Drugs/Viral Therapy Doses", "As, De, Hi, Wa", "1d6 x 0 + 1", "500000", "As +2, De +0, Hi +1, Wa +0", "Ri +2, Lt +1", "+2", "+3", "1d6 x 1"}

	tgrMap["432"] = []string{"Polymers", "Rubber/Vinyl Spooling", "In", "1d6 x 12", "1000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["433"] = []string{"Polymers", "Insulation/Polyurethane Foam", "In", "1d6 x 10", "3000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["434"] = []string{"Polymers", "Insulation/Polyurethane Foam", "In", "1d6 x 10", "3000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["435"] = []string{"Polymers", "Insulation/Polyurethane Foam", "In", "1d6 x 10", "3000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["436"] = []string{"Polymers", "Poured Plastics/Synthetic Fibre Spools", "In", "1d6 x 10", "7000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["437"] = []string{"Polymers", "Poured Plastics/Synthetic Fibre Spools", "In", "1d6 x 10", "7000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["438"] = []string{"Polymers", "Poured Plastics/Synthetic Fibre Spools", "In", "1d6 x 10", "7000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["439"] = []string{"Polymers", "Kevlar/Teflon", "In", "1d6 x 3", "9000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["4310"] = []string{"Polymers", "Kevlar/Teflon", "In", "1d6 x 3", "9000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["4311"] = []string{"Polymers", "Kevlar/Teflon", "In", "1d6 x 3", "9000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}
	tgrMap["4312"] = []string{"Polymers", "Advanced Ballistic Weave", "In", "1d6 x 1", "10000", "In +0", "Ri +2, Ni +1", "+1", "+0", "1d6 x 10"}

	tgrMap["442"] = []string{"Precious Metals", "Bismuth/Indium", "As, De, Ic, Fl", "1d6 x 1 + 2", "10000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["443"] = []string{"Precious Metals", "Beryllium/Silver", "As, De, Ic, Fl", "1d6 x 1 + 1", "25000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["444"] = []string{"Precious Metals", "Beryllium/Silver", "As, De, Ic, Fl", "1d6 x 1 + 1", "25000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["445"] = []string{"Precious Metals", "Beryllium/Silver", "As, De, Ic, Fl", "1d6 x 1 + 1", "25000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["446"] = []string{"Precious Metals", "Ruthenium/Rhenium", "As, De, Ic, Fl", "1d6 x 1", "50000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["447"] = []string{"Precious Metals", "Ruthenium/Rhenium", "As, De, Ic, Fl", "1d6 x 1", "50000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["448"] = []string{"Precious Metals", "Ruthenium/Rhenium", "As, De, Ic, Fl", "1d6 x 1", "50000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["449"] = []string{"Precious Metals", "Gold/Osmium/Iridium", "As, De, Ic, Fl", "1d6 x 1", "75000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["4410"] = []string{"Precious Metals", "Gold/Osmium/Iridium", "As, De, Ic, Fl", "1d6 x 1", "75000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["4411"] = []string{"Precious Metals", "Gold/Osmium/Iridium", "As, De, Ic, Fl", "1d6 x 1", "75000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}
	tgrMap["4412"] = []string{"Precious Metals", "Platinum/Rhodium", "As, De, Ic, Fl", "1d6 x 0", "100000", "As +3, De +1, Ic +2, Fl +0", "Ri +3, In +2, Ht +1", "+3", "+4", "1d6 x 1"}

	tgrMap["452"] = []string{"Radioactives", "Nuclear Waste/Deactivated Materials", "As, De, Lo", "1d6 x 1 + 3", "500000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["453"] = []string{"Radioactives", "Industrial Isotopes", "As, De, Lo", "1d6 x 1 + 2", "750000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["454"] = []string{"Radioactives", "Industrial Isotopes", "As, De, Lo", "1d6 x 1 + 2", "750000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["455"] = []string{"Radioactives", "Industrial Isotopes", "As, De, Lo", "1d6 x 1 + 2", "750000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["456"] = []string{"Radioactives", "Medical Isotopes/Reactor-Grade Uranium", "As, De, Lo", "1d6 x 1", "1000000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["457"] = []string{"Radioactives", "Medical Isotopes/Reactor-Grade Uranium", "As, De, Lo", "1d6 x 1", "1000000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["458"] = []string{"Radioactives", "Medical Isotopes/Reactor-Grade Uranium", "As, De, Lo", "1d6 x 1", "1000000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["459"] = []string{"Radioactives", "Weapons-Grade Plutonium/Fusion Cell Rods", "As, De, Lo", "1d6 x 0 + 1", "1250000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["4510"] = []string{"Radioactives", "Weapons-Grade Plutonium/Fusion Cell Rods", "As, De, Lo", "1d6 x 0 + 1", "1250000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["4511"] = []string{"Radioactives", "Weapons-Grade Plutonium/Fusion Cell Rods", "As, De, Lo", "1d6 x 0 + 1", "1250000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}
	tgrMap["4512"] = []string{"Radioactives", "Superweapon-grade Isotopes", "As, De, Lo", "1d6 x 0 + 1", "1500000", "Ass +2, De +0, Lo +2", "In +3, Ht +1, Ni -2, Ag -3", "+4", "+3", "1d6 x 1"}

	tgrMap["462"] = []string{"Robots", "Automated Robotics/Cargo Drones", "In", "1d6 x 7", "150000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["463"] = []string{"Robots", "Industrial or Personal Drones", "In", "1d6 x 6", "300000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["464"] = []string{"Robots", "Industrial or Personal Drones", "In", "1d6 x 6", "300000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["465"] = []string{"Robots", "Industrial or Personal Drones", "In", "1d6 x 6", "300000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["466"] = []string{"Robots", "Combat or Guardian Drones", "In", "1d6 x 5", "400000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["467"] = []string{"Robots", "Combat or Guardian Drones", "In", "1d6 x 5", "400000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["468"] = []string{"Robots", "Combat or Guardian Drones", "In", "1d6 x 5", "400000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["469"] = []string{"Robots", "Scout and Sensor Drones", "In", "1d6 x 2", "500000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["4610"] = []string{"Robots", "Scout and Sensor Drones", "In", "1d6 x 2", "500000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["4611"] = []string{"Robots", "Scout and Sensor Drones", "In", "1d6 x 2", "500000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}
	tgrMap["4612"] = []string{"Robots", "Advanced Robotics", "In", "1d6 x 1", "650000", "In +0", "Ag +2, Ht +1", "2", "1", "1d6 x 5"}

	tgrMap["512"] = []string{"Spices", "Table Salt/Black Pepper", "Ga, De, Wa", "1d6 x 6", "1000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["513"] = []string{"Spices", "Adobo/Basil/Sage", "Ga, De, Wa", "1d6 x 5", "3000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["514"] = []string{"Spices", "Adobo/Basil/Sage", "Ga, De, Wa", "1d6 x 5", "3000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["515"] = []string{"Spices", "Adobo/Basil/Sage", "Ga, De, Wa", "1d6 x 5", "3000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["516"] = []string{"Spices", "Aniseed/Curry/Fennel/White Pepper", "Ga, De, Wa", "1d6 x 5", "6000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["517"] = []string{"Spices", "Aniseed/Curry/Fennel/White Pepper", "Ga, De, Wa", "1d6 x 5", "6000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["518"] = []string{"Spices", "Aniseed/Curry/Fennel/White Pepper", "Ga, De, Wa", "1d6 x 5", "6000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["519"] = []string{"Spices", "Cinnamon/Marjoram/Wasabi", "Ga, De, Wa", "1d6 x 3", "9000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["5110"] = []string{"Spices", "Cinnamon/Marjoram/Wasabi", "Ga, De, Wa", "1d6 x 3", "9000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["5111"] = []string{"Spices", "Cinnamon/Marjoram/Wasabi", "Ga, De, Wa", "1d6 x 3", "9000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}
	tgrMap["5112"] = []string{"Spices", "Black Salt/Saffron/Alien Flavourings", "Ga, De, Wa", "1d6 x 1", "12000", "Ga +0, De +2, Wa +0", "Hi +2, Ri +3, Po +3", "+2", "-1", "1d6 x 5"}

	tgrMap["522"] = []string{"Textiles", "Yarn/Wool/Canvas", "Ag, Ni", "1d6 x 12", "1000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["523"] = []string{"Textiles", "Animal-based Fabrics", "Ag, Ni", "1d6 x 10", "2000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["524"] = []string{"Textiles", "Animal-based Fabrics", "Ag, Ni", "1d6 x 10", "2000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["525"] = []string{"Textiles", "Animal-based Fabrics", "Ag, Ni", "1d6 x 10", "2000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["526"] = []string{"Textiles", "Cotton or Flax-based Fabrics", "Ag, Ni", "1d6 x 10", "3000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["527"] = []string{"Textiles", "Cotton or Flax-based Fabrics", "Ag, Ni", "1d6 x 10", "3000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["528"] = []string{"Textiles", "Cotton or Flax-based Fabrics", "Ag, Ni", "1d6 x 10", "3000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["529"] = []string{"Textiles", "Synthetic Silks/Finished Common Clothing", "Ag, Ni", "1d6 x 6", "4000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["5210"] = []string{"Textiles", "Synthetic Silks/Finished Common Clothing", "Ag, Ni", "1d6 x 6", "4000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["5211"] = []string{"Textiles", "Synthetic Silks/Finished Common Clothing", "Ag, Ni", "1d6 x 6", "4000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}
	tgrMap["5212"] = []string{"Textiles", "Organic Silk/Satin/Finished Fine Clothing", "Ag, Ni", "1d6 x 3", "5000", "Ag +7, Ni +0", "Hi +3, Na +2", "+1", "-2", "1d6 x 10"}

	tgrMap["532"] = []string{"Uncommon Ore", "Lead/Zinc", "As, Ic", "1d6 x 10", "1000", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["533"] = []string{"Uncommon Ore", "Copper/Tin", "As, Ic", "1d6 x 10", "2500", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["534"] = []string{"Uncommon Ore", "Copper/Tin", "As, Ic", "1d6 x 10", "2500", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["535"] = []string{"Uncommon Ore", "Copper/Tin", "As, Ic", "1d6 x 10", "2500", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["536"] = []string{"Uncommon Ore", "Nickel/Sodium/Tungsten", "As, Ic", "1d6 x 10", "5000", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["537"] = []string{"Uncommon Ore", "Nickel/Sodium/Tungsten", "As, Ic", "1d6 x 10", "5000", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["538"] = []string{"Uncommon Ore", "Nickel/Sodium/Tungsten", "As, Ic", "1d6 x 10", "5000", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["539"] = []string{"Uncommon Ore", "Gold/Silver/Ilmenite", "As, Ic", "1d6 x 5", "7500", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["5310"] = []string{"Uncommon Ore", "Gold/Silver/Ilmenite", "As, Ic", "1d6 x 5", "7500", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["5311"] = []string{"Uncommon Ore", "Gold/Silver/Ilmenite", "As, Ic", "1d6 x 5", "7500", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}
	tgrMap["5312"] = []string{"Uncommon Ore", "Platinum/Uranium", "As, Ic", "1d6 x 2", "10000", "As +4, Ic +0", "In +3, Ni +1", "+2", "-2", "1d6 x 10"}

	tgrMap["542"] = []string{"Uncommon Raw Materials", "Aluminium/Brass/Calcium", "Ag, De, Wa", "1d6 x 14", "5000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["543"] = []string{"Uncommon Raw Materials", "Carbonate/Magnesium/Meteoric Iron", "Ag, De, Wa", "1d6 x 12", "10000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["544"] = []string{"Uncommon Raw Materials", "Carbonate/Magnesium/Meteoric Iron", "Ag, De, Wa", "1d6 x 12", "10000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["545"] = []string{"Uncommon Raw Materials", "Carbonate/Magnesium/Meteoric Iron", "Ag, De, Wa", "1d6 x 12", "10000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["546"] = []string{"Uncommon Raw Materials", "Marble/Potassium/Titanium", "Ag, De, Wa", "1d6 x 10", "20000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["547"] = []string{"Uncommon Raw Materials", "Marble/Potassium/Titanium", "Ag, De, Wa", "1d6 x 10", "20000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["548"] = []string{"Uncommon Raw Materials", "Marble/Potassium/Titanium", "Ag, De, Wa", "1d6 x 10", "20000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["549"] = []string{"Uncommon Raw Materials", "Stellite/Tombac", "Ag, De, Wa", "1d6 x 8", "35000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["5410"] = []string{"Uncommon Raw Materials", "Stellite/Tombac", "Ag, De, Wa", "1d6 x 8", "35000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["5411"] = []string{"Uncommon Raw Materials", "Stellite/Tombac", "Ag, De, Wa", "1d6 x 8", "35000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}
	tgrMap["5412"] = []string{"Uncommon Raw Materials", "Depleted Uranium/Ceramic-Alloy", "Ag, De, Wa", "1d6 x 3", "50000", "Ag +2, De +0, Wa +1", "In +2, Ht +1", "+2", "-2", "1d6 x 10"}

	tgrMap["552"] = []string{"Wood", "Low-grade Rough Cuts/Construction Scrap", "Ag, Ga", "1d6 x 12", "100", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["553"] = []string{"Wood", "High-Grade Rough-Cut", "Ag, Ga", "1d6 x 10", "500", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["554"] = []string{"Wood", "High-Grade Rough-Cut", "Ag, Ga", "1d6 x 10", "500", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["555"] = []string{"Wood", "High-Grade Rough-Cut", "Ag, Ga", "1d6 x 10", "500", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["556"] = []string{"Wood", "Construction-grade Timber", "Ag, Ga", "1d6 x 10", "1000", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["557"] = []string{"Wood", "Construction-grade Timber", "Ag, Ga", "1d6 x 10", "1000", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["558"] = []string{"Wood", "Construction-grade Timber", "Ag, Ga", "1d6 x 10", "1000", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["559"] = []string{"Wood", "Furniture-grade Timber/Rare Grades", "Ag, Ga", "1d6 x 6", "2000", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["5510"] = []string{"Wood", "Furniture-grade Timber/Rare Grades", "Ag, Ga", "1d6 x 6", "2000", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["5511"] = []string{"Wood", "Furniture-grade Timber/Rare Grades", "Ag, Ga", "1d6 x 6", "2000", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}
	tgrMap["5512"] = []string{"Wood", "Exotics (Pernambuco, White Mahogany, etc.)", "Ag, Ga", "1d6 x 2", "4000", "Ag +6, Ga +0", "Ri +2, In +1", "+1", "-4", "1d6 x 10"}

	tgrMap["562"] = []string{"Vehicles", "Wheeled Repair Components", "In, Ht", "1d6 x 14", "5000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["563"] = []string{"Vehicles", "Tracked Repair Components", "In, Ht", "1d6 x 12", "10000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["564"] = []string{"Vehicles", "Tracked Repair Components", "In, Ht", "1d6 x 12", "10000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["565"] = []string{"Vehicles", "Tracked Repair Components", "In, Ht", "1d6 x 12", "10000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["566"] = []string{"Vehicles", "Wheeled Components or Packages", "In, Ht", "1d6 x 10", "15000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["567"] = []string{"Vehicles", "Wheeled Components or Packages", "In, Ht", "1d6 x 10", "15000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["568"] = []string{"Vehicles", "Wheeled Components or Packages", "In, Ht", "1d6 x 10", "15000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["569"] = []string{"Vehicles", "Wheeled Vehicles/Tracked Components or Packages", "In, Ht", "1d6 x 6", "20000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["5610"] = []string{"Vehicles", "Wheeled Vehicles/Tracked Components or Packages", "In, Ht", "1d6 x 6", "20000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["5611"] = []string{"Vehicles", "Wheeled Vehicles/Tracked Components or Packages", "In, Ht", "1d6 x 6", "20000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}
	tgrMap["5612"] = []string{"Vehicles", "Tracked Vehicles", "In, Ht", "1d6 x 2", "30000", "In +2, Ht +1", "Ni +2, Hi +1", "+2", "-2", "1d6 x 10"}

	tgrMap["612"] = []string{"Biochemicals, Illegal", "Herbal Stimulants/Ultra-Caffeine", "Ag, Wa", "1d6 x 6", "10000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["613"] = []string{"Biochemicals, Illegal", "Raw Growth Hormones", "Ag, Wa", "1d6 x 5", "25000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["614"] = []string{"Biochemicals, Illegal", "Raw Growth Hormones", "Ag, Wa", "1d6 x 5", "25000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["615"] = []string{"Biochemicals, Illegal", "Raw Growth Hormones", "Ag, Wa", "1d6 x 5", "25000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["616"] = []string{"Biochemicals, Illegal", "Chemical Solvents/Protein Duplexer Steroids", "Ag, Wa", "1d6 x 5", "50000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["617"] = []string{"Biochemicals, Illegal", "Chemical Solvents/Protein Duplexer Steroids", "Ag, Wa", "1d6 x 5", "50000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["618"] = []string{"Biochemicals, Illegal", "Chemical Solvents/Protein Duplexer Steroids", "Ag, Wa", "1d6 x 5", "50000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["619"] = []string{"Biochemicals, Illegal", "Bio-Acid/Pheromone Extracts", "Ag, Wa", "1d6 x 0 + 2", "100000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["6110"] = []string{"Biochemicals, Illegal", "Bio-Acid/Pheromone Extracts", "Ag, Wa", "1d6 x 0 + 2", "100000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["6111"] = []string{"Biochemicals, Illegal", "Bio-Acid/Pheromone Extracts", "Ag, Wa", "1d6 x 0 + 2", "100000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}
	tgrMap["6112"] = []string{"Biochemicals, Illegal", "Genetic Mutagens/Organic Toxins", "Ag, Wa", "1d6 x 0 + 1", "200000", "Ag +0, Wa +2", "In +6", "+4", "+4", "1d6 x 5"}

	tgrMap["622"] = []string{"Cybernetics, Illegal", "Unlicensed Augment Tools and Parts", "Ht", "1d6 x 2", "100000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["623"] = []string{"Cybernetics, Illegal", "Physical Enhancement Tissues", "Ht", "1d6 x 2", "150000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["624"] = []string{"Cybernetics, Illegal", "Physical Enhancement Tissues", "Ht", "1d6 x 2", "150000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["625"] = []string{"Cybernetics, Illegal", "Physical Enhancement Tissues", "Ht", "1d6 x 2", "150000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["626"] = []string{"Cybernetics, Illegal", "Unlicensed Augmentatives/Combat Implant Additives", "Ht", "1d6 x 1", "250000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["627"] = []string{"Cybernetics, Illegal", "Unlicensed Augmentatives/Combat Implant Additives", "Ht", "1d6 x 1", "250000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["628"] = []string{"Cybernetics, Illegal", "Unlicensed Augmentatives/Combat Implant Additives", "Ht", "1d6 x 1", "250000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["629"] = []string{"Cybernetics, Illegal", "Combat Prosthetics/Surgical Duplications", "Ht", "1d6 x 0 + 2", "400000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["6210"] = []string{"Cybernetics, Illegal", "Combat Prosthetics/Surgical Duplications", "Ht", "1d6 x 0 + 2", "400000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["6211"] = []string{"Cybernetics, Illegal", "Combat Prosthetics/Surgical Duplications", "Ht", "1d6 x 0 + 2", "400000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}
	tgrMap["6212"] = []string{"Cybernetics, Illegal", "Mimicry Augmetics", "Ht", "1d6 x 0 + 1", "650000", "Ht +0", "As +4, Ic +4, Ri +8, AZ +6, RZ +6", "+5", "+5", "1d6 x 1"}

	tgrMap["632"] = []string{"Drugs, Illegal", "Herbal Stimulants/Biological Hallucinogens", "As, De, Hi, Wa", "1d6 x 1 + 2", "25000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["633"] = []string{"Drugs, Illegal", "Chemical Depressants/Natural Narcotics", "As, De, Hi, Wa", "1d6 x 1 + 1", "50000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["634"] = []string{"Drugs, Illegal", "Chemical Depressants/Natural Narcotics", "As, De, Hi, Wa", "1d6 x 1 + 1", "50000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["635"] = []string{"Drugs, Illegal", "Chemical Depressants/Natural Narcotics", "As, De, Hi, Wa", "1d6 x 1 + 1", "50000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["636"] = []string{"Drugs, Illegal", "Chemical Stimulants and Hallucinogens", "As, De, Hi, Wa", "1d6 x 1", "100000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["637"] = []string{"Drugs, Illegal", "Chemical Stimulants and Hallucinogens", "As, De, Hi, Wa", "1d6 x 1", "100000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["638"] = []string{"Drugs, Illegal", "Chemical Stimulants and Hallucinogens", "As, De, Hi, Wa", "1d6 x 1", "100000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["639"] = []string{"Drugs, Illegal", "Designer Narcotics", "As, De, Hi, Wa", "1d6 x 0 + 2", "200000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["6310"] = []string{"Drugs, Illegal", "Designer Narcotics", "As, De, Hi, Wa", "1d6 x 0 + 2", "200000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["6311"] = []string{"Drugs, Illegal", "Designer Narcotics", "As, De, Hi, Wa", "1d6 x 0 + 2", "200000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}
	tgrMap["6312"] = []string{"Drugs, Illegal", "Alien Synthetics/Psi-Drugs", "As, De, Hi, Wa", "1d6 x 0 + 1", "300000", "As +0, De +0, Ga +0, Wa +0", "Ri +6, Hi +6", "+4", "+6", "1d6 x 1"}

	tgrMap["642"] = []string{"Luxuries, Illegal", "Anti-Governmental Propaganda/Endangered Animal Products", "Ag, Ga, Wa", "1d6 x 1 + 1", "10000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["643"] = []string{"Luxuries, Illegal", "Black-data Recordings/Slaving Gear", "Ag, Ga, Wa", "1d6 x 1", "25000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["644"] = []string{"Luxuries, Illegal", "Black-data Recordings/Slaving Gear", "Ag, Ga, Wa", "1d6 x 1", "25000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["645"] = []string{"Luxuries, Illegal", "Black-data Recordings/Slaving Gear", "Ag, Ga, Wa", "1d6 x 1", "25000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["646"] = []string{"Luxuries, Illegal", "Extinct Animal Products", "Ag, Ga, Wa", "1d6 x 1", "50000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["647"] = []string{"Luxuries, Illegal", "Extinct Animal Products", "Ag, Ga, Wa", "1d6 x 1", "50000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["648"] = []string{"Luxuries, Illegal", "Extinct Animal Products", "Ag, Ga, Wa", "1d6 x 1", "50000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["649"] = []string{"Luxuries, Illegal", "BTL Devices/Cloning Equipment", "Ag, Ga, Wa", "1d6 x 0 + 2", "100000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["6410"] = []string{"Luxuries, Illegal", "BTL Devices/Cloning Equipment", "Ag, Ga, Wa", "1d6 x 0 + 2", "100000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["6411"] = []string{"Luxuries, Illegal", "BTL Devices/Cloning Equipment", "Ag, Ga, Wa", "1d6 x 0 + 2", "100000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}
	tgrMap["6412"] = []string{"Luxuries, Illegal", "Forbidden Pleasures", "Ag, Ga, Wa", "1d6 x 0 + 1", "200000", "Ag +2, Ga +0, Wa +1", "Ri +6, Hi +4", "+4", "+4", "1d6 x 1"}

	tgrMap["652"] = []string{"Weapons, Illegal", "Chain-drive Weaponry/Armour-Piercing Ammunition", "In, Ht", "1d6 x 6", "50000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["653"] = []string{"Weapons, Illegal", "Protected Technologies/Explosive or Incendiary Ammunition", "In, Ht", "1d6 x 5", "100000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["654"] = []string{"Weapons, Illegal", "Protected Technologies/Explosive or Incendiary Ammunition", "In, Ht", "1d6 x 5", "100000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["655"] = []string{"Weapons, Illegal", "Protected Technologies/Explosive or Incendiary Ammunition", "In, Ht", "1d6 x 5", "100000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["656"] = []string{"Weapons, Illegal", "Synthetic Poisons/Personal-scale Mass Trauma Explosives", "In, Ht", "1d6 x 5", "150000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["657"] = []string{"Weapons, Illegal", "Synthetic Poisons/Personal-scale Mass Trauma Explosives", "In, Ht", "1d6 x 5", "150000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["658"] = []string{"Weapons, Illegal", "Synthetic Poisons/Personal-scale Mass Trauma Explosives", "In, Ht", "1d6 x 5", "150000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["659"] = []string{"Weapons, Illegal", "Arclight Weaponry/Biological or Chemical Weaponry/Naval Starship Weaponry", "In, Ht", "1d6 x 2", "300000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["6510"] = []string{"Weapons, Illegal", "Arclight Weaponry/Biological or Chemical Weaponry/Naval Starship Weaponry", "In, Ht", "1d6 x 2", "300000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["6511"] = []string{"Weapons, Illegal", "Arclight Weaponry/Biological or Chemical Weaponry/Naval Starship Weaponry", "In, Ht", "1d6 x 2", "300000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	tgrMap["6512"] = []string{"Weapons, Illegal", "Disintegrators/Psi-Weaponry/Weapons of Mass Destruction", "In, Ht", "1d6 x 1", "450000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 5"}
	//return 6at, "max"a
	tgrMap["662"] = []string{"Exotic, Illegal", "Chain-drive Weaponry/Armour-Piercing Ammunition", "In, Ht", "1d6 x 6", "50000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["663"] = []string{"Exotic, Illegal", "Protected Technologies/Explosive or Incendiary Ammunition", "In, Ht", "1d6 x 5", "100000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["664"] = []string{"Exotic, Illegal", "Protected Technologies/Explosive or Incendiary Ammunition", "In, Ht", "1d6 x 5", "100000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["665"] = []string{"Exotic, Illegal", "Protected Technologies/Explosive or Incendiary Ammunition", "In, Ht", "1d6 x 5", "100000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["666"] = []string{"Exotic, Illegal", "Synthetic Poisons/Personal-scale Mass Trauma Explosives", "In, Ht", "1d6 x 5", "150000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["667"] = []string{"Exotic, Illegal", "Synthetic Poisons/Personal-scale Mass Trauma Explosives", "In, Ht", "1d6 x 5", "150000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["668"] = []string{"Exotic, Illegal", "Synthetic Poisons/Personal-scale Mass Trauma Explosives", "In, Ht", "1d6 x 5", "150000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["669"] = []string{"Exotic, Illegal", "Arclight Weaponry/Biological or Chemical Weaponry/Naval Starship Weaponry", "In, Ht", "1d6 x 2", "300000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["6610"] = []string{"Exotic, Illegal", "Arclight Weaponry/Biological or Chemical Weaponry/Naval Starship Weaponry", "In, Ht", "1d6 x 2", "300000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["6611"] = []string{"Exotic, Illegal", "Arclight Weaponry/Biological or Chemical Weaponry/Naval Starship Weaponry", "In, Ht", "1d6 x 2", "300000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	tgrMap["6612"] = []string{"Exotic, Illegal", "Disintegrators/Psi-Weaponry/Weapons of Mass Destruction", "In, Ht", "1d6 x 1", "450000", "In +0, Ht +2", "Po +6, AZ +8, RZ +10", "+5", "+6", "1d6 x 0 + 1"}
	return tgrMap
}

func parseTags(allTags string) []string {
	tags := strings.Split(allTags, ", ")
	// for i := range tags {
	// 	fmt.Println(tags[i])
	// }
	return tags
}

func tgrCategory(code string) string {

	return ""
}

// func (tgr *TradeGoodR) pickRandomDescription() string {
// 	descrAll := strings.Split(tgr.description, "/")
// 	dice := "d" + convert.ItoS(len(descrAll))
// 	pos := utils.RollDice(dice, -1)
// 	return descrAll[pos]
// }

func allTradeGoodsRCodes() (allCodesSlice []string) {
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			for k := 2; k <= 12; k++ {
				code := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k)
				allCodesSlice = append(allCodesSlice, code)
			}
		}
	}
	return allCodesSlice
}
