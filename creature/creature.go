package creature

import (
	"math/rand"
)

var animals = [...]string{"ox",
	"tiger",
	"bat",
	"lion",
	"lizard",
	"wolf",
	"spider",
	"eagle",
	"sheep",
	"fish",
	"man",
	"sloth",
}

var leg_num = [...]int{2, 4}

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

var prefix = [...]string{"ptero", "xae", "loax", "awe", "enon"}
var radix = [...]string{"at", "xer", "apros", "leas", "reas"}
var suffix = [...]string{"don", "mos", "lia", "as", "saur"}

var habitats = [...]string{
	"tundra",
	"evergreen forests",
	"seasonal forests",
	"grasslands",
	"deserts",
	"tropical rainforests",
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

func NewName() string {
	return prefix[rand.Intn(len(prefix))] + radix[rand.Intn(len(radix))] + suffix[rand.Intn(len(suffix))]
}

func NewCreature() Creature {
	var generatedCreature Creature
	generatedCreature.Name = NewName()
	generatedCreature.Head = animals[rand.Intn(len(animals))]
	generatedCreature.Body = animals[rand.Intn(len(animals))]
	generatedCreature.Legs = animals[rand.Intn(len(animals))]
	generatedCreature.Leg_num = leg_num[rand.Intn(2)]
	generatedCreature.Habitat = habitats[rand.Intn(6)]
	generatedCreature.Nature = natures[rand.Intn(24)]
	generatedCreature.Class = class[rand.Intn(4)]
	return generatedCreature
}
