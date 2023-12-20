package creature

import (
	"math/rand"
)

var animals = [10]string{"bull",
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

var radix = [...]string{"at", "xer", "apros", "leas"}
var prefix = [...]string{"ptero", "xae", "loax", "awe"}
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
	generatedCreature.Head = animals[rand.Intn(10)]
	generatedCreature.Body = animals[rand.Intn(10)]
	generatedCreature.Legs = animals[rand.Intn(10)]
	generatedCreature.Leg_num = leg_num[rand.Intn(2)]
	generatedCreature.Habitat = habitats[rand.Intn(6)]
	generatedCreature.Nature = natures[rand.Intn(24)]
	generatedCreature.Class = class[rand.Intn(4)]
	return generatedCreature
}
