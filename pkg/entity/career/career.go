package career

import "github.com/Galdoba/TR_Dynasty/pkg/dice"

const (
	AGENT       = "Agent"
	ARMY        = "Army"
	CITIZEN     = "Citizen"
	DRIFTER     = "Drifter"
	ENTERTAINER = "Entertainer"
	MARINE      = "Marine"
	MERCHANT    = "Merchant"
	NAVY        = "Navy"
	NOBLE       = "Noble"
	ROUGE       = "Rogue"
	SCHOLAR     = "Scholar"
	SCOUT       = "Scout"
	PRISONER    = "Prisoner"
	PSION       = "Psion"

	Imperial_Agent_LawEnforcement = 10101
)

/*
Imperial Agent Intelligence = 10102
Imperial Army Cavalery = 10203



карьера - это слайс термов.
терм - объект описывающий события и результаты бросков.
терм состоит из ряда бросков, наград и решений.
Бросок {
	описание теста = (INT 7+) проверка 2d6 + модификатор интелекта.
	эффект = результат броска минус сложность.
}

Итог Терма:
смена карьеры:
-1 запрещена
 0 разрешена
+1 вынуждена

Общая последовательность терма:
-Бросок квалификации
-Бросок выживания
-Событие(+) или Событие(-)
-Бросок продвижения

=========================================================
2.
LifeDecidionPoint - момент между Термами, в который оценивается доступность карьер

карьера - последовательность термов в одной линии.

-начало карьеры - проверка доступности карьер для персонажа.
-цикл термов
-конец карьеры


*/

type Career struct {
	Name                    string              //ID карьеры
	Assignment              []string            //доступные пути развития
	Rank                    map[string]int      //map[Assignment]rank - текущее состояние рангов в карьере
	QualificationConditions string              //Условие квалификации
	SurvivalConditions      map[string]string   //Условие выживания
	AdvancementConditions   map[string]string   //Условие продвежения
	CommisionConditions     map[string]string   //Условие комисии
	TermsSpent              int                 //общее колличество циклов проведенное в карьере
	EarnedBenefits          int                 //общее колличество бросков наград доступное тревеллеру
	QualificationPassed     bool                //Пройдена ли квалификация к данной карьере?
	CommisionPassed         bool                //Пройдена ли комиссия
	Ended                   bool                //Завершена ли данная карьера?
	TrainingTables          map[string][]string //доступные таблицы тренировки
	CashTables              map[int][]int       //таблицы денежных наград
	BenefitsTables          map[int][]string    //таблицы обычных наград
	EventMap                map[int]string      //map[int]Event - перечень возможных событий, которые могут произойти с тревеллером
	MishapMap               map[int]string      //map[int]Mishap - перечень возможных неудач, которые могут произойти с тревеллером
}

func NewCareer(name string) Career {
	cr := Career{}
	assignments := Assignment(name)
	for _, val := range assignments {
		cr.Assignment = append(cr.Assignment, name+" ("+val+")")
	}
	cr.QualificationConditions = QualificationCondition(name)
	cr.SurvivalConditions = make(map[string]string)

	return cr
}

func Assignment(career string) []string {
	switch career {
	default:
		return []string{"NO ASSIGNMENTS DEFINED"}
	case AGENT:
		return []string{
			"Law Enforcement",
			"Intelligence",
			"Corporate",
		}
	case ARMY:
		return []string{
			"Support",
			"Infantry",
			"Cavalry",
		}
	case CITIZEN:
		return []string{
			"Corporate",
			"Worker",
			"Colonist",
		}
	case DRIFTER:
		return []string{
			"Barbarian",
			"Wanderer",
			"Scavenger",
		}
	case ENTERTAINER:
		return []string{
			"Artist",
			"Journalist",
			"Performer",
		}
	case MARINE:
		return []string{
			"Support",
			"Star Marine",
			"Ground Assault",
		}
	case MERCHANT:
		return []string{
			"Merchant Marine",
			"Free Trader",
			"Broker",
		}
	case NAVY:
		return []string{
			"Line/Crew",
			"Engineer/Gunner",
			"Flight",
		}
	case NOBLE:
		return []string{
			"Administrator",
			"Diplomat",
			"Dilettante",
		}
	case ROUGE:
		return []string{
			"Thief",
			"Enforcer",
			"Pirate",
		}
	case SCHOLAR:
		return []string{
			"Field Researcher",
			"Scientist",
			"Physician",
		}
	case SCOUT:
		return []string{
			"Courier",
			"Surveyor",
			"Explorer",
		}
	case PRISONER:
		return []string{
			"Inmate",
			"Thug",
			"Fixer",
		}
	case PSION:
		return []string{
			"Wild Talent",
			"Adept",
			"Psi-Warrior",
		}
	}
}

func QualificationCondition(career string) string {
	switch career {
	default:
		return "NOTIMPLEMENTED"
	case AGENT:
		return "INT 6+"
	case ARMY:
		return "END 5+"
	case CITIZEN:
		return "EDU 5+"
	case DRIFTER:
		return "AUTOMATIC"
	case ENTERTAINER:
		return "DEX 5+ or INT 5+" //исключение с которым надо разобраться - две проверки? проверка на выбор?
	case MARINE:
		return "END 6+"
	case MERCHANT:
		return "INT 4+"
	case NAVY:
		return "INT 6+"
	case NOBLE:
		return "SOC 10+"
	case ROUGE:
		return "DEX 6+"
	case SCHOLAR:
		return "INT 6+"
	case SCOUT:
		return "INT 5+"
	case PRISONER:
		return "AUTOMATIC"
	case PSION:
		return "PSI 6+"
	}
}

///////////////////////

func RandomCareer() string {
	dp := dice.New()
	careers := []string{
		AGENT,
		ARMY,
		CITIZEN,
		DRIFTER,
		ENTERTAINER,
		MARINE,
		MERCHANT,
		NAVY,
		NOBLE,
		ROUGE,
		SCHOLAR,
		SCOUT,
		PRISONER,
		PSION,
	}
	return dp.RollFromList(careers)
}
