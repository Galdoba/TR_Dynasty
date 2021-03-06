package name

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

func firstName() string {
	names := []string{
		"Aamir",
		"Ayub",
		"Binyamin",
		"Efraim",
		"Ibrahim",
		"Ilyas",
		"Ismail",
		"Jibril",
		"Jumanah",
		"Kazi",
		"Lut",
		"Matta",
		"Mohammed",
		"Mubarak",
		"Mustafa",
		"Nazir",
		"Rahim",
		"Reza",
		"Sharif",
		"Taimur",
		"Usman",
		"Yakub",
		"Yusuf",
		"Zakariya",
		"Zubair",
		"Aisha",
		"Alimah",
		"Badia",
		"Bisharah",
		"Chanda",
		"Daliya",
		"Fatimah",
		"Ghania",
		"Halah",
		"Kaylah",
		"Khayrah",
		"Layla",
		"Mina",
		"Munisa",
		"Mysha",
		"Naimah",
		"Nissa",
		"Nura",
		"Parveen",
		"Rana",
		"Shalha",
		"Suhira",
		"Tahirah",
		"Yasmin",
		"Zulehka",
		"Adan",
		"Ahsa",
		"Andalus",
		"Asmara",
		"Asqlan",
		"Baqubah",
		"Basit",
		"Baysan",
		"Baytlahm",
		"Bursaid",
		"Dahilah",
		"Darasalam",
		"Dawhah",
		"Ganin",
		"Gebal",
		"Gibuti",
		"Giddah",
		"Harmah",
		"Hartum",
		"Hibah",
		"Hims",
		"Hubar",
		"Karbala",
		"Kut",
		"Lacant",
		"Magrit",
		"Masqat",
		"Misr",
		"Muruni",
		"Qabis",
		"Qina",
		"Rabat",
		"Ramlah",
		"Riyadh",
		"Sabtah",
		"Salalah",
		"Sana",
		"Sinqit",
		"Suqutrah",
		"Sur",
		"Tabuk",
		"Tangah",
		"Tarifah",
		"Tarrakunah",
		"Tisit",
		"Uman",
		"Urdunn",
		"Wasqah",
		"Yaburah",
		"Yaman",
		"Aiguo",
		"Bohai",
		"Chao",
		"Dai",
		"Dawei",
		"Duyi",
		"Fa",
		"Fu",
		"Gui",
		"Hong",
		"Jianyu",
		"Kang",
		"Li",
		"Niu",
		"Peng",
		"Quan",
		"Ru",
		"Shen",
		"Shi",
		"Song",
		"Tao",
		"Xue",
		"Yi",
		"Yuan",
		"Zian",
		"Biyu",
		"Changying",
		"Daiyu",
		"Huidai",
		"Huiliang",
		"Jia",
		"Jingfei",
		"Lan",
		"Liling",
		"Liu",
		"Meili",
		"Niu",
		"Peizhi",
		"Qiao",
		"Qing",
		"Ruolan",
		"Shu",
		"Suyin",
		"Ting",
		"Xia",
		"Xiaowen",
		"Xiulan",
		"Ya",
		"Ying",
		"Zhilan",
		"Andong",
		"Anqing",
		"Anshan",
		"Chaoyang",
		"Chaozhou",
		"Chifeng",
		"Dalian",
		"Dunhuang",
		"Fengjia",
		"Fengtian",
		"Fuliang",
		"Fushun",
		"Gansu",
		"Ganzhou",
		"Guizhou",
		"Hotan",
		"Hunan",
		"Jinan",
		"Jingdezhen",
		"Jinxi",
		"Jinzhou",
		"Kunming",
		"Liaoning",
		"Linyi",
		"Lushun",
		"Luzhou",
		"Ningxia",
		"Pingxiang",
		"Pizhou",
		"Qidong",
		"Qingdao",
		"Qinghai",
		"Rehe",
		"Shanxi",
		"Taiyuan",
		"Tengzhou",
		"Urumqi",
		"Weifang",
		"Wugang",
		"Wuxi",
		"Xiamen",
		"Xian",
		"Xikang",
		"Xining",
		"Xinjiang",
		"Yidu",
		"Yingkou",
		"Yuxi",
		"Zigong",
		"Zoige",
		"Adam",
		"Albert",
		"Alfred",
		"Allan",
		"Archibald",
		"Arthur",
		"Basil",
		"Charles",
		"Colin",
		"Donald",
		"Douglas",
		"Edgar",
		"Edmund",
		"Edward",
		"George",
		"Harold",
		"Henry",
		"Ian",
		"James",
		"John",
		"Lewis",
		"Oliver",
		"Philip",
		"Richard",
		"William",
		"Abigail",
		"Anne",
		"Beatrice",
		"Blanche",
		"Catherine",
		"Charlotte",
		"Claire",
		"Eleanor",
		"Elizabeth",
		"Emily",
		"Emma",
		"Georgia",
		"Harriet",
		"Joan",
		"Judy",
		"Julia",
		"Lucy",
		"Lydia",
		"Margaret",
		"Mary",
		"Molly",
		"Nora",
		"Rosie",
		"Sarah",
		"Victoria",
		"Aldington",
		"Appleton",
		"Ashdon",
		"Berwick",
		"Bramford",
		"Brimstage",
		"Carden",
		"Churchill",
		"Clifton",
		"Colby",
		"Copford",
		"Cromer",
		"Davenham",
		"Dersingham",
		"Doverdale",
		"Elsted",
		"Ferring",
		"Gissing",
		"Heydon",
		"Holt",
		"Hunston",
		"Hutton",
		"Inkberrow",
		"Inworth",
		"Isfield",
		"Kedington",
		"Latchford",
		"Leigh",
		"Leighton",
		"Maresfield",
		"Markshall",
		"Netherpool",
		"Newton",
		"Oxton",
		"Preston",
		"Ridley",
		"Rochford",
		"Seaford",
		"Selsey",
		"Stanton",
		"Stockham",
		"Stoke",
		"Sutton",
		"Thakeham",
		"Thetford",
		"Thorndon",
		"Ulting",
		"Upton",
		"Westhorpe",
		"Worcester",
		"Alexander",
		"Alexius",
		"Anastasius",
		"Christodoulos",
		"Christos",
		"Damian",
		"Dimitris",
		"Dysmas",
		"Elias",
		"Giorgos",
		"Ioannis",
		"Konstantinos",
		"Lambros",
		"Leonidas",
		"Marcos",
		"Miltiades",
		"Nestor",
		"Nikos",
		"Orestes",
		"Petros",
		"Simon",
		"Stavros",
		"Theodore",
		"Vassilios",
		"Yannis",
		"Alexandra",
		"Amalia",
		"Callisto",
		"Charis",
		"Chloe",
		"Dorothea",
		"Elena",
		"Eudoxia",
		"Giada",
		"Helena",
		"Ioanna",
		"Lydia",
		"Melania",
		"Melissa",
		"Nika",
		"Nikolina",
		"Olympias",
		"Philippa",
		"Phoebe",
		"Sophia",
		"Theodora",
		"Valentina",
		"Valeria",
		"Yianna",
		"Zoe",
		"Adramyttion",
		"Ainos",
		"Alikarnassos",
		"Avydos",
		"Dakia",
		"Dardanos",
		"Dekapoli",
		"Dodoni",
		"Efesos",
		"Efstratios",
		"Elefsina",
		"Ellada",
		"Epidavros",
		"Erymanthos",
		"Evripos",
		"Gavdos",
		"Gytheio",
		"Ikaria",
		"Ilios",
		"Illyria",
		"Iraia",
		"Irakleio",
		"Isminos",
		"Ithaki",
		"Kadmeia",
		"Kallisto",
		"Katerini",
		"Kithairon",
		"Kydonia",
		"Lakonia",
		"Leros",
		"Lesvos",
		"Limnos",
		"Lykia",
		"Megara",
		"Messene",
		"Milos",
		"Nikaia",
		"Orontis",
		"Parnasos",
		"Petro",
		"Samos",
		"Syros",
		"Thapsos",
		"Thessalia",
		"Thira",
		"Thiva",
		"Varvara",
		"Voiotia",
		"Vyvlos",
		"Amrit",
		"Ashok",
		"Chand",
		"Dinesh",
		"Gobind",
		"Harinder",
		"Jagdish",
		"Johar",
		"Kurien",
		"Lakshman",
		"Madhav",
		"Mahinder",
		"Mohal",
		"Narinder",
		"Nikhil",
		"Omrao",
		"Prasad",
		"Pratap",
		"Ranjit",
		"Sanjay",
		"Shankar",
		"Thakur",
		"Vijay",
		"Vipul",
		"Yash",
		"Amala",
		"Asha",
		"Chandra",
		"Devika",
		"Esha",
		"Gita",
		"Indira",
		"Indrani",
		"Jaya",
		"Jayanti",
		"Kiri",
		"Lalita",
		"Malati",
		"Mira",
		"Mohana",
		"Neela",
		"Nita",
		"Rajani",
		"Sarala",
		"Sarika",
		"Sheela",
		"Sunita",
		"Trishna",
		"Usha",
		"Vasanta",
		"Ahmedabad",
		"Alipurduar",
		"Alubari",
		"Anjanadri",
		"Ankleshwar",
		"Balarika",
		"Bhanuja",
		"Bhilwada",
		"Brahmaghosa",
		"Bulandshahar",
		"Candrama",
		"Chalisgaon",
		"Chandragiri",
		"Charbagh",
		"Chayanka",
		"Chittorgarh",
		"Dayabasti",
		"Dikpala",
		"Ekanga",
		"Gandhidham",
		"Gollaprolu",
		"Grahisa",
		"Guwahati",
		"Haridasva",
		"Indraprastha",
		"Jaisalmer",
		"Jharonda",
		"Kadambur",
		"Kalasipalyam",
		"Karnataka",
		"Kutchuhery",
		"Lalgola",
		"Mainaguri",
		"Nainital",
		"Nandidurg",
		"Narayanadri",
		"Panipat",
		"Panjagutta",
		"Pathankot",
		"Pathardih",
		"Porbandar",
		"Rajasthan",
		"Renigunta",
		"Sewagram",
		"Shakurbasti",
		"Siliguri",
		"Sonepat",
		"Teliwara",
		"Tinpahar",
		"Villivakkam",
		"Akira",
		"Daisuke",
		"Fukashi",
		"Goro",
		"Hiro",
		"Hiroya",
		"Hotaka",
		"Katsu",
		"Katsuto",
		"Keishuu",
		"Kyuuto",
		"Mikiya",
		"Mitsunobu",
		"Mitsuru",
		"Naruhiko",
		"Nobu",
		"Shigeo",
		"Shigeto",
		"Shou",
		"Shuji",
		"Takaharu",
		"Teruaki",
		"Tetsushi",
		"Tsukasa",
		"Yasuharu",
		"Aemi",
		"Airi",
		"Ako",
		"Ayu",
		"Chikaze",
		"Eriko",
		"Hina",
		"Kaori",
		"Keiko",
		"Kyouka",
		"Mayumi",
		"Miho",
		"Namiko",
		"Natsu",
		"Nobuko",
		"Rei",
		"Ririsa",
		"Sakimi",
		"Shihoko",
		"Shika",
		"Tsukiko",
		"Tsuzune",
		"Yoriko",
		"Yorimi",
		"Yoshiko",
		"Agrippa",
		"Appius",
		"Aulus",
		"Caeso",
		"Decimus",
		"Faustus",
		"Gaius",
		"Gnaeus",
		"Hostus",
		"Lucius",
		"Mamercus",
		"Manius",
		"Marcus",
		"Mettius",
		"Nonus",
		"Numerius",
		"Opiter",
		"Paulus",
		"Proculus",
		"Publius",
		"Quintus",
		"Servius",
		"Tiberius",
		"Titus",
		"Volescus",
		"Appia",
		"Aula",
		"Caesula",
		"Decima",
		"Fausta",
		"Gaia",
		"Gnaea",
		"Hosta",
		"Lucia",
		"Maio",
		"Marcia",
		"Maxima",
		"Mettia",
		"Nona",
		"Numeria",
		"Octavia",
		"Postuma",
		"Prima",
		"Procula",
		"Septima",
		"Servia",
		"Tertia",
		"Tiberia",
		"Titia",
		"Vibia",
		"Adesegun",
		"Akintola",
		"Amabere",
		"Arikawe",
		"Asagwara",
		"Chidubem",
		"Chinedu",
		"Chiwetei",
		"Damilola",
		"Esangbedo",
		"Ezenwoye",
		"Folarin",
		"Genechi",
		"Idowu",
		"Kelechi",
		"Ketanndu",
		"Melubari",
		"Nkanta",
		"Obafemi",
		"Olatunde",
		"Olumide",
		"Tombari",
		"Udofia",
		"Uyoata",
		"Uzochi",
		"Abike",
		"Adesuwa",
		"Adunola",
		"Anguli",
		"Arewa",
		"Asari",
		"Bisola",
		"Chioma",
		"Eduwa",
		"Emilohi",
		"Fehintola",
		"Folasade",
		"Mahparah",
		"Minika",
		"Nkolika",
		"Nkoyo",
		"Nuanae",
		"Obioma",
		"Olafemi",
		"Shanumi",
		"Sominabo",
		"Suliat",
		"Tariere",
		"Temedire",
		"Yemisi",
		"Aleksandr",
		"Andrei",
		"Arkady",
		"Boris",
		"Dmitri",
		"Dominik",
		"Grigory",
		"Igor",
		"Ilya",
		"Ivan",
		"Kiril",
		"Konstantin",
		"Leonid",
		"Nikolai",
		"Oleg",
		"Pavel",
		"Petr",
		"Sergei",
		"Stepan",
		"Valentin",
		"Vasily",
		"Viktor",
		"Yakov",
		"Yegor",
		"Yuri",
		"Aleksandra",
		"Anastasia",
		"Anja",
		"Catarina",
		"Devora",
		"Dima",
		"Ekaterina",
		"Eva",
		"Irina",
		"Karolina",
		"Katlina",
		"Kira",
		"Ludmilla",
		"Mara",
		"Nadezdha",
		"Nastassia",
		"Natalya",
		"Oksana",
		"Olena",
		"Olga",
		"Sofia",
		"Svetlana",
		"Tatyana",
		"Vilma",
		"Yelena",
		"Alejandro",
		"Alonso",
		"Amelio",
		"Armando",
		"Bernardo",
		"Carlos",
		"Cesar",
		"Diego",
		"Emilio",
		"Estevan",
		"Felipe",
		"Francisco",
		"Guillermo",
		"Javier",
		"Jose",
		"Juan",
		"Julio",
		"Luis",
		"Pedro",
		"Raul",
		"Ricardo",
		"Salvador",
		"Santiago",
		"Valeriano",
		"Vicente",
		"Adalina",
		"Aleta",
		"Ana",
		"Ascencion",
		"Beatriz",
		"Carmela",
		"Celia",
		"Dolores",
		"Elena",
		"Emelina",
		"Felipa",
		"Inez",
		"Isabel",
		"Jacinta",
		"Lucia",
		"Lupe",
		"Maria",
		"Marta",
		"Nina",
		"Paloma",
		"Rafaela",
		"Soledad",
		"Teresa",
		"Valencia",
		"Zenaida",
		//Drinaxian Companion
		"Ghathtagu",
		"Ansiesta",
		"Sam",
		"Thanihryihmbakepe",
		"Carneliana",
		"Carse",
		"Nils",
		"Bryn",
		"Frei",
		"Mannie",
		"Handow",
		"Lars",
		"Narmure",
		"Ehrae",
		"Kurgakikash",
		"Wilheim",
		"Joharn",
		"Arym",
		"Mitchell",
		"Jenaime",
		"Poia",
		"Harnon",
		"Raif",
		"Erech",
		"Vrine",
		"Aix",
		"Felix",
	}
	die := strconv.Itoa(len(names))
	return names[dice.Roll("1d"+die).Sum()-1]
}

