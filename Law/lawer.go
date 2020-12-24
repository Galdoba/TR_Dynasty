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

type lawReport struct {
	dp         *dice.Dicepool
	ulp        string //		L5-444444 [Wp Dr In Te Tr Ps]
	levelOf    map[string]int
	contraband [6]int
}

func NewLawReport(uwpStr string) (lawReport, error) {
	lr := lawReport{}
	uwp := profile.NewUWP(uwpStr)

	lr.dp = dice.New().SetSeed(uwpStr)
	lr.levelOf = make(map[string]int)
	lr.levelOf[lawOverall] = uwp.Laws().Value()
	lr.levelOf[goverment] = uwp.Govr().Value()
	lr.determineContraband()
	lr.determineActivity()
	return lr, nil
}

func (lr *lawReport) determineActivity() {
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

func (lr *lawReport) determineContraband() {
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

func (lr lawReport) ULP() string {
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

func describeOverall(lr lawReport) string {
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
		descr += "Heavy Limited: The first appearance of 'common sense' laws, this Law Level has thick restrictions on anything that might cause undue or irreparable harm to the common people."
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

func describeWeapon(lr lawReport) string {
	descr := ""
	switch lr.levelOf[lawWeapon] {
	default:
		descr += "Nothing that can be considered a weapon in any circumstance is allowed to be carried personally. From a stone tied to a stick or a shard of broken glass carried in a menacingly manner – all implements of inflicting harm are forbidden."
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

func describeDrugs(lr lawReport) string {
	descr := ""
	switch lr.levelOf[lawDrugs] {
	default:
		descr += "All chemical additives of any kind are restricted by the government. No medicines, no vitamins, no anagathics; nothing of the sort are considered legal in any way."
	case 0:
		descr += "There are no legal restrictions on Drugs."
	case 1:
		descr += "Only the most dangerous and physically addictive drugs are disallowed, deemed utterly unsafe by the governing medical minds."
	case 2:
		descr += "Drugs that are physically demanding or that can easily create dependencies are added to the list of what the government does not want accessible by the people."
	case 3:
		descr += "The possible medical ramifications of using combat enhancing drugs can turn regular citizens into serious problems for local law enforcement, which is why they are deemed illegal."
	case 4:
		descr += "Anything with an addictive property to its chemical structure is added to the illegal category of drugs by the government, who want to keep its citizens from becoming dependant on them."
	case 5:
		descr += "Powerful drugs that stem the effects of aging, anagathics are considered dangerous in the hands of common citizens. At this Law Level they are governmentally controlled and kept watch over to ensure 'safe' usage."
	case 6:
		descr += "The potential for recreational use of medicinal drugs like Fast and Slow are brought into suspicion of abuse at this Law Level. The government feels that only medical professionals should make use of such chemical stimulants and depressants."
	case 7:
		descr += "The government is in position where they want to keep all narcotics – both harmful and beneficial, medically speaking – out of the hands of the general populace. Only through governmentally sanctioned medical avenues can they be attained and only in minimum doses to avoid stockpiling."
	case 8:
		descr += "Drugs not even accessible through medical avenues, pharmaceuticals are considered to be governmentally controlled only."
	}
	return descr + "\n"
}

func describeInformation(lr lawReport) string {
	descr := ""
	switch lr.levelOf[lawInformation] {
	default:
		descr += "Total governmental censorship curbs all information from offplanet or even beyond the local level. All mass media is sponsored and fabricated by the controlling governmental agency."
	case 0:
		descr += "There are no legal restrictions on Information."
	case 1:
		descr += "The government only cares about the use of artificial intelligences and intellect programming; anything that can make an inorganic thing seem sentient. Robots with lifelike programming that still remain non-heuristic are generally viewed as 'optionally legal'."
	case 2:
		descr += "Software that hides codes, information and smuggled data in a normal computer is rarely used for upstanding purposes to begin with but this Law Level forces all Agent programming of that sort to be considered illegal. Even if being used to send encrypted messages or store classified data, these programs are not to be used."
	case 3:
		descr += "Software packages designed to break through other encryption software to the illegal list at this Law Level is supposed to help protect the government from hackers and data thieves. The government begins to set up censorship protocols to keep anti-government activists from learning too much about classified goings on."
	case 4:
		descr += "The government no longer wants the citizen public to have access to software that could keep them from seeing what the people are doing electronically. Security programs are solely owned by officials and spies, with everyone else`s data being public access."
	case 5:
		descr += "Information Technology is severely restricted. The usage of any professional software programs is tightly limited."
	case 6:
		descr += "Electronic information from other worlds has begun to be held as privileged data by the government at this Law Level, only allowing older facts that are less likely to truly affect their populace through the filters."
	case 7:
		descr += "True censorship begins to surface at this Law Level, with historic and off-world data suffering the greatest amount of editing by the government. The general populace also finds that they cannot spread their own data on mass media levels without governmental approval, which is always heavily edited."
	case 8:
		descr += "The government tightens its fingers around what the civilian populace can learn, filtering all information that can be researched by the people. Off-world data not deemed as necessary for everyday life is silenced and anything published or broadcast by the people is heavily edited and manipulated by the state."
	}
	return descr + "\n"
}

func describeTechnology(lr lawReport) string {
	descr := ""
	switch lr.levelOf[lawTechnology] {
	default:
		descr += "A state-imposed medieval culture, this Law Level only allows extremely simple mechanical equipment and the knowledge to use them, ceasing the legal use of steam works or industrial advances (TL 3+)."
	case 0:
		descr += "There are no legal restrictions on Technology."
	case 1:
		descr += "Making only the most dangerous and invasive technologies illegal, like full-body replacement cybernetics and self-replicating nanotech, the government keeps these devices out of the hands of those who could not handle the power or responsibility."
	case 2:
		descr += "Watching out for technologies that are designed and manufactured by alien species, the government sets laws against the use of alien sciences. This could be due to a protectionist view of their people's goods or perhaps on account of some kind of xenophobia."
	case 3:
		descr += "Fearing that the common citizen cannot handle the responsibility for themselves, the government claims that all ultra-science technology (TL 15) is restricted without their clearance."
	case 4:
		descr += "Advanced technologies (TL 13+) is in the list of unauthorised access removes several important medical procedures and useful pieces of equipment but helps keep the government in control of well-studied sciences."
	case 5:
		descr += "As the legal grip of the government closes to this Law Level, technologies thought of as general Imperial status quos (TL 11+) are now restricted. Some useful applications, like drone design and chemical engineering, are outlawed for common use."
	case 6:
		descr += "The government not allows most spacefaring sciences (TL 9+) to be practiced by their populations, leaving all jump technologies in the hands of the military and governmentally-sponsored officials. This is the beginning of the legal nullification of a learned populace."
	case 7:
		descr += "Officially keeping all important jump-theory science knowledge and equipment (TL 7+) behind locked doors along with anything else that shows a modern look at invention and design, this Law Level forces its people to remain in a combustion engine era under penalty of law."
	case 8:
		descr += "Forcing its people to remain in the industrial era of technological advancement, the equipment and instruments that might lead to higher science (TL 5+) are now forbidden. The government has managed to keep its citizens simple and likely blinded to the wonders of the galaxy."
	}
	return descr + "\n"
}

func describeTravellers(lr lawReport) string {
	descr := ""
	switch lr.levelOf[lawTravellers] {
	default:
		descr += "Complete planetary transit lockdown. The government not allows off-world traffic to or from their planet, enforcing their laws with force if they have to. Only Imperial agents with business on the ground will be allowed to land."
	case 0:
		descr += "There are no legal restrictions on Travellers."
	case 1:
		descr += "The government keeps open borders and allows free travel across them, only requesting travellers check in with local authorities when they land. Radio contact, even delayed response radios, is considered enough warning to allow a landing anywhere the ship can find room to do so."
	case 2:
		descr += "The government is checkinig passenger and cargo manifests before allowing a starship to land, normally doing so electronically while the ship is in orbit. The ship is not restricted in any other way."
	case 3:
		descr += "The transit authority system requires vessels to land at specified locations on the planet to ensure safe flights and inter-ship traffic. Starports, landing zones and berthing sites are registered with the government as available landing locations."
	case 4:
		descr += "The government keeps security ships and personnel around specified governmental starport locations, only allowing ships to make landings at these sites when cleared to do so. The starport monitors traffic and regulates its flow."
	case 5:
		descr += "The passage to and from the planet is now recorded for security reasons. Citizens going off-world must inform the government as to why and how long; incoming traffic must declare and register why they are visiting the planet. The government may or may not use this information to keep tabs on their populace."
	case 6:
		descr += "Making starport access stark, unattractive and business-only, the government creates laws to try and keep off-world visitation to a bare minimum. Starports are no longer commonly accessible by the local populace, only those with special permission or clearance."
	case 7:
		descr += "The government officially limits the mingling of outsiders and the local populace. Citizens can only leave the planet if doing so at the behest of the military or government and any off-worlders are only allowed access to the militarily-controlled and secured starport grounds. Only government agents are allowed to do business with the off-worlders at the starport or electronically only."
	case 8:
		descr += "The planet is closed to the vast majority of outside access. Only those carrying Imperial transit authority papers or other interstellar documentation are allowed to land but are still restricted to starport locations only. The government keeps a minimal staff at these starports, further limiting interactions even between these Imperial agents and the population."
	}
	return descr + "\n"
}

func describePsionics(lr lawReport) string {
	descr := ""
	switch lr.levelOf[lawPsionics] {
	default:
		descr += "The government has a strict 'no tolerance' policy for anything at all related to psionic ability, technology or genetics. Psions are sterilised, lobotomised or worse."
	case 0:
		descr += "There are no legal restrictions on Psionics."
	case 1:
		descr += "With only the basest regulations on psionic abilities, the government feels it necessary to keep track of convicted psi-offenders or psions with particularly dangerous talents. The government does not restrict these abilities but needs to know who has access to them."
	case 2:
		descr += "Governmental control over psionic abilities is increased to a standard registration of all psions, noting who has access to which abilities in case they are to be utilised, monitored or apprehended. A new set of 'psi-laws' is put into place that seriously punishes those who use dangerous or overly invasive psionic abilities."
	case 3:
		descr += "Government restricts the psionic members of a population by requiring any and all practicing psions with the telepathy talent to be employed by the government. Telepathic use outside of government-approved services will be punished by antipsi chemicals and inhibiting drug treatments."
	case 4:
		descr += "Due to the potential illicit use of such psionic abilities as teleportation and clairvoyance, the government has added these talents to the list of psionic restrictions. Only those approved by and registered with the government can use these talents without persecution."
	case 5:
		descr += "No longer seeing any distinction between the psionic talents, the government claims complete control over psions. Only those employed by the governing agency can use their talents, all others are chemically or surgically inhibited."
	case 6:
		descr += "Goverment treats psionic-enhancing drugs and chemicals like any other dangerous narcotic substance. Possessing such substances is considered reason enough to prosecute the individual for 'intention to break psi-restrictions'."
	case 7:
		descr += "The government no longer caters to psionic abilities in any way, seeing any use of them as a dangerous crime against the populace. Being a natural psion is likely to mean being forced to take chemical inhibitors by the state or being deported from the common population for 'mental safekeeping'."
	case 8:
		descr += "The government sees the use or manufacture of anything involving psionic technologies as a precursor to psionic revolution, banning their very existence under strict penalties. The only exception is often a government 'anti-psi' agency, which is likely armed with technologies aimed at nullifying psions."
	}
	return descr + "\n"
}

func (lr lawReport) Report() string {
	str := "Universal Law Profile: " + lr.ULP() + "\n\n"
	str += describeOverall(lr)
	activities := []string{lawWeapon, lawDrugs, lawInformation, lawTechnology, lawTravellers, lawPsionics}
	for i := range activities {
		switch activities[i] {
		case lawWeapon:
			str += describeWeapon(lr)
		case lawDrugs:
			str += describeDrugs(lr)
		case lawInformation:
			str += describeInformation(lr)
		case lawTechnology:
			str += describeTechnology(lr)
		case lawTravellers:
			str += describeTravellers(lr)
		case lawPsionics:
			str += describePsionics(lr)
		}
	}
	return str
}

func Describe(uwp string) string {
	lr, err := NewLawReport(uwp)
	if err != nil {
		fmt.Println(err.Error())
	}
	return lr.Report()
}
