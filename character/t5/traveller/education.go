package traveller

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
	"github.com/Galdoba/TR_Dynasty/T5/tasks"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

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
	TInst_Apprenticeship
	TInst_Mentor
	TInst_TrainingCourse
	Graduation_BA
	Graduation_BA_Honors
	Graduation_MA
)

type educationManager interface {
	EducationInstitutions(*TravellerT5) []int
}

type education struct {
	majorSpec        int
	minorSpec        int
	graduations      []int
	wafersUsed       int
	educationAtempts map[int]int
}

func (trv *TravellerT5) Educate() {
	edData := &education{}
	edData.educationAtempts = make(map[int]int)
	trv.education = edData
	basicEducationFinished := false
	//Basic
	for !basicEducationFinished {
		edProgs := pickValidBasicEducation(trv)
		for _, edPrg := range edProgs {
			if edPrg == EInst_ED5 && trv.education.educationAtempts[EInst_ED5] < 1 {
				trv.AddEvent("Apply for ED5 Education Program...")
				trv.AddEvent("Application: Automatic")
				trv.err = trv.ApplyED5Program()
			}
			if edPrg == EInst_TradeSchool {
				if dice.Roll1D() > 3 {
					trv.AddEvent("Apply for Trade School Education Program...")
					trv.err = trv.ApplyTradeSchool()
				}
			}
		}
		if dice.Roll1D() > 3 {
			basicEducationFinished = true
		}
	}

}

func pickValidBasicEducation(trv *TravellerT5) []int {
	validPrograms := []int{}
	testCharactiristics := []int{assets.Education, assets.Training, assets.Instinct}
	for _, char := range testCharactiristics {
		if _, ok := trv.characteristic[char]; !ok {
			continue
		}
		switch char {
		case assets.Education:
			edu := trv.characteristic[char]
			if edu.Value() <= 4 && trv.education.educationAtempts[EInst_ED5] < 1 {
				validPrograms = append(validPrograms, EInst_ED5)
			}
			if edu.Value() >= 5 {
				validPrograms = append(validPrograms, EInst_TradeSchool)
			}
		case assets.Training:
			tra := trv.characteristic[char]
			validPrograms = append(validPrograms, TInst_Apprenticeship)
			validPrograms = append(validPrograms, TInst_Mentor)
			if tra.Value() > 4 {
				validPrograms = append(validPrograms, TInst_TrainingCourse)
			}
		}
	}
	fmt.Println("Valid Programs:", validPrograms)
	return validPrograms
}

func (trv *TravellerT5) ApplyED5Program() error {
	ed5Test := tasks.Create()
	intelligence := trv.characteristic[assets.Intelligence]
	ed5Test.SetupAssets(intelligence)
	ed5Test.SetupEnviroment("finish ED5 Education Program", 2, 0)
	trv.AddEvent(ed5Test.TaskPhrase())
	ed5Test.Resolve()
	trv.AddEvent(ed5Test.Outcome())
	if ed5Test.Completed() {
		intelligence.SetValue(5)
		trv.AddEvent("Education is now = 5")
	}
	trv.education.educationAtempts[EInst_ED5]++
	trv.characteristic[assets.Intelligence] = intelligence
	return nil
}

func (trv *TravellerT5) ApplyTradeSchool() error {
	intel := trv.characteristic[assets.Intelligence]
	edu := trv.characteristic[assets.Education]
	soc := trv.characteristic[assets.Social]
	trv.education.majorSpec = dice.New().RollFromListInt(listTradeSchoolMajors())
	applyTask := tasks.Create()
	applyTask.SetupEnviroment("apply for Trade School Studying", 2, 0)
	applyTask.SetupAssets(intel)
	trv.AddEvent(applyTask.TaskPhrase())
	applyTask.Resolve()
	trv.AddEvent(applyTask.Outcome())
	trv.education.educationAtempts[EInst_TradeSchool]++
	if !applyTask.Completed() {
		trv.AddEvent("Enter exams were failed...")
		if dice.Roll1D() > 3 {
			return nil
		}
		trv.AddEvent("Apealing for retry using education wafer...")
		applyTask.SetupAssets(soc)
		applyTask.SetupEnviroment("apeal the exams", 2, 0)
		trv.AddEvent(applyTask.TaskPhrase())
		applyTask.Resolve()
		trv.AddEvent(applyTask.Outcome())
		if !applyTask.Completed() {
			return nil
		}
	}
	sturdyChar := edu

	return fmt.Errorf("Trade school not implemented")
}

func listTradeSchoolMajors() []int {
	return []int{
		assets.KNOWLEDGE_ACV,
		assets.KNOWLEDGE_Automotive,
		assets.KNOWLEDGE_GravDriver,
		assets.KNOWLEDGE_Legged,
		assets.KNOWLEDGE_Tracked,
		assets.KNOWLEDGE_Wheeled,
		assets.KNOWLEDGE_Blades,
		assets.KNOWLEDGE_SlugThrowers,
		assets.KNOWLEDGE_Unarmed,
		assets.KNOWLEDGE_JumpDrives,
		assets.KNOWLEDGE_LifeSupport,
		assets.KNOWLEDGE_ManeuverDrive,
		assets.KNOWLEDGE_PowerSystems,
		assets.KNOWLEDGE_Linguistics,
		assets.KNOWLEDGE_Robotics,
		assets.KNOWLEDGE_Aeronautics,
		assets.KNOWLEDGE_Flapper,
		assets.KNOWLEDGE_GravFlyer,
		assets.KNOWLEDGE_LTA,
		assets.KNOWLEDGE_Rotor,
		assets.KNOWLEDGE_Winged,
		assets.KNOWLEDGE_SmallCraft,
		assets.KNOWLEDGE_Rider,
		assets.KNOWLEDGE_Teamster,
		assets.KNOWLEDGE_Trainer,
		assets.KNOWLEDGE_Aquanautics,
		assets.KNOWLEDGE_GravSea,
		assets.KNOWLEDGE_Boat,
		assets.KNOWLEDGE_ShipSea,
		assets.KNOWLEDGE_Sub,
	}
}
