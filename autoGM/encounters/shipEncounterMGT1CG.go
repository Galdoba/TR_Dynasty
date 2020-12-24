package encounters

import "fmt"

const (
	ShipTypeS               = "Type S Scout"
	ShipCorsair             = "Corsair"
	ShipYacht               = "Yacht"
	ShipGazelleCloseEscort  = "Gazelle Close Escort"
	ShipXboat               = "X-Boat"
	ShipAssaultGunship      = "Assault Gunship (TG 94)"
	ShipBombardmentShip     = "Bombardment Ship (TG 96)"
	ShipSerpentPoliceCutter = "Serpent Police Cutter (CRB 131)"
	ShipFreeTrader          = "Free Trader"
	ShipHeavyFreighter      = "Heavy Freighter (CRB 125)"
	ShipLargeFreighter      = "Large Freighter (TG 44)"
)

func (e *encounterEvent) RollShipEncounterMGT1CG() {
	e.eventClass = "Ship Encounter"
	dRoll := e.dicepool.RollNext("2d6").ResultString()
	table := ""
	fmt.Println(dRoll)
	dRoll = "21"
	switch dRoll {
	case "11":
		table = "Abandoned ship Any "
		e.AbandonedShip()
	case "12":
		table = "Alien Scouts 1–3: Type S Scot (CR 114) 4–6: Corsair (CR 129) "
		e.AlienScouts()
	case "13":
		table = "Ambassador 1–3: Yacht (CR 126) + Gazelle Close Escort (CR 123) 4–6: Xboat (TG 40) "
		e.Ambassador()
	case "14":
		table = "Aslan battleship 1–2: Corsair (CR 129) 3–4: Assault Gunship (TG 94) 5: Bombardment Ship (TG 96) 6: Serpent Police Cutter (CR 131) "
		e.AslanBattleship()
	case "15":
		table = "Bounty Hunters 1–3: Free Trader (CR 117) 4–6: Corsair (CR 129) "
		e.BountyHunters()
	case "16":
		table = "Colonists 1–4: Heavy Freighter (CR 125) 5–6: Large Freighter (TG 44) "
		e.Colonists()
	case "21":
		table = "Cultists 1–3: Free Trader (CR 117) 4–6: Yacht (CR 126) "
		e.Cultists()
	case "22":
		table = "Debris See Page 49 "
	case "23":
		table = "Droyne Explorers Type S Scot (CR 114) "
	case "24":
		table = "Exiles Any "
	case "25":
		table = "Experimental Androids Any "
	case "26":
		table = "Fugitive(s) 1–3: Free Trader (CR 117) 4: Antique In-System Hauler (TG 37) 5: Xboat (TG 40) 6: Fast Smuggler (TG 58) "
	case "31":
		table = "Mail freight 1–3: Free Trader (CR 117) 4–5: Heavy Freighter (CR 125) 6: Xboat (TG 40) "
	case "32":
		table = "Hiver Degenerates 1–2: Yacht (CR 126) 3–4: Corsair (CR 129) 5–6: Assault Gunship (TG 94) "
	case "33":
		table = "Imperial Navy 1–2: Corsair (CR 129) 3–4: Assault Gunship (TG 94) 5: Bombardment Ship (TG 96) 6: Mercenary Cruiser (CR 127) "
	case "34":
		table = "Imperial Scouts Type S Scout (CR 114) "
	case "35":
		table = "K’kree Deserters/Escapees 1–3: Mercenary Cruiser (CR 127) 4–6: Assault Gunship (TG 94) "
	case "36":
		table = "Merchants 1–2: Free Trader (CR 117) 3–4: Fat Trader (CR 119) 5–6: Heavy Freighter (CR 125) "
	case "41":
		table = "Mining ship Type S Scot (CR 114) "
	case "42":
		table = "Passenger, luxury Yacht (CR 126) "
	case "43":
		table = "Passenger, standard 1–3: Free Trader (CR 117) 4–6: Subsidised Liner (TG 80) "
	case "44":
		table = "Pirates 1: Corsair (CR 129) 2: Fast Smuggler (TG 58) 3: Free Trader (CR 117) 4: Junk Fighter (TG 28) 5: Mercenary Cruiser (CR 127) 6: Pirate Raider (TG 55) "
	case "45":
		table = "Primitives Junk Fighter (TG 28) "
	case "46":
		table = "Prison Transport Sanatorium Hospice Boat (TG 52) "
	case "51":
		table = "Robot Rebels Any"
	case "52":
		table = "Rock Stars Yacht (CR 126) "
	case "53":
		table = "Scavengers 1–3: Free Trader (CR 117) 4–6: Fast Smuggler (TG 58) "
	case "54":
		table = "Scientists 1–3: Type S Scot (CR 114) 4–6: Laboratory Ship (CR 121) "
	case "55":
		table = "Self-Aware Any "
	case "56":
		table = "Space battle Any "
	case "61":
		table = "Mercenaries Mercenary Cruiser (CR 127) "
	case "62":
		table = "Vargr Raiding party 1–2: Free Trader (CR 117) 3–4: Mercenary Cruiser (CR 127) 5–6: Fast Smuggler (TG 58) "
	case "63":
		table = "Xenologists Same as Merchants "
	case "64":
		table = "Zhodani Thought Police 1–2: Corsair (CR 129) 3–4: Xboat (TG 40) 5: Assault Gunship (TG 94) 6: Mercenary Cruiser (CR 127) "
	case "65":
		table = "Zombie Armada Any "
	case "66":
		table = "Zombies Any "
	}
	fmt.Println(table)
}

