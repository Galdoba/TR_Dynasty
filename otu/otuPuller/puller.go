package puller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Galdoba/utils"
)

//PullOtuData -
func PullOtuData() {
	var client http.Client
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	totalSectors := len(allSectors())
	for i, val := range allSectors() {

		resp, err := client.Get(val)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			//log.Info(bodyString)
			fmt.Println(bodyString)

			l, err := f.WriteString(bodyString)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			fmt.Println(l, "bytes written successfully")

		}
		fmt.Println("Pulled", i, "sectors of", totalSectors)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := utils.LinesFromTXT("data.txt")
	fmt.Println(lines)
	for i, val := range lines {
		fmt.Println("Check entry", i)
		fmt.Println(val)
		val = strings.ReplaceAll(val, "’", "'")

		if val != "Sector	SS	Hex	Name	UWP	Bases	Remarks	Zone	PBG	Allegiance	Stars	{Ix}	(Ex)	[Cx]	Nobility	W	RU" {
			utils.AddLineToFile("dataClean.txt", val)
		}
	}
	fmt.Println("Formatting...")
	linesC := utils.LinesFromTXT("dataClean.txt")
	//Sector	SS	Hex	Name	UWP	Bases	Remarks	Zone	PBG	Allegiance	Stars	{Ix}	(Ex)	[Cx]	Nobility	W	RU
	mapLen := make(map[int]int)
	mapLen[0] = 6
	mapLen[1] = 2
	mapLen[2] = 4
	mapLen[3] = 28
	mapLen[4] = 9
	mapLen[5] = 5
	mapLen[6] = 44
	mapLen[7] = 4
	mapLen[8] = 3
	mapLen[9] = 10
	mapLen[10] = 29
	mapLen[11] = 7
	mapLen[12] = 7
	mapLen[13] = 6
	mapLen[14] = 8
	mapLen[15] = 2
	mapLen[16] = 5
	total := len(lines)
	for i, val := range linesC {

		if i == 1 {
			template := []string{}
			for col := 0; col < 17; col++ {
				template = append(template, "-")
				for len(template[col]) < mapLen[col] {
					template[col] += "-"
				}

			}
			tempF := strings.Join(template, " ")
			utils.AddLineToFile("dataFormatted.txt", tempF)
		}
		data := strings.Split(val, "	")
		//fmt.Print("Formatting Line ", i)
		for col := 0; col < 17; col++ {
			// if mapLen[col] < len(data[col]) {
			// 	mapLen[col] = len(data[col])
			// }
			// fmt.Print(mapLen[col], " ")
			// sum += mapLen[col]
			for len(data[col]) < mapLen[col] {
				data[col] += " "
				//		fmt.Print(".")
			}
		}
		dataF := strings.Join(data, " ")
		utils.AddLineToFile("dataFormatted.txt", dataF)
		//fmt.Print("ok\n")
		//fmt.Println(dataF)
		fmt.Print(i, " of ", total, " complete\r")
	}
	//6 2 4 28 9 5 44 4 3 10 29 7 7 6 8 2 5  - Line 66191 - max len(line):179
}

