package main

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/convert"

	"github.com/Galdoba/utils"
)

//Boon -
type Boon struct {
	name      string
	archetype string
	cost      []boonCost
	descr     string
	bType     string
}

type boonCost struct {
	char string
	qty  int
}

//Dynasty -
type Dynasty struct {
	characteristics    map[string]int
	traits             map[string]int
	aptitudes          map[string]int
	values             map[string]int
	aptTotalEffect     map[string]int
	misc               map[string]int
	powerBase          string
	mainCharacteristic string
	archetype          string
	firstGenBonus      string
	managementAssets   []string
	boons              []Boon
	nextDate           map[string]string //нужен интерфейс для работы со временем
}

//NewDynasty -
func NewDynasty() *Dynasty {

	/*
	   1. Roll Characteristics and determine Characteristic
	   modifiers.
	   2. a. Choose a Power Base. 							-ok
	   b. Gain Trait and Aptitude Modifiers.				-ok
	   3. a. Choose a Dynasty Archetype.					-ok
	   b. Determine Base Traits and Aptitudes.				-ok
	   c. Gain First Generation bonuses.					-ok
	   d. Determine Dynasty Boons and Hinders.				-ok
	   4. a. Determine Management Assets.					-ok
	   b. Gain Management Asset Benefits.					-onHold (это 6 действий/эффектов которые отыгрывают в игровом цикле)
	   5. Calculate First Generation Values.
	   a. Characteristic Practice							-ok
	   b. Trait Management									-ok
	   c. Aptitude Training									-ok
	   d. Final Value Adjustment							-ok
	   6. Move on to the Background and Historic Events
	   process.
	*/
	dyn := &Dynasty{}
	//1 Characteristics
	dyn.characteristics = make(map[string]int)
	charList := characteristicsLIST()
	max := 0
	for i := range charList {
		dyn.characteristics[charList[i]] = utils.RollDice("2d6")
		if dyn.characteristics[charList[i]] > max {
			max = dyn.characteristics[charList[i]]
			dyn.mainCharacteristic = charList[i]
		}
	}
	dyn.traits = make(map[string]int)
	dyn.aptitudes = make(map[string]int)
	dyn.values = make(map[string]int)
	dyn.aptTotalEffect = make(map[string]int)
	dyn.misc = make(map[string]int)
	//2 PowerBase
	dyn.powerBase = utils.RandomFromList(powerBaseList())

	//3 Archetype
	acrhList := archetypeLIST()
	var pickFrom []string
	for i := range acrhList {
		if dyn.qualifyForArchetype(acrhList[i]) {
			pickFrom = append(pickFrom, acrhList[i])
		}
	}
	fmt.Println(pickFrom)
	if len(pickFrom) == 0 {
		return nil
	}
	dyn.archetype = utils.RandomFromList(pickFrom)
	dyn.determineBaseTraits()
	dyn.gainBaseAptitudes()
	dyn.applyEffectPowerBase()

	dyn.boonsAndHinders()
	dyn.applyBoonEffects()
	//4 First Generation
	dyn.gainFirstGenerationBonuses()
	//5 management Assets
	dyn.managementAssets = determineManagementAssets(dyn.archetype)
	dyn.trainCharacteristics()
	dyn.manageTraits()
	dyn.trainAptitudes()
	dyn.finalValuesAdjustmen()
	//6 BackGround Events:
	dyn.misc["historicEvents"] = 0
	backGrEvents := utils.RollDice("1d3", 2)
	for i := 0; i < backGrEvents; i++ {

		eventCode := randomEventCode(dyn.archetype)
		fmt.Println("eventCode:", eventCode)
		dyn.doBackGroundEvent(eventCode)
	}
	//7 end

	for i := range listTRAITS {
		if dyn.pickVal(listTRAITS[i]) < 0 {
			dyn.setBase(listTRAITS[i], 0)
		}
		if dyn.pickVal(listTRAITS[i]) > 10 {
			dyn.setBase(listTRAITS[i], 10)
		}
	}
	return dyn
}

// type BoonMap struct {
// 	BoonByName map[string]*Boon
// }
func newBoon(name string, archetype string, cost []boonCost) Boon {
	boon := Boon{}
	boon.name = name
	boon.archetype = archetype
	boon.cost = cost
	boon.descr = "TODO: " + name + " DESCRIPTION"
	boon.bType = "Boon"
	return boon
}

func newHinder(name string, archetype string, cost []boonCost) Boon {
	boon := Boon{}
	boon.name = name
	boon.archetype = archetype
	boon.cost = cost
	boon.descr = "TODO: " + name + " DESCRIPTION"
	boon.bType = "Hinder"
	return boon
}

func payCost(dyn *Dynasty, char string, mod int) {
	dyn.setBonus(char, mod)

}

//BoonHinderByName -
func BoonHinderByName() map[string]Boon {
	BoonHinderByName := make(map[string]Boon)
	BoonHinderByName["Commercial Psions"] = newBoon("Commercial Psions", archConglomerate, []boonCost{boonCost{charPopularity, -1}})
	BoonHinderByName["Endless Funds"] = newBoon("Endless Funds", archConglomerate, []boonCost{boonCost{traitFiscalDefence, -2}})
	BoonHinderByName["Governmental Backing"] = newBoon("Governmental Backing", archConglomerate, []boonCost{boonCost{charTradition, -1}})
	BoonHinderByName["Military Contracts"] = newBoon("Military Contracts", archConglomerate, []boonCost{boonCost{charPopularity, -1}, boonCost{charGreed, -1}})
	BoonHinderByName["Total Control"] = newBoon("Total Control", archConglomerate, []boonCost{boonCost{traitTerritorialDefence, -2}})

	BoonHinderByName["Alien Extortions"] = newHinder("Alien Extortions", archConglomerate, []boonCost{boonCost{charGreed, 1}})
	BoonHinderByName["Market Mercenaries"] = newHinder("Market Mercenaries", archConglomerate, []boonCost{boonCost{charCleverness, 1}, boonCost{charMilitarism, 1}})
	BoonHinderByName["Spies in the Network"] = newHinder("Spies in the Network", archConglomerate, []boonCost{boonCost{charScheming, 1}})
	BoonHinderByName["Underworld Loans"] = newHinder("Underworld Loans", archConglomerate, []boonCost{boonCost{traitFiscalDefence, 2}})

	BoonHinderByName["Bureaucratic Roots"] = newBoon("Bureaucratic Roots", archMediaEmpire, []boonCost{boonCost{charGreed, -1}})
	BoonHinderByName["Gossip Rags"] = newBoon("Gossip Rags", archMediaEmpire, []boonCost{boonCost{traitCulture, -1}})
	BoonHinderByName["Politics Engine (Lty)"] = newBoon("Politics Engine (Lty)", archMediaEmpire, []boonCost{boonCost{charLoyalty, -1}})
	BoonHinderByName["Politics Engine (Sch)"] = newBoon("Politics Engine (Sch)", archMediaEmpire, []boonCost{boonCost{charScheming, -1}})
	BoonHinderByName["Sports Contracts (Pop)"] = newBoon("Sports Contracts (Pop)", archMediaEmpire, []boonCost{boonCost{charPopularity, -1}})
	BoonHinderByName["Sports Contracts (Culture)"] = newBoon("Sports Contracts (Culture)", archMediaEmpire, []boonCost{boonCost{traitCulture, -1}})
	BoonHinderByName["Voice of a Generation"] = newBoon("Voice of a Generation", archMediaEmpire, []boonCost{boonCost{charPopularity, -1}})

	BoonHinderByName["Hostile Paparazzi"] = newHinder("Hostile Paparazzi", archMediaEmpire, []boonCost{boonCost{traitCulture, 2}})
	BoonHinderByName["Pirate Comms Station (Pop)"] = newHinder("Pirate Comms Station (Pop)", archMediaEmpire, []boonCost{boonCost{charPopularity, 1}})
	BoonHinderByName["Pirate Comms Station (Tcy)"] = newHinder("Pirate Comms Station (Tcy)", archMediaEmpire, []boonCost{boonCost{charTenacity, 1}})
	BoonHinderByName["Rumours of Corruption"] = newHinder("Rumours of Corruption", archMediaEmpire, []boonCost{boonCost{charCleverness, 1}})
	BoonHinderByName["Translation Troubles"] = newHinder("Translation Troubles", archMediaEmpire, []boonCost{boonCost{traitTerritorialDefence, 1}})

	BoonHinderByName["Commercial Psions"] = newBoon("Commercial Psions", archMerchantMarket, []boonCost{boonCost{charPopularity, -1}})
	BoonHinderByName["Interstellar Funding (Tra)"] = newBoon("Interstellar Funding (Tra)", archMerchantMarket, []boonCost{boonCost{charTradition, -1}})
	BoonHinderByName["Interstellar Funding (Culture)"] = newBoon("Interstellar Funding (Culture)", archMerchantMarket, []boonCost{boonCost{traitCulture, -2}})
	BoonHinderByName["Naval Escorts"] = newBoon("Naval Escorts", archMerchantMarket, []boonCost{boonCost{charMilitarism, -1}})
	BoonHinderByName["Secure Production (FD)"] = newBoon("Secure Production (FD)", archMerchantMarket, []boonCost{boonCost{traitFiscalDefence, -1}})
	BoonHinderByName["Secure Production (TD)"] = newBoon("Secure Production (TD)", archMerchantMarket, []boonCost{boonCost{traitTerritorialDefence, -1}})
	BoonHinderByName["Vaulted Technologies"] = newBoon("Vaulted Technologies", archMerchantMarket, []boonCost{boonCost{traitTechnology, -1}})

	BoonHinderByName["Charitable Causes"] = newHinder("Charitable Causes", archMerchantMarket, []boonCost{boonCost{traitCulture, 1}})
	BoonHinderByName["Depression Debts"] = newHinder("Depression Debts", archMerchantMarket, []boonCost{boonCost{charGreed, 1}})
	BoonHinderByName["Pirate Problems"] = newHinder("Pirate Problems", archMerchantMarket, []boonCost{boonCost{traitTerritorialDefence, 2}})
	BoonHinderByName["Resource Mercenaries"] = newHinder("Resource Mercenaries", archMerchantMarket, []boonCost{boonCost{charCleverness, 1}, boonCost{charMilitarism, 1}})

	BoonHinderByName["Aggressive Politics"] = newBoon("Aggressive Politics", archMilitaryCharter, []boonCost{boonCost{charPopularity, -1}})
	BoonHinderByName["Homeland Foundation"] = newBoon("Homeland Foundation", archMilitaryCharter, []boonCost{boonCost{traitFleet, -1}})
	BoonHinderByName["Laurels of Victory"] = newBoon("Laurels of Victory", archMilitaryCharter, []boonCost{boonCost{charTenacity, -1}})
	BoonHinderByName["Martial Law"] = newBoon("Martial Law", archMilitaryCharter, []boonCost{boonCost{charLoyalty, -1}, boonCost{traitCulture, -1}})
	BoonHinderByName["War Hero"] = newBoon("War Hero", archMilitaryCharter, []boonCost{boonCost{charScheming, -1}})

	BoonHinderByName["Enemies on All Fronts"] = newHinder("Enemies on All Fronts", archMilitaryCharter, []boonCost{boonCost{charCleverness, 1}, boonCost{charMilitarism, 1}})
	BoonHinderByName["Gun Runner Gambles"] = newHinder("Gun Runner Gambles", archMilitaryCharter, []boonCost{boonCost{traitTechnology, 1}})
	BoonHinderByName["Tech Problems"] = newHinder("Tech Problems", archMilitaryCharter, []boonCost{boonCost{charTenacity, 1}})
	BoonHinderByName["War Eternal"] = newHinder("War Eternal", archMilitaryCharter, []boonCost{boonCost{traitTerritorialDefence, 2}})

	BoonHinderByName["Breeding Eugenics"] = newBoon("Breeding Eugenics", archNobleLine, []boonCost{boonCost{traitTechnology, -1}})
	BoonHinderByName["Inherited Fortunes"] = newBoon("Inherited Fortunes", archNobleLine, []boonCost{boonCost{traitFiscalDefence, -1}})
	BoonHinderByName["Pocket Goverment"] = newBoon("Pocket Goverment", archNobleLine, []boonCost{boonCost{traitFleet, -1}})
	BoonHinderByName["Royal Family"] = newBoon("Royal Family", archNobleLine, []boonCost{boonCost{charLoyalty, -1}, boonCost{traitTechnology, -1}})
	BoonHinderByName["Secret Society"] = newBoon("Secret Society", archNobleLine, []boonCost{boonCost{charScheming, -1}})

	BoonHinderByName["Desease in the Genes"] = newHinder("Desease in the Genes", archNobleLine, []boonCost{boonCost{charTradition, 1}})
	BoonHinderByName["Inbred Rumors"] = newHinder("Inbred Rumors", archNobleLine, []boonCost{boonCost{traitCulture, 1}})
	BoonHinderByName["Primitive Subjects"] = newHinder("Primitive Subjects", archNobleLine, []boonCost{boonCost{traitTerritorialDefence, 2}})
	BoonHinderByName["Revolution in the Future"] = newHinder("Revolution in the Future", archNobleLine, []boonCost{boonCost{charScheming, 1}, boonCost{charMilitarism, 1}})

	BoonHinderByName["Alien Congregation"] = newBoon("Alien Congregation", archReligiousFaith, []boonCost{boonCost{charPopularity, -1}, boonCost{traitCulture, -1}})
	BoonHinderByName["Defenders of Faith"] = newBoon("Defenders of Faith", archReligiousFaith, []boonCost{boonCost{charScheming, -1}})
	BoonHinderByName["Holy Missionaries"] = newBoon("Holy Missionaries", archReligiousFaith, []boonCost{boonCost{charMilitarism, -1}})
	BoonHinderByName["Tithes and Donations"] = newBoon("Tithes and Donations", archReligiousFaith, []boonCost{boonCost{traitCulture, -1}})
	BoonHinderByName["Words of Gods"] = newBoon("Words of Gods", archReligiousFaith, []boonCost{boonCost{charTradition, -1}})

	BoonHinderByName["Atheist Coalition"] = newHinder("Atheist Coalition", archReligiousFaith, []boonCost{boonCost{charTenacity, 1}, boonCost{traitCulture, 1}})
	BoonHinderByName["Controversial Clergy"] = newHinder("Controversial Clergy", archReligiousFaith, []boonCost{boonCost{charLoyalty, 1}})
	BoonHinderByName["Superstitions Abound"] = newHinder("Superstitions Abound", archReligiousFaith, []boonCost{boonCost{traitCulture, 1}})
	BoonHinderByName["War Between Heavens (TD)"] = newHinder("War Between Heavens (TD)", archReligiousFaith, []boonCost{boonCost{traitTerritorialDefence, 2}})
	BoonHinderByName["War Between Heavens (Mil)"] = newHinder("War Between Heavens (Mil)", archReligiousFaith, []boonCost{boonCost{charMilitarism, 1}})

	BoonHinderByName["Deadly Reputation"] = newBoon("Deadly Reputation", archSyndicate, []boonCost{boonCost{charPopularity, -1}})
	BoonHinderByName["Family of Crime"] = newBoon("Family of Crime", archSyndicate, []boonCost{boonCost{charLoyalty, -1}})
	BoonHinderByName["Law Enforcement Spies"] = newBoon("Law Enforcement Spies", archSyndicate, []boonCost{boonCost{charMilitarism, -1}})
	BoonHinderByName["Pirate Shipyard"] = newBoon("Pirate Shipyard", archSyndicate, []boonCost{boonCost{charGreed, -1}, boonCost{traitFiscalDefence, -1}})
	BoonHinderByName["Rule Through Fear"] = newBoon("Rule Through Fear", archSyndicate, []boonCost{boonCost{charLoyalty, -1}})

	BoonHinderByName["Bounty Hunters"] = newHinder("Bounty Hunters", archSyndicate, []boonCost{boonCost{charCleverness, 1}, boonCost{charMilitarism, 1}})
	BoonHinderByName["Grudges and Vendettas"] = newHinder("Grudges and Vendettas", archSyndicate, []boonCost{boonCost{charLoyalty, 1}})
	BoonHinderByName["Most Wanted"] = newHinder("Most Wanted", archSyndicate, []boonCost{boonCost{traitCulture, 1}, boonCost{charScheming, 1}})
	BoonHinderByName["Question of Authority"] = newHinder("Question of Authority", archSyndicate, []boonCost{boonCost{charGreed, 1}})

	return BoonHinderByName
}

func (dyn *Dynasty) gainFirstGenerationBonuses() {
	dyn.firstGenBonus = rollFirstGenBonus(dyn.archetype)
	switch dyn.firstGenBonus {
	case "University Board Members":
		dyn.threeAptitudesToLevel1()
	case "Monopoly":
		dyn.setBonus(traitFiscalDefence, 1)
	case "Shipyard Access":
		dyn.setBonus(traitFleet, 1)
	case "Noble Investors":
		dyn.setBonus(valueWealth, 1)
	case "Inherited Pride":
		dyn.setBonus(traitCulture, 1)
	case "Multi-Stellar Benefactor":
		dyn.setBonus(utils.RandomFromList(listTRAITS), 1)
		dyn.setBonus(utils.RandomFromList(listTRAITS), 1)
	case "Royal Backing":
		dyn.add1d6ToValues()
	case "Psionic Investigators":
		dyn.setBonus(utils.RandomFromList(aptitudeLIST()), 1)
		dyn.setBonus(utils.RandomFromList(aptitudeLIST()), 1)
	case "Pyramid Structure":
		if utils.RandomBool() {
			r := utils.RollDice("d6")
			dyn.setBonus(valueWealth, r)
		} else {
			dyn.setBonus(charGreed, 1)
		}
	case "High-Tech Communications":
		dyn.setBonus(traitTechnology, 1)
	case "Military Reporters":
		if dyn.aptitudes[apttConquest] < 1 {
			dyn.setBase(apttConquest, 1)
		}
		if dyn.aptitudes[apttSecurity] < 1 {
			dyn.setBase(apttSecurity, 1)
		}
	case "Interstellar Cover Story":
		if dyn.aptitudes[apttExpression] < 1 {
			dyn.setBase(apttExpression, 1)
		}
		if dyn.aptitudes[apttPolitics] < 1 {
			dyn.setBase(apttPolitics, 1)
		}
		if dyn.aptitudes[apttPosturing] < 1 {
			dyn.setBase(apttPosturing, 1)
		}
	case "Total Media Monopoly":
		dyn.add1d6ToValues()
	case "Barter Over Sales":
		loop := 2
		for k, v := range dyn.aptitudes {
			if loop == 0 {
				continue
			}
			if v == 0 {
				dyn.aptitudes[k] = 1
				loop--
			}
		}
	case "Intense Collegiate Training":
		dyn.setBonus(utils.RandomFromList(aptitudeLIST()), 1)
	case "Patents Upon Patents":
		dyn.setBonus(valueWealth, 1)
	case "Govermental Acquisitions":
		dyn.setBonus(traitFleet, 1)
	case "A Republic in Good Fortune":
		r := utils.RollDice("d6")
		dyn.setBonus(valueWealth, r)
	case "Perfect Economy":
		dyn.add1d6ToValues()
	case "Intense Generational Training":
		dyn.threeAptitudesToLevel1()
	case "War Coffers":
		dyn.setBonus(valueWealth, 1)
	case "Naval Partners":
		dyn.setBonus(traitFleet, 1)
	case "An Armed Populance":
		dyn.setBonus(traitTerritorialDefence, 1)
	case "Victory Over Invasion":
		if utils.RandomBool() {
			dyn.setBonus(traitCulture, 2)
		} else {
			dyn.setBonus(charTradition, 1)
		}
	case "War Colleges":
		if utils.RandomBool() {
			dyn.threeAptitudesToLevel1()
		} else {
			i := 2
			for i > 0 {
				apt := utils.RandomFromList(aptitudeLIST())
				if dyn.aptitudes[apt] < 1 {
					dyn.setBase(apt, 1)
					i--
				}
			}
		}
	case "Noble Armada":
		dyn.setBonus(charTradition, 1)
		dyn.setBonus(traitFleet, 2)
	case "Royal Guard":
		dyn.setBonus(charMilitarism, 1)
		dyn.setBonus(traitTerritorialDefence, 1)
	case "Of Pawns and Kings":
		dyn.setBonus(charScheming, 1)
	case "The Love of the People":
		dyn.setBonus(valueMorale, 1)
	case "Order of Protectors":
		dyn.setBonus(traitTerritorialDefence, 1)
	case "Military Honour":
		dyn.setBonus(charMilitarism, 1)
	case "Interstellar Marriages":
		dyn.setBonus(traitCulture, 1)
		dyn.setBonus(traitFleet, 1)
	case "No Peers In Sight":
		dyn.add1d6ToValues()
	case "Clergy Scholars":
		dyn.threeAptitudesToLevel1()
	case "Knights and Templars":
		dyn.setBonus(traitTerritorialDefence, 1)
	case "Holy Treasures":
		dyn.setBonus(valueWealth, 1)
	case "Family Comes First":
		dyn.setBonus(valuePopulace, 1)
	case "Online Scripture":
		dyn.setBonus(traitTechnology, 1)
	case "Blessings from Beyond":
		dyn.setBonus(utils.RandomFromList(listTRAITS), 1)
		dyn.setBonus(utils.RandomFromList(listTRAITS), 1)
	case "Living Legends":
		if utils.RandomBool() {
			dyn.add1d6ToValues()
		} else {
			dyn.setBonus(utils.RandomFromList(characteristicsLIST()), 1)
			dyn.setBonus(utils.RandomFromList(characteristicsLIST()), 1)
		}
	case "Undeniable Success":
		dyn.threeAptitudesToLevel1()
	case "Art Thieves and Extortions":
		dyn.setBonus(traitFiscalDefence, 1)
	case "Pirate Captains":
		if utils.RandomBool() {
			dyn.setBonus(traitFleet, 1)
		} else {
			dyn.setBonus(charMilitarism, 1)
		}
	case "Gangs Upon Gangs":
		dyn.setBonus(valuePopulace, 1)
	case "Tougher than Street":
		dyn.setBonus(traitTerritorialDefence, 1)
		dyn.setBonus(traitTechnology, 1)
	case "Empire of Crime":
		dyn.setBonus(utils.RandomFromList(listTRAITS), 1)
		dyn.setBonus(utils.RandomFromList(listTRAITS), 1)
	case "Intergalactic Mafia":
		dyn.add1d6ToValues()
	}
}

