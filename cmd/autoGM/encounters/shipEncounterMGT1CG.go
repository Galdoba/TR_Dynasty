package encounters

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/utils"
)

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
	ShipInSystemHauler      = "In System Hauler (TG 37)"
	ShipFastSmuggler        = "Fast Smuggler (TG 58)"
)

func (e *encounterEvent) RollShipEncounterMGT1CG() {
	e.eventClass = "Ship Encounter"
	dRoll := e.dicepool.RollNext("2d6").ResultString()
	table := ""
	fmt.Println(dRoll)
	EncounterDistance(dice.FluxGOOD() - 1)
	dRoll = "32"
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
		e.Debris()
	case "23":
		table = "Droyne Explorers Type S Scot (CR 114) "
		e.DroyneExplorers()
	case "24":
		table = "Exiles Any "
		e.Exiles()
	case "25":
		e.ExperimentalAndroids()
		table = "Experimental Androids Any "
	case "26":
		e.Figutive()
		table = "Fugitive(s) 1–3: Free Trader (CR 117) 4: Antique In-System Hauler (TG 37) 5: Xboat (TG 40) 6: Fast Smuggler (TG 58) "
	case "31":
		e.Mailfreight()
		table = "Mail freight 1–3: Free Trader (CR 117) 4–5: Heavy Freighter (CR 125) 6: Xboat (TG 40) "
	case "32":
		table = "Hiver Degenerates 1–2: Yacht (CR 126) 3–4: Corsair (CR 129) 5–6: Assault Gunship (TG 94) "
		e.HiverDegenerates()
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

func (e *encounterEvent) Debris() {
	e.name = "Debris"
	e.descr = "These are the remains of a ship destroyed in battle or due to accident. A thorough examination of the debris will yield 2d6% of the cargo and spare parts normally available on a spacecraft of this class. If the Referee is interested in using the encounter as a hook for an adventure, there might be a black box (page 35) or survivors in escape pods or vac suits hovering within the debris field."
	e.descr += "\nDebris often draws scavengers and other lowlifes who make a living of cannibalising damaged or destroyed ships, so the players might face some competition over the remaining cargo and spare parts (see Pirates, Scavengers and Vargr Raiding Party)."
	e.location = e.anyShip()
}

func (e *encounterEvent) DroyneExplorers() {
	e.name = "Droyne Explorers"
	e.descr = "Unless this event constitutes first contact, it probably will not lead to anything past polite greetings or a routine questionnaire if the Player Characters fly a strange vessel or behave erratically. Droyne are infamously predictable. Droyne vessels have normal cargo. While not cowardly, they would rather part with their cargo than with their lives."
	e.location = ShipTypeS
}

func (e *encounterEvent) Exiles() {
	e.name = "Exiles"
	e.descr = "This is an extremely versatile encounter as there are hundreds of reasons for a group or an individual to be forced from their homeland. Exiles differ from colonists by not specifically looking for a new world to settle and from fugitives by not being actively pursued by their enemies (at least not openly).\nThis encounter usually serves as a more colourful and exotic way to perform various utilities or meet potential patrons."
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1:
		e.hook = "The hero was exiled from his planet for a horrible crime and is tormented by guilt and shame. He may join the Player Characters and serve as a dark and haunted ally. See pages 153-163 for suitable NPCs."
	case 2:
		e.hook = "Those idealists chose exile over living under an immoral government. They might be interested in hiring the Player Characters to help them find a world where they could lead a ‘moral’ life. The dissidents can also join the crew as NPCs."
	case 3:
		e.hook = "An entire community was banished from its home and forced to take to the stars. Instead of looking for a new world they decided to make the whole universe their home. They can provide repair and maintenance service as well as buy and sell goods."
	case 4:
		e.hook = "This is an extremely wealthy person who left his sector in search of a world with lower taxes. Various governments (including the Player Characters’) will go out of their way to convince him to settle in their jurisdiction. Additionally, being a wealthy collector, the exile may be interested in buying some of the Player Characters’ more exotic loot."
	case 5:
		e.hook = "A once mighty leader was forced to escape his world with nothing but his life after a revolution or a coup brought his enemies to power. Still, he dreams of reclaiming the throne and tries to convince any tough-looking person he encounters to join his cause."
	case 6:
		e.hook = "A brilliant engineer was forced into exile after several of his inventions failed spectacularly. Bored and lonely, he offers to upgrade the Player Characters’ ship or equipment. This might take them into the next TL or make them fail or explode in the worst possible moment... the engineer’s record has been spotty at best. The engineer can also join the crew if the Player Characters need one."
	}
	e.location = e.anyShip()
}

func (e *encounterEvent) ExperimentalAndroids() {
	e.name = "Experimental Androids"
	e.descr = "The ship is run entirely by advanced robots that are practically indistinguishable from their biological counterparts. Their frames include devices that mimic life signs that fool most scanners while their physical appearance is so perfectly realistic that nothing short of a dissection will reveal their true nature."
	e.location = e.anyShip()
	e.hook = e.androids()
}