/////////////////////////
//SHIP ENCOUNTER EVENTS//
/////////////////////////

//AbandonedShip -
func (e *encounterEvent) AbandonedShip() {
	e.name = "Abandoned ship"
	e.descr = "Disease, technical malfunction or some other misfortune have killed or driven away anyone on board and the ship is ripe for the taking... or is it?\n"
	e.location = e.anyShip()
	reason := "Unknown"
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1:
		reason = "This is a trap created by pirates to capture Travellers. As soon as the Player Characters board the abandoned ship few Pirate Ships emerge (TODO: RollPirateEncounter())."
	case 2:
		reason = "The ship’s crew was en route to commit the perfect heist when an argument over sharing the loot led to a shootout that left no survivors. The plans for the heist, as well as a list of contact people and passwords are still onboard."
	case 3:
		reason = "This is a smuggler ship. The smugglers got more than they bargained for when the ‘fossils’ they dug up on a restricted planet turned out to be living and deadly creatures, now scouring the ship for food."
	case 4:
		reason = "The ship is self-aware (CG p.57) but not particularly intelligent. After years of wandering in space it ran out of fuel and shut down. As soon as the Player Characters power it, it ‘awakens’ and begins to act according to its personality (page 43)."
	case 5:
		reason = "The ship is fully manned by a crew sleeping in a cryobitrh, but their vital signs do not register on the bioscanner. If the Player Characters board their ship, the crew will think they are being invaded and fight back. Finding a peaceful solution to this misunderstanding is possible."
	case 6:
		reason = "All crewmen on the ship died from an alien virus contracted on a newly discovered planet. Any Player Character who boards the ship contracts the disease as well (page 108)."
	}
	e.descr += reason
}

//AlienScouts -
func (e *encounterEvent) AlienScouts() {
	e.name = "Alien Scouts"
	e.descr = "These are scouts from a faraway alien minor race.\n"
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2:
		e.descr += "Scouts question Player Characters about their world and customs in a respectful manner."
	case 3, 4:
		e.descr += "Scouts fascinated with the Player Characters and will not leave them alone."
	case 5, 6:
		e.descr += "Scouts treat the Player Characters as test subjects to be studied, dissected and discarded."
	}
	e.descr += "\nScout ships rarely have any cargo save a small amount of basic consumables. However, they often carry exotic gifts for newly-met alien rulers, which can be anything from colourful beads to ancient technology (TODO: Exotic Trade Good)."
	e.descr += "\nAlien scout ships are of special interest to governments in whose space they operate as the way a first encounter plays out can determine whether the visited race will gain a new trade partner or be drawn into a bloody galactic war. For this reason, this encounter can have far-reaching consequences, especially if the Player Characters rout the scouts. Even if the scouts attack first, the local government may wish to sacrifice the Player Characters to appease the aliens."
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2, 3:
		e.location = ShipTypeS
	case 4, 5, 6:
		e.location = ShipCorsair
	}
}

