package skills

//Skill - are the primary means by which characters do things in Traveller.
//	Each character has a variety of skills, and the higher a skill rating,
//	the more expert the character is with that skill. With training, any
//	character can eventually become proficient at any skill
type Skill struct {
	alias      string
	rating     int
	knowledges []string
}
