package autoGM

import (
	"fmt"

	"github.com/Galdoba/convert"
)

func step2Complication() {
	complication := complication()
	fmt.Println(complication)
}

func complication() string {
	switch rolld6() {
	case 1:
		return rollAmbush()
	case 2:
		return rollCompetition()
	case 3:
		return rollContactVictimised()
	case 4:
		return rollMisunderstanding()
	case 5:
		return rollOddity()
	case 6:
		return rollVillain()
	default:
		return "Error"
	}
}

func rollAmbush() string {
	text := "The meeting with the patron or service provider turns out to be a set up! See page 161 for NPCs and pages 139-148 for location maps.\n"
	add := ""
	textRoll := rolld6()
	switch textRoll {
	case 1:
		add = "Slavers want to kidnap the Player Characters and force them to participate in some sordid affair."
	case 2:
		add = "Criminals intend to rob the Player Characters. If the ambush is successful, the criminals "
		switch rolld6() {
		case 1, 2, 3:
			add = add + "kill the characters"
		case 4, 5:
			add = add + "sell the characters to alien slavers"
		case 6:
			add = add + "leave the characters bleeding on the floor."
		}
	case 3:
		add = "Strange aliens are interested in the Player Characters for some reason and set this meeting in order to study them. This is the same encounter as Alien Scouts (page 44) except that it occurs planetside. "
		switch rolld6() {
		case 1, 2, 3, 4:
			add = add + "The alien ship is hidden somewhere in the area."
		case 5, 6:
			add = add + "The alien ship is orbiting the planet with a skeleton crew."

		}
	case 4:
		add = "An insane self-proclaimed behavioural scientist needs extraordinary people for his illegal experiments. During the job interview the Player Characters are drugged and wake up in an illegal research laboratory"
	case 5:
		add = "The Player Characters arrive just in time to witness a firefight between the patron and his men and "
		switch rolld6() {
		case 1, 2:
			add = add + "security forces."
		case 3, 4:
			add = add + "competitors."
		case 5, 6:
			add = add + "mysterious cultists."
		}
	case 6:
		add = "The patron was meddling in affairs beyond his reach and is now paying the ultimate price. Instead of a job offer, the Player Characters find a horde of ravenous zombies or mutants (page 164) and must now flee for their lives and warn others before the monsters break free and start a Zombie Apocalypse (page 16)."
	}
	return text + add
}

func rollCompetition() string {
	//return "The Player Characters are not the only ones looking for the contact. If this is a patron they will have to convince him to hire them instead of the competitors by demonstrating various skills and showing credentials. In case of success there is also a 1 in 6 chance the competition will turn ugly...\n
	//In case of a service provider, the Player Characters will have to outbid the competitors, which increases the price of the service by "+ convert.ItoS(rolld6()*10)+"%."
	price := convert.ItoS(rolld6() * 10)
	return "The Player Characters are not the only ones looking for the contact. If this is a patron they will have to convince him to hire them instead of the competitors by demonstrating various skills and showing credentials. In case of success there is also a 1 in 6 chance the competition will turn ugly...\nIn case of a service provider, the Player Characters will have to outbid the competitors, which increases the price of the service by " + price + "%."
}

func rollContactVictimised() string {
	return "The patron was kidnapped or murdered by his enemies. A successful Investigation Check will reveal clues leading to his assailant’s identity. As the Player Characters investigate the murder, unseen forces repeatedly target them."
}

func rollMisunderstanding() string {
	text := "The Patron mistakes the Player Characters for his enemies and either attempts to flee or fight back. Smart and coolheaded Player Characters may resolve this situation peacefully through Persuade, Diplomacy and proof that they are not the enemy. Otherwise this encounter is likely to end in a bloodbath.\n\n"
	reaction := ""
	switch rolld6() {
	case 1:
		reaction = "Patron runs away on foot from the Player Characters."
	case 2:
		reaction = "Patron hops into a vehicle and speeds away, screaming ‘so long, suckers!’"
	case 3:
		numberOfGuards := convert.ItoS(roll2d6())
		reaction = "The patron has " + numberOfGuards + " armed guards (page 161) with him. They attack the Player Characters without warning, shooting to neutralise. Captured Player Characters will be interrogated very roughly and will have to prove they are not part of a conflict they do not know anything about."
	case 4:
		numberOfGuards := convert.ItoS(roll2d6())
		reaction = "The patron has " + numberOfGuards + " armed guards (page 161) with him. They attack the Player Characters without warning, shooting to kill. Captured Player Characters will be interrogated very roughly and will have to prove they are not part of a conflict they do not know anything about."
	case 5:
		reaction = "The patron informs the Player Characters via a recorded message that the house is booby trapped. See page 166 for traps and page 139-148 for maps."
	case 6:
		reaction = "The patron quickly pulls out a gun and blows his brains out. "
		switch rolld6() {
		case 1, 2, 3, 4:
			reaction = reaction + "Nothing in the room holds any clues as to why he acted in such an extreme manner."
		case 5, 6:
			reaction = reaction + "A successful Investigation Check will reveal the Patron was being haunted by a demonic cult he had incriminating evidence against. For more information on cults see page 49."
		}
	}
	return text + reaction
}
func rollOddity() string {
	text := "For service providers: The service provider had just developed some experimental technology and is willing to give it to the Player Characters for free in return for reporting on how it fared in field conditions upon the Player Characters’ return... unless, of course, the new technology ends up killing them.\n If the Player Characters agree to this deal, they receive an item three TL higher than what would normally be available on the planet. However, there is a one in six chance per adventure that the item will malfunction in the worst possible moment.\n For patrons: The patron is kosher but the job he offers is extremely weird, insanely dangerous or both. Some examples:\n\n"
	reaction := ""
	switch rolld6() {
	case 1:
		reaction = "An aging roboticist dreams of dying in the company of the most advanced computer in the universe, a device so powerful and complex it is worshipped as a space god by many races. He is willing to pay the Player Characters an astronomic sum of money for transporting him to the heart of this vast AI."
	case 2:
		reaction = "The high priest of a small and backward nation hires the Player Characters to find his people a new God after the previous one has not answered their prayers for more than a decade. The Player Characters need to locate a suitable higher entity and convince it to follow them to the Godless world."
	case 3:
		reaction = "A wealthy industrialist cannot get over the death of his beloved wife in a space accident. Every night he dreams of her screaming for help from inside the belly of a mighty beast. The dreams are not just a product of his imagination. His wife has been swallowed, along with her ship, by a colossal alien (page 163). Part of her persona still lives in its mind and communicates with the world through its powerful psionic talent."
	case 4:
		reaction = "As for previous only the industrialist is in fact a creation of the woman’s imagination, a desperate attempt to attract help from the outside."
	case 5:
		reaction = "A unique and lonely alien is looking for a date. His species is extinct but maybe somewhere out there, there is some being more or less compatible with the alien."
	case 6:
		reaction = "The same alien hires the Player Characters again to babysit its child, a cute little thing with psionic abilities undreamed of by any of the major races and no comprehension of the fragility of the human body."
	}
	return text + reaction
}

func rollVillain() string {
	return "A villain is a patron with hostile ulterior motives, such as the corrupt noble sending merchants to their deaths at the hands of pirates in return for some of the shares or a deranged psychopath playing a twisted game with human lives. The difference between this complication and an ambush is that in the latter the hostile action is immediate while the former will wait for the right moment to strike.\n For more information on villains, as well as some ready to use NPCs, see pages 153-163."
}