func familyName() string {
	names := []string{
		//Drinaxian Companion
		"Carmaichel",
		"Leto",
		"Karmalli",
		"Zumczizcy",
		"Zentoulli",
		"Pallix",
		"Pellique",
		"Muuru",
		"Ygrant",
		"Argane",
		"Essden",
		"Azse",
		"Iabl",
		"Dedherileh",
		"Hilfssen",
		"Vaasirn",
		"Alfney",
		"Sawnenson",
		"Hadsen",
		"Tsaiboud",
		"Venquist",
		"Idais",
		"Recheille",
		"Awolr",
		"Mila",
		"Amuiinzier",
		"Ishinko",
		//SWN
		"Arellano",
		"Arispana",
		"Borrego",
		"Carderas",
		"Carranzo",
		"Cordova",
		"Enciso",
		"Espejo",
		"Gavilan",
		"Guerra",
		"Guillen",
		"Huertas",
		"Illan",
		"Jurado",
		"Moretta",
		"Motolinia",
		"Pancorbo",
		"Paredes",
		"Quesada",
		"Roma",
		"Rubiera",
		"Santoro",
		"Torrillas",
		"Vera",
		"Vivero",
		"Aguascebas",
		"Alcazar",
		"Barranquete",
		"Bravatas",
		"Cabezudos",
		"Calderon",
		"Cantera",
		"Castillo",
		"Delgadas",
		"Donablanca",
		"Encinetas",
		"Estrella",
		"Faustino",
		"Fuentebravia",
		"Gafarillos",
		"Gironda",
		"Higueros",
		"Huelago",
		"Humilladero",
		"Illora",
		"Isabela",
		"Izbor",
		"Jandilla",
		"Jinetes",
		"Limones",
		"Loreto",
		"Lujar",
		"Marbela",
		"Matagorda",
		"Nacimiento",
		"Niguelas",
		"Ogijares",
		"Ortegicar",
		"Pampanico",
		"Pelado",
		"Quesada",
		"Quintera",
		"Riguelo",
		"Ruescas",
		"Salteras",
		"Santopitar",
		"Taberno",
		"Torres",
		"Umbrete",
		"Valdecazorla",
		"Velez",
		"Vistahermosa",
		"Yeguas",
		"Zahora",
		"Zumeta",
		"Abdel",
		"Awad",
		"Dahhak",
		"Essa",
		"Hanna",
		"Harbi",
		"Hassan",
		"Isa",
		"Kasim",
		"Katib",
		"Khalil",
		"Malik",
		"Mansoor",
		"Mazin",
		"Musa",
		"Najeeb",
		"Namari",
		"Naser",
		"Rahman",
		"Rasheed",
		"Saleh",
		"Salim",
		"Shadi",
		"Sulaiman",
		"Tabari",
		"Bai",
		"Cao",
		"Chen",
		"Cui",
		"Ding",
		"Du",
		"Fang",
		"Fu",
		"Guo",
		"Han",
		"Hao",
		"Huang",
		"Lei",
		"Li",
		"Liang",
		"Liu",
		"Long",
		"Song",
		"Tan",
		"Tang",
		"Wang",
		"Wu",
		"Xing",
		"Yang",
		"Zhang",
		"Barker",
		"Brown",
		"Butler",
		"Carter",
		"Chapman",
		"Collins",
		"Cook",
		"Davies",
		"Gray",
		"Green",
		"Harris",
		"Jackson",
		"Jones",
		"Lloyd",
		"Miller",
		"Roberts",
		"Smith",
		"Taylor",
		"Thomas",
		"Turner",
		"Watson",
		"White",
		"Williams",
		"Wood",
		"Young",
		"Andrea",
		"Andreas",
		"Argyros",
		"Dimitriou",
		"Floros",
		"Gavras",
		"Ioannidis",
		"Katsaros",
		"Kyrkos",
		"Leventis",
		"Makris",
		"Metaxas",
		"Nikolaidis",
		"Pallis",
		"Pappas",
		"Petrou",
		"Raptis",
		"Simonides",
		"Spiros",
		"Stavros",
		"Stephanidis",
		"Stratigos",
		"Terzis",
		"Theodorou",
		"Vasiliadis",
		"Yannakakis",
		"Achari",
		"Banerjee",
		"Bhatnagar",
		"Bose",
		"Chauhan",
		"Chopra",
		"Das",
		"Dutta",
		"Gupta",
		"Johar",
		"Kapoor",
		"Mahajan",
		"Malhotra",
		"Mehra",
		"Nehru",
		"Patil",
		"Rao",
		"Saxena",
		"Shah",
		"Sharma",
		"Singh",
		"Trivedi",
		"Venkatesan",
		"Verma",
		"Yadav",
		"Abe",
		"Arakaki",
		"Endo",
		"Fujiwara",
		"Goto",
		"Ito",
		"Kikuchi",
		"Kinjo",
		"Kobayashi",
		"Koga",
		"Komatsu",
		"Maeda",
		"Nakamura",
		"Narita",
		"Ochi",
		"Oshiro",
		"Saito",
		"Sakamoto",
		"Sato",
		"Suzuki",
		"Takahashi",
		"Tanaka",
		"Watanabe",
		"Yamamoto",
		"Yamasaki",
		"Bando",
		"Chikuma",
		"Chikusei",
		"Chino",
		"Hitachi",
		"Hitachinaka",
		"Hitachiomiya",
		"Hitachiota",
		"Iida",
		"Iiyama",
		"Ina",
		"Inashiki",
		"Ishioka",
		"Itako",
		"Kamisu",
		"Kasama",
		"Kashima",
		"Kasumigaura",
		"Kitaibaraki",
		"Kiyose",
		"Koga",
		"Komagane",
		"Komoro",
		"Matsumoto",
		"Mito",
		"Mitsukaido",
		"Moriya",
		"Nagano",
		"Naka",
		"Nakano",
		"Ogi",
		"Okaya",
		"Omachi",
		"Ryugasaki",
		"Saku",
		"Settsu",
		"Shimotsuma",
		"Shiojiri",
		"Suwa",
		"Suzaka",
		"Takahagi",
		"Takeo",
		"Tomi",
		"Toride",
		"Tsuchiura",
		"Tsukuba",
		"Ueda",
		"Ushiku",
		"Yoshikawa",
		"Yuki",
		"Antius",
		"Aurius",
		"Barbatius",
		"Calidius",
		"Cornelius",
		"Decius",
		"Fabius",
		"Flavius",
		"Galerius",
		"Horatius",
		"Julius",
		"Juventius",
		"Licinius",
		"Marius",
		"Minicius",
		"Nerius",
		"Octavius",
		"Pompeius",
		"Quinctius",
		"Rutilius",
		"Sextius",
		"Titius",
		"Ulpius",
		"Valerius",
		"Vitellius",
		"Adegboye",
		"Adeniyi",
		"Adeyeku",
		"Adunola",
		"Agbaje",
		"Akpan",
		"Akpehi",
		"Aliki",
		"Asuni",
		"Babangida",
		"Ekim",
		"Ezeiruaku",
		"Fabiola",
		"Fasola",
		"Nwokolo",
		"Nzeocha",
		"Ojo",
		"Okonkwo",
		"Okoye",
		"Olaniyan",
		"Olawale",
		"Olumese",
		"Onajobi",
		"Soyinka",
		"Yamusa",
		"Abadan",
		"Ador",
		"Agatu",
		"Akamkpa",
		"Akpabuyo",
		"Ala",
		"Askira",
		"Bakassi",
		"Bama",
		"Bayo",
		"Bekwara",
		"Biase",
		"Boki",
		"Buruku",
		"Calabar",
		"Chibok",
		"Damboa",
		"Dikwa",
		"Etung",
		"Gboko",
		"Gubio",
		"Guzamala",
		"Gwoza",
		"Hawul",
		"Ikom",
		"Jere",
		"Kalabalge",
		"Katsina",
		"Knoduga",
		"Konshishatse",
		"Kukawa",
		"Kwande",
		"Kwayakusar",
		"Logo",
		"Mafa",
		"Makurdi",
		"Nganzai",
		"Obanliku",
		"Obi",
		"Obubra",
		"Obudu",
		"Odukpani",
		"Ogbadibo",
		"Ohimini",
		"Okpokwu",
		"Otukpo",
		"Shani",
		"Ugep",
		"Vandeikya",
		"Yala",
		"Abelev",
		"Bobrikov",
		"Chemerkin",
		"Gogunov",
		"Gurov",
		"Iltchenko",
		"Kavelin",
		"Komarov",
		"Korovin",
		"Kurnikov",
		"Lebedev",
		"Litvak",
		"Mekhdiev",
		"Muraviov",
		"Nikitin",
		"Ortov",
		"Peshkov",
		"Romasko",
		"Shvedov",
		"Sikorski",
		"Stolypin",
		"Turov",
		"Volokh",
		"Zaitsev",
		"Zhukov",
	}
	die := strconv.Itoa(len(names))
	return names[dice.Roll("1d"+die).Sum()-1]
}

func RandomNew() string {
	return firstName() + " " + familyName()
}
