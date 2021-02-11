package threat

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/dice"
)

type Risk struct {
	probability, severity, imminence int
}

func NewRisk() Risk {
	r := Risk{}
	r.probability = dice.Flux()
	r.severity = dice.Flux()
	r.imminence = dice.Flux()
	return r
}

func (r Risk) Express() string {
	descr := "The scale of this Risk is " + strconv.Itoa(r.probability+r.severity+r.severity)
	descr += ". It can be described as " + probibilityStr(r.probability) + " and " + severenityStr(r.severity) + ", with impact will be felt in " + imminenceStr(r.imminence) + "."
	return descr
}

func probibilityStr(i int) string {
	if i < -5 {
		i = -6
	}
	if i > 5 {
		i = 6
	}
	list := []string{"Impossible", "Highly Improbable", "Improbable", "Highly Unlikely", "Unlikely", "Not Likely", "Either Way", "Possible", "Likely", "Probable", "Very Probable", "Almost Certain", "Certain"}
	return list[i+6]
}

func severenityStr(i int) string {
	if i < -5 {
		i = -6
	}
	if i > 5 {
		i = 6
	}
	list := []string{"None", "Trivial", "Negligible", "Very Minor", "Minor", "Mild", "Temporary", "Strong", "Major", "Severe", "Very Severe", "Devastating", "Total "}
	return list[i+6]
}

func imminenceStr(i int) string {
	if i < -5 {
		i = -6
	}
	if i > 5 {
		i = 6
	}
	list := []string{"Far Future", "Centuries", "a Lifetime", "Generation", "Decades", "Years", "Months", "Weeks", "Days", "Hours", "Minutes", "Seconds", "Now"}
	return list[i+6]
}
