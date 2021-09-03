package traveller

const (
	EInst_ED5 = iota
	EInst_TradeSchool
	EInst_College
	EInst_University
	EInst_Academy
	EInst_MastersProgram
	EInst_ProfessorsProgram
	EInst_LawSchool
	EInst_MedicalSchool
	EInst_OTC
	EInst_NOTC
	Graduation_BA
	Graduation_BA_Honors
	Graduation_MA
)

type educationManager interface {
	EducationInstitutions(*TravellerT5) []int
}

type education struct {
	majorSpec int
	minorSpec int
}
