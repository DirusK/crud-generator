package app

import (
	"crud-generator-gui/internal/generators"
	vipcoin "crud-generator-gui/internal/generators/vipcoin/generator"
	"crud-generator-gui/internal/models"
)

type (
	GeneratorType string

	GeneratorInit func(entity models.Entity, settings models.Settings) generators.Generator
)

// Block which stores all generator types.
const (
	GeneratorDefault GeneratorType = "default"
	GeneratorVipCoin GeneratorType = "VipCoin"
)

var generatorsInit = map[GeneratorType]GeneratorInit{
	GeneratorVipCoin: vipcoin.NewGenerator,
}

// GeneratorsString array of generators in string representation.
var generatorsString = []string{
	GeneratorDefault.String(),
	GeneratorVipCoin.String(),
}

// String returns string representation of generator type.
func (g GeneratorType) String() string {
	return string(g)
}