func (dyn *Dynasty) add1d6ToValues() {
	r := utils.RollDice("d6")
	for i := 0; i < r; i++ {
		dyn.setBonus(utils.RandomFromList(valuesLIST()), 1)
	}
}

func (dyn *Dynasty) threeAptitudesToLevel1() {
	i := 3
	for i > 0 {
		apt := utils.RandomFromList(aptitudeLIST())
		if dyn.aptitudes[apt] < 1 {
			dyn.setBase(apt, 1)
			i--
		}
	}
}

func rollFirstGenBonus(archetype string) string {
	fgb := "Error"
	switch archetype {
	case archConglomerate:
		fgb = rollConglamerateFGB()
	case archMediaEmpire:
		fgb = rollMediaEmpireFGB()
	case archMerchantMarket:
		fgb = rollMerchantMarketFGB()
	case archMilitaryCharter:
		fgb = rollMilitaryCharterFGB()
	case archNobleLine:
		fgb = rollNobleLineFGB()
	case archReligiousFaith:
		fgb = rollReligiousFaithFGB()
	case archSyndicate:
		fgb = rollSyndicateFGB()
	}
	return fgb
}

func rollConglamerateFGB() string {
	r := utils.RollDice("2d6")
	if r < 3 {
		return "University Board Members"
	}
	if r < 5 {
		return "Monopoly"
	}
	if r < 7 {
		return "Shipyard Access"
	}
	if r < 8 {
		return "Noble Investors"
	}
	if r < 10 {
		return "Inherited Pride"
	}
	if r < 12 {
		return "Multi-Stellar Benefactor"
	}
	return "Royal Backing"
}

func rollMediaEmpireFGB() string {
	r := utils.RollDice("2d6")
	if r < 3 {
		return "Psionic Investigators"
	}
	if r < 5 {
		return "Pyramid Structure"
	}
	if r < 7 {
		return "High-Tech Communications"
	}
	if r < 8 {
		return "Noble Investors"
	}
	if r < 10 {
		return "Military Reporters"
	}
	if r < 12 {
		return "Interstellar Cover Story"
	}
	return "Total Media Monopoly"
}

func rollMerchantMarketFGB() string {
	r := utils.RollDice("2d6")
	if r < 3 {
		return "Barter Over Sales"
	}
	if r < 5 {
		return "Monopoly"
	}
	if r < 7 {
		return "Intense Collegiate Training"
	}
	if r < 8 {
		return "Patents Upon Patents"
	}
	if r < 10 {
		return "Govermental Acquisitions"
	}
	if r < 12 {
		return "A Republic in Good Fortune"
	}
	return "Perfect Economy"
}

func rollMilitaryCharterFGB() string {
	r := utils.RollDice("2d6")
	if r < 3 {
		return "Intense Generational Training"
	}
	if r < 5 {
		return "War Coffers"
	}
	if r < 7 {
		return "Naval Partners"
	}
	if r < 8 {
		return "An Armed Populance"
	}
	if r < 10 {
		return "Victory Over Invasion"
	}
	if r < 12 {
		return "War Colleges"
	}
	return "Noble Armada"
}

func rollNobleLineFGB() string {
	r := utils.RollDice("2d6")
	if r < 3 {
		return "Royal Guard"
	}
	if r < 5 {
		return "Of Pawns and Kings"
	}
	if r < 7 {
		return "The Love of the People"
	}
	if r < 8 {
		return "Order of Protectors"
	}
	if r < 10 {
		return "Military Honour"
	}
	if r < 12 {
		return "Interstellar Marriages"
	}
	return "No Peers In Sight"
}

func rollReligiousFaithFGB() string {
	r := utils.RollDice("2d6")
	if r < 3 {
		return "Clergy Scholars"
	}
	if r < 5 {
		return "Knights and Templars"
	}
	if r < 7 {
		return "Holy Treasures"
	}
	if r < 8 {
		return "Family Comes First"
	}
	if r < 10 {
		return "Online Scripture"
	}
	if r < 12 {
		return "Blessings from Beyond"
	}
	return "Living Legends"
}

func rollSyndicateFGB() string {
	r := utils.RollDice("2d6")
	if r < 3 {
		return "Undeniable Success"
	}
	if r < 5 {
		return "Art Thieves and Extortions"
	}
	if r < 7 {
		return "Pirate Captains"
	}
	if r < 8 {
		return "Gangs Upon Gangs"
	}
	if r < 10 {
		return "Tougher than Street"
	}
	if r < 12 {
		return "Empire of Crime"
	}
	return "Intergalactic Mafia"
}

func determineManagementAssets(archetype string) []string {
	var assets []string
	var possibleOptions []string
	switch archetype {
	case archConglomerate:
		possibleOptions = []string{
			mAssetHeroicLeaders,
			mAssetOverlord,
			mAssetBoardOfDirectors,
			mAssetCommandStaff,
		}
	case archMediaEmpire:
		possibleOptions = []string{
			mAssetOverlord,
			mAssetMatriarchPatriarch,
			mAssetBoardOfDirectors,
			mAssetTheocrat,
		}
	case archMerchantMarket:
		possibleOptions = []string{
			mAssetHeroicLeaders,
			mAssetMatriarchPatriarch,
			mAssetBoardOfDirectors,
			mAssetCommandStaff,
		}
	case archMilitaryCharter:
		possibleOptions = []string{
			mAssetTheocrat,
			mAssetHeroicLeaders,
			mAssetCommandStaff,
			mAssetOverlord,
		}
	case archNobleLine:
		possibleOptions = []string{
			mAssetHeroicLeaders,
			mAssetOverlord,
			mAssetMatriarchPatriarch,
			mAssetTheocrat,
		}
	case archReligiousFaith:
		possibleOptions = []string{
			mAssetCommandStaff,
			mAssetHeroicLeaders,
			mAssetTheocrat,
			mAssetOverlord,
		}
	case archSyndicate:
		possibleOptions = []string{
			mAssetCommandStaff,
			mAssetBoardOfDirectors,
			mAssetOverlord,
			mAssetMatriarchPatriarch,
		}
	}
	assets = rollManagementAsset(possibleOptions)
	return assets
}

func selectFromTable(possibleOptions []string, r int) string {
	switch r {
	default:
		return utils.RandomFromList(possibleOptions)
	case 3, 4:
		return possibleOptions[0]
	case 5, 6:
		return possibleOptions[1]
	case 7, 8, 9:
		return possibleOptions[2]
	case 10, 11:
		return possibleOptions[3]
	}
	return "Error"
}

func rollManagementAsset(possibleOptions []string) []string {
	var result []string
	r := utils.RollDice("2d6")
	result = utils.AppendUniqueStr(result, selectFromTable(possibleOptions, r))
	if r != 12 {
		return result
	}
	r = utils.RollDice("2d6")
	utils.AppendUniqueStr(result, selectFromTable(possibleOptions, r))
	return result
}

func (dyn *Dynasty) trainCharacteristics() {
	//выбери 3 характеристики
	//для каждой проведи тест 2d6 + DM против char
	var charsToTrain []string
	for len(charsToTrain) < 3 {
		charsToTrain = utils.AppendUniqueStr(charsToTrain, utils.RandomFromList(characteristicsLIST()))
	}
	//для каждой проведи тест 2d6 + DM против char
	for i := range charsToTrain {
		tn := dyn.characteristics[charsToTrain[i]]
		dmMod := charDM(dyn.characteristics[charsToTrain[i]])
		r := utils.RollDice("2d6", dmMod)
		switch r {
		default:
			if r > tn {
				dyn.setBonus(charsToTrain[i], 1)
			}
		case 2:
			dyn.setBonus(charsToTrain[i], -1)
		case 12:
			dyn.setBonus(charsToTrain[i], 1)
		}

	}
}

func (dyn *Dynasty) manageTraits() {
	managementAtempts := charDM(dyn.characteristics[charCleverness]) + 1

	fmt.Println("managementAtempts =", managementAtempts)
	for managementAtempts > 0 {
		managementAtempts--
		if utils.RandomBool() {
			canBeManaged := defineManagebleTraits(dyn)
			if len(canBeManaged) == 0 {
				continue
			}
			donor := utils.RandomFromList(canBeManaged)
			dyn.setBonus(donor, -1)
			reciver := utils.RandomFromList(listTRAITS)
			dyn.setBonus(reciver, 1)
			fmt.Println("Manage from", donor, "to", reciver)
		}
	}
}

func defineManagebleTraits(dyn *Dynasty) []string {
	var manageble []string
	for key := range dyn.traits {
		if dyn.traits[key] > 1 {
			manageble = append(manageble, key)
		}
	}
	return manageble
}

func (dyn *Dynasty) trainAptitudes() {
	//выбери 3 характеристики
	//для каждой проведи тест 2d6 + DM против char
	var aptsToTrain []string
	for len(aptsToTrain) < 5 {
		aptsToTrain = utils.AppendUniqueStr(aptsToTrain, utils.RandomFromList(aptitudeLIST()))
	}
	//для каждой проведи тест 2d6 + DM против char
	for i := range aptsToTrain {
		bonus := dyn.aptitudes[aptsToTrain[i]]
		if bonus < 0 {
			bonus = -2
		}

		r := utils.RollDice("2d6", bonus)
		effect := r - 8
		fmt.Println("train", aptsToTrain[i], r, bonus, effect, dyn.aptitudes[aptsToTrain[i]])
		if effect > dyn.aptitudes[aptsToTrain[i]] {
			dyn.setBonus(aptsToTrain[i], 1)
			fmt.Println("train successful")
		}
	}
}

func (dyn *Dynasty) finalValuesAdjustmen() {
	dyn.setBonus(valueMorale, dyn.characteristics[charLoyalty]+charDM(dyn.characteristics[charPopularity])+dyn.traits[traitCulture])
	dyn.setBonus(valuePopulace, dyn.characteristics[charTenacity]+charDM(dyn.characteristics[charTradition])+dyn.traits[traitTechnology])
	dyn.setBonus(valueWealth, dyn.characteristics[charGreed]+charDM(dyn.characteristics[charCleverness])+dyn.traits[traitFiscalDefence])
	vals := valuesLIST()
	for i := range vals {
		if utils.RandomBool() {
			donor := vals[i]
			dyn.setBonus(donor, -1)
			rcvr := utils.RandomFromList(valuesLIST())
			dyn.setBonus(rcvr, 1)
			fmt.Println("Adjust", donor, "to", rcvr)
		}
	}
}

func (dyn *Dynasty) toString() string {
	str := "Main Characteristic: " + dyn.mainCharacteristic + "\n"
	str += "Characteristics:" + "\n"
	charList := characteristicsLIST()
	for i := range charList {
		str += charList[i] + ": " + strconv.Itoa(dyn.characteristics[charList[i]]) + "\n"
	}
	str += "Traits:" + "\n"
	traits := listTRAITS
	for j := range traits {
		str += traits[j] + ": " + strconv.Itoa(dyn.traits[traits[j]]) + "\n"
	}
	str += "Aptitude:" + "\n"
	aptitude := aptitudeLIST()
	for k := range aptitude {
		str += aptitude[k] + ": " + strconv.Itoa(dyn.aptitudes[aptitude[k]]) + "\n"
	}
	str += "Values:" + "\n"
	value := valuesLIST()
	for k := range value {
		str += value[k] + ": " + strconv.Itoa(dyn.values[value[k]]) + "\n"
	}
	// for k := range dyn.characteristics {
	// 	str += k + " " + strconv.Itoa(dyn.characteristics[k].statVal()) + "\n"
	// }
	str += "Power Base:  " + dyn.powerBase + "\n"
	str += "Archetype:  " + dyn.archetype + "\n"
	for i := range dyn.boons {
		str += "Boon:  " + dyn.boons[i].name + "\n"
	}
	for i := range dyn.managementAssets {
		str += "Management Asset:  " + dyn.managementAssets[i] + "\n"
	}
	return str
}

func (dyn *Dynasty) boonsAndHinders() {
	b := utils.RollDice("d3", -1)
	h := utils.RollDice("d3", -1)
	fmt.Println("boons", b, "hinders", h)

	//dyn.boons = append(dyn.boons, *BoonHinderByName["Commercial Psions"])
	bhMap := BoonHinderByName()
	for k, v := range bhMap {
		if b == 0 {
			break
		}
		if v.archetype == dyn.archetype && v.bType == "Boon" && canPay(dyn, v.cost) {
			dyn.boons = append(dyn.boons, bhMap[k])
			b--
		}
	}
	for k, v := range bhMap {
		if h == 0 {
			break
		}
		if v.archetype == dyn.archetype && v.bType == "Hinder" && canPay(dyn, v.cost) {
			dyn.boons = append(dyn.boons, bhMap[k])
			h--
		}
	}
}

func (dyn *Dynasty) determineBaseTraits() {
	switch dyn.archetype {
	case archConglomerate:
		dyn.setBase(traitCulture, charDM(dyn.characteristics[charGreed])+charDM(dyn.characteristics[charTradition]))
		dyn.setBase(traitFiscalDefence, charDM(dyn.characteristics[charGreed])+charDM(dyn.characteristics[charTenacity])+1)
		dyn.setBase(traitFleet, charDM(dyn.characteristics[charMilitarism])+1)
		dyn.setBase(traitTechnology, charDM(dyn.characteristics[charGreed])+charDM(dyn.characteristics[charLoyalty]))
		dyn.setBase(traitTerritorialDefence, charDM(dyn.characteristics[charLoyalty])+charDM(dyn.characteristics[charPopularity]))
	case archMediaEmpire:
		dyn.setBase(traitCulture, charDM(dyn.characteristics[charPopularity])+charDM(dyn.characteristics[charTradition]))
		dyn.setBase(traitFiscalDefence, charDM(dyn.characteristics[charLoyalty])+2)
		dyn.setBase(traitFleet, charDM(dyn.characteristics[charMilitarism])+1)
		dyn.setBase(traitTechnology, charDM(dyn.characteristics[charGreed])+charDM(dyn.characteristics[charPopularity])+1)
		dyn.setBase(traitTerritorialDefence, charDM(dyn.characteristics[charCleverness])+charDM(dyn.characteristics[charLoyalty]))
	case archMerchantMarket:
		dyn.setBase(traitCulture, charDM(dyn.characteristics[charGreed])+charDM(dyn.characteristics[charPopularity]))
		dyn.setBase(traitFiscalDefence, charDM(dyn.characteristics[charGreed])+charDM(dyn.characteristics[charLoyalty])+1)
		dyn.setBase(traitFleet, charDM(dyn.characteristics[charLoyalty])+charDM(dyn.characteristics[charMilitarism]))
		dyn.setBase(traitTechnology, charDM(dyn.characteristics[charCleverness])+charDM(dyn.characteristics[charTradition])+1)
		dyn.setBase(traitTerritorialDefence, charDM(dyn.characteristics[charLoyalty])+2)
	case archMilitaryCharter:
		dyn.setBase(traitCulture, charDM(dyn.characteristics[charTradition])+1)
		dyn.setBase(traitFiscalDefence, charDM(dyn.characteristics[charGreed])+charDM(dyn.characteristics[charMilitarism]))
		dyn.setBase(traitFleet, charDM(dyn.characteristics[charMilitarism])+charDM(dyn.characteristics[charTenacity])+1)
		dyn.setBase(traitTechnology, charDM(dyn.characteristics[charMilitarism])+1)
		dyn.setBase(traitTerritorialDefence, charDM(dyn.characteristics[charMilitarism])+charDM(dyn.characteristics[charPopularity])+1)
	case archNobleLine:
		dyn.setBase(traitCulture, charDM(dyn.characteristics[charLoyalty])+charDM(dyn.characteristics[charTradition])+2)
		dyn.setBase(traitFiscalDefence, charDM(dyn.characteristics[charGreed])+charDM(dyn.characteristics[charTenacity]))
		dyn.setBase(traitFleet, charDM(dyn.characteristics[charMilitarism])+1)
		dyn.setBase(traitTechnology, charDM(dyn.characteristics[charTenacity])+1)
		dyn.setBase(traitTerritorialDefence, charDM(dyn.characteristics[charLoyalty])+charDM(dyn.characteristics[charMilitarism])+1)
	case archReligiousFaith:
		dyn.setBase(traitCulture, charDM(dyn.characteristics[charLoyalty])+charDM(dyn.characteristics[charTradition])+2)
		dyn.setBase(traitFiscalDefence, charDM(dyn.characteristics[charGreed])+1)
		dyn.setBase(traitFleet, charDM(dyn.characteristics[charLoyalty])+1)
		dyn.setBase(traitTechnology, charDM(dyn.characteristics[charCleverness])+charDM(dyn.characteristics[charTenacity]))
		dyn.setBase(traitTerritorialDefence, charDM(dyn.characteristics[charLoyalty])+charDM(dyn.characteristics[charMilitarism])+1)
	case archSyndicate:
		dyn.setBase(traitCulture, charDM(dyn.characteristics[charGreed])+charDM(dyn.characteristics[charScheming]))
		dyn.setBase(traitFiscalDefence, charDM(dyn.characteristics[charLoyalty])+1)
		dyn.setBase(traitFleet, charDM(dyn.characteristics[charMilitarism])+charDM(dyn.characteristics[charScheming]))
		dyn.setBase(traitTechnology, charDM(dyn.characteristics[charMilitarism])+2)
		dyn.setBase(traitTerritorialDefence, charDM(dyn.characteristics[charLoyalty])+charDM(dyn.characteristics[charMilitarism])+1)
	}
}

func (dyn *Dynasty) gainBaseAptitudes() {
	switch dyn.archetype {
	case archConglomerate:
		dyn.setBase(apttAcquisition, 1)
		dyn.setBase("Bureaucracy", 1)
		dyn.setBase("Conquest", -1)
		dyn.setBase("Economics", 2)
		dyn.setBase("Entertain", -1)
		dyn.setBase("Expression", -1)
		dyn.setBase("Hostility", -1)
		dyn.setBase("Illicit", -1)
		dyn.setBase("Intel", 0)
		dyn.setBase("Maintenance", 0)
		dyn.setBase("Politics", -1)
		dyn.setBase("Posturing", -1)
		dyn.setBase("Propaganda", 0)
		dyn.setBase("Public Relations", 0)
		dyn.setBase("Recruitment", 0)
		dyn.setBase("Research", -1)
		dyn.setBase("Sabotage", -1)
		dyn.setBase("Security", -1)
		dyn.setBase("Tactical", -1)
		dyn.setBase("Tutelage", 1)
	case archMediaEmpire:
		dyn.setBase(apttAcquisition, -1)
		dyn.setBase("Bureaucracy", -1)
		dyn.setBase("Conquest", -1)
		dyn.setBase("Economics", 0)
		dyn.setBase("Entertain", 1)
		dyn.setBase("Expression", 0)
		dyn.setBase("Hostility", -1)
		dyn.setBase("Illicit", 0)
		dyn.setBase("Intel", 1)
		dyn.setBase("Maintenance", -1)
		dyn.setBase("Politics", 0)
		dyn.setBase("Posturing", 0)
		dyn.setBase("Propaganda", 2)
		dyn.setBase("Public Relations", 1)
		dyn.setBase("Recruitment", -1)
		dyn.setBase("Research", -1)
		dyn.setBase("Sabotage", -1)
		dyn.setBase("Security", -1)
		dyn.setBase("Tactical", -1)
		dyn.setBase("Tutelage", -1)
	case archMerchantMarket:
		dyn.setBase(apttAcquisition, 1)
		dyn.setBase("Bureaucracy", 0)
		dyn.setBase("Conquest", -1)
		dyn.setBase("Economics", 2)
		dyn.setBase("Entertain", -1)
		dyn.setBase("Expression", 0)
		dyn.setBase("Hostility", -1)
		dyn.setBase("Illicit", -1)
		dyn.setBase("Intel", 0)
		dyn.setBase("Maintenance", -1)
		dyn.setBase("Politics", -1)
		dyn.setBase("Posturing", -1)
		dyn.setBase("Propaganda", 1)
		dyn.setBase("Public Relations", 1)
		dyn.setBase("Recruitment", 0)
		dyn.setBase("Research", 0)
		dyn.setBase("Sabotage", -1)
		dyn.setBase("Security", -1)
		dyn.setBase("Tactical", -1)
		dyn.setBase("Tutelage", -1)
	case archMilitaryCharter:
		dyn.setBase(apttAcquisition, -1)
		dyn.setBase("Bureaucracy", -1)
		dyn.setBase("Conquest", 1)
		dyn.setBase("Economics", -1)
		dyn.setBase("Entertain", -1)
		dyn.setBase("Expression", -1)
		dyn.setBase("Hostility", -1)
		dyn.setBase("Illicit", 0)
		dyn.setBase("Intel", 1)
		dyn.setBase("Maintenance", 0)
		dyn.setBase("Politics", 0)
		dyn.setBase("Posturing", -1)
		dyn.setBase("Propaganda", 1)
		dyn.setBase("Public Relations", -1)
		dyn.setBase("Recruitment", 1)
		dyn.setBase("Research", -1)
		dyn.setBase("Sabotage", -1)
		dyn.setBase("Security", 0)
		dyn.setBase("Tactical", 2)
		dyn.setBase("Tutelage", -1)
	case archNobleLine:
		dyn.setBase(apttAcquisition, -1)
		dyn.setBase("Bureaucracy", 1)
		dyn.setBase("Conquest", -1)
		dyn.setBase("Economics", 0)
		dyn.setBase("Entertain", 0)
		dyn.setBase("Expression", 1)
		dyn.setBase("Hostility", -1)
		dyn.setBase("Illicit", 0)
		dyn.setBase("Intel", -1)
		dyn.setBase("Maintenance", -1)
		dyn.setBase("Politics", 2)
		dyn.setBase("Posturing", -1)
		dyn.setBase("Propaganda", -1)
		dyn.setBase("Public Relations", -1)
		dyn.setBase("Recruitment", 0)
		dyn.setBase("Research", -1)
		dyn.setBase("Sabotage", -1)
		dyn.setBase("Security", 0)
		dyn.setBase("Tactical", -1)
		dyn.setBase("Tutelage", 1)
	case archReligiousFaith:
		dyn.setBase(apttAcquisition, -1)
		dyn.setBase("Bureaucracy", -1)
		dyn.setBase("Conquest", 0)
		dyn.setBase("Economics", -1)
		dyn.setBase("Entertain", 0)
		dyn.setBase("Expression", 1)
		dyn.setBase("Hostility", -1)
		dyn.setBase("Illicit", -1)
		dyn.setBase("Intel", -1)
		dyn.setBase("Maintenance", -1)
		dyn.setBase("Politics", 0)
		dyn.setBase("Posturing", -1)
		dyn.setBase("Propaganda", 1)
		dyn.setBase("Public Relations", -1)
		dyn.setBase("Recruitment", 2)
		dyn.setBase("Research", -1)
		dyn.setBase("Sabotage", -1)
		dyn.setBase("Security", 0)
		dyn.setBase("Tactical", -1)
		dyn.setBase("Tutelage", -1)
	case archSyndicate:
		dyn.setBase(apttAcquisition, -1)
		dyn.setBase("Bureaucracy", -1)
		dyn.setBase("Conquest", 1)
		dyn.setBase("Economics", -1)
		dyn.setBase("Entertain", 0)
		dyn.setBase("Expression", 0)
		dyn.setBase("Hostility", 0)
		dyn.setBase("Illicit", 2)
		dyn.setBase("Intel", 0)
		dyn.setBase("Maintenance", -1)
		dyn.setBase("Politics", -1)
		dyn.setBase("Posturing", 1)
		dyn.setBase("Propaganda", -1)
		dyn.setBase("Public Relations", -1)
		dyn.setBase("Recruitment", -1)
		dyn.setBase("Research", -1)
		dyn.setBase("Sabotage", 1)
		dyn.setBase("Security", 0)
		dyn.setBase("Tactical", -1)
		dyn.setBase("Tutelage", -1)
	}
}

