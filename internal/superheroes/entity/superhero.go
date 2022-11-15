//Package entity holds the database entity representation for all the domains
package entity

import (
	"github.com/OmgAbear/aoe/internal/superheroes/value_object"
)

type (
	//Superhero is the aggregate root (entity)
	Superhero struct {
		Name        string                `json:"name"`
		Identity    value_object.Identity `json:"identity"`
		Birthday    string                `json:"birthday"`
		Superpowers []string              `json:"superpowers"`
	}
)
