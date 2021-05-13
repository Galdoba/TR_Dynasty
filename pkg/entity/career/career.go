package career

const (
	//Imperial_Agent                = 10100
	Imperial_Agent_LawEnforcement = 10101
	Imperial_Agent_Intelligence   = 10102
	Imperial_Agent_Corporate      = 10103
	Imperial_Army_Support         = 10201
	Imperial_Army_Infantry        = 10202
	Imperial_Army_Cavalery        = 10203
)

func Allterms() []int {
	return []int{
		Imperial_Agent_LawEnforcement,
		Imperial_Agent_Intelligence,
		Imperial_Agent_Corporate,
		Imperial_Army_Support,
		Imperial_Army_Infantry,
		Imperial_Army_Cavalery,
	}
}

func Data(code int) Career {
	cr := Career{}
	switch code {
	case Imperial_Agent_LawEnforcement:
		cr = Career{
			Name:                    "Agent",
			Assignment:              []string{"Law Enforcement", "Intelligence", "Corporate"},
			RankBonuses:             map[string]string{},
			QualificationConditions: "",
			SurvivalConditions:      map[string]string{},
			AdvancementConditions:   map[string]string{},
			CommisionConditions:     map[string]string{},
			OpenedOnStart:           false,
			TrainingTables:          map[string][]string{},
			CashTables:              map[int][]int{},
			BenefitsTables:          map[int][]string{},
			EventMap:                map[int]string{},
			MishapMap:               map[int]string{},
		}

	}
	return cr
}

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
	RankBonuses             map[string]string   //map[Name+Assignment+i]skill+i||atr+i||atr - бонус который должен получить персонаж
	QualificationConditions string              //Условие квалификации
	SurvivalConditions      map[string]string   //Условие выживания
	AdvancementConditions   map[string]string   //Условие продвежения
	CommisionConditions     map[string]string   //Условие комисии
	OpenedOnStart           bool                //Открыта ли данная карьера?
	TrainingTables          map[string][]string //доступные таблицы тренировки
	CashTables              map[int][]int       //таблицы денежных наград
	BenefitsTables          map[int][]string    //таблицы обычных наград
	EventMap                map[int]string      //map[int]Event - перечень возможных событий, которые могут произойти с тревеллером
	MishapMap               map[int]string      //map[int]Mishap - перечень возможных неудач, которые могут произойти с тревеллером
}
