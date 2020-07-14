package universe

/*
	Sector										Astergoth
		-system									Astergoth B0323
			--name
			--position
			--stars								Astergoth B0323 Alpha
				---name
				---Classification
				---position
				---hab zone						Astergoth B0323 Alpha 2
					----name
					----planet					Astergoth B0323 Alpha 2
						-----worldData
						-----name
						-----position
						-----satellite			Astergoth B0323 Alpha 2 Ay
							------worldData
							------name
							------position

Astergoth B0323 Alpha 2 Ay
*/

/*
Hex(coords)		Hex
Name			MnWorld
UWP				MnWorld
Trade Codes		EvrWorld
{Ix}			System
(Ex)			System
[Cx]			System
Nobility		System
Bases			System
Zone			System
PBG				System
Worlds			System
Stellar			Hex

Sector										Astergoth
	-system									Astergoth B0323
		--name
		--position
		--stars								Astergoth B0323 Alpha
			---name
			---Classification
			---position
			---hab zone						Astergoth B0323 Alpha 2
				----name
				----planet					Astergoth B0323 Alpha 2
					-----worldData
					-----name
					-----position
					-----satellite			Astergoth B0323 Alpha 2 Ay
						------worldData
						------name
						------position

Astergoth B0323 Alpha 2 Ay



все звёзды - знаем звезды и кол-во планет (всего или для каждой?)
размещаем ГГ
размещаем Астеройды
размещаем Планеты
для каждой кланеты ПЕРВЫЙ ОБЗОР
выбираем ГЛАВНЫЙ МИР
Населяем главный Мир
для главного мира ВТОРОЙ ОБЗОР(или системы в Общем?)
Населяем каждую планету

каждая звезда имеет ограниченное кол-во доступных планет

*/

type GasGigant struct {
	name     string
	alias    string
	size     string
	diameter int
	gee      float64
	//satelite map[systemName]*world
}
