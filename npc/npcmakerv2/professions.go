package npcmakerv2

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/entity"
	"github.com/Galdoba/utils"
)

const (
	CareerB0AgentLawenforcement = "B0=1=1==Law Enforcement Agent"
	CareerB0Agentintelligence   = "B0=1=2==Intelligence Agent"
	CareerB0AgentCorporate      = "B0=1=3==Corporate Agent"
	////
	STR          = entity.CharCodeTrvC1STRENGTH
	DEX          = entity.CharCodeTrvC2DEXTERITY
	END          = entity.CharCodeTrvC3ENDURANCE
	INT          = entity.CharCodeTrvC4INTELLIGENCE
	EDU          = entity.CharCodeTrvC5EDUCATION
	Advocate     = entity.SCTrvAdvocate
	Athletics    = entity.SCTrvAthletics
	Comms        = entity.SCTrvElectronicsComms
	Computers    = entity.SCTrvElectronicsComputers
	Deception    = entity.SCTrvDeception
	Drive        = entity.SCTrvDrive
	Gambler      = entity.SCTrvGambler
	GunCombat    = entity.SCTrvGuncombat
	Investigate  = entity.SCTrvInvestigate
	Medic        = entity.SCTrvMedic
	Melee        = entity.SCTrvMelee
	MeleeUnarmed = entity.SCTrvMeleeUnarmed
	MeleeBlade   = entity.SCTrvMeleeBlade
	Persuade     = entity.SCTrvPersuade
	Recon        = entity.SCTrvRecon
	RemoteOps    = entity.SCTrvElectronicsRemoteops
	Stealth      = entity.SCTrvStealth
	Streetwise   = entity.SCTrvStreetwise
)

func ListCareers() []string {
	return []string{
		CareerB0AgentLawenforcement,
		CareerB0Agentintelligence,
		CareerB0AgentCorporate,
	}
}

func SearchCareers(input string) []string {
	valid := []string{}
	for _, val := range ListCareers() {
		if strings.Contains(val, input) {
			valid = append(valid, val)
		}
	}
	if len(valid) == 0 {
		return ListCareers()
	}
	return valid
}

var SaT map[string][]string

func init() {
	SaT = make(map[string][]string)
	SaT["B0=1pd"] = []string{GunCombat, DEX, END, Melee, INT, Athletics}
	SaT["B0=1ss"] = []string{Streetwise, Drive, Investigate, Computers, Recon, GunCombat}
	SaT["B0=1ae"] = []string{Advocate, Drive, Computers, Medic, Stealth, RemoteOps}
	SaT["B0=1=1"] = []string{Investigate, Recon, Streetwise, Stealth, Melee, Advocate}
	SaT["B0=1=2"] = []string{Investigate, Recon, Comms, Stealth, Persuade, Deception}
	SaT["B0=1=3"] = []string{Investigate, Computers, Stealth, GunCombat, Deception, Streetwise}
}

func (trv *traveller) getCareerCodes() {
	var sat []string
	data := strings.Split(trv.career, "=")
	book, career, spec := data[0], data[1], data[2]
	ae := ""
	edu, err := trv.characteristics.GetValue(EDU)
	if err != nil {
		fmt.Println(err)
	}
	if edu >= 8 {
		ae = "ae"
	}
	sat = append(sat, SaT[book+"="+career+"pd"]...)
	sat = append(sat, SaT[book+"="+career+"ss"]...)
	sat = append(sat, SaT[book+"="+career+ae]...)
	sat = append(sat, SaT[book+"="+career+"="+spec]...)
	//for i := range sat {
	//fmt.Println(sat[i])
	//trv.Learn(sat[i])
	//}
	utils.RandomSeed()
	for i := 0; i < dice.Roll("3d6").Sum(); i++ {
		trv.Learn(utils.RandomFromList(sat))
	}
}

func (trv *traveller) Learn(code string) {
	fmt.Println("Learn:", code)
	for _, val := range entity.SkillCodesList() {
		if strings.Contains(val, code) {
			trv.skills.Train(code)
			return
		}
	}
	for _, val := range entity.CharacteristicsCodesList() {
		if strings.Contains(val, code) {
			trv.characteristics.Train(code)
			return
		}
	}
}
