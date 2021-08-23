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
	for _, code := range bg.worldTC {
		fmt.Print(code + "\r   ")
	}
	return []string{}
}
