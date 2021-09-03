package traveller

import "github.com/Galdoba/TR_Dynasty/T5/assets"

func (trv *TravellerT5) Characteristic(code int) *assets.Characteristic {
	return trv.characteristic[code]
}

func (trv *TravellerT5) Skill(code int) *assets.Skill {
	return trv.skills[code]
}
