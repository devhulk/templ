package services

var Options = map[string]string{
	"C":          "c.png",
	"Golang":     "golang.png",
	"Python":     "python.png",
	"Java":       "java.png",
	"C++":        "cpp.png",
	"C#":         "csharp.png",
	"Haskell":    "haskell.png",
	"Zig":        "zig.png",
	"Rust":       "rust.png",
	"Javascript": "js.png",
}

type TierItem struct {
	Location string
}

type Tiers struct {
	STier      []TierItem
	VeryUseful []TierItem
	Average    []TierItem
	Bupkis     []TierItem
	DogWater   []TierItem
	Options    map[string]string
}

var MyTier = &Tiers{
	STier:      []TierItem{},
	VeryUseful: []TierItem{},
	Average:    []TierItem{},
	Bupkis:     []TierItem{},
	DogWater:   []TierItem{},
	Options:    Options,
}
