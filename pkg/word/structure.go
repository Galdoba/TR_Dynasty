package word

import (
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	V   = 1
	CV  = 2
	VC  = 3
	CVC = 4
	//////////////
	Language_Vilani  = 1
	Language_Aslan   = 2
	Language_Vargr   = 3
	Language_Droyne  = 4
	Language_Zhodani = 5
	/////////////
	basicSylable = "Basic"
	alterSylable = "Alternate"
)

type Word struct {
	syl      []syllable
	dp       *dice.Dicepool
	language int
	meaning  string
	spelling string
}

func New(languageCode, seed int) (Word, error) {
	w := Word{}
	w.dp = dice.New().SetSeed(seed)
	switch languageCode {
	case 1, 2, 3, 4, 5:
		w.language = languageCode
	default:
		w.language = w.dp.RollNext("1d5").Sum()
	}
	lenth := 2
	switch w.dp.RollNext("2d6").Sum() {
	case 2, 12:
		lenth = 5
	case 3, 11:
		lenth = 4
	case 4, 5, 9, 10:
		lenth = 3
	case 6, 7, 8:
		lenth = 2
	}
	for i := 0; i < lenth; i++ {
		switch i {
		case 0:
			w.syl = append(w.syl, syllable{basicSyllable(w.language, w.dp), i, "[undefined]"})
		default:
			w.syl = append(w.syl, syllable{alterSyllable(w.language, w.dp), i, "[undefined]"})
		}
	}
	for i := range w.syl {
		w.syl[i].generateSound(w.dp, languageCode)
		w.spelling += w.syl[i].sound
	}
	return w, nil
}

func (w *Word) Shout() string {
	return w.spelling
}

type syllable struct {
	sType    int
	position int
	sound    string
}

func (s *syllable) generateSound(dp *dice.Dicepool, languageCode int) {
	syl := ""
	switch s.sType {
	case V:
		syl = callVowel(dp, languageCode)
	case CV:
		syl = callIConsonant(dp, languageCode)
		syl += callVowel(dp, languageCode)
	case VC:
		syl = callVowel(dp, languageCode)
		syl += callFConsonant(dp, languageCode)
	case CVC:
		syl = callIConsonant(dp, languageCode)
		syl += callVowel(dp, languageCode)
		syl += callFConsonant(dp, languageCode)
	}
	s.sound = syl
}

func callVowel(dp *dice.Dicepool, languageCode int) string {
	r1 := dp.RollNext("1d6").Sum()
	r2 := dp.RollNext("1d6").Sum()
	r3 := dp.RollNext("1d6").Sum()
	switch languageCode {
	case Language_Vilani:
		return vilaniVowel(r1, r2, r3)
	}
	return "?"
}

func callIConsonant(dp *dice.Dicepool, languageCode int) string {
	r1 := dp.RollNext("1d6").Sum()
	r2 := dp.RollNext("1d6").Sum()
	r3 := dp.RollNext("1d6").Sum()
	switch languageCode {
	case Language_Vilani:
		return vilaniIConsonant(r1, r2, r3)
	}
	return "?"
}

func callFConsonant(dp *dice.Dicepool, languageCode int) string {
	r1 := dp.RollNext("1d6").Sum()
	r2 := dp.RollNext("1d6").Sum()
	r3 := dp.RollNext("1d6").Sum()
	switch languageCode {
	case Language_Vilani:
		return vilaniFConsonant(r1, r2, r3)
	}
	return "?"
}

func basicSyllable(languageCode int, dp *dice.Dicepool) int {
	rd := dp.RollNext("1d6").Sum()
	wd := dp.RollNext("1d6").Sum()
	sylMap := make(map[int][]int)
	switch languageCode {
	case Language_Vilani:
		sylMap = basicVilaniSyllables()
	}
	return sylMap[rd][wd-1]
}

func alterSyllable(languageCode int, dp *dice.Dicepool) int {
	rd := dp.RollNext("1d6").Sum()
	wd := dp.RollNext("1d6").Sum()
	sylMap := make(map[int][]int)
	switch languageCode {
	case Language_Vilani:
		sylMap = alterVilaniSyllables()
	}
	return sylMap[rd][wd-1]
}