func characteristicsMAP() map[string]int {
	chaMap := make(map[string]int)
	for _, val := range characteristicsLIST() {
		chaMap[val] = -1
	}
	return chaMap
}

func charDM(char int) int {
	if char < 1 {
		return -3
	}
	if char < 3 {
		return -2
	}
	if char < 6 {
		return -1
	}
	if char < 9 {
		return 0
	}
	if char < 12 {
		return 1
	}
	if char < 15 {
		return 2
	}
	if char < 18 {
		return 3
	}
	if char < 21 {
		return 4
	}
	return 5
}

func (dyn *Dynasty) applyEffectPowerBase() {
	switch dyn.powerBase {
	case "Colony/Setlement":
		dyn.setBonus(traitCulture, 1)
		dyn.setBonus(traitTerritorialDefence, -1)
		dyn.setBonus("Expression", 1)
		dyn.setBonus("Recruitment", 1)
		dyn.setBonus("Maintenance", 1)
		dyn.setBonus("Propaganda", 1)
		dyn.setBonus("Tutelage", 1)
	case "Conflict Zone":
		dyn.setBonus(traitTerritorialDefence, 2)
		dyn.setBonus(traitFiscalDefence, -1)
		dyn.setBonus(traitFleet, -1)
		dyn.setBonus("Hostility", 2)
		dyn.setBonus("Posturing", 1)
		dyn.setBonus("Security", 1)
		dyn.setBonus("Tactical", 1)
	case "Megalopolis":
		dyn.setBonus(traitFiscalDefence, 1)
		dyn.setBonus(traitTechnology, 1)
		dyn.setBonus(traitCulture, -2)
		dyn.setBonus("Bureaucracy", 2)
		dyn.setBonus("Economics", 1)
		dyn.setBonus("Public Relations", 1)
		dyn.setBonus("Research", 1)
	case "Military Compound":
		dyn.setBonus(traitTerritorialDefence, 2)
		dyn.setBonus(traitFleet, 1)
		dyn.setBonus(traitFiscalDefence, -2)
		dyn.setBonus("Conquest", 2)
		dyn.setBonus("Tactical", 2)
		dyn.setBonus("Politics", 1)
		dyn.setBonus("Posturing", 1)
		dyn.setBonus("Security", 1)
	case "Noble Estate":
		dyn.setBonus(traitCulture, 1)
		dyn.setBonus(traitFiscalDefence, 1)
		dyn.setBonus(traitTerritorialDefence, -2)
		dyn.setBonus(traitFleet, -1)
		dyn.setBonus("Bureaucracy", 2)
		dyn.setBonus("Politics", 2)
		dyn.setBonus("Expression", 1)
		dyn.setBonus("Posturing", 1)
		dyn.setBonus("Security", 1)
	case "Starship/Flotilla":
		dyn.setBonus(traitFleet, 2)
		dyn.setBonus(traitTechnology, 1)
		dyn.setBonus(traitTerritorialDefence, -2)
		dyn.setBonus("Intel", 2)
		dyn.setBonus("Conquest", 1)
		dyn.setBonus("Economics", 1)
		dyn.setBonus("Maintenance", 1)
		dyn.setBonus("Posturing", 1)
		dyn.setBonus("Research", 1)
		dyn.setBonus("Tactical", 1)
	case "Temple/Holy Land":
		dyn.setBonus(traitCulture, 2)
		dyn.setBonus(traitTechnology, -2)
		dyn.setBonus("Expression", 2)
		dyn.setBonus("Recruitment", 2)
		dyn.setBonus("Maintenance", 1)
		dyn.setBonus("Propaganda", 1)
		dyn.setBonus("Public Relations", 1)
		dyn.setBonus("Tutelage", 1)
	case "Uncharted Wilderness":
		dyn.setBonus(traitTerritorialDefence, 2)
		dyn.setBonus(traitTechnology, -1)
		dyn.setBonus("Security", 2)
		dyn.setBonus("Entertain", 1)
		dyn.setBonus("Illicit", 1)
		dyn.setBonus("Security", 1)
	case "Underworld Slum":
		dyn.setBonus(traitFiscalDefence, 1)
		dyn.setBonus(traitTerritorialDefence, 1)
		dyn.setBonus(traitCulture, -2)
		dyn.setBonus("Illicit", 2)
		dyn.setBonus("Sabotage", 2)
		dyn.setBonus("Entertain", 1)
		dyn.setBonus("Intel", 1)
		dyn.setBonus("Posturing", 1)
		dyn.setBonus("Security", 1)
	case "Urban Offices":
		dyn.setBonus(traitCulture, 1)
		dyn.setBonus(traitFiscalDefence, 1)
		dyn.setBonus(traitFleet, -1)
		dyn.setBonus(apttAcquisition, 2)
		dyn.setBonus("Economics", 2)
		dyn.setBonus("Bureaucracy", 1)
		dyn.setBonus("Intel", 1)
		dyn.setBonus("Public Relations", 1)
		dyn.setBonus("Tutelage", 1)
	}
}

///////////////////////////LISTS

func validKey(key string) bool {
	chars := characteristicsLIST()
	for i := range chars {
		if chars[i] == key {
			return true
		}
	}
	trts := listTRAITS
	for i := range trts {
		if trts[i] == key {
			return true
		}
	}
	aptList := aptitudeLIST()
	for i := range aptList {
		if aptList[i] == key {
			return true
		}
	}
	valList := valuesLIST()
	for i := range valList {
		if valList[i] == key {
			return true
		}
	}
	return false
}

func characteristicsLIST() []string {
	chaList := []string{
		charCleverness,
		charGreed,
		charLoyalty,
		charMilitarism,
		charPopularity,
		charScheming,
		charTenacity,
		charTradition,
	}
	return chaList
}

func traitsLIST() []string {
	trtsList := []string{
		traitCulture,
		traitFiscalDefence,
		traitFleet,
		traitTechnology,
		traitTerritorialDefence,
	}
	return trtsList
}

func valuesLIST() []string {
	trtsList := []string{
		valueMorale,
		valuePopulace,
		valueWealth,
	}
	return trtsList
}

func managementAssetLIST() []string {
	trtsList := []string{
		mAssetBoardOfDirectors,
		mAssetCommandStaff,
		mAssetHeroicLeaders,
		mAssetMatriarchPatriarch,
		mAssetOverlord,
		mAssetTheocrat,
	}
	return trtsList
}

func aptitudeLIST() []string {
	aptList := []string{
		apttAcquisition,
		"Bureaucracy",
		"Conquest",
		"Economics",
		"Entertain",
		"Expression",
		"Hostility",
		"Illicit",
		"Intel",
		"Maintenance",
		"Politics",
		"Posturing",
		"Propaganda",
		"Public Relations",
		"Recruitment",
		"Research",
		"Sabotage",
		"Security",
		"Tactical",
		"Tutelage",
	}
	return aptList
}

func powerBaseList() []string {
	pb := []string{
		"Colony/Setlement",
		"Conflict Zone",
		"Megalopolis",
		"Military Compound",
		"Noble Estate",
		"Starship/Flotilla",
		"Temple/Holy land",
		"Uncharted Wilderness",
		"Underworld Slum",
		"Urban Offices",
	}
	return pb
}

func archetypeLIST() []string {
	al := []string{
		archConglomerate,
		archMediaEmpire,
		archMerchantMarket,
		archMilitaryCharter,
		archNobleLine,
		archReligiousFaith,
		archSyndicate,
	}
	return al
}

func (dyn *Dynasty) qualifyForArchetype(arch string) bool {
	switch arch {
	case archConglomerate:
		if dyn.characteristics[charGreed] < 8 || dyn.characteristics[charPopularity] < 6 || dyn.characteristics[charTenacity] < 5 {
			return false
		}
	case archMediaEmpire:
		if dyn.characteristics[charCleverness] < 6 || dyn.characteristics[charPopularity] < 8 || dyn.characteristics[charScheming] < 5 {
			return false
		}
	case archMerchantMarket:
		if dyn.characteristics[charCleverness] < 6 || dyn.characteristics[charGreed] < 8 || dyn.characteristics[charPopularity] < 5 {
			return false
		}
	case archMilitaryCharter:
		if dyn.characteristics[charLoyalty] < 5 || dyn.characteristics[charMilitarism] < 8 || dyn.characteristics[charTradition] < 6 {
			return false
		}
	case archNobleLine:
		if dyn.characteristics[charCleverness] < 6 || dyn.characteristics[charPopularity] < 8 || dyn.characteristics[charScheming] < 5 {
			return false
		}
	case archReligiousFaith:
		if dyn.characteristics[charLoyalty] < 8 || dyn.characteristics[charPopularity] < 5 || dyn.characteristics[charTradition] < 6 {
			return false
		}
	case archSyndicate:
		if dyn.characteristics[charGreed] < 6 || dyn.characteristics[charScheming] < 8 || dyn.characteristics[charTenacity] < 5 {
			return false
		}
	}
	return true
}

func (dyn *Dynasty) setBonus(key string, changeVal int) {
	if !validKey(key) {
		panic("Not valid key: " + key)
	}
	//fmt.Println("Bonus:", key, changeVal)
	for i := range characteristicsLIST() {
		if characteristicsLIST()[i] == key {
			dyn.characteristics[key] = dyn.characteristics[key] + changeVal
		}
	}
	for i := range listTRAITS {
		if listTRAITS[i] == key {
			dyn.traits[key] = dyn.traits[key] + changeVal
		}
	}
	for i := range aptitudeLIST() {
		if aptitudeLIST()[i] == key {
			dyn.aptitudes[key] = dyn.aptitudes[key] + changeVal
		}
	}
	for i := range valuesLIST() {
		if valuesLIST()[i] == key {
			dyn.values[key] = dyn.values[key] + changeVal
		}
	}
}

func (dyn *Dynasty) setBase(key string, val int) {
	if !validKey(key) {
		panic("Not valid key: " + key)
	}
	for i := range characteristicsLIST() {
		if characteristicsLIST()[i] == key {
			dyn.characteristics[key] = val
		}
	}
	for i := range listTRAITS {
		if listTRAITS[i] == key {
			dyn.traits[key] = val
		}
	}
	for i := range aptitudeLIST() {
		if aptitudeLIST()[i] == key {
			dyn.aptitudes[key] = val
		}
	}
	for i := range valuesLIST() {
		if valuesLIST()[i] == key {
			dyn.values[key] = val
		}
	}
}

func (dyn *Dynasty) pickVal(char string) int {
	chars := characteristicsLIST()
	for i := range chars {
		if chars[i] == char {
			return dyn.characteristics[char]
		}
	}

	for i := range listTRAITS {
		if listTRAITS[i] == char {
			return dyn.traits[char]
		}
	}
	aptList := aptitudeLIST()
	for i := range aptList {
		if aptList[i] == char {
			return dyn.aptitudes[char]
		}
	}
	return -999
}

func canPay(dyn *Dynasty, cost []boonCost) bool {
	for i := range cost {
		stat := cost[i].char
		qty := cost[i].qty
		dynCopy := dyn
		dynCopy.setBonus(stat, qty)
		if dynCopy.pickVal(stat) < 0 {
			return false
		}
	}
	return true
}

func (dyn *Dynasty) applyBoonEffects() {
	for i := range dyn.boons {
		cost := dyn.boons[i].cost
		for j := range cost {
			dyn.setBonus(cost[j].char, cost[j].qty)
		}
	}
}

func randomEventCode(archetype string) string {
	code := ""
	switch archetype {
	default:
		code = "0"
	case archConglomerate:
		code = "1"
	case archMediaEmpire:
		code = "2"
	case archMerchantMarket:
		code = "3"
	case archMilitaryCharter:
		code = "4"
	case archNobleLine:
		code = "5"
	case archReligiousFaith:
		code = "6"
	case archSyndicate:
		code = "7"
	}
	code = code + rollD66()
	return code
}

func rollD66() string {
	d1 := convert.ItoS(utils.RollDice("d6"))
	d2 := convert.ItoS(utils.RollDice("d6"))
	return d1 + d2
}

func (dyn *Dynasty) RollHistoricEvent() {
	r := utils.RollDice("2d6", dyn.misc["historicEvents"])
	eventDescr := ""
	eventOutcome := ""
	switch r {
	case 2:
		//There is an interstellar war between planetary forces, sweeping them into the dangerous realm of battles and destruction.
		//The Dynasty must roll Conquest, Hostility and Security 8+ each; each successful check adds +1 to all Traits, each failure reduces all Values by –1.
		eventDescr = "There is an interstellar war between planetary forces, sweeping them into the dangerous realm of battles and destruction."
		suc1 := dyn.failureCheck(apttConquest, 8)
		suc2 := dyn.failureCheck(apttHostility, 8)
		suc3 := dyn.failureCheck(apttSecurity, 8)
		tSuc := 0
		tFail := 0
		if suc1 {
			tSuc++
		} else {
			tFail++
		}
		if suc2 {
			tSuc++
		} else {
			tFail++
		}
		if suc3 {
			tSuc++
		} else {
			tFail++
		}
		eventOutcome = "(+" + convert.ItoS(tSuc) + " to all Traits, -" + convert.ItoS(tFail) + " to all Values)"
		for i := range listVALUES {
			dyn.setBonus(listVALUES[i], -1)
		}
		for i := range listTRAITS {
			dyn.setBonus(listTRAITS[i], 1)
		}
	case 3, 4:
		//A consolidation of enemies have targeted the Dynasty and are coming at them from all directions.
		//The Dynasty must roll Intel, Posturing and Security 8+;
		//each successful check adds +1 to all Values, each failure reduces Fiscal Defence, Territorial Defence and Fleet Traits by –1 each.
		eventDescr = "A consolidation of enemies have targeted the Dynasty and are coming at them from all directions."
		suc1 := dyn.failureCheck(apttIntel, 8)
		suc2 := dyn.failureCheck(apttPosturing, 8)
		suc3 := dyn.failureCheck(apttSecurity, 8)
		tSuc := 0
		tFail := 0
		if suc1 {
			tSuc++
		} else {
			tFail++
		}
		if suc2 {
			tSuc++
		} else {
			tFail++
		}
		if suc3 {
			tSuc++
		} else {
			tFail++
		}
		eventOutcome = "(+" + convert.ItoS(tSuc) + " to all Values, -" + convert.ItoS(tFail) + " to Fiscal Defence, Territorial Defence and Fleet Traits)"
		down := []string{traitFiscalDefence, traitFleet, traitTerritorialDefence}
		for i := range down {
			dyn.setBonus(down[i], -1)
		}
		for i := range listVALUES {
			dyn.setBonus(listVALUES[i], 1)
		}
	case 5, 6:
		//One of the Dynasty’s inner members is given an opportunity to do something truly amazing – and does.
		//The Dynasty can increase any two Characteristics by +1, as well as +1 to its Morale Value.
		eventDescr = "One of the Dynasty’s inner members is given an opportunity to do something truly amazing – and does."
		var bonuses []string
		for len(bonuses) < 2 {
			bonuses = utils.AppendUniqueStr(bonuses, utils.RandomFromList(listCHARS))
		}
		dyn.setBonus(bonuses[0], 1)
		dyn.setBonus(bonuses[1], 1)
		dyn.setBonus(valueMorale, 1)
		eventOutcome = "(+1 " + bonuses[0] + ", +1 " + bonuses[1] + ", +1 Morale)"
	case 7:
		//A long stretch of time without any conflicts has given the Dynasty the perfect opportunity to focus on growth, expansion and self-indulgence.
		//The Dynasty can add +1 to any single Characteristic, +1 to any Trait and +1 to any two Aptitudes.
		eventDescr = "A long stretch of time without any conflicts has given the Dynasty the perfect opportunity to focus on growth, expansion and self-indulgence."
		var bonuses []string
		for len(bonuses) < 2 {
			bonuses = utils.AppendUniqueStr(bonuses, utils.RandomFromList(listAPTITUDES))
		}
		dyn.setBonus(bonuses[0], 1)
		dyn.setBonus(bonuses[1], 1)
		trat := utils.RandomFromList(listTRAITS)
		chr := utils.RandomFromList(listCHARS)
		dyn.setBonus(bonuses[0], 1)
		dyn.setBonus(bonuses[1], 1)
		dyn.setBonus(trat, 1)
		dyn.setBonus(chr, 1)
		eventOutcome = "(+1 " + bonuses[0] + ", +1 " + bonuses[1] + ", +1 " + chr + ", +1 " + trat + ")"
	case 8, 9:
		//Travellers from beyond Chartered Space have come into the Dynasty’s territory and are setting up a lifelong colony with their help.
		//The Dynasty may roll a Public Relations Aptitude check modified by their Loyalty DM. For every point the result is over 7,
		//add +1 to any Trait or Value (but not more than +1 per Trait or Value).
		eventDescr = "Travellers from beyond Chartered Space have come into the Dynasty’s territory and are setting up a lifelong colony with their help."
		checkEffect := dyn.aptitideCheck(apttPublicRelations, charLoyalty, -1)
		if checkEffect < 0 {
			checkEffect = 0
			eventOutcome = "No benefits was received"
		}
		var bonuses []string
		possib := listTRAITS
		possib = append(possib, listVALUES...)
		for len(bonuses) < checkEffect {
			bonuses = utils.AppendUniqueStr(bonuses, utils.RandomFromList(possib))
		}
		for i := range bonuses {
			dyn.setBonus(bonuses[i], 1)
			eventOutcome += "+1 " + bonuses[i] + " "
		}
	case 10, 11:
		//A major powerhouse in the interstellar politicking circles vanishes suddenly, leaving an easily filled vacuum that the Dynasty can take advantage of.
		//The Dynasty may roll an Acquisition Aptitude check modified by their Greed DM. For every point the result is over 7,
		//add +1 to any Aptitude, Trait or Value (but not more than +1 per Aptitude, Trait or Value).
		eventDescr = "A major powerhouse in the interstellar politicking circles vanishes suddenly, leaving an easily filled vacuum that the Dynasty can take advantage of."
		checkEffect := dyn.aptitideCheck(apttAcquisition, charGreed, -1)
		if checkEffect < 0 {
			checkEffect = 0
			eventOutcome = "No benefits was received"
		}
		var bonuses []string
		possib := listTRAITS
		possib = append(possib, listVALUES...)
		possib = append(possib, listAPTITUDES...)
		for len(bonuses) < checkEffect {
			bonuses = utils.AppendUniqueStr(bonuses, utils.RandomFromList(possib))
		}
		for i := range bonuses {
			dyn.setBonus(bonuses[i], 1)
			eventOutcome += "+1 " + bonuses[i] + " "
		}
	case 12:
		//A former enemy or rival suddenly sets aside its differences and comes forward to the Dynasty and wants to be steadfast allies.
		//The Dynasty may roll Cleverness or Scheming 9+ to ensure the alliance is real before committing to it.
		//If successful, the Dynasty can add +1 to all Traits and add +1 Populace. If failed, the Dynasty adds +1 to Hostility instead.
		eventDescr = "A former enemy or rival suddenly sets aside its differences and comes forward to the Dynasty and wants to be steadfast allies."
		eventChar := dyn.safestOption(9, charCleverness, charScheming)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		if eventSuccess {
			dyn.setBonus(valuePopulace, 1)
			for i := range listTRAITS {
				dyn.setBonus(listTRAITS[i], 1)
			}
			eventOutcome = "Negotiations was massive success! (+1 to all Traits, +1 Populance)"
		} else {
			eventOutcome = "Negotiations failed (+1 Hostility)"
			dyn.setBonus(apttHostility, 1)
		}
	default:
		//A truly ancient and powerful being with abilities bordering on ‘magic’ comes forward to lend aid to the Dynasty for its own mysterious reasons.
		//The Dynasty rolls 1d6 and consults the following:
		eventDescr = "A truly ancient and powerful being with abilities bordering on 'magic' comes forward to lend aid to the Dynasty for its own mysterious reasons."
		r := utils.RollDice("d6")
		//1 – Gain +1 to all Traits.
		//2 – Gain +1 to all Values.
		//3 – Gain +1 to any three Aptitudes.
		//4 – Gain +1 to any Characteristic.
		//5 – Gain +2 to any Characteristic.
		//6 – Gain +3 to any Characteristic.
		switch r {
		case 1:
			for i := range listTRAITS {
				dyn.setBonus(listTRAITS[i], 1)
			}
			eventOutcome = "(+1 to all Traits)"
		case 2:
			for i := range listVALUES {
				dyn.setBonus(listVALUES[i], 1)
			}
			eventOutcome = "(+1 to all Values)"
		case 3:
			var bonus []string
			for len(bonus) < 3 {
				bonus = utils.AppendUniqueStr(bonus, utils.RandomFromList(listAPTITUDES))
			}
			for i := range bonus {
				dyn.setBonus(bonus[i], 1)
			}
			eventOutcome = "(+1 " + bonus[0] + ", +1 " + bonus[1] + ", +1 " + bonus[2] + ")"
		case 4:
			bonus := utils.RandomFromList(listCHARS)
			dyn.setBonus(bonus, 1)
			eventOutcome = "(+1 " + bonus + ")"
		case 5:
			bonus := utils.RandomFromList(listCHARS)
			dyn.setBonus(bonus, 2)
			eventOutcome = "(+2 " + bonus + ")"
		case 6:
			bonus := utils.RandomFromList(listCHARS)
			dyn.setBonus(bonus, 3)
			eventOutcome = "(+3 " + bonus + ")"
		}

	}
	eventLog := eventDescr + " " + eventOutcome
	fmt.Println(eventLog)
	dyn.misc["historicEvents"] = dyn.misc["historicEvents"] + 1
}

