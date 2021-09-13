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
	//Basic
	edProgs := pickValidEducation(trv)
	applyTo := dice.New().RollFromListInt(edProgs)
	if dice.Roll1D() > 4 {
		fmt.Println("Education skipped...")
		return
	}
	switch applyTo {
	case EInst_ED5:
		trv.err = trv.ED5Program()
	case EInst_TradeSchool:
		trv.err = trv.TradeSchool()
	case EInst_College:
		trv.err = trv.College()
	}

}

func pickValidEducation(trv *TravellerT5) []int {
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
			if edu.Value() >= 5 {
				validPrograms = append(validPrograms, EInst_College)
			}
			if edu.Value() >= 6 {
				validPrograms = append(validPrograms, EInst_Academy)
			}
			if edu.Value() >= 7 {
				validPrograms = append(validPrograms, EInst_University)
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

	return validPrograms
}

func (trv *TravellerT5) ED5Program() error {
	trv.AddEvent("Apply for ED5 Education Program...")
	trv.AddEvent("Application: Automatic")
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

func (trv *TravellerT5) TradeSchool() error {
	trv.AddEvent("Apply for Trade School Education Program...")
	intel := trv.characteristic[assets.Intelligence]
	edu := trv.characteristic[assets.Education]
	trv.education.majorSpec = dice.New().RollFromListInt(listTradeSchoolMajors())
	applyTask := tasks.Create()
	applyTask.SetupEnviroment("apply for Trade School Studying", 2, 0)
	applyTask.SetupAssets(intel)
	applyTask.Resolve()
	trv.AddTaskEvent(applyTask)
	trv.education.educationAtempts[EInst_TradeSchool]++
	if !applyTask.Completed() {
		if !trv.UseWaifer("re apply for Trade School") {
			return nil
		}
	}
	sturdyChar := edu
	if intel.Value() > edu.Value() {
		sturdyChar = intel
	}
	studyTask := tasks.Create()
	studyTask.SetupEnviroment("graduate from Trade School", 2, 0)
	studyTask.SetupAssets(sturdyChar)
	studyTask.Resolve()
	trv.AddTaskEvent(studyTask)
	if !studyTask.Completed() {
		if !trv.UseWaifer("use waifer to graduate from Trade School") {
			return nil
		}
	}
	trv.AddEvent("graduation from Trade School Completed...")
	major := trv.education.majorSpec
	trv.knowledges[major] = assets.NewKnowledge(major)
	trv.knowledges[major].Train()
	trv.knowledges[major].Train()
	fmt.Println("major is", trv.knowledges[major].Name())
	trv.AddAge(1)
	return nil
}

func (trv *TravellerT5) College() error {
	studyChar := trv.characteristic[assets.Education]
	if trv.characteristic[assets.Intelligence].Value() > studyChar.Value() {
		studyChar = trv.characteristic[assets.Intelligence]
	}
	major, minor := 0, 0
	for major == minor {
		major = dice.New().RollFromListInt(listColledgeMajors())
		minor = dice.New().RollFromListInt(listColledgeMajors())
	}
	trv.AddEvent("Apply for College Education Program...")
	collegeTask := tasks.Create()
	collegeTask.SetupAssets(studyChar)
	collegeTask.SetupEnviroment("pass the entering exams", 2, 0)
	collegeTask.Resolve()
	trv.AddTaskEvent(collegeTask)
	if !collegeTask.Completed() {
		if !trv.UseWaifer("Re-apply enterings exams using wafer") {
			return nil
		}
	}
	for year := 1; year < 5; year++ {
		trv.AddEvent(fmt.Sprintf("Attending College year %v...\n", year))
		collegeTask.Resolve()
		trv.AddTaskEvent(collegeTask)
		if !collegeTask.Completed() {
			if !trv.UseWaifer("Re-apply exams using wafer") {
				trv.AddEvent("busted from College")
				return nil
			}
		}

		switch year {
		case 2, 4:
		}
	}
	return nil
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

func listColledgeMajors() []int {
	return []int{
		assets.SKILL_Athlete,
		assets.SKILL_Broker,
		assets.SKILL_Bureaucrat,
		assets.SKILL_Counsellor,
		assets.SKILL_Designer,
		assets.SKILL_Language,
		assets.SKILL_Teacher,
		assets.SKILL_Astrogator,
		assets.ART_Actor,
		assets.ART_Artist,
		assets.ART_Author,
		assets.ART_Chef,
		assets.ART_Dancer,
		assets.ART_Musician,
		assets.KNOWLEDGE_Biology,
		assets.KNOWLEDGE_SmallCraft,
		assets.TRADE_Electronics,
		assets.TRADE_Craftsman,
		assets.TRADE_Fluidics,
		assets.TRADE_Gravitics,
		assets.TRADE_Magnetics,
		assets.TRADE_Mechanic,
		assets.TRADE_Photonics,
		assets.TRADE_Polymers,
		assets.TRADE_Programmer,
		assets.KNOWLEDGE_Automotive,
		assets.KNOWLEDGE_Archeology,
		assets.KNOWLEDGE_Biology,
		assets.KNOWLEDGE_Chemistry,
		assets.KNOWLEDGE_History,
		assets.KNOWLEDGE_Linguistics,
		assets.KNOWLEDGE_Philosophy,
		assets.KNOWLEDGE_Physics,
		assets.KNOWLEDGE_Planetology,
		assets.KNOWLEDGE_Psionicology,
		assets.KNOWLEDGE_Psychohistory,
		assets.KNOWLEDGE_Psychology,
		assets.KNOWLEDGE_Robotics,
		assets.KNOWLEDGE_Sophontology,
		assets.KNOWLEDGE_Aeronautics,
		assets.KNOWLEDGE_Aquanautics,
	}
}