//Ambassador -
func (e *encounterEvent) Ambassador() {
	e.name = "Ambassador"
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2, 3:
		e.location = ShipYacht + " + " + ShipGazelleCloseEscort
	case 4, 5, 6:
		e.location = ShipXboat
	}
	e.descr = "The ship contains an ambassador or a special envoy from a random world. It is heavily guarded and the defenders are jumpy, making violent confrontation due to misunderstanding very likely."
	e.descr += "\nThis encounter has the chance to alter the entire campaign. The murder of an ambassador can lead to war, even if this was nothing but a tragic misunderstanding. On the other hand, players befriending an ambassador can be sucked into a world of intrigue and treachery as their government tries to use them to further national agendas or to get rid of them to ensure they do not cloud the ambassador’s judgement. Ambassadors often bear expensive gifts to the governments they are visiting."
}

//AslanBattleship -
func (e *encounterEvent) AslanBattleship() {
	e.name = "Aslan Battleship"
	e.descr = "\nAslans are an honourable race and will not attack anyone without a good cause. On the other hand, they are also hotheaded and often overact to the smallest perceived offences."
	e.descr += "\nUnless patrolling the borders of the Aslan Hierate, the Aslan battleship is on a mission."
	e.descr += "\nAslans are obsessed with owning as much wealth as possible, as well as defending it. If the ship is on a mission, roll twice on the cargo table (Traveller Core Rulebook page 165). The Aslans will never surrender this cargo and fight to the death to protect it."
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2:
		e.location = ShipCorsair
	case 3, 4:
		e.location = ShipAssaultGunship
	case 5:
		e.location = ShipBombardmentShip
	case 6:
		e.location = ShipSerpentPoliceCutter
	}
}

//BountyHunters -
func (e *encounterEvent) BountyHunters() {
	e.name = "Bounty Hunters"
	e.descr = "Bounty hunters differ from pirates and criminals by current occupation, not personality. A small minority are ex-lawmen or noble avengers hunting criminals across the stars but the majority are immoral brutes looking for easy cash. "
	e.descr += "\nWhile bounty hunters love to ask for other travellers’ assistance, they are loath to share their bounty and would go as far as murdering their former allies to avoid sharing the Credits."
	e.descr += "\nAn insane space chase that involves fighting dangerous criminals and working with the same sort of people can be quite entertaining."
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2, 3:
		e.location = ShipFreeTrader
	case 4, 5, 6:
		e.location = ShipCorsair
	}
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1:
		e.hook = "Bounty Hunter is a genocide survivor is looking for a war criminal posing as a wealthy philanthropist."
	case 2:
		e.hook = "Bounty Hunter is a police investigator has quit the force to catch a serial killer with powerful connections who has eluded him for too long."
	case 3:
		e.hook = "Bounty Hunter is a farmer whose wife was dishonoured by an aristocrat while the farmer was away fighting on alien worlds. Equipped with a shabby ship and a battered laser rifle, he is going to take on the most powerful man in the sector."
	case 4:
		e.hook = "Bounty Hunter is a deeply religious man character only hunts ‘sinners’ is pursuing a brilliant thief who has stolen his church’s most cherished relic... but for what end?"
	case 5:
		e.hook = "Bounty Hunter is a dangerous criminal has enlisted and is presently serving as a commissioned officer in a deadly war zone. On the eve of an interstellar war, he is going to play a deadly game with the bounty hunter who came to take him back to face the music."
	case 6:
		e.hook = "Bounty Hunter is a decorated war hero became a bounty hunter to finance an expensive procedure for his dying wife. His target is a sadistic psychopath whose quest is to hurt one female of every species."
	}
}

//Colonists -
func (e *encounterEvent) Colonists() {
	e.name = "Colonists"
	e.descr = "Colonists pose a moral challenge more often than they pose a combat challenge.\nAlthough the peaceful colonist ship packed with huddled women and children attacked by pirates or savages is a common sci-fi conceit, this is not always the case. Colonists, unlike refugees, are rarely simple people looking for a new start. Instead they are...\n" + e.colonists()
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2, 3, 4:
		e.location = ShipHeavyFreighter
	case 5, 6:
		e.location = ShipLargeFreighter
	}
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1:
		e.hook = "The colonists send a distress call asking for technical and medical assistance after running into some dangerous space event (roll on Space Events table, page 32)."
	case 2:
		e.hook = "The colonists are attacked by a superior enemy force that will destroy them without the Player Characters’ assistance. The attackers are likely to have a good reason for that, such as defending their world, trying to destroy an infected ship or opposing the colonists’ twisted religion."
	case 3:
		e.hook = "The Player Characters’ vessel suffers a technical malfunction which they cannot fix without the assistance of the colonists, who invite them to stay on-board during the repairs."
	case 4:
		e.hook = "The Player Characters were hired under a false pretence to stop the colonists from reaching their destination."
	case 5:
		e.hook = "The Player Characters realise that the colonists are heading towards their world. This hook only works if the Player Characters hail from a primitive independent world."
	case 6:
		e.hook = "The captain of one of the colonists’ ships is a friend or former commander of one or more of the Player Characters and invites them for dinner onboard his ship."
	}
}

