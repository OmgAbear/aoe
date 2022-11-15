package superheroes

import (
	"github.com/OmgAbear/aoe/internal/superheroes/entity"
)

type superheroValidator struct{}

var allowedSuperpowers = map[string]struct{}{
	"strength": {}, "speed": {}, "flight": {}, "invulnerability": {}, "healing": {},
}

//newSuperheroValidator - creates a new "instance" of the superheroValidator
func newSuperheroValidator() superheroValidator {
	return superheroValidator{}
}

//ValidateSuperpowers - checks
func (validator superheroValidator) ValidateSuperpowers(superhero entity.Superhero) error {
	for _, superpower := range superhero.Superpowers {
		if _, ok := allowedSuperpowers[superpower]; !ok {
			return InvalidSuperpowerError{}
		}
	}

	return nil
}
