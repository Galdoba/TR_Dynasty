package background

import "fmt"

type background struct {
	worldName string
	worldTC   []string
}

func NewBackground(name string, codes []string) *background {
	bg := background{}
	bg.worldName = name
	bg.worldTC = codes
	return &bg
}

func (bg *background) BasicSkills() []string {
	list := []string{}
	for _, code := range bg.worldTC {
		switch code {
		case "Ag", "Po":
			list = append(list, "  ")
		case "As", "Va":
			list = append(list, "  ")
		case "De", "Lt":
			list = append(list, "  ")
		case "Fl", "Wa":
			list = append(list, "  ")
		case "Ht":
			list = append(list, "  ")
		case "Hi":
			list = append(list, "  ")
		case "Ic":
			list = append(list, "  ")
		case "In":
			list = append(list, "  ")
		case "Ri":
			list = append(list, "  ")
		}
		fmt.Print(code + "\r   ")
	}
	return []string{}
}