func (dyn *Dynasty) doBackGroundEvent(eventCode string) string {
	fmt.Println("Do event:", eventCode)
	//eventCode = "266"
	eventDescr := ""
	eventOutcome := ""
	fmt.Println("Event:", eventCode)
	switch eventCode {
	default:
		fmt.Println(eventCode + " not defined")
	///CONGLOMERATE
	case "111":
		eventDescr = "Stocks are falling all over the galaxy for years."
		eventSuccess := dyn.failureCheck(charGreed, 8)
		eventOutcome = "Conglomerate averted damage."
		if !eventSuccess {
			dyn.setBonus(valueWealth, -1)
			eventOutcome = "Conglomerate lost 1 point of Wealth."
		}

	case "112":
		eventDescr = "Scandal rocks the shareholders’ memo meetings and prices hit an all time low."
		eventSuccess := dyn.failureCheck(charLoyalty, 8)
		eventOutcome = "Conglomerate averted damage."
		if !eventSuccess {
			dyn.setBonus(valueMorale, -1)
			eventOutcome = "Conglomerate lost 1 point of Morale."
		}

	case "113":
		eventDescr = "Scandal rocks the shareholders’ memo meetings and prices hit an all time low."
		eventChar := dyn.safestOption(8, charTenacity, apttBureaucracy)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Conglomerate averted damage."
		if !eventSuccess {
			dyn.setBonus(valueWealth, -1)
			eventOutcome = "Conglomerate lost 1 point of Wealth."
		}

	case "114":
		dyn.RollHistoricEvent()

	case "115":
		//New ideas on the market test the Conglomerate’s ingenuity and adaptability; roll Economics or Research 7+ or lose 1 point of Fiscal Defence.
		eventDescr = "New ideas on the market test the Conglomerate’s ingenuity and adaptability."
		eventChar := dyn.safestOption(7, apttEconomics, apttResearch)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "Conglomerate adapted successfuly."
		if !eventSuccess {
			dyn.setBonus(traitFiscalDefence, -1)
			eventOutcome = "Conglomerate lost 1 point of Fiscal Defence."
		}

	case "116":
		//A rival has been moving in on your workers, roll Security 8+ or lose 1 point of Populace.
		eventDescr = "A rival has been moving in on Conglomerate's workers."
		eventChar := dyn.safestOption(8, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Conglomerate averted damage."
		if !eventSuccess {
			dyn.setBonus(valuePopulace, -1)
			eventOutcome = "Conglomerate lost 1 point of Populace."
		}

	case "121":
		//Heavy competition in the interplanetary market has really toughened things up around the power base; roll 1d6: 1–4, Gain +1 Territorial Defence; 5–6: Gain +1 Tenacity.
		eventDescr = "Heavy competition in the interplanetary market has really toughened things up around the power base."
		r := utils.RollDice("d6")
		switch r {
		case 1, 2, 3, 4:
			dyn.setBonus(traitTerritorialDefence, 1)
			eventOutcome = "Conglomerate gained 1 point of Territorial Defence."
		case 5, 6:
			dyn.setBonus(charTenacity, 1)
			eventOutcome = "Conglomerate gained 1 point of Tenacity."
		}

	case "122":
		//A massive media event provides management with a chance to make a name for itself; Gain +1 Popularity or +2 Morale.
		eventDescr = "A massive media event provides management with a chance to make a name for itself."
		r := utils.RollDice("d6")
		switch r { //TODO: ПРИНЯТЬ РЕШЕНИЕ (Что нужнее получить?)
		case 1, 2, 3:
			dyn.setBonus(charPopularity, 1)
			eventOutcome = "Conglomerate gained 1 point of Popularity."
		case 4, 5, 6:
			dyn.setBonus(valueMorale, 2)
			eventOutcome = "Conglomerate gained 2 points of Morale."
		}

	case "123":
		//Big business is good business these days; roll Bureaucracy 7+ to gain one Level in Wealth.
		eventDescr = "Big business is good business these days."
		eventChar := dyn.safestOption(7, apttBureaucracy)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "Conglomerate is not big enough.."
		if eventSuccess {
			dyn.setBonus(valueWealth, 1)
			eventOutcome = "Conglomerate gained 1 point of Wealth."
		}

	case "124":
		dyn.RollHistoricEvent()

	case "125":
		//The territories are wearing your logo and there are not many locals who do not know your name. Gain +1 Popularity.
		eventDescr = "The territories are wearing your logo and there are not many locals who do not know your name."
		eventOutcome = "Gain +1 Popularity."
		dyn.setBonus(charPopularity, 1)

	case "126":
		//A major coup in the local government risks sweeping in the Conglomerate. Join in and roll Conquest 8+ to help the new regime. Avoid the conflict and roll Security 8+ to keep out of the line of fire. Succeed in either roll and gain +1 to any Value; fail and lose 1 point of Loyalty and Militarism.
		eventDescr = "A major coup in the local government risks sweeping in the Conglomerate."
		eventChar := dyn.safestOption(8, apttConquest, apttSecurity)
		if eventChar == apttConquest {
			eventOutcome = "Conglomerate joined the coup. "
		} else {
			eventOutcome = "Conglomerate avoided the coup. "
		}
		eventSuccess := dyn.failureCheck(eventChar, 8)

		if eventSuccess {
			value := utils.RandomFromList(listVALUES) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
			dyn.setBonus(value, 1)
			eventOutcome = "Conglomerate gained 1 point of " + value
		} else {
			eventOutcome = "Conglomerate lost 1 point of Loyalty and Militarism"
			dyn.setBonus(charLoyalty, -1)
			dyn.setBonus(charMilitarism, -1)
		}

	case "131":
		//Labour unions are not happy about the solidification of the management entities through the Conglomerate. Roll Propaganda and Security 7+; succeed in both and gain +1 to the Characteristic of the Player’s choice.
		eventDescr = "Labour unions are not happy about the solidification of the management entities through the Conglomerate."
		suc1 := dyn.failureCheck(apttPropaganda, 7)
		suc2 := dyn.failureCheck(apttSecurity, 7)
		eventOutcome = "Conglomerate avoided damage"
		if suc1 == true && suc2 == true {
			char := utils.RandomFromList(listCHARS) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
			dyn.setBonus(char, 1)
			eventOutcome = "Conglomerate gained 1 point of " + char
		}

	case "132":
		//The management of the Conglomerate are contacted by tremendously powerful alien benefactors; Gain +1 Bureaucracy, Expression or Recruit.
		eventDescr = "The management of the Conglomerate are contacted by tremendously powerful alien benefactors."
		apt := utils.RandomFromList([]string{apttBureaucracy, apttExpression, apttRecruitment}) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
		dyn.setBonus(apt, 1)
		eventOutcome = "Conglomerate gained 1 point of " + apt + "."

	case "133":
		//Everything goes as planned for decades; add +1 to any Trait or Value.
		eventDescr = "Everything goes as planned for decades."
		apt := utils.RandomFromList(listTRAITS) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
		dyn.setBonus(apt, 1)
		eventOutcome = "Conglomerate gained 1 point of " + apt + "."

	case "134":
		dyn.RollHistoricEvent()

	case "135":
		//The power base suffers a major natural disaster and the Conglomerate can lend charitable aid; you may spend 1 point of Wealth to increase Popularity by +1.
		eventDescr = "The power base suffers a major natural disaster."
		if utils.RandomBool() {
			eventOutcome = "Conglomerate lended help and gained 1 point of Popularity, but lost 1 point of Wealth"
			dyn.setBonus(charPopularity, 1)
			dyn.setBonus(valueWealth, -1)
		}

	case "136":
		//A sickness plagues the population and the workforce, putting the Conglomerate at risk but giving them a good idea to back medical resources. Roll Acquisition 8+ to gain 1 point of any Value.
		eventDescr = "A sickness plagues the population and the workforce, putting the Conglomerate at risk but giving them a good idea to back medical resources."
		eventChar := dyn.safestOption(8, apttAcquisition)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit."
		if eventSuccess {
			value := utils.RandomFromList(listVALUES) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
			dyn.setBonus(value, 1)
			eventOutcome = "Conglomerate gained 1 point of " + value
		}

	case "141":
		//A powerful client puts the Conglomerate through a vicious courtroom drama that lasts months, if not years; roll Politics or Security 8+ to avoid losing 1 point of Wealth.
		eventDescr = "A powerful client puts the Conglomerate through a vicious courtroom drama that lasts months, if not years."
		eventChar := dyn.safestOption(8, apttPolitics, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit."
		if !eventSuccess {
			dyn.setBonus(valueWealth, 1)
			eventOutcome = "Conglomerate lost 1 point of Wealth."
		}

	case "142":
		//A university grant is created in the Conglomerate’s honour; Gain +1 Popularity.
		eventDescr = "A university grant is created in the Conglomerate’s honour."
		eventOutcome = "Conglomerate gained 1 point of Popularity"
		dyn.setBonus(charPopularity, 1)

	case "143":
		//Everything goes as planned for decades; add +1 to any Aptitude or Trait.
		eventDescr = "Everything goes as planned for decades."
		var things []string
		things = append(things, listAPTITUDES...)
		things = append(things, listTRAITS...)
		thing := utils.RandomFromList(things)
		dyn.setBonus(thing, 1)
		eventOutcome = "Conglomerate gained 1 point of " + thing + "."

	case "144":
		dyn.RollHistoricEvent()

	case "145":
		//High-credit gambling establishments become not only legal but encouraged among big businesses; Roll Illicit 7+ to gain +1 Wealth.
		eventDescr = "High-credit gambling establishments become not only legal but encouraged among big businesses."
		eventChar := dyn.safestOption(7, apttIllicit)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "No significant profit."
		if !eventSuccess {
			dyn.setBonus(valueWealth, 1)
			eventOutcome = "Conglomerate lost 1 point of Wealth."
		}

	case "146":
		//Primitives are in great supply to be exploited. If the Conglomerate treats them with respect, it gains +1 Popularity. If it uses them harshly, gain +1 Wealth and +1 Populace.
		eventDescr = "Primitives are in great supply to be exploited."
		if utils.RandomBool() {
			eventOutcome = "Conglomerate treats them with respect. (+1 Popularity)"
			dyn.setBonus(charPopularity, 1)
		} else {
			eventOutcome = "Conglomerate treats them harshly. (+1 Wealth, +1 Populance)"
			dyn.setBonus(valueWealth, 1)
			dyn.setBonus(valuePopulace, 1)
		}

	case "151":
		//Industrial sabotage is rumoured to be targeting the Conglomerate; roll Security 8+ to protect itself, gaining +1 Territorial Defence.
		eventDescr = "Industrial sabotage is rumoured to be targeting the Conglomerate."
		eventChar := dyn.safestOption(8, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant loses."
		if eventSuccess {
			dyn.setBonus(traitTerritorialDefence, 1)
			eventOutcome = "Conglomerate protected itself (+1 Territorial Defence)."
		}
	case "152":
		//Advanced aliens have chosen the Conglomerate to fabricate their devices, adding their tooling to their own; Roll Maintenance 8+ to gain +1 Technology.
		eventDescr = "Advanced aliens have chosen the Conglomerate to fabricate their devices, adding their tooling to their own."
		eventChar := dyn.safestOption(8, apttMaintenance)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "It turns out to be no significant Profit."
		if eventSuccess {
			dyn.setBonus(traitTechnology, 1)
			eventOutcome = "Conglomerate gained 1 point of Technology."
		}

	case "153":
		//Everything goes as planned for decades; add +1 to any Aptitude or Trait.
		eventDescr = "Everything goes as planned for decades."
		var things []string
		things = append(things, listAPTITUDES...)
		things = append(things, listVALUES...)
		thing := utils.RandomFromList(things)
		dyn.setBonus(thing, 1)
		eventOutcome = "Conglomerate gained 1 point of " + thing + "."

	case "154":
		dyn.RollHistoricEvent()

	case "155":
		//War profiteers are looking to launder their ill-gotten gains through the Conglomerate; you may spend 1 point of Loyalty to gain 1 point of Scheming before rolling Illicit 9+; succeed in the Aptitude check to gain 1d6-4 Wealth (minimum of 1).
		eventDescr = "War profiteers are looking to launder their ill-gotten gains through the Conglomerate."
		if utils.RandomBool() && dyn.pickVal(charLoyalty) > 0 { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			eventOutcome = "Some were not found of the fact that offer was considered (-1 Loyalty, +1 Scheming)."
			dyn.setBonus(charLoyalty, 1)
			dyn.setBonus(charScheming, 1)
		}
		eventChar := dyn.safestOption(9, apttIllicit)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = eventOutcome + " It turns out to be no significant Profit."
		if eventSuccess {
			profit := utils.RollDice("d6", 4)
			if profit < 1 {
				profit = 1
			}
			dyn.setBonus(valueWealth, profit)
			eventOutcome = eventOutcome + " Conglomerate gained +" + convert.ItoS(profit) + " Wealth"
		}

	case "156":
		//A celebrity enjoys associating on a business level with the Conglomerate; gain +1 Morale or blackmail the Celebrity with Illicit 8+ to gain +1 Wealth and +1 Scheming.
		eventDescr = "A celebrity enjoys associating on a business level with the Conglomerate."
		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			eventOutcome = "Blackmail was released."
			eventOutcome = eventOutcome + "It turns out to be no significant Profit."
			eventChar := dyn.safestOption(8, apttIllicit)
			eventSuccess := dyn.failureCheck(eventChar, 8)
			if eventSuccess {
				eventOutcome = eventOutcome + " Conglomerate gains +1 Wealth and +1 Scheming."
			}
		} else {
			eventOutcome = "Blackmail was not used (+1 Morale)."
			dyn.setBonus(valueMorale, 1)
		}

	case "161":
		//The government names a holiday after the Conglomerate’s founder(s); Gain +1 Loyalty, Popularity or Tradition.
		eventDescr = "The management of the Conglomerate are contacted by tremendously powerful alien benefactors."
		apt := utils.RandomFromList([]string{charLoyalty, charPopularity, traitCulture}) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
		dyn.setBonus(apt, 1)
		eventOutcome = "Conglomerate gained 1 point of " + apt + "."

	case "162":
		//An interstellar sports team needs a sponsor right before a major multi-planet tournament; buy the team by spending 1 point of Wealth, gaining +1 Culture and +1 Morale.
		eventDescr = "An interstellar sports team needs a sponsor right before a major multi-planet tournament."
		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			eventOutcome = "Conglomerate desided to buy them. (-1 Wealth, +1 Culture, +1 Morale)"
			dyn.setBonus(valueWealth, -1)
			dyn.setBonus(traitCulture, 1)
			dyn.setBonus(valueMorale, 1)
		} else {
			eventOutcome = "Conglomerate refused the offer."
		}

	case "163":
		//Things could not go any better for the Dynasty; add +1 to any Characteristic, Aptitude, Trait or Value.
		eventDescr = "Things could not go any better for the Dynasty."
		var things []string
		things = append(things, listAPTITUDES...)
		things = append(things, listVALUES...)
		things = append(things, listCHARS...)
		things = append(things, listTRAITS...)
		thing := utils.RandomFromList(things)
		dyn.setBonus(thing, 1)
		eventOutcome = "Conglomerate gained 1 point of " + thing + "."

	case "164":
		dyn.RollHistoricEvent()

	case "165":
		//An unexpected territory shift puts a new planet in the Conglomerate’s control, adding +1 to all Values.
		eventDescr = "An unexpected territory shift puts a new planet in the Conglomerate’s control."
		for i := range listVALUES {
			dyn.setBonus(listVALUES[i], 1)
		}
		eventOutcome = "Conglomerate gained 1 point to all Values."

	case "166":
		//A formerly powerful Conglomerate folds, leaving its resources and assets for the new one to claim unchallenged; Gain +1 to any one Characteristic and +1 to any two Aptitudes.
		eventDescr = "A formerly powerful Conglomerate folds, leaving its resources and assets for the new one to claim unchallenged."
		var things []string
		for len(things) < 2 {
			things = utils.AppendUniqueStr(things, utils.RandomFromList(listAPTITUDES))
		}
		things = append(things, utils.RandomFromList(listCHARS))
		for i := range things {
			dyn.setBonus(things[i], 1)
		}
		eventOutcome = "Conglomerate gained 1 point to " + things[0] + ", " + things[1] + " and " + things[2] + "."

	case "211":
		//War forces a media blackout for years; roll Tenacity 8+ or lose 1 point of Morale.
		eventDescr = "War forces a media blackout for years."
		eventChar := dyn.safestOption(8, charTenacity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant loses."
		if !eventSuccess {
			dyn.setBonus(valueMorale, -1)
			eventOutcome = "The Media Empire lost 1 point of Morale."
		}

	case "212":
		//The Empire is labelled as spreading lies; roll Propaganda 9+ or lose 1 point of Wealth.
		eventDescr = "The Empire is labelled as spreading lies."
		eventChar := dyn.safestOption(9, apttPropaganda)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "Damage averted."
		if !eventSuccess {
			dyn.setBonus(valueMorale, -1)
			eventOutcome = "Legal expenses was considerable (-1 Wealth)"
		}

	case "213":
		//One of the largest stories of the generation goes to the competition; roll Acquisition or Sabotage 9+ or lose 1 point of Popularity.
		eventDescr = "One of the largest stories of the generation goes to the competition."
		eventChar := dyn.safestOption(9, apttAcquisition, apttSabotage)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "Damage averted."
		if !eventSuccess {
			dyn.setBonus(charPopularity, -1)
			eventOutcome = "(-1 Popularity)"
		}

	case "214":
		dyn.RollHistoricEvent()

	case "215":
		//A new information-distribution device becomes available and the Empire needs to gain it or lose huge standing; roll Acquisition or Intel 7+ or lose 1 point of Technology.
		eventDescr = "A new information-distribution device becomes available."
		eventChar := dyn.safestOption(7, apttAcquisition, apttIntel)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "Media Empire could get it in time."
		if !eventSuccess {
			dyn.setBonus(traitTechnology, -1)
			eventOutcome = "Media Empire couldn't get it in time (-1 Popularity)."
		}

	case "216":
		//Computer virus strikes the data stores; roll Research 8+ or lose 1 point from both the Research Aptitude and the Tradition Characteristic.
		eventDescr = "Computer virus strikes the data stores."
		eventChar := dyn.safestOption(8, apttResearch)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Damage was averted."
		if !eventSuccess {
			dyn.setBonus(apttResearch, -1)
			dyn.setBonus(charTradition, -1)
			eventOutcome = "Damage was tremendous (-1 Research, -1 Tradition)."
		}

	case "221":
		//There are plentiful stories to be had during a major governmental shift; Gain +1 Greed or +2 Wealth.
		eventDescr = "There are plentiful stories to be had during a major governmental shift."
		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			eventOutcome = "Media Empire is thriving (+1 Greed)"
			dyn.setBonus(charGreed, 1)
		} else {
			eventOutcome = "Media Empire is thriving (+2 Wealth)"
			dyn.setBonus(valueWealth, 2)
		}

	case "222":
		//Small-time competition has been growing steadily for years all around the home territory; roll 1d6: 1–2, Gain +1 Tenacity; 3–6: Gain +1 Fiscal Defence.
		eventDescr = "Small-time competition has been growing steadily for years all around the home territory."
		r := utils.RollDice("d6")
		switch r {
		case 1, 2:
			eventOutcome = "Competition increased (+1 Tenacity)"
			dyn.setBonus(charTenacity, 1)
		case 3, 4, 5, 6:
			eventOutcome = "Competition increased (+1 Fiscal Defence)"
			dyn.setBonus(traitFiscalDefence, 1)
		}

	case "223":
		//There is nothing like a leadership scandal to get people to pay attention; roll Expression or Propaganda 8+ to gain +1 Wealth.
		eventDescr = "There is nothing like a leadership scandal to get people to pay attention."
		eventChar := dyn.safestOption(8, apttExpression, apttPropaganda)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			dyn.setBonus(valueWealth, 1)
			eventOutcome = "Profit was made (+1 Wealth)."
		}

	case "224":
		dyn.RollHistoricEvent()

	case "225":
		//Through priority of coverage, the Empire is listed as the most-connected information source in the subsector. Gain +1 Culture.
		eventDescr = "Through priority of coverage, the Empire is listed as the most-connected information source in the subsector. (+1 Culture)."
		dyn.setBonus(traitCulture, 1)

	case "226":
		//War has ripped the subsector into several mini-governments with their own rules and regulations. Stay neutral and roll Intel 8+ to simply learn from what is going on. Help set the new leadership and roll Politics 8+. Succeed in either of these rolls and gain +1 to any Trait or Value; fail and lose 1 point of Militarism or Popularity.
		eventDescr = "War has ripped the subsector into several mini-governments with their own rules and regulations."
		eventChar := dyn.safestOption(8, apttIntel, apttPolitics)
		if eventChar == apttIntel {
			eventOutcome = "Media Empire stayed neutral. "
		} else {
			eventOutcome = "Media Empire tryed to bring new leader. "
		}
		eventSuccess := dyn.failureCheck(eventChar, 8)
		if eventSuccess {
			things := listTRAITS
			things = append(things, listVALUES...)
			bonus := utils.RandomFromList(things)
			dyn.setBonus(bonus, 1)
			eventOutcome += "(+1 " + bonus + ")"
		} else {
			things := []string{charMilitarism, charPopularity}
			bonus := utils.RandomFromList(things)
			dyn.setBonus(bonus, -1)
			eventOutcome += "(-1 " + bonus + ")"
		}

	case "231":
		//The Empire has a chance at exclusivity over the, arguably, largest event in recent history. Roll Expression and Propaganda 7+; succeed in both and gain +1 to Greed, Loyalty or Tradition.
		eventDescr = "The Empire has a chance at exclusivity over the, arguably, largest event in recent history."
		suc1 := dyn.failureCheck(apttExpression, 7)
		suc2 := dyn.failureCheck(apttPropaganda, 7)
		eventOutcome = "No significant profit was made."
		if suc1 == true && suc2 == true {
			char := utils.RandomFromList([]string{charGreed, charLoyalty, charTradition}) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
			dyn.setBonus(char, 1)
			eventOutcome = "Media Empire gained 1 point of " + char
		}

	case "232":
		//A neighbouring corporate powerhouse funnels money into the Media Empire as long as they run favourable stories about them; Gain +1 Economics, Entertain or Politics.
		eventDescr = "A neighbouring corporate powerhouse funnels money into the Media Empire as long as they run favourable stories about them."
		char := utils.RandomFromList([]string{apttEconomics, apttEntertain, apttPolitics}) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
		dyn.setBonus(char, 1)
		eventOutcome = "Media Empire gained 1 point of " + char

	case "233":
		//Everything goes as planned for decades; add +1 to any Trait or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listTRAITS
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "234":
		dyn.RollHistoricEvent()

	case "235":
		//A stellar event blacks out entire populations within the Empire’s territory, causing months of lost revenue without any kind of advance in power sources; roll Research 8+ to increase Technology by +1.
		eventDescr = "A stellar event blacks out entire populations within the Empire’s territory, causing months of lost revenue without any kind of advance in power sources."
		eventChar := dyn.safestOption(8, apttResearch)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant damage received."
		if eventSuccess {
			dyn.setBonus(traitTechnology, 1)
			eventOutcome = "(+1 Technology)."
		}

	case "236":
		//A power-hungry noble is hiding something huge from the Media Empire. Roll Intel 9+ to gain 1 point of any Value.
		eventDescr = "A power-hungry noble is hiding something huge from the Media Empire."
		eventChar := dyn.safestOption(9, apttIntel)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "No significant damage received."
		if eventSuccess {
			val := utils.RandomFromList(listVALUES)
			dyn.setBonus(val, 1)
			eventOutcome = "Data recovered (+1 " + val + ")."
		}

	case "241":
		//A medical epidemic of planetary scale needed serious coordination to help make aid efforts effective; roll Public Relations or Tutelage 8+ to avoid losing 1 point of Populace.
		eventDescr = "A medical epidemic of planetary scale needed serious coordination to help make aid efforts effective."
		eventChar := dyn.safestOption(8, apttPublicRelations, apttTutelage)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Major disaster was averted."
		if !eventSuccess {
			dyn.setBonus(valuePopulace, 1)
			eventOutcome = "Dosens of thousands died (-1 Populance)."
		}

	case "242":
		//An art school begins to focus on multi-media training, creating an entire class of reporters and paparazzi; Gain +1 to any Trait.
		eventDescr = "An art school begins to focus on multi-media training, creating an entire class of reporters and paparazzi."
		char := utils.RandomFromList(listTRAITS) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
		dyn.setBonus(char, 1)
		eventOutcome = "Media Empire gained 1 point of " + char

	case "243":
		//Everything goes as planned for decades; add +1 to any Aptitude or Trait.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listTRAITS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "244":
		dyn.RollHistoricEvent()

	case "245":
		//There is a heavy market for ‘illegal’ broadcasts and edgy personal videos; Roll Scheming 7+ to make professional versions and gain +1 Wealth.
		eventDescr = "There is a heavy market for ‘illegal’ broadcasts and edgy personal videos."
		eventChar := dyn.safestOption(7, charScheming)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			dyn.setBonus(valueWealth, 1)
			eventOutcome = "(+1 Wealth)."
		}

	case "246":
		//A new politician is well known for his scandalous sound-bytes and loves to be in front of the camera. If the Empire exploits this fact, it gains +1 Morale. If it ignores him in favour of ‘good news reporting’, gain +1 Popularity.
		eventDescr = "A new politician is well known for his scandalous sound-bytes and loves to be in front of the camera."
		if utils.RandomBool() {
			eventOutcome = "Media Empire exploited the fact (+1 Morale)."
			dyn.setBonus(valueMorale, 1)
		} else {
			eventOutcome = "Media Empire ignored ."
			dyn.setBonus(valueMorale, 1)
		}

	case "251":
		//A bad story riles up some very angry and powerful people; Roll Security 8+ to protect the Empire from hackers, gaining +1 Fiscal Defence.
		eventDescr = "A bad story riles up some very angry and powerful people."
		eventChar := dyn.safestOption(8, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Major disaster was averted."
		if eventSuccess {
			dyn.setBonus(traitFiscalDefence, 1)
			eventOutcome = "Hacker protection was enchanced (+1 Fiscal Defence)."
		}

	case "252":
		//The leader of an advanced alien race wants to add the Media Empire to its far-reaching network of affiliates, if they can convince them of the Empire’s skill; Roll Expression 8+ to gain +1 Technology and Culture.
		eventDescr = "The leader of an advanced alien race wants to add the Media Empire to its far-reaching network of affiliates, if they can convince them of the Empire’s skill."
		eventChar := dyn.safestOption(8, apttExpression)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Negotiations failed."
		if eventSuccess {
			dyn.setBonus(traitTechnology, 1)
			dyn.setBonus(traitCulture, 1)
			eventOutcome = "Negotiations successeed (+1 Technology, +1 Culture)."
		}

	case "253":
		//Everything goes as planned for decades; add +1 to any Aptitude or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "254":
		dyn.RollHistoricEvent()

	case "255":
		//Powerful people are building a new government and want the Media Empire to be the foundation of this growth; either roll Illicit 9+ to be the stereotypical muckraker they might want or roll Public Relations 9+ to be a proper flagwaver; succeed in the Aptitude check to gain +1 to any two Traits.
		eventDescr = "Powerful people are building a new government and want the Media Empire to be the foundation of this growth."
		eventChar := dyn.safestOption(9, apttIllicit, apttPublicRelations)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			var things []string
			for len(things) < 2 {
				things = utils.AppendUniqueStr(things, utils.RandomFromList(listTRAITS))
			}
			for i := range things {
				dyn.setBonus(things[i], 1)
			}
			eventOutcome = "Media Empire joined. (+1 " + things[0] + ", +1 " + things[1] + ")."
		}

	case "256":
		//An elderly business mogul gives the Media Empire the rights to produce and publish his life story; gain +1 Culture or wait and tell popular lies about him after his passing with Propaganda 8+ to gain +1 Greed but –1 Loyalty.
		eventDescr = "An elderly business mogul gives the Media Empire the rights to produce and publish his life story."
		if utils.RandomBool() {
			eventOutcome = "Media Empire published memoirs (+1 Culture)."
			dyn.setBonus(traitCulture, 1)
		} else {
			eventOutcome = "Media Empire waited and told popular lies about him after his passing."
			eventChar := dyn.safestOption(8, apttPropaganda)
			eventSuccess := dyn.failureCheck(eventChar, 8)
			if eventSuccess {
				dyn.setBonus(charGreed, 1)
				dyn.setBonus(charLoyalty, -1)
				eventOutcome += "(+1 Greed, -1 Loyalty)"
			}
			dyn.setBonus(valueMorale, 1)
		}

	case "261":
		//There is a mercenary charter that wants to trade its services for a good public relations campaign; roll Public Relations 8+ to gain +1 Militarism or +2 Territorial Defence.
		eventDescr = "There is a mercenary charter that wants to trade its services for a good public relations campaign."
		eventChar := dyn.safestOption(8, apttPublicRelations)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Negotiations failed."
		if eventSuccess {
			dyn.setBonus(charMilitarism, 1)
			dyn.setBonus(traitTerritorialDefence, 2)
			eventOutcome = "Negotiations successeed (+1 Militarism, +2 Territorial Defence)."
		}

	case "262":
		//New hololithic and broadcasting flotillas have been built by private contractors; buy use of them by spending 1 point of Wealth, gaining +1 Technology and +1 Fleet.
		eventDescr = "New hololithic and broadcasting flotillas have been built by private contractors."
		if utils.RandomBool() {
			eventOutcome = "Offer to buy them was dissmissed."
		} else {
			dyn.setBonus(valueWealth, -1)
			dyn.setBonus(traitTechnology, 1)
			dyn.setBonus(traitFleet, 1)
			eventOutcome = "Offer to buy them was accepted. (-1 Wealth, +1 Technology, +1 Fleet)"
		}

	case "263":
		//Things could not go any better for the Dynasty; add +1 to any Characteristic, Aptitude, Trait or Value.
		eventDescr = "Things could not go any better for the Dynasty."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		things = append(things, listTRAITS...)
		things = append(things, listCHARS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "264":
		dyn.RollHistoricEvent()

	case "265":
		//The only rival in the area has gone bankrupt and its workers and clientele need a new Media Empire to maintain their lifestyles, adding +1 to all Values.
		eventDescr = "The only rival in the area has gone bankrupt and its workers and clientele need a new Media Empire to maintain their lifestyles."
		for i := range listVALUES {
			dyn.setBonus(listVALUES[i], 1)
			eventOutcome = "(+1 to all Values)."
		}

	case "266":
		//The leader of an interplanetary government pledges the Media Empire to be the only service his people will use; Gain +1 to any one Characteristic and +1 to any two Aptitudes.
		eventDescr = "The leader of an interplanetary government pledges the Media Empire to be the only service his people will use."
		var things []string
		for len(things) < 2 {
			things = utils.AppendUniqueStr(things, utils.RandomFromList(listCHARS))
		}
		things = append(things, utils.RandomFromList(listAPTITUDES))
		for i := range things {
			dyn.setBonus(things[i], 1)
		}
		eventOutcome = "(+1 to " + things[0] + ", " + things[1] + " and " + things[2] + ")"

		////MERCHANT MARKET

	case "311":
		//An economic depression is killing the local businesses; roll Cleverness 8+ or lose 1 point of Wealth.
		eventDescr = "An economic depression is killing the local businesses."
		eventChar := dyn.safestOption(8, charCleverness)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant loses."
		if !eventSuccess {
			dyn.setBonus(valueWealth, -1)
			eventOutcome = "The Merchant market lost 1 point of Wealth."
		}

	case "312":
		//Someone started the rumour that the Merchant Market is cheating all of its distributors; roll Public Relations 8+ or lose 1 point of Popularity.
		eventDescr = "Someone started the rumour that the Merchant Market is cheating all of its distributors."
		eventChar := dyn.safestOption(8, apttPublicRelations)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Damage averted."
		if !eventSuccess {
			dyn.setBonus(charPopularity, -1)
			eventOutcome = "Reputation damage was considerable (-1 Popularity)"
		}

	case "313":
		//A bad choice sends a shockwave through the economy, forcing the Merchant Market to think outside the box to come out unscathed; roll Illicit or Sabotage 8+ or lose 1 point of Morale.
		eventDescr = "A bad choice sends a shockwave through the economy, forcing the Merchant Market to think outside the box to come out unscathed."
		eventChar := dyn.safestOption(8, apttIllicit, apttSabotage)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Damage averted."
		if !eventSuccess {
			dyn.setBonus(charPopularity, -1)
			eventOutcome = "(-1 Morale)"
		}

	case "314":
		dyn.RollHistoricEvent()

	case "315":
		//A seller’s union moves into the area and must be scared off; roll Hostility or Posturing 8+ or lose 1 point of Fiscal Defence.
		eventDescr = "A seller’s union moves into the area and must be scared off."
		eventChar := dyn.safestOption(8, apttHostility, apttPosturing)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Damage averted."
		if !eventSuccess {
			dyn.setBonus(traitFiscalDefence, -1)
			eventOutcome = "Financial damage was considerable (-1 Fiscal Defence)."
		}

	case "316":
		//A new product became available but carried some risks along with it; roll Intel 8+ or lose 1 point from both the Research Aptitude and the Popularity Characteristic.
		eventDescr = "A new product became available but carried some risks along with it."
		eventChar := dyn.safestOption(8, apttIntel)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Damage was averted."
		if !eventSuccess {
			dyn.setBonus(apttResearch, -1)
			dyn.setBonus(charPopularity, -1)
			eventOutcome = "Damage was tremendous (-1 Research, -1 Popularity)."
		}

	case "321":
		//The economy took a sudden post-war upturn; Gain +1 Fiscal Defence or +2 Wealth.
		eventDescr = "The economy took a sudden post-war upturn."
		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			eventOutcome = "Merchant Market is thriving (+1 Fiscal Defence)"
			dyn.setBonus(traitFiscalDefence, 1)
		} else {
			eventOutcome = "Merchant Market is thriving (+2 Wealth)"
			dyn.setBonus(valueWealth, 2)
		}

	case "322":
		//The space lane authorities have been busy cleaning up pirate and raider cells in the area; Gain +1 Fleet.
		eventDescr = "The space lane authorities have been busy cleaning up pirate and raider cells in the area."
		eventOutcome = "(+1 Fleet)"
		dyn.setBonus(traitFleet, 1)

	case "323":
		//Undercutting prices can be a good way to gain market advantage over rival Dynasties; Roll Economics or Sabotage 8+ to gain +1 Wealth.
		eventDescr = "Undercutting prices can be a good way to gain market advantage over rival Dynasties."
		eventChar := dyn.safestOption(8, apttEconomics, apttSabotage)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			dyn.setBonus(valueWealth, 1)
			eventOutcome = "Profit was made (+1 Wealth)."
		}

	case "324":
		dyn.RollHistoricEvent()

	case "325":
		//There is nothing that cannot be sold by the agents of the Merchant Market. Gain +1 Fiscal Defence.
		eventDescr = "There is nothing that cannot be sold by the agents of the Merchant Market."
		eventOutcome = "(+1 Fiscal Defence)"
		dyn.setBonus(traitFiscalDefence, 1)

	case "326":
		//War is forcing the Merchant Market to steer its products toward military and combat-related industries. Evolve along with these tendencies and roll Tactical 8+ to maintain these efforts. Try to steer things back toward peaceful endeavours and roll Propaganda 8+. Succeed in either of these rolls and gain +2 to Militarism; fail and lose 1 point of Territorial Defence or Fleet.
		eventDescr = "War is forcing the Merchant Market to steer its products toward military and combat-related industries."
		eventChar := dyn.safestOption(8, apttTactical, apttPropaganda)
		if eventChar == apttTactical {
			eventOutcome = "Merchant market evolved along with these tendencies. "
		} else {
			eventOutcome = "Merchant market tryed to steer things back toward peaceful endeavours. "
		}
		eventSuccess := dyn.failureCheck(eventChar, 8)
		if eventSuccess {
			dyn.setBonus(charMilitarism, 2)
			eventOutcome += "(+2 " + charMilitarism + ")"
		} else {
			things := []string{traitTerritorialDefence, traitFleet}
			bonus := utils.RandomFromList(things)
			dyn.setBonus(bonus, -1)
			eventOutcome += "(-1 " + bonus + ")"
		}

	case "331":
		//A mass marketing plan could put the Merchant Market at the top of the pyramid. Roll Posturing and Propaganda 7+; succeed in both and gain +1 to Greed, Popularity or Wealth.
		eventDescr = "A mass marketing plan could put the Merchant Market at the top of the pyramid."
		suc1 := dyn.failureCheck(apttPosturing, 7)
		suc2 := dyn.failureCheck(apttPropaganda, 7)
		eventOutcome = "No significant profit was made."
		if suc1 == true && suc2 == true {
			char := utils.RandomFromList([]string{charGreed, charPopularity, valueWealth}) //TODO: ПРИНЯТЬ РЕШЕНИЕ (AnyValue - что нужнее получить?)
			dyn.setBonus(char, 1)
			eventOutcome = "The plan was Successful (+1 " + char + ")"
		}

	case "332":
		//The Merchant Market has the chance to buy a majority holding in a local university; spend 1 Wealth to gain +1 in any Trait or Aptitude.
		eventDescr = "The Merchant Market has the chance to buy a majority holding in a local university."
		if utils.RandomBool() {
			eventOutcome = "Offer to buy it was dissmissed."
		} else {
			things := listTRAITS
			things = append(things, listAPTITUDES...)
			bonus := utils.RandomFromList(things)
			dyn.setBonus(bonus, 1)
			eventOutcome = "Offer to buy it was accepted. (-1 Wealth, +1 " + bonus + ")"
		}

	case "333":
		//Everything goes as planned for decades; add +1 to any Trait or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listTRAITS
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "334":
		dyn.RollHistoricEvent()

	case "335":
		//The people are clamouring for new and improved products that only the Merchant Market might be able to appropriate; roll Acquisition or Research 8+ to increase Technology by +1.
		eventDescr = "The people are clamouring for new and improved products that only the Merchant Market might be able to appropriate."
		eventChar := dyn.safestOption(8, apttAcquisition, apttResearch)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			dyn.setBonus(traitTechnology, 1)
			eventOutcome = "(+1 Technology)."
		}

	case "336":
		//The economy is ripe with possibilities, if the Merchant Market can discover what those are. Roll Intel 9+ to gain 1 point of any Value.
		eventDescr = "The economy is ripe with possibilities, if the Merchant Market can discover what those are."
		eventChar := dyn.safestOption(9, apttIntel)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "No significant damage received."
		if eventSuccess {
			val := utils.RandomFromList(listVALUES)
			dyn.setBonus(val, 1)
			eventOutcome = "Data recovered (+1 " + val + ")."
		}

	case "341":
		//Planetary disaster gives the Merchant Market the opportunity to do major charity work – or exploit the needy; roll Expression or Scheming 8+ to gain +1 Wealth or Populace.
		eventDescr = "Planetary disaster gives the Merchant Market the opportunity to do major charity work – or exploit the needy."
		eventChar := dyn.safestOption(8, apttExpression, charScheming)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Market failed to use the opportunity."
		if !eventSuccess {
			bonus := utils.RandomFromList([]string{valueWealth, valuePopulace})
			dyn.setBonus(valuePopulace, 1)
			eventOutcome = "Market reacted to opportunity (+1 " + bonus + ")."
		}

	case "342":
		//A union of shipping labourers petitions the Merchant Market for long term contracting; Gain +1 Fleet.
		eventDescr = "A union of shipping labourers petitions the Merchant Market for long term contracting."
		dyn.setBonus(traitFleet, 1)
		eventOutcome = "(+1 Fleet)"

	case "343":
		//Everything goes as planned for decades; add +1 to any Aptitude or Trait.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listTRAITS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "344":
		dyn.RollHistoricEvent()

	case "345":
		//A dangerous product risks putting the Merchant Market in actionable danger if they do not cover their legal options tightly; Roll Bureaucracy 6+ to avoid losing –1 Wealth. If the result of this check is 8+, gain +1 Morale instead!
		eventDescr = "A dangerous product risks putting the Merchant Market in actionable danger if they do not cover their legal options tightly."
		eventChar := dyn.safestOption(7, apttBureaucracy)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "No significant profit received."
		if !eventSuccess {
			dyn.setBonus(valueWealth, -1)
			eventOutcome = "(-1 Wealth)."
		}

	case "346":
		//A new type of star craft engine has become available to the mercantile shipyards; Gain +1 Fleet or +1 Technology.
		eventDescr = "A new type of star craft engine has become available to the mercantile shipyards."
		if utils.RandomBool() {
			eventOutcome = "(+1 Fleet)."
			dyn.setBonus(traitFleet, 1)
		} else {
			eventOutcome = "(+1 Technology)."
			dyn.setBonus(traitTechnology, 1)
		}

	case "351":
		//An alien sales force wants to join forces for the future of both Dynasties; roll Acquisition 9+ to gain +1 to any two Values.
		eventDescr = "An alien sales force wants to join forces for the future of both Dynasties."
		eventChar := dyn.safestOption(9, apttAcquisition)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "Negotiations failed."
		if eventSuccess {
			var things []string
			for len(things) < 2 {
				things = utils.AppendUniqueStr(things, utils.RandomFromList(listVALUES))
			}
			for i := range things {
				dyn.setBonus(things[i], 1)
			}
			eventOutcome = "Negotiations successeded (+1 " + things[0] + ", +1 " + things[1] + ")."
		}

	case "352":
		//A death in a noble family leaves hundreds of thousands of Credits to the Merchant Market out of some misplaced loyalty; Gain +1 Wealth or +1 Culture.
		eventDescr = "A death in a noble family leaves hundreds of thousands of Credits to the Merchant Market out of some misplaced loyalty."
		if utils.RandomBool() {
			eventOutcome = "(+1 Wealth)"
			dyn.setBonus(valueWealth, 1)
		} else {
			eventOutcome = "(+1 Culture)"
			dyn.setBonus(traitCulture, 1)
		}

	case "353":
		//Everything goes as planned for decades; add +1 to any Aptitude or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "354":
		dyn.RollHistoricEvent()

	case "355":
		//Purposeful shorting of inventories can artificially create massive demand in the populace; you may roll Economics or Illicit 8+ to gain +1 Wealth but lose –1 Popularity if you are not successful.
		eventDescr = "Purposeful shorting of inventories can artificially create massive demand in the populace."
		eventChar := dyn.safestOption(8, apttIllicit, apttEconomics)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		if utils.RandomBool() {
			if eventSuccess {
				eventOutcome = "Opportunity was exploited (+1 Wealth)"
				dyn.setBonus(valueWealth, 1)
			} else {
				eventOutcome = "Opportunity failed (-1 Popularity)"
				dyn.setBonus(charPopularity, -1)
			}
		} else {
			eventOutcome = "Opportunity was dismissed."
		}

	case "356":
		//A terrible fire tears through a multi-million Credit investment; roll Bureaucracy 7+ to make sure insurance covers the loss and results in +1 Morale.
		eventDescr = "A terrible fire tears through a multi-million Credit investment."
		eventChar := dyn.safestOption(7, apttBureaucracy)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "Insurance covers the loss"
		if eventSuccess {
			eventOutcome = "Insurance covers the loss (+1 Morale)"
			dyn.setBonus(valueMorale, 1)
		}

	case "361":
		//Freelance psionic specialists make themselves available to the Merchant Market for commercial use; roll Acquisition 9+ or spend 1 Wealth to increase any two Traits by +1 point each.
		eventDescr = "Freelance psionic specialists make themselves available to the Merchant Market for commercial use."
		if utils.RandomBool() {
			var things []string
			for len(things) < 2 {
				things = utils.AppendUniqueStr(things, utils.RandomFromList(listTRAITS))
			}
			for i := range things {
				dyn.setBonus(things[i], 1)
			}
			eventOutcome = "Negotiations successeded (-1 Wealth, +1" + things[0] + ", +1 " + things[1] + ")."
		} else {
			eventChar := dyn.safestOption(9, apttAcquisition)
			eventSuccess := dyn.failureCheck(eventChar, 9)
			eventOutcome = "Negotiations failed."
			if eventSuccess {
				var things []string
				for len(things) < 2 {
					things = utils.AppendUniqueStr(things, utils.RandomFromList(listTRAITS))
				}
				for i := range things {
					dyn.setBonus(things[i], 1)
				}
				eventOutcome = "Negotiations successeded (+1" + things[0] + ", +1 " + things[1] + ")."
			}
		}

	case "362":
		//A class-action lawsuit fails miserably against your legal teams; Gain +1 Fiscal Defence.
		eventDescr = "A class-action lawsuit fails miserably against your legal teams."
		eventOutcome = "(+1 Fiscal Defence)."
		dyn.setBonus(traitFiscalDefence, 1)

	case "363":
		//Things could not go any better for the Dynasty; add +1 to any Characteristic, Aptitude, Trait or Value.
		eventDescr = "Things could not go any better for the Dynasty."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		things = append(things, listTRAITS...)
		things = append(things, listCHARS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "364":
		dyn.RollHistoricEvent()

	case "365":
		//The armed forces of the local government want to buy all of their major supplies from the Merchant Market. Roll Maintenance 8+ to keep the lines running smoothly, gaining +1 to Fleet, Territorial Defence or Militarism.
		eventDescr = "The armed forces of the local government want to buy all of their major supplies from the Merchant Market."
		eventChar := dyn.safestOption(8, apttMaintenance)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Merchant Market could not keep the lines running smoothly"
		if eventSuccess {
			bonus := utils.RandomFromList([]string{traitFleet, traitTerritorialDefence, charMilitarism})
			eventOutcome = "Merchant Market kept the lines running smoothly (+1 " + bonus + ")."
			dyn.setBonus(bonus, 1)
		}

	case "366":
		//A true sales monopoly on an entire economy; Gain +1 to any three Aptitudes, or +1 to any two Traits.
		eventDescr = "A true sales monopoly on an entire economy."
		var things []string
		if utils.RandomBool() {
			for len(things) < 3 {
				things = utils.AppendUniqueStr(things, utils.RandomFromList(listAPTITUDES))
			}
			for i := range things {
				dyn.setBonus(things[i], 1)
			}
			eventOutcome = "(+1 to " + things[0] + ", " + things[1] + " and " + things[2] + ")"
		} else {
			for len(things) < 2 {
				things = utils.AppendUniqueStr(things, utils.RandomFromList(listTRAITS))
			}
			for i := range things {
				dyn.setBonus(things[i], 1)
			}
			eventOutcome = "(+1 to " + things[0] + " and " + things[1] + ")"
		}

		////MILITARY CHARTER

	case "411":
		//Skirmishes plague the borders for decades; roll Militarism 8+ or lose 1 point of Populace.
		eventDescr = "Skirmishes plague the borders for decades."
		eventChar := dyn.safestOption(8, charMilitarism)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant loses."
		if !eventSuccess {
			dyn.setBonus(valuePopulace, -1)
			eventOutcome = "Populance migration is on the rise (-1 Populance)."
		}

	case "412":
		//Arms dealers are acting particularly difficult for many years; roll Acquisition 8+ or lose 1 point of Territorial Defence.
		eventDescr = "Arms dealers are acting particularly difficult for many years."
		eventChar := dyn.safestOption(8, apttAcquisition)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = ""
		if !eventSuccess {
			dyn.setBonus(traitTerritorialDefence, -1)
			eventOutcome = "Shortage of arms ocured (-1 Territorial Defence)."
		}

	case "413":
		//The unexpected use of nuclear arms has forced the Military Charter to seek higher grades of personnel protection; roll Research 8+ or lose 1 point of Populace.
		eventDescr = "The unexpected use of nuclear arms has forced the Military Charter to seek higher grades of personnel protection."
		eventChar := dyn.safestOption(8, apttResearch)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Damage averted."
		if !eventSuccess {
			dyn.setBonus(valuePopulace, -1)
			eventOutcome = "Populance migration is on the rise (-1 Populance)."
		}

	case "414":
		dyn.RollHistoricEvent()

	case "415":
		//There are numerous mercenary companies that are constantly moving in on the Military Charter; roll Hostility or Posturing 9+ or lose 1 point of Territorial Defence.
		eventDescr = "There are numerous mercenary companies that are constantly moving in on the Military Charter."
		eventChar := dyn.safestOption(9, apttHostility, apttPosturing)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "No Significant danger."
		if !eventSuccess {
			dyn.setBonus(traitTerritorialDefence, -1)
			eventOutcome = "Independence become an issue (-1 Territorial Defence)."
		}

	case "416":
		//Weapon technology in neighbouring cultures is vastly superior to those currently available to the Military Charter; roll Research 8+ or lose 1 point from both the Security Aptitude and the Technology Trait.
		eventDescr = "Weapon technology in neighbouring cultures is vastly superior to those currently available to the Military Charter."
		eventChar := dyn.safestOption(8, apttResearch)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Dynasty managed to keep up."
		if !eventSuccess {
			dyn.setBonus(apttSecurity, -1)
			dyn.setBonus(traitTechnology, -1)
			eventOutcome = "Danger is significant (-1 Security, -1 Technology)."
		}

	case "421":
		//Victory over an affluent and advanced target; Gain +1 Technology or +2 Wealth.
		eventDescr = "Victory over an affluent and advanced target."
		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			eventOutcome = "Dynasty is thriving (+1 Technology)"
			dyn.setBonus(traitTechnology, 1)
		} else {
			eventOutcome = "Dynasty is thriving (+2 Wealth)"
			dyn.setBonus(valueWealth, 2)
		}

	case "422":
		//Shipyards have been cranking out fighter craft for years; Gain +1 Territorial Defence or +1 Fleet.
		eventDescr = "Shipyards have been cranking out fighter craft for years."
		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			eventOutcome = "Dynasty is thriving (+1 Territorial Defence)"
			dyn.setBonus(traitTerritorialDefence, 1)
		} else {
			eventOutcome = "Dynasty is thriving (+1 Fleet)"
			dyn.setBonus(traitFleet, 2)
		}

	case "423":
		//Good routes are established and smugglers bring the Military Charter resources far more safely; Roll Acquisition or Security 8+ to gain +1 Wealth or Technology.
		eventDescr = "Good routes are established and smugglers bring the Military Charter resources far more safely."
		eventChar := dyn.safestOption(8, apttAcquisition, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			bonus := utils.RandomFromList([]string{valueWealth, traitTechnology})
			dyn.setBonus(bonus, 1)
			eventOutcome = "Profit was made (+1 " + bonus + ")."
		}

	case "424":
		dyn.RollHistoricEvent()

	case "425":
		//The local populace look to the Military Charter for leadership in their societal roles and ideals. Gain +1 Tradition or +1 Culture.
		eventDescr = "The local populace look to the Military Charter for leadership in their societal roles and ideals."
		bonus := utils.RandomFromList([]string{charTradition, traitCulture})
		dyn.setBonus(bonus, 1)
		eventOutcome = "Dynasty is thriving (+1 " + bonus + ")."

	case "426":
		//War threatens to tear the government apart. Stay out of it and let them work things through on their own by rolling Posturing 8+. Choose to back one side or the other and roll Politics 8+. Succeed in either of these rolls and gain +1 to Militarism or Popularity.
		eventDescr = "War threatens to tear the government apart."
		eventChar := dyn.safestOption(8, apttPosturing, apttPolitics)
		if eventChar == apttTactical {
			eventOutcome = "Dynasty tryed to not get involved. "
		} else {
			eventOutcome = "Dynasty joined one of the sides. "
		}
		eventSuccess := dyn.failureCheck(eventChar, 8)
		if eventSuccess {
			bonus := utils.RandomFromList([]string{charMilitarism, charPopularity})
			dyn.setBonus(bonus, 1)
			eventOutcome += "(+1 " + bonus + ")"
		}

	case "431":
		//The Military Charter can focus on showing its ‘lighter side’ to the population for many months or even years. Roll Expression and Public Relations 8+; succeed in both and gain +1 to all Values.
		eventDescr = "The Military Charter focuses on showing its ‘lighter side’ to the population for many months or even years."
		suc1 := dyn.failureCheck(apttExpression, 8)
		suc2 := dyn.failureCheck(apttPublicRelations, 8)
		eventOutcome = "Opportunity failed."
		if suc1 == true && suc2 == true {
			for i := range listVALUES {
				dyn.setBonus(listVALUES[i], 1)
			}
			eventOutcome = "The plan was Successful (+1 to all values)"
		}

	case "432":
		//Local media services can be tapped to relay positive information about the Military Charter; Gain +1 Entertain, Expression or Posturing.
		eventDescr = "Local media services tapped to relay positive information about the Military Charter."
		bonus := utils.RandomFromList([]string{apttEntertain, apttExpression, apttPosturing})
		dyn.setBonus(bonus, 1)
		eventOutcome += "(+1 " + bonus + ")"

	case "433":
		//Everything goes as planned for decades; add +1 to any Trait or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listTRAITS
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "434":
		dyn.RollHistoricEvent()

	case "435":
		//War is commonplace and war is big business for the Military Charter; you may spend 1 point of Populace to increase any two Traits by +1.
		eventDescr = "War is commonplace and war is big business for the Military Charter."
		if utils.RandomBool() {
			var things []string
			for len(things) < 2 {
				things = utils.AppendUniqueStr(things, utils.RandomFromList(listTRAITS))
			}
			for i := range things {
				dyn.setBonus(things[i], 1)
			}
			dyn.setBonus(valuePopulace, -1)
			eventOutcome = "(-1 Populance, +1 " + things[0] + ", +1 " + things[1] + ")."
		}

	case "436":
		//The Military Charter has vanquished the enemy from all over the sector, sweeping up their spoils wherever possible. Roll Acquisition 8+ to gain 1 point of any Value.
		eventDescr = "The Military Charter has vanquished the enemy from all over the sector, sweeping up their spoils wherever possible."
		eventChar := dyn.safestOption(8, apttAcquisition)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			val := utils.RandomFromList(listVALUES)
			dyn.setBonus(val, 1)
			eventOutcome = "(+1 " + val + ")."
		}

	case "441":
		//Governmental contracts come up empty but commercial ones skyrocket; roll Bureaucracy or Politics 8+ to avoid losing 1 point of Wealth.
		eventDescr = "Governmental contracts come up empty but commercial ones skyrocket."
		eventChar := dyn.safestOption(8, apttBureaucracy, apttPolitics)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Dynasty adopted."
		if !eventSuccess {
			dyn.setBonus(valueWealth, -1)
			eventOutcome = "(-1 Wealth)"
		}

	case "442":
		//A local holiday is created in the Military Charter’s honour due to one of its victories; Gain +1 Popularity.
		eventDescr = "A local holiday is created in the Military Charter’s honour due to one of its victories."
		dyn.setBonus(charPopularity, 1)
		eventOutcome = "(+1 Popularity)"

	case "443":
		//Everything goes as planned for decades; add +1 to any Aptitude or Trait.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listTRAITS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "444":
		dyn.RollHistoricEvent()

	case "445":
		//Several small targets present themselves in local space, allowing the Military Charter to possibly increase territory; Roll Conquest 8+ to gain +1 Morale.
		eventDescr = "Several small targets present themselves in local space, allowing the Military Charter to possibly increase territory."
		eventChar := dyn.safestOption(8, apttConquest)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Dynasty missed the opportunity."
		if !eventSuccess {
			dyn.setBonus(valueMorale, 1)
			eventOutcome = "(+1 Morale)."
		}

	case "446":
		//The local militia forces can be tapped to bolster the ranks. If the Military Charter uses them sparingly, it gains +1 Popularity. If it pushes them too hard, gain +1 Populace and Wealth but –1 Morale.
		eventDescr = "The local militia forces can was tapped to bolster the ranks."
		if utils.RandomBool() {
			eventOutcome = "Dynasty was using them sparingly (+1 Popularity)."
			dyn.setBonus(charPopularity, 1)
		} else {
			eventOutcome = "Dynasty pushed them too hard (+1 Populace, +1 Wealth, -1 Morale)."
			dyn.setBonus(valuePopulace, 1)
			dyn.setBonus(valueWealth, 1)
			dyn.setBonus(valueMorale, -1)
		}

	case "451":
		//Professional soldiers are being called in to bolster the defences of the Military Charter; spend 1 point of Wealth to gain +1 Fleet and Territorial Defence.
		eventDescr = "Professional soldiers are being called in to bolster the defences of the Military Charter."
		eventOutcome = "(-1 Wealth, +1 Fleet, +1 Territirial Defence)."
		dyn.setBonus(valueWealth, -1)
		dyn.setBonus(traitFleet, 1)
		dyn.setBonus(traitTerritorialDefence, 1)

	case "452":
		//An engineering corps opens up within the Military Charter; Gain + 1 level in Intel, Research or Technology.
		eventDescr = "An engineering corps opens up within the Military Charter"
		bonus := utils.RandomFromList([]string{apttIntel, apttResearch, traitTechnology})
		dyn.setBonus(bonus, 1)
		eventOutcome += "(+1 " + bonus + ")"

	case "453":
		//Everything goes as planned for decades; add +1 to any Aptitude or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "454":
		dyn.RollHistoricEvent()

	case "455":
		//A superior foe of which the Military Charter could not defeat has left itself open to underhanded tactics; roll Sabotage 8+ to increase any Value or Trait by +1.
		eventDescr = "A superior foe of which the Military Charter could not defeat has left itself open to underhanded tactics."
		eventChar := dyn.safestOption(8, apttSabotage)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Opportunity was missed"
		if eventSuccess {
			things := listTRAITS
			things = append(things, listVALUES...)
			bonus := utils.RandomFromList(things)
			dyn.setBonus(bonus, 1)
			eventOutcome = "Opportunity was exploited (+1 " + bonus + ")"

		}

	case "456":
		//Cybernetics and true bionic soldiers are all the craze; Gain +1 Conquest, Hostility or Security.
		eventDescr = "Cybernetics and true bionic soldiers are all the craze."
		bonus := utils.RandomFromList([]string{apttConquest, apttHostility, apttSecurity})
		dyn.setBonus(bonus, 1)
		eventOutcome += "(+1 " + bonus + ")"

	case "461":
		//A massive barbarian horde must be dealt with before it can become more advanced to deal with the Military Charter; Roll Conquest 9+ and gain +1 Morale and +1 Territorial Defence.
		eventDescr = "A massive barbarian horde must be dealt with before it can become more advanced to deal with the Military Charter."
		eventChar := dyn.safestOption(9, apttConquest)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "Opportunity was missed"
		if eventSuccess {
			dyn.setBonus(apttConquest, 1)
			dyn.setBonus(valueMorale, 1)
			eventOutcome = "Opportunity was exploited (+1 Conquest, +1 Morale)"
		}

	case "462":
		//The Military Charter’s management has the opportunity to leave a lasting memoir about the area’s conflicts and contests; Gain +1 Popularity if the story is glorious, gain +1 Loyalty if the story is hard-edged, or gain +2 Culture if the story is locally flattering.
		eventDescr = "The Military Charter’s management left a lasting memoir about the area’s conflicts and contests."
		bonus := utils.RandomFromList([]string{charPopularity, charLoyalty, traitCulture})
		switch bonus {
		case charPopularity:
			eventOutcome += "Story was glorious (+1 " + bonus + ")"
			dyn.setBonus(bonus, 1)
		case charLoyalty:
			eventOutcome += "Story was hard-edged (+1 " + bonus + ")"
			dyn.setBonus(bonus, 1)
		case traitCulture:
			eventOutcome += "Story was locally flattering (+2 " + bonus + ")"
			dyn.setBonus(bonus, 2)
		}

	case "463":
		//Things could not go any better for the Dynasty; add +1 to any Characteristic, Aptitude, Trait or Value.
		eventDescr = "Things could not go any better for the Dynasty."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		things = append(things, listTRAITS...)
		things = append(things, listCHARS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "464":
		dyn.RollHistoricEvent()

	case "465":
		//A truly glorious battle ends a war that had been raging for a long time, leaving the Military Charter as the heroes of the people; Gain +2 Popularity or +1 to all Values.
		eventDescr = "A truly glorious battle ends a war that had been raging for a long time, leaving the Military Charter as the heroes of the people."
		if utils.RandomBool() {
			eventOutcome = "(+2 Popularity)"
			dyn.setBonus(charPopularity, 2)
		} else {
			for i := range listVALUES {
				dyn.setBonus(listVALUES[i], 1)
			}
			eventOutcome = "(+1 to all values)"

		}

	case "466":
		//An ancient and powerful alien race allies with the Military Charter; Gain +1 to Territorial Defence, Fleet and Technology.
		eventDescr = "An ancient and powerful alien race allies with the Military Charter."
		eventOutcome = "(+1 to Territorial Defence, Fleet and Technology)"
		dyn.setBonus(traitTerritorialDefence, 1)
		dyn.setBonus(traitFleet, 1)
		dyn.setBonus(traitTechnology, 1)

		////NOBLE LINE

	case "511":
		//Arguments within the family plague them greatly; roll Maintenance 8+ or lose 1 point of Morale.
		eventDescr = "Arguments within the family plague them greatly."
		eventChar := dyn.safestOption(8, apttMaintenance)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = ""
		if !eventSuccess {
			dyn.setBonus(valueMorale, -1)
			eventOutcome = "(-1 Morale)."
		}

	case "512":
		//An estranged splinter bloodline rises to lay claim to the family fortune. The Noble Line must discourage them or lose millions. Roll Hostility, Posturing or Security 8+ or lose 2 points of Wealth.
		eventDescr = "An estranged splinter bloodline rises to lay claim to the family fortune. The Noble Line must discourage them or lose millions."
		eventChar := dyn.safestOption(8, apttHostility, apttPosturing, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "The Noble Line discouraged them."
		if !eventSuccess {
			dyn.setBonus(valueWealth, -2)
			eventOutcome = "The Noble Line lost millions (-2 Wealth)."
		}

	case "513":
		//A rival family has sided with a dangerous criminal Syndicate; roll Security 9+ or lose 1 point from a Value of the player’s choice.
		eventDescr = "A rival family has sided with a dangerous criminal Syndicate."
		eventChar := dyn.safestOption(9, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "Rival's new tactics was not effective."
		if !eventSuccess {
			bonus := utils.RandomFromList(listVALUES)
			dyn.setBonus(bonus, -1)
			eventOutcome = "Rival's new tactics was effective (-1 " + bonus + ")."
		}

	case "514":
		dyn.RollHistoricEvent()

	case "515":
		//There are numerous mercenary companies that are constantly moving in on the Military Charter; roll Hostility or Posturing 9+ or lose 1 point of Territorial Defence.
		eventDescr = "There are numerous mercenary companies that are constantly moving in on the Noble Line."
		eventChar := dyn.safestOption(9, apttHostility, apttPosturing)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "No Significant danger."
		if !eventSuccess {
			dyn.setBonus(traitTerritorialDefence, -1)
			eventOutcome = "Independence become an issue (-1 Territorial Defence)."
		}

	case "516":
		//Genetic research and applied eugenics are dangerous when first applied; roll Research 8+ or lose 1 point from both the Technology Trait and the Populace Value.
		eventDescr = "Genetic research and applied eugenics are dangerous when first applied."
		eventChar := dyn.safestOption(8, apttResearch)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Dynasty managed secure the research."
		if !eventSuccess {
			dyn.setBonus(valuePopulace, -1)
			dyn.setBonus(traitTechnology, -1)
			eventOutcome = "Major outbreack happen (-1 Populance, -1 Technology)."
		}

	case "521":
		//Birth and marriage rates in the family are far higher than normal for many years; Gain +1 Tradition or +2 Populace.
		eventDescr = "Birth and marriage rates in the family are far higher than normal for many years."
		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			eventOutcome = "Dynasty is thriving (+1 Tradition)"
			dyn.setBonus(charTradition, 1)
		} else {
			eventOutcome = "Dynasty is thriving (+2 Populance)"
			dyn.setBonus(valuePopulace, 2)
		}

	case "522":
		//A local government asks the family for monetary aid; Spend up to 2 points of Wealth to gain an equal amount of additional points in the Politics Aptitude and Territorial Defence Traits.
		eventDescr = "A local government asks the family for monetary aid."
		i := 0

		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			i++
			dyn.setBonus(apttPolitics, 1)
			dyn.setBonus(traitTerritorialDefence, 1)
			dyn.setBonus(valueWealth, -1)
		}
		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			i++
			dyn.setBonus(apttPolitics, 1)
			dyn.setBonus(traitTerritorialDefence, 1)
			dyn.setBonus(valueWealth, -1)
		}
		if i > 0 {
			eventOutcome = "Noble line lended some aid (+" + convert.ItoS(i) + " Politics and Territirial Defence, -" + convert.ItoS(i) + " Wealth)."
		} else {
			eventOutcome = "Dynasty rejected the plea."
		}

	case "523":
		//A family scandal can easily be turned around into something positive for the Noble Line as a whole; Roll Propaganda or Politics 8+ to gain +1 Culture or Morale.
		eventDescr = "A family scandal can easily be turned around into something positive for the Noble Line as a whole."
		eventChar := dyn.safestOption(8, traitCulture, apttPolitics)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			bonus := utils.RandomFromList([]string{valueMorale, traitCulture})
			dyn.setBonus(bonus, 1)
			eventOutcome = "(+1 " + bonus + ")."
		}

	case "524":
		dyn.RollHistoricEvent()

	case "525":
		//The Noble Line must resort to some unsavoury practices to snatch victory from the jaws of defeat; roll Illicit 8+ to gain +1 Wealth.
		eventDescr = "The Noble Line resorted to some unsavoury practices to snatch victory from the jaws of defeat."
		eventChar := dyn.safestOption(8, apttIllicit)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			dyn.setBonus(valueWealth, 1)
			eventOutcome = "(+1 Wealth)."
		}

	case "526":
		//A scandalous marriage with a similar but still alien species brings new sciences into their possession for understanding; roll Research 7+ to gain +1 Technology.
		eventDescr = "A scandalous marriage with a similar but still alien species brings new sciences into their possession for understanding."
		eventChar := dyn.safestOption(7, apttResearch)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "Theoriries later proved to be wrong"
		if eventSuccess {
			dyn.setBonus(traitTechnology, 1)
			eventOutcome = "(+1 Technology)."
		}

	case "531":
		//Uncharted territory falls into the Noble Line’s area of influence, requiring new assets to expand there; Gain +1 Fleet.
		eventDescr = "Uncharted territory falls into the Noble Line’s area of influence, requiring new assets to expand there."
		dyn.setBonus(traitFleet, 1)
		eventOutcome = "(+1 Fleet)"

	case "532":
		//The curing of a genetic birth defect in the family extends life expectancy; Gain +1 Populace.
		eventDescr = "The curing of a genetic birth defect in the family extends life expectancy."
		dyn.setBonus(valuePopulace, 1)
		eventOutcome = "(+1 Populance)"

	case "533":
		//Everything goes as planned for decades; add +1 to any Trait or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listTRAITS
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "534":
		dyn.RollHistoricEvent()

	case "535":
		//Uncovered treasures are bickered and fought over for decades between the family’s branches and inner-groups; roll Maintenance 8+ to gain +1 Wealth or +1 Morale.
		eventDescr = "Uncovered treasures are bickered and fought over for decades between the family’s branches and inner-groups."
		eventChar := dyn.safestOption(8, apttMaintenance)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant profit received."
		if eventSuccess {
			bonus := utils.RandomFromList([]string{valueWealth, valueMorale})
			dyn.setBonus(bonus, 1)
			eventOutcome = "(+1 " + bonus + ")"
		}

	case "536":
		//An artiste wants to immortalise the Noble Line’s successes through a series of personalised sculptures, so long as the family can afford his services. Roll Greed 8+ to increase Culture by +1.
		eventDescr = "An artiste wants to immortalise the Noble Line’s successes through a series of personalised sculptures, so long as the family can afford his services."
		eventChar := dyn.safestOption(8, charGreed)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Dynasty turned down the offer."
		if eventSuccess {

			dyn.setBonus(traitCulture, 1)
			eventOutcome = "(+1 Culture)."
		}

	case "541":
		//The younger generation is in danger of forgetting the ways of the elders, requiring additional teaching to remain ‘pure’ in the original ideals; roll Tutelage 8+ to gain +1 Tradition.
		eventDescr = "The younger generation is in danger of forgetting the ways of the elders, requiring additional teaching to remain 'pure' in the original ideals."
		eventChar := dyn.safestOption(8, apttTutelage)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Mentors were not found"
		if !eventSuccess {
			dyn.setBonus(charTradition, 1)
			eventOutcome = "Best mentors were hired (+1 Traditions)"
		}

	case "542":
		//Superstition runs rampant in the people controlled by the Noble Line, allowing manipulative nobles to gain a stranglehold on them that much easier; roll Propaganda 7+ to gain +1 to any Value.
		eventDescr = "Superstition runs rampant in the people controlled by the Noble Line, allowing manipulative nobles to gain a stranglehold on them that much easier."
		eventChar := dyn.safestOption(7, apttPropaganda)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "Dynasty missed the opportunity."
		if !eventSuccess {
			bonus := utils.RandomFromList(listVALUES)
			dyn.setBonus(bonus, 1)
			eventOutcome = "(+1 " + bonus + ")"
		}

	case "543":
		//Everything goes as planned for decades; add +1 to any Aptitude or Trait.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listTRAITS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "544":
		dyn.RollHistoricEvent()

	case "545":
		//Rising tempers require the Noble Line to either stay more advanced than local rivals or train that much harder. Roll Cleverness 8+ to increase Territorial Defence or Technology by +1.
		eventDescr = "Rising tempers require the Noble Line to either stay more advanced than local rivals or train that much harder."
		eventChar := dyn.safestOption(8, charCleverness)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Dynasty missed the opportunity."
		if eventSuccess {
			bonus := utils.RandomFromList([]string{traitTerritorialDefence, traitTechnology})
			dyn.setBonus(bonus, 1)
			eventOutcome = "(+1 " + bonus + ")"
		}

	case "546":
		//A major catastrophe kills off a dozen important members of the heirs apparent, causing some major upheaval from within. Roll Loyalty 8+ to keep the arguments and in-fighting to a minimum, gaining +1 Tradition or +2 Morale.
		eventDescr = "A major catastrophe kills off a dozen important members of the heirs apparent, causing some major upheaval from within."
		eventChar := dyn.safestOption(8, charLoyalty)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Crisis was long-lasting."
		if !eventSuccess {
			if utils.RandomBool() {
				dyn.setBonus(charTradition, 1)
				eventOutcome = "Arguments kept at minimum (+1 Tradition)"
			} else {
				dyn.setBonus(valueMorale, 2)
				eventOutcome = "Arguments kept at minimum (+2 Morale)"
			}

		}

	case "551":
		//An interplanetary scholastic service receives a massive donation from the family; spend –1 Wealth to gain +1 to any three Aptitudes.
		eventDescr = "An interplanetary scholastic service receives a massive donation from the family."

		dyn.setBonus(valueWealth, -1)
		var things []string
		for len(things) < 3 {
			things = utils.AppendUniqueStr(things, utils.RandomFromList(listAPTITUDES))
		}
		for i := range things {
			dyn.setBonus(things[i], 1)
		}
		eventOutcome = "(-1 Wealth, +1 " + things[0] + ", +1 " + things[1] + ", +1" + things[2] + ")"

	case "552":
		//War breaks out near the family territory, drafting several young men into the dangerous ranks; roll Militarism 8+ or lose –1 Populace but gain +1 Hostility regardless of the outcome.
		eventDescr = "War breaks out near the family territory, drafting several young men into the dangerous ranks."
		eventChar := dyn.safestOption(8, charMilitarism)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "("
		if !eventSuccess {
			dyn.setBonus(valuePopulace, -1)
			eventOutcome += "-1 " + valuePopulace + ","
		}
		dyn.setBonus(apttHostility, 1)
		eventOutcome += "+1 Hostility)"

	case "553":
		//Everything goes as planned for decades; add +1 to any Aptitude or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "554":
		dyn.RollHistoricEvent()

	case "555":
		//Assassins and hit men are targeting leaders in the family; roll Security 8+ to gain +1 Territorial Defence.
		eventDescr = "Assassins and hit men are targeting leaders in the family."
		eventChar := dyn.safestOption(8, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		if eventSuccess {
			dyn.setBonus(traitTerritorialDefence, -1)
			eventOutcome = "Protective measures enchanced (+1 Territorial Defence)"
		}

	case "556":
		//The family’s primary territory is at risk from unconquerable sources, forcing the Noble Line to relocate; gain +1 Fleet and +1 Intel but lose –1 Territorial Defence.
		eventDescr = "The family’s primary territory is at risk from unconquerable sources, forcing the Noble Line to relocate."
		dyn.setBonus(traitFleet, 1)
		dyn.setBonus(apttIntel, 1)
		dyn.setBonus(traitTerritorialDefence, -1)
		eventOutcome = "(+1 Fleet, +1 Intel, –1 Territorial Defence)"

	case "561":
		//The games between forms of royalty are thick and filled with risks; roll Scheming or Cleverness 9+ to gain +1 to any Characteristic.
		eventDescr = "The games between forms of royalty are thick and filled with risks."
		eventChar := dyn.safestOption(9, charCleverness, charScheming)
		eventSuccess := dyn.failureCheck(eventChar, 9)
		eventOutcome = "Opportunity was missed"
		if eventSuccess {
			bonus := utils.RandomFromList(listCHARS)
			dyn.setBonus(bonus, 1)
			eventOutcome = "Opportunity was exploited (+1 " + bonus + ")."
		}

	case "562":
		//Victories in a long-lasting vendetta fill the Noble Line with joy and pride; Gain +1 to any Trait or Value.
		eventDescr = "Victories in a long-lasting vendetta fill the Noble Line with joy and pride."
		things := listTRAITS
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "563":
		//Things could not go any better for the Dynasty; add +1 to any Characteristic, Aptitude, Trait or Value.
		eventDescr = "Things could not go any better for the Dynasty."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		things = append(things, listTRAITS...)
		things = append(things, listCHARS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "564":
		dyn.RollHistoricEvent()

	case "565":
		//There is a dangerous disease that strikes at the youngest members of the Noble Line but it is overcome through persistence and increased breeding habits; Gain +2 Tenacity or +1 to all Values.
		eventDescr = "There is a dangerous disease that strikes at the youngest members of the Noble Line but it is overcome through persistence and increased breeding habits."
		if utils.RandomBool() {
			eventOutcome = "(+2 Popularity)"
			dyn.setBonus(charTenacity, 2)
		} else {
			for i := range listVALUES {
				dyn.setBonus(listVALUES[i], 1)
			}
			eventOutcome = "(+1 to all values)"

		}

	case "566":
		//A true genius assumes control of the Noble Line’s accounts and finances sector-wide; Gain +1 to Fiscal Defence, Fleet and Culture.
		eventDescr = "A true genius assumes control of the Noble Line’s accounts and finances sector-wide."
		eventOutcome = "(+1 to Fiscal Defence, Fleet and Culture)"
		dyn.setBonus(traitFiscalDefence, 1)
		dyn.setBonus(traitFleet, 1)
		dyn.setBonus(traitCulture, 1)

	case "611":
		//Atheists are in control for a long time; lose –1 from the Value of your choice.
		eventDescr = "Atheists are in control for a long time."
		val := utils.RandomFromList(listVALUES)
		dyn.setBonus(val, -1)
		eventOutcome = "(-1 " + val + ")"

	case "612":
		//Surrounded by dangerous planets, missionaries are at great risk; roll Conquest 7+ or lose –1 Populace.
		eventDescr = "Surrounded by dangerous planets, missionaries are at great risk."
		eventChar := dyn.safestOption(7, apttConquest)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "Faithful withstanded."
		if !eventSuccess {
			dyn.setBonus(valuePopulace, -1)
			eventOutcome = "Many faithful are lost (-1 Populance)."
		}

	case "613":
		//Convincing arguments against the faith have put several clergy centres on edge, having to try very hard to keep the congregation faithful. Roll Loyalty 8+ or lose –1 Populace and –1 Morale.
		eventDescr = "Convincing arguments against the faith have put several clergy centres on edge, having to try very hard to keep the congregation faithful."
		eventChar := dyn.safestOption(8, charLoyalty)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Apostates were punished."
		if !eventSuccess {
			dyn.setBonus(valuePopulace, -1)
			dyn.setBonus(valueMorale, -1)
			eventOutcome = "Some faithful were corrupt (-1 Populance, -1 Morale)."
		}

	case "614":
		dyn.RollHistoricEvent()

	case "615":
		//Anti-theocracy rebels are targeting temples and shrines; roll Security 8+ or lose –1 Wealth.
		eventDescr = "Anti-theocracy rebels are targeting temples and shrines."
		eventChar := dyn.safestOption(8, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Apostates were captured."
		if !eventSuccess {
			dyn.setBonus(valueWealth, -1)
			eventOutcome = "Many temples were looted (-1 Wealth)."
		}

	case "616":
		//Scandals rock the media concerning the Religious Faith; lose –1 Popularity.
		eventDescr = "Scandals rock the media concerning the Religious Faith."
		dyn.setBonus(charPopularity, -1)
		eventOutcome = "(-1 popularity)."

	case "621":
		//Aliens desire to learn about the Faith; Roll Tutelage 8+ to gain +1 Culture.
		eventDescr = "Aliens desire to learn about the Faith."
		eventChar := dyn.safestOption(8, apttTutelage)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Few could grasp the Truth."
		if eventSuccess {
			dyn.setBonus(traitCulture, 1)
			eventOutcome = "Many of them touched the Truth (+1 Culture)."
		}

	case "622":
		//A holy crusade pushes the boundaries to the neighbouring sectors; gain +1 Fleet or +1 Technology.
		eventDescr = "A holy crusade pushes the boundaries to the neighbouring sectors."
		if utils.RandomBool() { //TODO: ПРИНЯТЬ РЕШЕНИЕ
			eventOutcome = "A Holy Fleet is assembled. (+1 Fleet)"
			dyn.setBonus(traitFleet, 1)
		} else {
			eventOutcome = "Many knowledge came to faithful. (+1 Technology)"
			dyn.setBonus(traitTechnology, 1)
		}

	case "623":
		//An interstellar celestial event is prophesised in the ancient scriptures, giving great credence to the Religious Faith; Roll Public Relations 8+ to gain +1 Populace.
		eventDescr = "An interstellar celestial event is prophesised in the ancient scriptures, giving great credence to the Religious Faith."
		eventChar := dyn.safestOption(8, apttPublicRelations)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Few joined to our Faith."
		if eventSuccess {
			dyn.setBonus(valuePopulace, 1)
			eventOutcome = "Many joined to our Faith. (+1 Populance)"
		}

	case "624":
		dyn.RollHistoricEvent()

	case "625":
		//There is a swell in fanatics and zealots in the congregation. Roll Hostility 8+ to incite them against all foes, gaining +1 Militarism if successful
		//; or roll Maintenance 8+ to keep them under better control, gaining +1 Loyalty instead. Fail either roll and lose –1 Popularity as the zealots run rampant.
		eventDescr = "There is a swell in fanatics and zealots in the congregation."
		eventChar := dyn.safestOption(8, apttHostility, apttMaintenance)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		if eventSuccess {
			if eventChar == apttHostility {
				eventOutcome = "We incited them against all foes (+1 Militarism)."
				dyn.setBonus(charMilitarism, 1)
			} else {
				eventOutcome = "We kept them under better control (+1 Loyalty)."
				dyn.setBonus(charLoyalty, 1)
			}
		} else {
			eventOutcome = "The zealots run rampant (-1 Popularity)."
			dyn.setBonus(charPopularity, -1)
		}

	case "626":
		//A cosmic ‘sign’ draws new followers from all over; Roll Recruit 8+ to gain +1 Populace.
		eventDescr = "A cosmic 'sign' draws new followers from all over."
		eventChar := dyn.safestOption(8, apttRecruitment)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "But most are left."
		if eventSuccess {
			dyn.setBonus(valuePopulace, 1)
			eventOutcome = "And many stayed (+1 Populance)."
		}

	case "631":
		//There is a lot of money to be had from within the church; Roll Economics 7+ to gain +1 Fiscal Defence.
		eventDescr = "There is a lot of money to be had from within the church."
		eventChar := dyn.safestOption(7, apttEconomics)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = ""
		if eventSuccess {
			dyn.setBonus(traitFiscalDefence, 1)
			eventOutcome = "We should put them in good use (+1 Fiscal Defence)."
		}

	case "632":
		//True believers question the source of the Religious Faith, giving the clergy an opportunity to shine against obstacles; roll Tradition 8+ or lose –1 Morale.
		eventDescr = "True believers question the source of the Religious Faith."
		eventChar := dyn.safestOption(8, charTradition)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "The clergy shines against the heresy."
		if !eventSuccess {
			dyn.setBonus(valueMorale, -1)
			eventOutcome = "The clergy could not shine the doubts (-1 Morale)."
		}

	case "633":
		//Everything goes as planned for decades; add +1 to any Trait or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listTRAITS
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "634":
		dyn.RollHistoricEvent()

	case "635":
		//The Religious Faith opens a private school that does not require membership in the church to be a student;
		//spend one point of Culture or Morale to gain +1 Wealth and +1 to an Aptitude of the player’s choice.
		eventDescr = "The Religious Faith opens a private school that does not require membership in the church to be a student."
		eventOutcome = "Atheists are spreading within the dynasty."
		if utils.RandomBool() {
			eventOutcome += " (-1 Culture, "
			dyn.setBonus(traitCulture, -1)
		} else {
			eventOutcome += " (-1 Morale, "
			dyn.setBonus(valueMorale, -1)
		}
		if utils.RandomBool() {
			eventOutcome += "+1 Wealth)."
			dyn.setBonus(valueWealth, 1)
		} else {
			apt := utils.RandomFromList(listAPTITUDES)
			eventOutcome += "+1 " + apt + ")."
			dyn.setBonus(apt, -1)
		}

	case "636":
		//Science and religion meld into one belief; Gain +1 Technology or +1 Research.
		eventDescr = "Science and religion meld into one belief."
		if utils.RandomBool() {
			dyn.setBonus(traitTechnology, 1)
			eventOutcome = "(+1 Technology)"
		} else {
			dyn.setBonus(apttResearch, 1)
			eventOutcome = "(+1 Research)"
		}

	case "641":
		//The opportunity to create a true martyr presents itself;
		//allow this to happen and roll Propaganda 7+, gaining +1 Popularity if successful, losing –1 Loyalty if not. Save the martyr instead and gain +1 Morale.
		eventDescr = "The opportunity to create a true martyr presents itself."
		eventChar := dyn.safestOption(7, apttPropaganda)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		if utils.RandomBool() {
			if eventSuccess {
				dyn.setBonus(charPopularity, 1)
				eventOutcome = "Many zealots joined the church (+1 Popularity)"
			} else {
				dyn.setBonus(charLoyalty, -1)
				eventOutcome = "Faithful saw it as betrayal (-1 Loyalty)"
			}
		} else {
			dyn.setBonus(valueMorale, 1)
			eventOutcome = "Church saved the martyr (+1 Morale)"
		}

	case "642":
		//The government is thoroughly mired in the Religious Faith’s doctrines and dogma; gain +1 Militarism, +1 Popularity or +1 to any two Traits.
		eventDescr = "The government is thoroughly mired in the Religious Faith’s doctrines and dogma."
		r := utils.RollDice("d3")
		switch r {
		case 1:
			dyn.setBonus(charMilitarism, 1)
			eventOutcome = "(+1 Militarism)"
		case 2:
			dyn.setBonus(charPopularity, 1)
			eventOutcome = "(+1 Popularity)"
		case 3:
			var traits []string
			for len(traits) < 2 {
				traits = utils.AppendUniqueStr(traits, utils.RandomFromList(listTRAITS))
			}
			dyn.setBonus(traits[0], 1)
			dyn.setBonus(traits[1], 1)
			eventOutcome = "(+1 " + traits[0] + ", +1 " + traits[1] + ")"
		}

	case "643":
		//Everything goes as planned for decades; add +1 to any Aptitude or Trait.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listTRAITS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "644":
		dyn.RollHistoricEvent()

	case "645":
		//The Religious Faith can reach out to the masses through art; roll Entertain 8+ to gain +1 Popularity. Fail and lose –1 Populace.
		eventDescr = "The Religious Faith can reach out to the masses through art."
		eventChar := dyn.safestOption(8, apttEntertain)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		if eventSuccess {
			dyn.setBonus(charPopularity, 1)
			eventOutcome = "New followers are comming (+1 Popularity)."
		} else {
			dyn.setBonus(valuePopulace, -1)
			eventOutcome = "Some faithful were discusted (-1 Populance)."
		}

	case "646":
		//The church has grown far too large to not be considered a business at this point; Gain +1 Greed or increase Bureaucracy by +1.
		eventDescr = "The church has grown far too large to not be considered a business at this point."
		if utils.RandomBool() {
			dyn.setBonus(apttBureaucracy, 1)
			eventOutcome = "(+1 Bureacracy)"
		} else {
			dyn.setBonus(charGreed, 1)
			eventOutcome = "(+1 Greed)"
		}

	case "651":
		//Hateful warmongers seek to start a war with the church by any means they can; Roll Security 8+ to increase Territorial Defence by +1.
		eventDescr = "Hateful warmongers seek to start a war with the church by any means they can."
		eventChar := dyn.safestOption(8, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "The church is ready!"
		if eventSuccess {
			dyn.setBonus(traitTerritorialDefence, 1)
			eventOutcome = "The church is ready! (+1 Territorial Defence)"
		}

	case "652":
		//The galaxy begins to see the power behind the Religious Faith’s congregation; Gain +1 Posturing.
		eventDescr = "The galaxy begins to see the power behind the Religious Faith's congregation."
		eventOutcome = "(+1 Posturing)"
		dyn.setBonus(apttPosturing, 1)

	case "653":
		//Everything goes as planned for decades; add +1 to any Aptitude or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "654":
		dyn.RollHistoricEvent()

	case "655":
		//A neighbouring culture is in possession of several important holy artefacts; Roll Acquisition 8+ to increase Wealth or Culture by +1.
		eventDescr = "A neighbouring culture is in possession of several important holy artefacts."
		eventChar := dyn.safestOption(8, apttAcquisition)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "The church could not retrieve them."
		if eventSuccess {
			if utils.RandomBool() {
				dyn.setBonus(valueWealth, 1)
				eventOutcome = "Artifacts retrived (+1 Wealth)"
			} else {
				dyn.setBonus(traitCulture, 1)
				eventOutcome = "Artifacts retrived (+1 Culture)"
			}
		}

	case "656":
		//A major celestial event brings about a holiday celebration that lasts weeks, putting a new face on the Religious Faith if hosted well.
		//Roll Entertain 7+ or Public Relations 9+; success adds +1 to Popularity or Loyalty.
		eventDescr = "A major celestial event brings about a holiday celebration that lasts weeks, putting a new face on the Religious Faith if hosted well."
		chance1 := successOf2d6(difficultyAverage+dyn.pickVal(apttEntertain), 7)
		chance2 := successOf2d6(difficultyHard+dyn.pickVal(apttPublicRelations), 9)
		eventSuccess := false
		if chance1 > chance2 {
			eventSuccess = dyn.failureCheck(apttEntertain, 7)
		} else {
			eventSuccess = dyn.failureCheck(apttPublicRelations, 9)
		}
		eventOutcome = "Holiday was failure."
		if eventSuccess {
			bonus := utils.RandomFromList([]string{charPopularity, charLoyalty})
			dyn.setBonus(bonus, 1)
			eventOutcome = "Holyday was massive success (+1 " + bonus + ")."
		}

	case "661":
		//The enemies of the church have pushed things too far, forcing the Religious Faith to train ‘holy warriors’ and ‘godly assassins.’
		//Roll Militarism 7+ to increase Conquest, Hostility or Security by +1.
		eventDescr = "The enemies of the church have pushed things too far, forcing the Religious Faith to train 'holy warriors' and 'godly assassins'."
		eventChar := dyn.safestOption(7, charMilitarism)
		eventSuccess := dyn.failureCheck(eventChar, 7)
		eventOutcome = "Yet those were not effective."
		if eventSuccess {
			bonus := utils.RandomFromList([]string{apttConquest, apttHostility, apttSecurity})
			dyn.setBonus(bonus, 1)
			eventOutcome = "Those were effective (+1 " + bonus + ")."
		}

	case "662":
		//Sometimes it takes unsavoury behaviour to do god’s work; Gain +1 Illicit or Sabotage.
		eventDescr = "Sometimes it takes unsavoury behaviour to do god's work."
		bonus := utils.RandomFromList([]string{apttIllicit, apttSabotage})
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "663":
		//Things could not go any better for the Dynasty; add +1 to any Characteristic, Aptitude, Trait or Value.
		eventDescr = "Things could not go any better for the Dynasty."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		things = append(things, listTRAITS...)
		things = append(things, listCHARS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "664":
		dyn.RollHistoricEvent()

	case "665":
		//Monastic learners focus on the Religious Faith’s only weaknesses; raise any two Aptitudes from ‘—’ to 1.
		eventDescr = "Monastic learners focus on the Religious Faith’s only weaknesses."
		var possible []string
		for i := range listAPTITUDES {
			if dyn.pickVal(listAPTITUDES[i]) < 0 {
				possible = append(possible, listAPTITUDES[i])
			}
		}
		if len(possible) > 2 {
			var bonus []string
			for len(bonus) < 2 {
				bonus = utils.AppendUniqueStr(bonus, utils.RandomFromList(possible))
			}
			dyn.setBase(bonus[0], 1)
			dyn.setBase(bonus[1], 1)
			eventOutcome = "(+1 " + bonus[0] + ", +1 " + bonus[1] + ")"
		} else {
			for i := range possible {
				dyn.setBase(possible[i], 1)
				eventOutcome += "(+1 " + possible[i] + ")"
			}
		}

	case "666":
		//Witness to an undeniable miracle; Gain +1 to any two Characteristics.
		eventDescr = "Witness to an undeniable miracle."
		var bonus []string
		for len(bonus) < 2 {
			bonus = utils.AppendUniqueStr(bonus, utils.RandomFromList(listCHARS))
		}
		dyn.setBonus(bonus[0], 1)
		dyn.setBonus(bonus[1], 1)
		eventOutcome = "(+1 " + bonus[0] + ", +1 " + bonus[1] + ")"

		//SYNDICATE

	case "711":
		//Interstellar authorities are bent on bringing the Syndicate down; Lose –1 Wealth and –1 Morale.
		eventDescr = "Interstellar authorities are bent on bringing the Syndicate down."
		dyn.setBonus(valueWealth, -1)
		dyn.setBonus(valueMorale, -1)
		eventOutcome = "(-1 Wealth, -1 Morale)"

	case "712":
		//Street-level criminals have started to organise into mini-syndicates; Roll Hostility 8+ to scare them back into complacency or lose –1 Loyalty.
		eventDescr = "Street-level criminals have started to organise into mini-syndicates."
		eventChar := dyn.safestOption(8, apttHostility)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Opposition suppressed."
		if !eventSuccess {
			dyn.setBonus(charLoyalty, -1)
			eventOutcome = "Opposition arised (-1 Loyalty)."
		}

	case "713":
		//The Syndicate has double-agents, spies and backstabbers in its midst; Roll Security 8+ or lose –1 Territorial Defence.
		eventDescr = "The Syndicate has double-agents, spies and backstabbers in its midst."
		eventChar := dyn.safestOption(8, apttSecurity)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Double agent killed."
		if !eventSuccess {
			dyn.setBonus(traitTerritorialDefence, -1)
			eventOutcome = "Agent revealed some data before was killed (-1 Territorial Defence)."
		}

	case "714":
		dyn.RollHistoricEvent()

	case "715":
		//The locals are tired of being preyed upon by petty crime; Roll Public Relations 8+ or lose –1 Popularity.
		eventDescr = "The locals are tired of being preyed upon by petty crime."
		eventChar := dyn.safestOption(8, apttPublicRelations)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "PR relived the tensions."
		if !eventSuccess {
			dyn.setBonus(valueWealth, -1)
			eventOutcome = "Tentions cost considerable resources (-1 Wealth)."
		}

	case "716":
		//Bounty hunters have taken a renewed interest in the Syndicate’s leadership; Roll Illicit 8+ or lose –1 from a Characteristic of the Player’s choice.
		eventDescr = "Bounty hunters have taken a renewed interest in the Syndicate’s leadership."
		eventChar := dyn.safestOption(8, apttIllicit)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Bounty was removed."
		if !eventSuccess {
			chr := utils.RandomFromList(listCHARS)
			dyn.setBonus(chr, -1)
			eventOutcome = "Some key members were killed (-1 " + chr + ")"
		}

	case "721":
		//A pyramid scam promises to be extremely successful; Roll Economics 8+ to gain +1 Wealth.
		eventDescr = "A pyramid scam promises to be extremely successful."
		eventChar := dyn.safestOption(8, apttEconomics)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Scam was revealed to quickly."
		if eventSuccess {
			dyn.setBonus(valueWealth, 1)
			eventOutcome = "Scam was successful (+1 Wealth)."
		}

	case "722":
		//There is an army storage surplus near the Syndicate’s power base, begging to be pilfered by professionals!
		//Roll Illicit 8+ to gain +1 Militarism or +2 Territorial Defence.
		eventDescr = "There is an army storage surplus near the Syndicate’s power base, begging to be pilfered by professionals!"
		eventChar := dyn.safestOption(8, apttIllicit)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No significant loot was found."
		if eventSuccess {
			if utils.RandomBool() {
				dyn.setBonus(charMilitarism, 1)
				eventOutcome = "Some new weaponry retrived (+1 Militarism)."
			} else {
				dyn.setBonus(traitTerritorialDefence, 2)
				eventOutcome = "Some new weaponry retrived (+2 Territorial Defence)."
			}

		}

	case "723":
		//A pirate cell leader wants to join the Syndicate but for a price; spend 1 point of Wealth to gain +1 Fleet.
		eventDescr = "A pirate cell leader wants to join the Syndicate but for a price."
		dyn.setBonus(traitFleet, 1)
		dyn.setBonus(valueWealth, -1)
		eventOutcome = "(+1 Fleet, -1 Wealth)"

	case "724":
		dyn.RollHistoricEvent()

	case "725":
		//A marriage between crime families can solidify efforts and increase the strength of both; Gain +1 Loyalty, +1 Culture or +2 Morale.
		eventDescr = "A marriage between crime families can solidify efforts and increase the strength of both."
		r := utils.RollDice("d3")
		switch r {
		case 1:
			dyn.setBonus(charLoyalty, 1)
			eventOutcome = "(+1 Loyalty)"
		case 2:
			dyn.setBonus(traitCulture, 1)
			eventOutcome = "(+1 Culture)"
		case 3:

			dyn.setBonus(valueMorale, 2)
			eventOutcome = "(+2 Morale)"
		}

	case "726":
		//Arms smuggling has just been targeted by trade authorities, making it a lucrative endeavour. Roll Illicit 8+ or Tactics 7+ to gain +1 Militarism or +1 Fleet.
		eventDescr = "Arms smuggling has just been targeted by trade authorities, making it a lucrative endeavour."
		chance1 := successOf2d6(difficultyAverage+dyn.pickVal(apttIllicit), 8)
		chance2 := successOf2d6(difficultyHard+dyn.pickVal(apttTactical), 7)
		eventSuccess := false
		if chance1 > chance2 {
			eventSuccess = dyn.failureCheck(apttIllicit, 8)
		} else {
			eventSuccess = dyn.failureCheck(apttTactical, 7)
		}
		eventOutcome = "Syndicate adopted new ways."
		if eventSuccess {
			bonus := utils.RandomFromList([]string{charMilitarism, traitFleet})
			dyn.setBonus(bonus, 1)
			eventOutcome = "Syndicate adopted new ways (+1 " + bonus + ")."
		}

	case "731":
		//Interplanetary sports gambling rings are good business, especially for those who can rig the games.
		//If you just run the numbers, roll Illicit 8+; if you rig the events, roll Sabotage 8+. Succeed and gain +1 Wealth or Morale; fail and lose –1 Popularity.
		eventDescr = "Interplanetary sports gambling rings are good business, especially for those who can rig the games."
		eventChar := dyn.safestOption(8, apttIllicit, apttSabotage)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Fraud was revealed (-1 Popularity)"
		if eventSuccess {
			if eventChar == apttIllicit {
				eventOutcome = "Syndicate just run the numbers."
			} else {
				eventOutcome = "Syndicate rig the events."
			}
			bonus := utils.RandomFromList([]string{valueMorale, valueWealth})
			dyn.setBonus(bonus, 1)
			eventOutcome += "(+1 " + bonus + ")"
		}

	case "732":
		//Even a Syndicate has some legitimate businesses with which it makes some profits; Roll Bureaucracy or Economics 8+ to gain +1 Greed.
		eventDescr = "Even a Syndicate has some legitimate businesses with which it makes some profits."
		eventChar := dyn.safestOption(8, apttBureaucracy, apttEconomics)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Profit is not huge though."
		if !eventSuccess {
			dyn.setBonus(charGreed, 1)
			eventOutcome = "Profit is good enough. (+1 Greed)"
		}

	case "733":
		//Everything goes as planned for decades; add +1 to any Trait or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listTRAITS
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "734":
		dyn.RollHistoricEvent()

	case "735":
		//There is a minor coup within the ranks of the Syndicate’s subordinates; Roll Loyalty 8+ or lose –1 Populace.
		eventDescr = "There is a minor coup within the ranks of the Syndicate's subordinates."
		eventChar := dyn.safestOption(8, charLoyalty)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Unrest was suppressed."
		if !eventSuccess {
			dyn.setBonus(valuePopulace, -1)
			eventOutcome = "Some syndicate members left. (-1 Populance)"
		}

	case "736":
		//Powerful enemies of the Syndicate are ready to wage an open war against it; Roll Hostility 8+ to fight them successfully, gaining +1 Militarism or +2 Morale.
		eventDescr = "Powerful enemies of the Syndicate are ready to wage an open war against it."
		eventChar := dyn.safestOption(8, apttHostility)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "War draines a lot of resourses."
		if eventSuccess {
			dyn.setBonus(valueMorale, 2)
			dyn.setBonus(charMilitarism, 1)
			eventOutcome = "Opposition was fought off (+1 Militarism, +2 Morale)"
		}

	case "741":
		//A naturally stealthy alien species wants to sell its services to the Syndicate; spend 1 Wealth to gain +1 Illicit and Sabotage.
		eventDescr = "A naturally stealthy alien species wants to sell its services to the Syndicate."
		eventOutcome = "These recruits proved to be quite effective (+1 Illicit, +1 Sabotage, -1 Wealth)."
		dyn.setBonus(valueWealth, -1)
		dyn.setBonus(apttSabotage, 1)
		dyn.setBonus(apttIllicit, 1)

	case "742":
		//A primitive species protects vast amounts of precious metals to be exploited.
		//Trade with them only slightly in your favour with Public Relations 8+ and gain +1 Loyalty; take extreme advantage of their naivety with Posturing 8+ and gain +1 Wealth.
		eventDescr = "A primitive species protects vast amounts of precious metals to be exploited."
		eventChar := dyn.safestOption(8, apttPublicRelations, apttPosturing)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Syndicate's enemys took advantage faster."
		if eventSuccess {
			switch eventChar {
			case apttPublicRelations:
				eventOutcome = "Syndicate trade with them only slightly in it's favour (+1 Loyalty)"
				dyn.setBonus(charLoyalty, 1)
			case apttPosturing:
				eventOutcome = "Syndicate took extreme advantage of their naivety (+1 Wealth)"
				dyn.setBonus(valueWealth, 1)
			}
		}

	case "743":
		//Everything goes as planned for decades; add +1 to any Aptitude or Trait.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listTRAITS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "744":
		dyn.RollHistoricEvent()

	case "745":
		//A law enforcement agency is willing to feed the Syndicate information in order to avoid falling victim to their schemes; Gain +1 Intel, Politics or Research.
		eventDescr = "A law enforcement agency is willing to feed the Syndicate information in order to avoid falling victim to their schemes."
		bonus := utils.RandomFromList([]string{apttIntel, apttPolitics, apttResearch})
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "746":
		//The local government has asked the Syndicate to send hit men after one of their targets – as long as they do not leave any clues behind.
		//Roll Hostility 8+ or Illicit 9+ to gain +1 to any Trait.
		eventDescr = "The local government has asked the Syndicate to send hit men after one of their targets – as long as they do not leave any clues behind."
		chance1 := successOf2d6(difficultyAverage+dyn.pickVal(apttHostility), 8)
		chance2 := successOf2d6(difficultyHard+dyn.pickVal(apttIllicit), 9)
		eventSuccess := false
		if chance1 > chance2 {
			eventSuccess = dyn.failureCheck(apttHostility, 8)
		} else {
			eventSuccess = dyn.failureCheck(apttIllicit, 9)
		}
		eventOutcome = "Deal was a set up."
		if eventSuccess {
			bonus := utils.RandomFromList(listTRAITS)
			dyn.setBonus(bonus, 1)
			eventOutcome = "Deal was very beneficial (+1 " + bonus + ")."
		}

	case "751":
		//In order to pick up the pieces, the Syndicate must first smash the local businesses; Roll Conquest or Economics 8+ to gain +1 Wealth.
		eventDescr = "In order to pick up the pieces, the Syndicate must first smash the local businesses."
		eventChar := dyn.safestOption(8, apttConquest, apttEconomics)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Resistiance was met."
		if eventSuccess {
			dyn.setBonus(valueWealth, 1)
			eventOutcome = "Some property was aquired (+1 Wealth)"
		}

	case "752":
		//The Syndicate has a chance to manipulate the government to make a lot of their activities ‘less-illegal’ and therefore safer for its members to undertake.
		//Roll Politics or Posturing 8+ to gain +1 Populace.
		eventDescr = "The Syndicate has a chance to manipulate the government to make a lot of their activities 'less-illegal' and therefore safer for its members to undertake."
		eventChar := dyn.safestOption(8, apttPolitics, apttPosturing)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Lobbying failed."
		if eventSuccess {
			dyn.setBonus(valuePopulace, 1)
			eventOutcome = "Some bribes were properly placed (+1 Populance)"
		}

	case "753":
		//Everything goes as planned for decades; add +1 to any Aptitude or Value.
		eventDescr = "Everything goes as planned for decades."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "754":
		dyn.RollHistoricEvent()

	case "755":
		//An infamous crime family wants to join efforts with the Syndicate; Roll Maintenance 8+ to gain +1 Tradition or +2 Culture.
		eventDescr = "An infamous crime family wants to join efforts with the Syndicate."
		eventChar := dyn.safestOption(8, apttMaintenance)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "Negotiations failed."
		if eventSuccess {
			if utils.RandomBool() {
				dyn.setBonus(charTradition, 1)
				eventOutcome = "New family was quite resourceful (+1 Tradition)"
			} else {
				dyn.setBonus(traitCulture, 2)
				eventOutcome = "New family was quite resourceful (+2 Culture)"
			}
		}

	case "756":
		//The best way to beat the authorities is to stay one step ahead of them scientifically-speaking. Roll Research 8+ to gain +1 Technology.
		eventDescr = "The best way to beat the authorities is to stay one step ahead of them scientifically-speaking."
		eventChar := dyn.safestOption(8, apttResearch)
		eventSuccess := dyn.failureCheck(eventChar, 8)
		eventOutcome = "No breakthrough was made."
		if eventSuccess {
			dyn.setBonus(traitTechnology, 1)
			eventOutcome = "New devices were invented (+1 Technology)"
		}

	case "761":
		//A rival is teetering on the edge of existence. Wipe them out with Hostility 7+ to gain +1 Wealth and +1 Morale.
		//Bring them into the fold as a show of mercy with Acquisition 9+ to gain +1 to any Characteristic.
		eventDescr = "A rival is teetering on the edge of existence."
		chance1 := successOf2d6(difficultyAverage+dyn.pickVal(apttHostility), 7)
		chance2 := successOf2d6(difficultyHard+dyn.pickVal(apttAcquisition), 9)
		eventSuccess := false
		eventChar := ""
		if chance1 > chance2 {
			eventSuccess = dyn.failureCheck(apttHostility, 7)
			eventChar = apttHostility
		} else {
			eventSuccess = dyn.failureCheck(apttAcquisition, 9)
			eventChar = apttAcquisition
		}
		eventOutcome = "Rival could survive."
		if eventSuccess {
			switch eventChar {
			case apttHostility:
				eventOutcome = "Rival was wiped out (+1 Wealth, +1 Morale)."
				dyn.setBonus(valueWealth, 1)
				dyn.setBonus(valueMorale, 1)
			case apttAcquisition:
				bonus := utils.RandomFromList(listCHARS)
				dyn.setBonus(bonus, 1)
				eventOutcome = "Rival was merged (+1 " + bonus + ")."
			}

		}

	case "762":
		//Psions willing to use their talents for crime and profit approach the Syndicate for work; Gain +1 to Territorial Defence or +1 Wealth.
		eventDescr = "Psions willing to use their talents for crime and profit approach the Syndicate for work."
		if utils.RandomBool() {
			eventOutcome = "recruits proved to be useful (+1 Wealth)."
			dyn.setBonus(valueWealth, 1)
		} else {
			eventOutcome = "recruits proved to be useful (+1 Territorial Defence)."
			dyn.setBonus(traitTerritorialDefence, 1)
		}

	case "763":
		//Things could not go any better for the Dynasty; add +1 to any Characteristic, Aptitude, Trait or Value.
		eventDescr = "Things could not go any better for the Dynasty."
		things := listAPTITUDES
		things = append(things, listVALUES...)
		things = append(things, listTRAITS...)
		things = append(things, listCHARS...)
		bonus := utils.RandomFromList(things)
		dyn.setBonus(bonus, 1)
		eventOutcome = "(+1 " + bonus + ")"

	case "764":
		dyn.RollHistoricEvent()

	case "765":
		//The Syndicate’s leader is considered a true supervillain by the forces of authority; Gain +1 to any three Aptitudes or Traits.
		eventDescr = "The Syndicate’s leader is considered a true supervillain by the forces of authority."
		things := listAPTITUDES
		things = append(things, listTRAITS...)
		var bonuses []string
		for len(bonuses) < 3 {
			bonuses = utils.AppendUniqueStr(bonuses, utils.RandomFromList(things))
		}
		for i := range bonuses {
			dyn.setBonus(bonuses[i], 1)
		}
		eventOutcome = "Syndicate is thriving (+1 " + bonuses[0] + ", +1" + bonuses[1] + ", +1" + bonuses[2] + ", +1)"

	case "766":
		//The greatest crime of its time! Gain +1 to any two Characteristics.
		eventDescr = "The greatest crime of its time!"
		var bonus []string
		for len(bonus) < 2 {
			bonus = utils.AppendUniqueStr(bonus, utils.RandomFromList(listCHARS))
		}
		dyn.setBonus(bonus[0], 1)
		dyn.setBonus(bonus[1], 1)
		eventOutcome = "(+1 " + bonus[0] + ", +1 " + bonus[1] + ")"

	}
	eventLog := eventDescr + " " + eventOutcome
	fmt.Println(eventLog)
	return eventLog
}

func (dyn *Dynasty) safestOption(tn int, atr ...string) string {
	eventMap := make(map[string]float64)
	for i := range atr {
		eventMap[atr[i]] = dyn.probeFailureCheck(atr[i], tn)
		fmt.Println(atr[i], ":", eventMap[atr[i]])
	}
	bestKey := "nothing"
	bestVal := -1.0
	for key, val := range eventMap {
		if val >= bestVal {
			bestKey = key
			bestVal = val
		}
	}
	return bestKey
}

func isCharacteristic(val string) bool {
	for i := range listCHARS {
		if val == listCHARS[i] {
			return true
		}
	}
	return false
}
