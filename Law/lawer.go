package law

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/utils"
)

const (
	lawOverall     = "Overall"
	lawWeapon      = "Weapons"
	lawDrugs       = "Drugs"
	lawInformation = "Information"
	lawTechnology  = "Technology"
	lawTravellers  = "Travellers"
	lawPsionics    = "Psionics"
	goverment      = "Goverment"
)

type LawReport struct {
	dp         *dice.Dicepool
	ulp        string //L5-444444 [Wp Dr In Te Tr Ps]
	levelOf    map[string]int
	contraband [6]int
}

func New(uwpStr string) (LawReport, error) {
	lr := LawReport{}
	uwp, err := profile.NewUWP(uwpStr)
	if err != nil {
		return lr, err
	}
	lr.dp = dice.New(utils.SeedFromString(uwpStr))
	lr.levelOf = make(map[string]int)
	//lr.contraband = [6]int{2, 2, 2, 2, 2, 2}

	lr.levelOf[lawOverall] = TrvCore.EhexToDigit(uwp.Laws())
	lr.levelOf[goverment] = TrvCore.EhexToDigit(uwp.Govr())
	lr.determineContraband()
	lr.determineActivity()
	return lr, nil
}

func (lr *LawReport) determineActivity() {
	activities := []string{lawWeapon, lawDrugs, lawInformation, lawTechnology, lawTravellers, lawPsionics}
	for i, activitie := range activities {
		newLevel := lr.dp.RollNext("2d6").DM(lr.levelOf[lawOverall] - 7).Sum()
		if lr.contraband[i] == 0 {
			lr.levelOf[activitie] = utils.BoundInt(newLevel, 0, lr.levelOf[lawOverall])
		} else {
			lr.levelOf[activitie] = utils.BoundInt(newLevel, lr.levelOf[lawOverall], 18)
		}
	}
}

func (lr *LawReport) determineContraband() {
	contrSl := [6]int{}
	switch lr.levelOf[goverment] {
	case 0:
		contrSl = [6]int{0, 0, 0, 0, 0, 0}
	case 1:
		contrSl = [6]int{1, 1, 0, 0, 1, 0}
	case 2:
		contrSl = [6]int{0, 1, 0, 0, 0, 0}
	case 3:
		contrSl = [6]int{1, 0, 0, 1, 1, 0}
	case 4:
		contrSl = [6]int{1, 1, 0, 0, 0, 1}
	case 5:
		contrSl = [6]int{1, 0, 1, 1, 0, 0}
	case 6:
		contrSl = [6]int{1, 0, 0, 1, 1, 0}
	case 7:
		contrSl = [6]int{2, 2, 2, 2, 2, 2}
	case 8:
		contrSl = [6]int{1, 1, 0, 0, 0, 0}
	case 9:
		contrSl = [6]int{1, 1, 0, 1, 1, 1}
	case 10:
		contrSl = [6]int{0, 0, 0, 0, 0, 0}
	case 11:
		contrSl = [6]int{1, 0, 1, 1, 0, 0}
	case 12:
		contrSl = [6]int{1, 0, 0, 0, 0, 0}
	default:
		contrSl = [6]int{2, 2, 2, 2, 2, 2}
	}
	for i := range contrSl {
		if contrSl[i] != 0 && contrSl[i] != 1 {
			r := lr.dp.RollNext("1d2").DM(-1).Sum()
			contrSl[i] = r
		}
	}
	lr.contraband = contrSl
}

func (lr LawReport) ULP() string {
	if lr.ulp != "" {
		return lr.ulp
	}
	ulp := "L"
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawOverall]) + "-"
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawWeapon])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawDrugs])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawInformation])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawTechnology])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawTravellers])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawPsionics]) + " "
	//L5-444444 [Wp Dr In Te Tr Ps]
	if lr.contraband[0] == 1 {
		ulp += "Wp "
	}
	if lr.contraband[1] == 1 {
		ulp += "Dr "
	}
	if lr.contraband[2] == 1 {
		ulp += "In "
	}
	if lr.contraband[3] == 1 {
		ulp += "Te "
	}
	if lr.contraband[4] == 1 {
		ulp += "Tr "
	}
	if lr.contraband[5] == 1 {
		ulp += "Ps "
	}
	ulp = strings.TrimSuffix(ulp, " ")
	lr.ulp = ulp
	return lr.ulp
}