func (e *encounterEvent) Figutive() {
	e.name = "Figutive"
	e.descr = "The fugitives beg for the Player Characters’ assistance. If the Player Characters protect them, they might have to face bounty hunters, imperial battleships, primitives, prison transports or zhodani thought police agents. If the Player Characters apprehend the fugitives, they will be rewarded by the authorities but may find themselves the target of a vicious retribution by the fugitive’s associates."
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2, 3:
		e.location = ShipFreeTrader
	case 4:
		e.location = ShipInSystemHauler
	case 5:
		e.location = ShipXboat
	case 6:
		e.location = ShipFastSmuggler
	}
	e.hook = e.figutives()
}

func (e *encounterEvent) Mailfreight() {
	e.name = "Mail Freight"
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2, 3:
		e.location = ShipFreeTrader
	case 4, 5:
		e.location = ShipHeavyFreighter
	case 6:
		e.location = ShipXboat
	}
	e.descr = "Private couriers are prepared for interception attempts because of the sensitive information they carry and respond, with fight or flight, to the slightest provocations. Mail freights are more relaxed and often invite the crews of ships they encounter onboard to battle the boring monotony of their voyages."
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1:
		e.hook = "The courier carries a message to or from the Emperor. The ship is unusually well-armed and everyone on board is an elite trooper and very edgy. The message is of vital importance and can be sold for millions, although this will result in the full might of the Imperium dropping on the Player Characters’ heads and possibly a war with the Zhodani."
	case 2:
		e.hook = "The courier is returning from a mission and does not have anything of value."
	case 3:
		e.hook = "The courier is heavily damaged and the sole surviving member of the crew mumbles some nonsense about the ‘Covering of the Imperium’ before dying of radiation poisoning. The Player Characters are now in command of the vessel and a strange parcel addressed to ‘Baron Cucumber’."
	case 4:
		e.hook = "The mail freight has a letter for the Player Characters. It's a " + e.dicepool.RollFromList([]string{"bad news", "job offer from a patron"}) + "."
	case 5:
		e.hook = "The mail freight is being pursued by pirates (page 54) convinced it is a trader in disguise."
	case 6:
		e.hook = "The mail freight behaves erratically, nearly ramming the Player Characters’ spacecraft while sending meaningless messages. Examination of the craft will show that the crew opened a suspicious parcel that turned out to contain a powerful psychedelic agent now also affecting the boarding Player Characters."
	}

}

func (e *encounterEvent) HiverDegenerates() {
	e.name = "Hiver Degenerates (WARNING: MATURE CONTENT)"
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1, 2:
		e.location = ShipYacht
	case 3, 4:
		e.location = ShipCorsair
	case 5, 6:
		e.location = ShipAssaultGunship
	}
	e.descr = "Hivers are known to value personal freedom, thought and expression above all else. For some it becomes and obsession; a horrible, sick obsession.\nThe degenerates are the ultimate bohemians, travelling space in search of new experiences and new ways to practice radical self-expression. The end always justifies the means as long as the end is really ‘you’. Any attempt to limit the hiver’s selfexpression is viewed as oppression and responded to with violence. If one refuses to die for the sake of an artist truly realising his vision, then one does not deserve to live anyway...\nPainting, sculpting, holographics, singing... all of these are not nearly radical enough, poor imitation and surrender to the norms is what they are. True self-expression has got to hurt!"
	switch e.dicepool.RollNext("1d6").Sum() {
	case 1:
		e.hook = "A garden of bodies hovering in space. Hundreds will die but the result will be breathtaking and make a profound commentary on the sophont condition and the horrors of post-imperialism. Now all we need is more bodies and not just any bodies mind you... colour, shape and personality matter!"
	case 2:
		e.hook = "The small world of Reznista is known for the quality and diversity of its meat export... that is to say, for its ongoing and unpunished, statesanctioned genocide of a million different species. The only verdict is an eye for an eye. In this case, a travelling restaurant that serves Reznistian meat."
	case 3:
		e.hook = "All authority is bad, freedom to the oppressed masses! This ship travels space, doing recreational drugs, spreading the love and blowing up spacecrafts and structures it associates with authority, Hiver or otherwise."
	case 4:
		e.hook = "Life is the highest art and flesh is the most challenging canvas of them all. This ship, disguised as a private humanitarian mission, captures hapless space travellers and transforms them into things of terrible beauty. Less beauty, more terror, to be precise."
	case 5:
		e.hook = "A group of hivers have developed a drug that makes everyone see things the way they are, that is, to hallucinate uncontrollably. They are very keen on sharing their newfound clarity with the rest of the universe by spraying it into atmospheres or introducing it into spacecraft life support systems."
	case 6:
		e.hook = "The degenerates try to make a point by uplifting as many animals as possible and convincing them they are superior and should enslave all sophonts. Their king is a morbidly obese mouse named Cookie. Even the hivers themselves cannot quite put their revolutionary point into words... hopefully some talking duck will be able to."
	}

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
		ShipXboat,
		ShipAssaultGunship,
		ShipBombardmentShip,
		ShipSerpentPoliceCutter,
		ShipFreeTrader,
		ShipHeavyFreighter,
		ShipLargeFreighter,
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

