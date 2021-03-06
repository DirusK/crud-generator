package app

import (
	"crud-generator/internal/generators"
	vipcoin "crud-generator/internal/generators/vipcoin/generator"
	"crud-generator/internal/models"
)

type (
	// GeneratorType represents the type of generator.
	GeneratorType string

	// GeneratorInit init function to create new generator.
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
	// GeneratorDefault.String(),
	GeneratorVipCoin.String(),
}

// String returns string representation of generator type.
func (g GeneratorType) String() string {
	return string(g)
}