//Cultists -
func (e *encounterEvent) Cultists() {
	e.name = "Cultists"
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2, 3:
		e.location = ShipFreeTrader
	case 4, 5, 6:
		e.location = ShipYacht
	}
	e.descr = "The cultists are not simply extreme religious people looking for a place where they could practice their faith in peace. They are brainwashed, violent and led by a charismatic but positively insane ‘prophet’. Sadly, by the time the Player Characters realise this subtle difference, it might already be too late. In this case:\n" + e.cultists()
}

//////////////////////////
//SHIP ENCOUNTER HELPERS//
//////////////////////////

func (e *encounterEvent) anyShip() string {
	ships := []string{
		ShipTypeS,
		ShipCorsair,
		ShipYacht,
		ShipGazelleCloseEscort,
	}
	return e.dicepool.RollFromList(ships)
}

func (e *encounterEvent) colonists() string {
	switch e.dicepool.RollNext("1d6").Sum() {
	default:
		return "Unknown"
	case 1:
		return "Well-armed zealots who believe it is their religious duty to reclaim a lost world meanwhile colonised by a different race."
	case 2:
		return "Soldiers and their families dispatched by the Imperium to settle a contested world."
	case 3:
		return "A bizarre religious community looking for an isolated planet to practice their unspeakable rites (see ‘Cultists’ on page 49)."
	case 4:
		return "Criminals given a last chance at rehabilitation by being exiled to a newly discovered mineral-rich world. The criminals are eager to escape via any means available."
	case 5:
		return "Poor people fooled into thinking they are starting a new life. In truth they are taken to Glorious Empire's working camp... and not as workers."
	case 6:
		return "The colonists have secretly been given a deadly disease to infect the troublesome natives of the world they were sent to settle. The true colonists will arrive in a few years to find a depopulated world ready for a fresh start."
	}
}

func (e *encounterEvent) cultists() string {
	switch e.dicepool.RollNext("1d6").Sum() {
	default:
		return "Unknown"
	case 1:
		return "The head of the cult is a bully and a pervert who is fleeing the law after a journalist has exposed the true nature of his cult. His followers are completely brainwashed and will ignore all but the most damning evidence against their guru. A few people want to escape the ship but will find it difficult to convey their intention to the Player Characters due to fear of their fanatic brethren."
	case 2:
		return "The cultists are on their way to ‘awaken their sleeping God’ and would love to have the Player Characters onboard for this great mission. The God is " + e.god() + "."
	case 3:
		return "Despite much negative press and constant police harassment, the cultists are actually a decent and honest people, even if somewhat eccentric. They have decided to leave civilisation and start a colony somewhere faraway after the persecution became unbearable. Because they fully put their faith in God, they are presently lost, penniless and framed for crimes actually committed by the corrupt policemen investigating the cult."
	case 4:
		return "The cultists, along with their families, are on a holy war against one of the major races (see table on page 45). Poorly armed and untrained, they are not likely to survive an encounter even with the weakest of military vessels."
	case 5:
		return "The cultists are normal people with control chips installed in their brains by the guru – an android fugitive (page 50). The android believes in the ‘God in the Machine’, a colossal AI (page 43) a few parsecs away and intends to sacrifice the humans to his God to gain acceptance to the electronic heaven."
	case 6:
		return "Everyone on the ship is dead following a mass suicide. Treat this encounter as ‘Abandoned ship’."
	}
}

func (e *encounterEvent) god() string {
	switch e.dicepool.RollNext("1d6").Sum() {
	default:
		return "Unknown"
	case 1, 2:
		return "colossal alien beast"
	case 3, 4:
		return "super AI"
	case 5, 6:
		return "a charlatan"
	}
}
