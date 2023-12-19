package creature

import (
	"math/rand"
)

var head = [10]string{"bull",
	"tiger",
	"bat",
	"lion",
	"snake",
	"wolf",
	"spider",
	"eagle",
	"sheep",
	"fish",
}

var body = [10]string{
	"bull",
	"tiger",
	"bat",
	"lion",
	"snake",
	"wolf",
	"spider",
	"eagle",
	"sheep",
	"fish",
}

var legs = [10]string{
	"bull",
	"tiger",
	"bat",
	"lion",
	"snake",
	"wolf",
	"spider",
	"eagle",
	"sheep",
	"fish",
}

var leg_num = [...]int{2, 4, 6, 8, 10}

var natures = [...]string{
	"hardy",
	"lonely",
	"adamant",
	"naughty",
	"brave",
	"bold",
	"docile",
	"impish",
	"lax",
	"relaxed",
	"mild",
	"bashful",
	"rash",
	"quiet",
	"calm",
	"gentle",
	"careful",
	"quirky",
	"sassy",
	"timid",
	"hasty",
	"jolly",
	"naive",
	"serious",
}
var habitats = [...]string{
	"tundra",
	"evergreen forests",
	"seasonal forests",
	"grasslands",
	"deserts",
	"tropical Rainforests",
}

var class = [...]string{
	"predator",
	"prey",
	"scavenger",
	"parasite",
}

type Creature struct {
	Name    string
	Head    string
	Body    string
	Leg_num int
	Legs    string
	Nature  string
	Habitat string
	Class   string
}

func NewCreature() Creature {
	var generatedCreature Creature
	generatedCreature.Name = "PLACEHOLDER"
	generatedCreature.Head = head[rand.Intn(10)]
	generatedCreature.Body = body[rand.Intn(10)]
	generatedCreature.Legs = legs[rand.Intn(10)]
	generatedCreature.Leg_num = leg_num[rand.Intn(5)]
	generatedCreature.Habitat = habitats[rand.Intn(6)]
	generatedCreature.Nature = natures[rand.Intn(24)]
	generatedCreature.Class = class[rand.Intn(4)]
	return generatedCreature
}
