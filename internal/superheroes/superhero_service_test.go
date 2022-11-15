package superheroes

import (
	"github.com/OmgAbear/aoe/internal/infrastructure"
	"github.com/OmgAbear/aoe/internal/superheroes/entity"
	"github.com/OmgAbear/aoe/internal/superheroes/value_object"
	"reflect"
	"testing"
)

type RepoMock struct {
	createReturnData       entity.Superhero
	loadBySuperPowerReturn []entity.Superhero
}

func (repoMock RepoMock) Create(superhero entity.Superhero) entity.Superhero {
	return repoMock.createReturnData
}

func (repoMock RepoMock) LoadBySuperpowers(map[string]struct{}) []entity.Superhero {
	return repoMock.loadBySuperPowerReturn
}

func TestSuperheroService_Encrypt(t *testing.T) {

	type fields struct {
		validator     superheroValidator
		superheroRepo infrastructure.SuperheroCreatorRepoI
	}
	type args struct {
		identity value_object.Identity
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   value_object.Identity
	}{
		{
			"Check encryption functionality with default shift",
			fields{
				validator:     newSuperheroValidator(),
				superheroRepo: RepoMock{},
			},
			args{
				identity: value_object.Identity{
					FirstName: "a",
					LastName:  "b",
				},
			},
			value_object.Identity{
				FirstName: "d",
				LastName:  "e",
			},
		},
		{
			"Check that we're restarting when running out of letters",
			fields{
				validator:     newSuperheroValidator(),
				superheroRepo: RepoMock{},
			},
			args{
				identity: value_object.Identity{
					FirstName: "y",
					LastName:  "z",
				},
			},
			value_object.Identity{
				FirstName: "c",
				LastName:  "d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := SuperheroService{
				validator:     tt.fields.validator,
				superheroRepo: tt.fields.superheroRepo,
			}
			if got := service.Encrypt(tt.args.identity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuperheroService_LoadHeroes(t *testing.T) {
	type fields struct {
		validator     superheroValidator
		superheroRepo infrastructure.SuperheroCreatorRepoI
	}
	type args struct {
		params        map[string]struct{}
		shouldEncrypt bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []entity.Superhero
	}{
		{
			"Check loaded superheroes with encryption",
			fields{
				validator: newSuperheroValidator(),
				superheroRepo: RepoMock{
					loadBySuperPowerReturn: []entity.Superhero{
						{
							Name: "Data1",
							Identity: value_object.Identity{
								FirstName: "a",
								LastName:  "b",
							},
							Birthday:    "14-12-1989",
							Superpowers: []string{"stuff"},
						},
					},
				},
			},
			args{nil, true},
			[]entity.Superhero{
				{
					Name: "Data1",
					Identity: value_object.Identity{
						FirstName: "d",
						LastName:  "e",
					},
					Birthday:    "14-12-1989",
					Superpowers: []string{"stuff"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := SuperheroService{
				validator:     tt.fields.validator,
				superheroRepo: tt.fields.superheroRepo,
			}
			if got := service.LoadHeroesByParams(tt.args.params, tt.args.shouldEncrypt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadHeroesByParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