func allSectors() []string {
	return []string{
		"https://travellermap.com/data/Caesillian/tab",
		"https://travellermap.com/data/Chtierabl/tab",
		"https://travellermap.com/data/Stinj%20Tianz/tab",
		"https://travellermap.com/data/Bliardlie/tab",
		"https://travellermap.com/data/Zhiensh/tab",
		"https://travellermap.com/data/Savria/tab",
		"https://travellermap.com/data/Datsatl/tab",
		"https://travellermap.com/data/Gakghang/tab",
		"https://travellermap.com/data/Thaku%20Fung/tab",
		"https://travellermap.com/data/Rzakki/tab",
		"https://travellermap.com/data/Listanaya/tab",
		"https://travellermap.com/data/Veg%20Fergakh/tab",
		"https://travellermap.com/data/Dfotseth/tab",
		"https://travellermap.com/data/Irugangog/tab",
		"https://travellermap.com/data/Finggvakhou/tab",
		"https://travellermap.com/data/Zortakh/tab",
		"https://travellermap.com/data/Spangele/tab",
		"https://travellermap.com/data/Nadir/tab",
		"https://travellermap.com/data/Harbinger/tab",
		"https://travellermap.com/data/Extremus/tab",
		"https://travellermap.com/data/Viajlefliez/tab",
		"https://travellermap.com/data/Bleblqansh/tab",
		"https://travellermap.com/data/Driasera/tab",
		"https://travellermap.com/data/Dalchie%20Jdatl/tab",
		"https://travellermap.com/data/Chit%20Botshti/tab",
		"https://travellermap.com/data/Ghoekhnael/tab",
		"https://travellermap.com/data/Ksinanirz/tab",
		"https://travellermap.com/data/Zao%20Kfeng%20Ig%20Grilokh/tab",
		"https://travellermap.com/data/Knaeleng/tab",
		"https://travellermap.com/data/Kharrthon/tab",
		"https://travellermap.com/data/Oeghz%20Vaerrghr/tab",
		"https://travellermap.com/data/Kfazz%20Ghik/tab",
		"https://travellermap.com/data/Angfutsag/tab",
		"https://travellermap.com/data/Rfigh/tab",
		"https://travellermap.com/data/Tar'G'kell'p/tab",
		"https://travellermap.com/data/Kteex!/tab",
		"https://travellermap.com/data/Koog/tab",
		"https://travellermap.com/data/Xeeleer/tab",
		"https://travellermap.com/data/Unthaank/tab",
		"https://travellermap.com/data/Brieplanz/tab",
		"https://travellermap.com/data/Sidiadl/tab",
		"https://travellermap.com/data/Zdiedeiant/tab",
		"https://travellermap.com/data/Stiatlchepr/tab",
		"https://travellermap.com/data/Itvikiastaf/tab",
		"https://travellermap.com/data/Knoellighz/tab",
		"https://travellermap.com/data/Dhuerorrg/tab",
		"https://travellermap.com/data/Ngathksirz/tab",
		"https://travellermap.com/data/Fa%20Dzaets/tab",
		"https://travellermap.com/data/Gzaekfueg/tab",
		"https://travellermap.com/data/Gashikan/tab",
		"https://travellermap.com/data/Trenchans/tab",
		"https://travellermap.com/data/Ktiin'gzat/tab",
		"https://travellermap.com/data/Mugheen't/tab",
		"https://travellermap.com/data/Grikr!ng/tab",
		"https://travellermap.com/data/Ukaarriit!!b/tab",
		"https://travellermap.com/data/Kring%20Noor/tab",
		"https://travellermap.com/data/Mbil!!gh/tab",
		"https://travellermap.com/data/Kweenexis/tab",
		"https://travellermap.com/data/Pliabriebl/tab",
		"https://travellermap.com/data/Eiaplial/tab",
		"https://travellermap.com/data/Zhdant/tab",
		"https://travellermap.com/data/Tienspevnekr/tab",
		"https://travellermap.com/data/Ziafrplians/tab",
		"https://travellermap.com/data/Gvurrdon/tab",
		"https://travellermap.com/data/Tuglikki/tab",
		"https://travellermap.com/data/Provence/tab",
		"https://travellermap.com/data/Windhorn/tab",
		"https://travellermap.com/data/Meshan/tab",
		"https://travellermap.com/data/Mendan/tab",
		"https://travellermap.com/data/Amdukan/tab",
		"https://travellermap.com/data/Arzul/tab",
		"https://travellermap.com/data/Gn'hk'r/tab",
		"https://travellermap.com/data/Gur/tab",
		"https://travellermap.com/data/Un'k!!k'ng/tab",
		"https://travellermap.com/data/Xaagr/tab",
		"https://travellermap.com/data/Eekrookrigz/tab",
		"https://travellermap.com/data/Hoboreth/tab",
		"https://travellermap.com/data/Tsadra%20Davr/tab",
		"https://travellermap.com/data/Tsadra/tab",
		"https://travellermap.com/data/Yiklerzdanzh/tab",
		"https://travellermap.com/data/Far%20Frontiers/tab",
		"https://travellermap.com/data/Foreven/tab",
		"https://travellermap.com/data/Spinward%20Marches/tab",
		"https://travellermap.com/data/Deneb/tab",
		"https://travellermap.com/data/Corridor/tab",
		"https://travellermap.com/data/Vland/tab",
		"https://travellermap.com/data/Lishun/tab",
		"https://travellermap.com/data/Antares/tab",
		"https://travellermap.com/data/Empty%20Quarter/tab",
		"https://travellermap.com/data/Star's%20End/tab",
		"https://travellermap.com/data/Gh!hken/tab",
		"https://travellermap.com/data/Ruupiin/tab",
		"https://travellermap.com/data/Raakaan/tab",
		"https://travellermap.com/data/Uuk/tab",
		"https://travellermap.com/data/Gnaa%20Iimb'kr/tab",
		"https://travellermap.com/data/Jiti/tab",
		"https://travellermap.com/data/Chiep%20Zhez/tab",
		"https://travellermap.com/data/Astron/tab",
		"https://travellermap.com/data/Fulani/tab",
		"https://travellermap.com/data/Vanguard%20Reaches/tab",
		"https://travellermap.com/data/The%20Beyond/tab",
		"https://travellermap.com/data/Trojan%20Reach/tab",
		"https://travellermap.com/data/Reft/tab",
		"https://travellermap.com/data/Gushemege/tab",
		"https://travellermap.com/data/Dagudashaag/tab",
		"https://travellermap.com/data/Core/tab",
		"https://travellermap.com/data/Fornast/tab",
		"https://travellermap.com/data/Ley/tab",
		"https://travellermap.com/data/Gateway/tab",
		"https://travellermap.com/data/Luretiir!girr/tab",
		"https://travellermap.com/data/X'kug/tab",
		"https://travellermap.com/data/Kilong/tab",
		"https://travellermap.com/data/Bar'kakr/tab",
		"https://travellermap.com/data/Mighabohk/tab",
		"https://travellermap.com/data/Sing/tab",
		"https://travellermap.com/data/Mavuzog/tab",
		"https://travellermap.com/data/Theta%20Borealis/tab",
		"https://travellermap.com/data/Theron/tab",
		"https://travellermap.com/data/Iphigenaia/tab",
		"https://travellermap.com/data/Touchstone/tab",
		"https://travellermap.com/data/Riftspan%20Reaches/tab",
		"https://travellermap.com/data/Verge/tab",
		"https://travellermap.com/data/Ilelish/tab",
		"https://travellermap.com/data/Zarushagar/tab",
		"https://travellermap.com/data/Massilia/tab",
		"https://travellermap.com/data/Delphi/tab",
		"https://travellermap.com/data/Glimmerdrift%20Reaches/tab",
		"https://travellermap.com/data/Crucis%20Margin/tab",
		"https://travellermap.com/data/Kaa%20G!'kul/tab",
		"https://travellermap.com/data/Gzirr!k'l/tab",
		"https://travellermap.com/data/K'trekreer/tab",
		"https://travellermap.com/data/Nuughe/tab",
		"https://travellermap.com/data/N!!krumbiix/tab",
		"https://travellermap.com/data/Uidlexna/tab",
		"https://travellermap.com/data/Harea/tab",
		"https://travellermap.com/data/Khaeaw/tab",
		"https://travellermap.com/data/Faoheiroi'iyhao/tab",
		"https://travellermap.com/data/Ftaoiyekyu/tab",
		"https://travellermap.com/data/Afawahisa/tab",
		"https://travellermap.com/data/Hlakhoi/tab",
		"https://travellermap.com/data/Ealiyasiyw/tab",
		"https://travellermap.com/data/Reaver's%20Deep/tab",
		"https://travellermap.com/data/Daibei/tab",
		"https://travellermap.com/data/Diaspora/tab",
		"https://travellermap.com/data/Old%20Expanses/tab",
		"https://travellermap.com/data/Hinterworlds/tab",
		"https://travellermap.com/data/Leonidae/tab",
		"https://travellermap.com/data/Extolian/tab",
		"https://travellermap.com/data/Ricenden/tab",
		"https://travellermap.com/data/Blaskon/tab",
		"https://travellermap.com/data/Nooq/tab",
		"https://travellermap.com/data/Gzektixk/tab",
		"https://travellermap.com/data/Googalesh/tab",
		"https://travellermap.com/data/Tlyasea/tab",
		"https://travellermap.com/data/Hkakhaeaw/tab",
		"https://travellermap.com/data/Esai'yo/tab",
		"https://travellermap.com/data/Waroatahe/tab",
		"https://travellermap.com/data/Karleaya/tab",
		"https://travellermap.com/data/Staihaia'yo/tab",
		"https://travellermap.com/data/Iwahfuah/tab",
		"https://travellermap.com/data/Dark%20Nebula/tab",
		"https://travellermap.com/data/Magyar/tab",
		"https://travellermap.com/data/Solomani%20Rim/tab",
		"https://travellermap.com/data/Alpha%20Crucis/tab",
		"https://travellermap.com/data/Spica/tab",
		"https://travellermap.com/data/Phlask/tab",
		"https://travellermap.com/data/Centrax/tab",
		"https://travellermap.com/data/Wrenton/tab",
		"https://travellermap.com/data/Folgore/tab",
		"https://travellermap.com/data/Avereguar/tab",
		"https://travellermap.com/data/Kolire/tab",
		"https://travellermap.com/data/Taezohm/tab",
		"https://travellermap.com/data/Khuaryakh/tab",
		"https://travellermap.com/data/Yahehwe/tab",
		"https://travellermap.com/data/Kefiykhta/tab",
		"https://travellermap.com/data/Heakhafaw/tab",
		"https://travellermap.com/data/Etakhasoa/tab",
		"https://travellermap.com/data/Aktifao/tab",
		"https://travellermap.com/data/Uistilrao/tab",
		"https://travellermap.com/data/Ustral%20Quadrant/tab",
		"https://travellermap.com/data/Canopus/tab",
		"https://travellermap.com/data/Aldebaran/tab",
		"https://travellermap.com/data/Neworld/tab",
		"https://travellermap.com/data/Langere/tab",
		"https://travellermap.com/data/Drakken/tab",
		"https://travellermap.com/data/Lorspane/tab",
		"https://travellermap.com/data/Porlock/tab",
		"https://travellermap.com/data/Kidunal/tab",
		"https://travellermap.com/data/Treece/tab",
		"https://travellermap.com/data/Genfert/tab",
		"https://travellermap.com/data/The%20Roast/tab",
		"https://travellermap.com/data/Aftailr/tab",
		"https://travellermap.com/data/Ohieraoi/tab",
		"https://travellermap.com/data/Fahreahluis/tab",
		"https://travellermap.com/data/Hfiywitir/tab",
		"https://travellermap.com/data/Irlaftalea/tab",
		"https://travellermap.com/data/Teahloarifu/tab",
		"https://travellermap.com/data/Ahkiweahi'/tab",
		"https://travellermap.com/data/Banners/tab",
		"https://travellermap.com/data/Hanstone/tab",
		"https://travellermap.com/data/Malorn/tab",
		"https://travellermap.com/data/Hadji/tab",
		"https://travellermap.com/data/Storr/tab",
		"https://travellermap.com/data/Mikhail/tab",
		"https://travellermap.com/data/Darret/tab",
		"https://travellermap.com/data/Ataurre/tab",
		"https://travellermap.com/data/Katoonah/tab",
		"https://travellermap.com/data/Uytal/tab",
		"https://travellermap.com/data/Sporelex/tab",
		"https://travellermap.com/data/Olantar/tab",
		"https://travellermap.com/data/Tahahroal/tab",
		"https://travellermap.com/data/A'yosea/tab",
		"https://travellermap.com/data/Usoirarloiau/tab",
		"https://travellermap.com/data/Oiah/tab",
		"https://travellermap.com/data/Eahyaw/tab",
		"https://travellermap.com/data/Ftyer/tab",
		"https://travellermap.com/data/Elyetleisiyea/tab",
		"https://travellermap.com/data/Ahriman/tab",
		"https://travellermap.com/data/Holowon/tab",
		"https://travellermap.com/data/Amderstun/tab",
		"https://travellermap.com/data/RimReach/tab",
		"https://travellermap.com/data/Phlange/tab",
		"https://travellermap.com/data/Tracerie/tab",
		"https://travellermap.com/data/Wrence/tab",
		"https://travellermap.com/data/Muarne/tab",
		"https://travellermap.com/data/Lancask/tab",
		"https://travellermap.com/data/Tensk/tab",
		"https://travellermap.com/data/Aphlent/tab",
		"https://travellermap.com/data/Randtred/tab",
		"https://travellermap.com/data/Uroac/tab",
		"https://travellermap.com/data/Iphl/tab",
		"https://travellermap.com/data/Qona/tab",
		"https://travellermap.com/data/Qiask/tab",
		"https://travellermap.com/data/Uaftdual/tab",
		"https://travellermap.com/data/Orna/tab",
		"https://travellermap.com/data/Noontask/tab",
		"https://travellermap.com/data/Onoo/tab",
		"https://travellermap.com/data/Dranon/tab",
		"https://travellermap.com/data/Geerphli/tab",
		"https://travellermap.com/data/Lila/tab",
		"https://travellermap.com/data/Koirtrua/tab",
		"https://travellermap.com/data/Ultret/tab",
		"https://travellermap.com/data/Binaurie/tab",
		"https://travellermap.com/data/Grevanne/tab",
		"https://travellermap.com/data/Femengaal/tab",
		"https://travellermap.com/data/Wobanna/tab",
		"https://travellermap.com/data/Poinmaali/tab",
		"https://travellermap.com/data/Hintrab/tab",
		"https://travellermap.com/data/Ebollam/tab",
		"https://travellermap.com/data/Tumluun/tab",
		"https://travellermap.com/data/Yasktier/tab",
		"https://travellermap.com/data/Uumpallu/tab",
		"https://travellermap.com/data/Quavas/tab",
		"https://travellermap.com/data/Indierri/tab",
		"https://travellermap.com/data/Alila/tab",
		"https://travellermap.com/data/Gudala/tab",
		"https://travellermap.com/data/Giskakii/tab",
		"https://travellermap.com/data/Incognita%20Ulterior/tab",
		"https://travellermap.com/data/Coreward%20Shore/tab",
		"https://travellermap.com/data/Greenwald's%20Beach/tab",
		"https://travellermap.com/data/Narrow%20Transit/tab",
		"https://travellermap.com/data/Sisyphus%20Ulterior/tab",
		"https://travellermap.com/data/Pytheus%20Interior/tab",
		"https://travellermap.com/data/Idirda/tab",
		"https://travellermap.com/data/Kaalin%20Ulterior/tab",
		"https://travellermap.com/data/Near%20Side%20of%20Yonder/tab",
		"https://travellermap.com/data/Deepnight/tab",
		"https://travellermap.com/data/Vilaakasii/tab",
		"https://travellermap.com/data/Incognita%20Citerior/tab",
		"https://travellermap.com/data/Central%20Bay/tab",
		"https://travellermap.com/data/Far%20Shore%20of%20Yonder/tab",
		"https://travellermap.com/data/Colaeus/tab",
		"https://travellermap.com/data/Sisyphus%20Citerior/tab",
		"https://travellermap.com/data/Pytheas/tab",
		"https://travellermap.com/data/Diamond%20Scatter/tab",
		"https://travellermap.com/data/Kaalin%20Citerior/tab",
		"https://travellermap.com/data/First%20Prospect/tab",
		"https://travellermap.com/data/Few%20Stars/tab",
		"https://travellermap.com/data/No%20Shore/tab",
		"https://travellermap.com/data/Big%20Empty/tab",
		"https://travellermap.com/data/Open%20Rift/tab",
		"https://travellermap.com/data/Last%20Prospect/tab",
		"https://travellermap.com/data/Best%20Prospect/tab",
		"https://travellermap.com/data/No%20Prospect/tab",
		"https://travellermap.com/data/Black%20Night/tab",
		"https://travellermap.com/data/Few%20Glimmers/tab",
		"https://travellermap.com/data/Just%20Empty/tab",
		"https://travellermap.com/data/Crossing/tab",
		"https://travellermap.com/data/Far%20Shore/tab",
		"https://travellermap.com/data/Far%20Shore%202/tab",
		"https://travellermap.com/data/Annwn/tab",
		"https://travellermap.com/data/Baltia/tab",
		"https://travellermap.com/data/Brittia/tab",
		"https://travellermap.com/data/Lemuria/tab",
		"https://travellermap.com/data/Muspelheim/tab",
		"https://travellermap.com/data/Hades/tab",
		"https://travellermap.com/data/Harappa/tab",
		"https://travellermap.com/data/Atlantis/tab",
		"https://travellermap.com/data/Wilderlands/tab",
		"https://travellermap.com/data/Ifingr/tab",
		"https://travellermap.com/data/Styx/tab",
		"https://travellermap.com/data/Tir%20NaNog/tab",
		"https://travellermap.com/data/Olympus/tab",
		"https://travellermap.com/data/Bifrost/tab",
		"https://travellermap.com/data/Valhol/tab",
		"https://travellermap.com/data/Hvergelmir/tab",
		"https://travellermap.com/data/Acheron/tab",
		"https://travellermap.com/data/Cathago/tab",
		"https://travellermap.com/data/Arcadia/tab",
		"https://travellermap.com/data/Troy/tab",
		"https://travellermap.com/data/Gimli/tab",
		"https://travellermap.com/data/Niflheim/tab",
		"https://travellermap.com/data/Lethe/tab",
		"https://travellermap.com/data/Hubur/tab",
		"https://travellermap.com/data/Celadon/tab",
		"https://travellermap.com/data/Naraka/tab",
		"https://travellermap.com/data/Malvam/tab",
		"https://travellermap.com/data/Elivagar/tab",
		"https://travellermap.com/data/Riftsedge%20Coreward/tab",
		"https://travellermap.com/data/Riftsedge%20Bridge/tab",
		"https://travellermap.com/data/FSD%20Transition%20Spinward/tab",
		"https://travellermap.com/data/LSD%20Transition%20One/tab",
		"https://travellermap.com/data/ZSD%20Transition%20Two/tab",
		"https://travellermap.com/data/LSD%20Transition%20Three/tab",
		"https://travellermap.com/data/LSD%20Transition%20Four/tab",
		"https://travellermap.com/data/LSD%20Transition%20Trailing/tab",
		"https://travellermap.com/data/Big%20Island/tab",
		"https://travellermap.com/data/Riftsedge%20Central/tab",
		"https://travellermap.com/data/Voidshore%20One/tab",
		"https://travellermap.com/data/Voidshore%20Two/tab",
		"https://travellermap.com/data/Voidshore%20Three/tab",
		"https://travellermap.com/data/LSD%20Transition%20Rimward/tab",
		"https://travellermap.com/data/Voidshore%20Four/tab",
		"https://travellermap.com/data/Voidshore%20Five/tab",
		"https://travellermap.com/data/Voidshore%20Six/tab",
		"https://travellermap.com/data/Voidshore%20Trailing/tab",
		"https://travellermap.com/data/Riftsedge%20Rimward/tab",
		"https://travellermap.com/data/VS-X/tab",
		"https://travellermap.com/data/Voidshore%20Seven/tab",
		"https://travellermap.com/data/Orion's%20Spit%20Coreward/tab",
		"https://travellermap.com/data/Voidshore%20Eight/tab",
		"https://travellermap.com/data/Voidshore%20Nine/tab",
		"https://travellermap.com/data/Voidshore%20Ten/tab",
		"https://travellermap.com/data/Voidshore%20Eleven/tab",
		"https://travellermap.com/data/Voidshore%20Rimward/tab",
		"https://travellermap.com/data/Orion's%20Spit%20Rimward/tab",
	}
}