func describeOverall(lr LawReport) string {
	descr := "Overall Law Level can be considered " //text \033[1mbold\033[0m text
	switch lr.levelOf[lawOverall] {
	default:
		descr += "Stifling: The epitome of legal tyranny. Nothing is accessible by the people and everything is kept under governmental control. Punishments for breaking these laws are likely the harshest possible; better to keep the populace in line with the legal regime."
	case 0:
		descr += "Lawless: This is the absence of legal authority. Either through anarchy, barbarism or other assorted social fractures, this culture does not keep a set of laws to govern the indicated items."
	case 1:
		descr += "Light Limited: The culture really only keeps restrictions upon the most extreme examples of item or action. Legal action is likely strict about the punishments concerning these light laws, however."
	case 2:
		descr += "Moderate Limited: Increasing the limitations on particular situations or items, a culture at this Law Level has begun to monitor things for the general well-being of their populace."
	case 3:
		descr += "Standard Limited: Safety laws begin to appear at this level, with the governing power trying to limit what is readily available to protect the general populace."
	case 4:
		descr += "Heavy Limited: The first appearance of ‘common sense’ laws, this Law Level has thick restrictions on anything that might cause undue or irreparable harm to the common people."
	case 5:
		descr += "Strict: The governing power has begun to set arbitrary limitations on things and situations; likely based on personal politics rather than the good of the whole."
	case 6:
		descr += "Controlled: This law level represents the shift of freedoms from the people to the governing agencies. Some specialised items and services are made illegal, forcing the people to come to the government for aid in these areas."
	case 7:
		descr += "Tight: The legal codes of this culture are designed to take away many specific options from the people as a whole. Most civilian items and services are now regulated by the government, ensuring that they are not readily available without legal means or authority."
	case 8:
		descr += "Enforced: Total governmental control over most things; this Law Level represents most military states or areas under strict martial law."
	}
	return descr + "\n"
}

func describeWeapon(lr LawReport) string {
	descr := ""
	switch lr.levelOf[lawWeapon] {
	default:
		descr += "Nothing that can be considered a weapon in any circumstance is allowed to be carried personally. From a stone tied to a stick or a shard of broken glass carried in a menacingly manner – all implements of inflicting harm are forbidden."
		fmt.Println(lr.levelOf[lawWeapon])
	case 0:
		descr += "There are no legal restrictions on Weapons."
	case 1:
		descr += "Weapons (and other combat-oriented technologies) that are designed for massive indiscriminate losses of life are outlawed at this level. Weapons of mass destruction, chemical or biological weapons and the sorts of things that terrorists use to wreak havoc upon civilian targets are considered contraband."
	case 2:
		descr += "Weapon systems that can inflict massive bodily harm upon a target and likely generate radiation on a localised level are illegal. Lasers, fusion weaponry and plasma weapons cause enough visceral damage to be considered contraband."
	case 3:
		descr += "Weaponry that requires special training and military access, often with a remarkable rate of fire that can injure multiple targets in one volley. Squad-level support weaponry like heavy machine guns and anti-tank rifles are too dangerous for casual citizens to use and the government tries to make sure they do not."
	case 4:
		descr += "Personal weaponry with high rates of automatic fire such as light assault guns and submachine guns are, at this level, thought of as too easily acquired and abused to be in the hands of the common citizen."
	case 5:
		descr += "Government restricts all weaponry that could be hidden on the average person, making it much harder for non-authoritarian figures to be lethally armed."
	case 6:
		descr += "Government restricts all manners of firearms. Only projectile weapons that are nonlethal or originally designed for hunting are permited."
	case 7:
		descr += "The government does not recognises the hunting applications of shotguns or low-impact black powder firearms, placing all slug throwers in the illegal category."
	case 8:
		descr += "All manufactured weaponry removed from the hands on non-authoritarians. Knives, primitive projectiles and even stunning equipment become restricted. Without special consideration, being armed with something designed to harm or incapacitate another is not accessible to citizens."
	}
	return descr + "\n"
}

func (lr LawReport) Report() string {
	str := "Universal Law Profile: " + lr.ULP() + "\n"
	str += describeOverall(lr)
	activities := []string{lawWeapon, lawDrugs, lawInformation, lawTechnology, lawTravellers, lawPsionics}
	for i := range activities {
		switch activities[i] {
		case lawWeapon:
			str += describeWeapon(lr)
		}
	}
	return str
}