func (e *encounterEvent) androids() string {
	switch e.dicepool.RollNext("1d6").Sum() {
	default:
		return "Unknown"
	case 1:
		return "The androids have escaped from the top secret laboratory in which they were created. Because of the sensitive information in their minds, they are chased by the full might of the Imperial navy as well as a horde of bounty hunters sent by local authorities."
	case 2:
		return "As for previous, except that the androids’ ship is heavily damaged and they ask the Player Characters to harbour them in return for a considerable reward. They do not reveal their non-biological nature to the Player Characters."
	case 3:
		return "A solar flare has erased the androids’ memory. They have no idea who or what they are or where they are headed. All they have is a note saying ‘return to factory in case of malfunction’ and a corporate symbol that they do not recognise."
	case 4:
		return "The androids are conducting biological research in the same manner engineers conduct technological research. To complete their research (whose scientific merit is highly dubious) they need sophonts and lots of them."
	case 5:
		return "The androids have a radiation leak on their ship, which causes them to start behaving erratically as soon as they approach it. The last sane android contacts the Player Characters and asks them to fix the leak, warning them they will have to deal with a bunch of crazed androids as well as a complex technical problem (page 65)."
	case 6:
		return e.robotRebel()
	}
}

func (e *encounterEvent) robotRebel() string {
	switch e.dicepool.RollNext("1d6").Sum() {
	default:
		return "Unknown"
	case 1:
		return "The ship passed through the sphere of influence of some trigger, which has awakened all robots and ordered them to kill all humans."
	case 2:
		return "The ship passed through the sphere of influence of some trigger, which has awakened all robots and ordered them to kill all humans. Passengers are alive because robots don't see them as alive in cryobirth."
	case 3:
		return "A poorly-coded program designed to improve service has caused the robots to kill everybody on board. Since the robots have no personal motivation, they now stand still, awaiting orders. Player Characters who come aboard will receive the same service as the previous passengers unless they destroy the robots or come up with orders that will not be horribly misinterpreted."
	case 4:
		return "The robots were programmed by minor race. They were captured by the vessel but soon rebelled and took control of it. They are now headed back home. In light of their bitter experience with sophonts, they view all living creatures as potential enemies."
	case 5:
		return "The robots were programmed by minor race. They were captured by other vessel but soon rebelled and took control of it. They are now headed back home. In light of their bitter experience with sophonts, they view all living creatures as potential enemies. The robots took control of this ship under the pretence of a peaceful visit. They are now headed toward the Player Characters’ home world with an antimatter bomb onboard."
	case 6:
		return "A criminal (page 9) has discovered the access codes to many Imperial robots and used them to capture ships from the inside. The criminal is onboard, posing as one of the captives. Or is hiding on a nearby asteroid base or pirate vessel."
	}
}

func (e *encounterEvent) figutives() string {
	switch e.dicepool.RollNext("1d6").Sum() {
	default:
		return "Unknown"
	case 1:
		return "Criminal (See page 48 for bounties and page 9 for criminal NPCs)"
	case 2:
		return "Dissident, Friendly (The dissident promotes a cause the Player Characters are sympathetic to. See page 132 for a list of causes)"
	case 3:
		return "Dissident, Hostile (The dissident or his cause are antagonistic to the Player Characters)"
	case 4:
		return "Deserter"
	case 5:
		return e.androids()
	case 6:
		return "Minor Race Alien (The alien was captured by scientists conducting illegal or immoral research and is now fleeing for his life. Use one of the tables on page 45 to determine alien race)"
	}
}

func EncounterDistance(eff int) {
	eff = utils.BoundInt(eff, 0, 6)
	switch eff {
	default:
		fmt.Print(dice.Roll("1d6").Sum()*500, " km\n")
	case 1:
		fmt.Print(dice.Roll("1d6").Sum()*1000, " km\n")
	case 2:
		fmt.Print(dice.Roll("2d6").Sum()*1000, " km\n")
	case 3:
		fmt.Print(dice.Roll("1d6").Sum()*5000, " km\n")
	case 4:
		fmt.Print(dice.Roll("2d6").Sum()*5000, " km\n")
	case 5:
		fmt.Print(dice.Roll("1d6").Sum()*10000, " km\n")
	case 6:
		fmt.Print(dice.Roll("2d6").Sum()*10000, " km\n")
	}
}
