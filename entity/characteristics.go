package entity

const (
	CharCodeEntity            = 0
	CharCodeType              = 1
	CharCodePosition          = 2
	CharCodeShortName         = 3
	CharCodeFullName          = 4
	CharCodeTrvC1STRENGTH     = "1=1=C1=STR=Strength"
	CharCodeTrvC2AGILITY      = "1=1=C2=AGI=Agility"
	CharCodeTrvC2DEXTERITY    = "1=1=C2=DEX=Dexterity"
	CharCodeTrvC2Grace        = "1=1=C2=GRA=Grace"
	CharCodeTrvC3ENDURANCE    = "1=1=C3=END=Endurance"
	CharCodeTrvC3STAMINA      = "1=1=C3=STA=Stamina"
	CharCodeTrvC3VIGOR        = "1=1=C3=STA=Vigor"
	CharCodeTrvC4INTELLIGENCE = "1=1=C4=INT=Intelligence"
	CharCodeTrvC5EDUCATION    = "1=1=C5=EDU=Education"
	CharCodeTrvC5TRAINING     = "1=1=C5=TRA=Training"
	CharCodeTrvC5INSTINCT     = "1=1=C5=INS=Instinct"
	CharCodeTrvC6SOCIAL       = "1=1=C6=SOC=Social"
	CharCodeTrvC6CHARISMA     = "1=1=C6=CHA=Charisma"
	CharCodeTrvC6CASTE        = "1=1=C6=CAS=Caste"
	CharCodeTrvC6NOBILITY     = "1=1=C6=NOB=Nobility"
	CharCodeTrvC6TERRITORY    = "1=1=C6=TER=Territory"
	CharCodeTrvCSSANITY       = "1=1=CS=SAN=Sanity"
)

//skill -
type characteristic struct {
	entity      string //Ненужно?
	group       string //Ненужно?
	speciality  string //Ненужно?
	description string //Ненужно?
	value       int
}

// type Characteristic interface {
// 	Set()
// }
