package superheroes

import (
	"github.com/OmgAbear/aoe/internal/superheroes/entity"
	"github.com/OmgAbear/aoe/internal/superheroes/value_object"
	"testing"
)

func Test_superheroValidator_ValidateSuperpowers(t *testing.T) {
	type args struct {
		superhero entity.Superhero
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Check not allowed superpower",
			args: args{
				superhero: entity.Superhero{
					Name:        "",
					Identity:    value_object.Identity{},
					Birthday:    "",
					Superpowers: []string{"c"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := superheroValidator{}
			if err := validator.ValidateSuperpowers(tt.args.superhero); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSuperpowers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
