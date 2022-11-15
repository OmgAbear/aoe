package application

import (
	"github.com/OmgAbear/aoe/internal/application/dto"
	"github.com/OmgAbear/aoe/internal/superheroes/entity"
	"github.com/OmgAbear/aoe/internal/superheroes/value_object"
)

func SuperHeroDtoToEntity(inputDto dto.SuperheroInputDto) entity.Superhero {
	return entity.Superhero{
		Name: inputDto.Name,
		Identity: value_object.Identity{
			FirstName: inputDto.Identity.FirstName,
			LastName:  inputDto.Identity.LastName,
		},
		Birthday:    inputDto.Birthday,
		Superpowers: inputDto.Superpowers,
	}
}

func SuperheroEntityToDto(domainEntity entity.Superhero) dto.SuperheroOutputDto {
	return dto.SuperheroOutputDto{
		Name: domainEntity.Name,
		//Could have added something for the marshalling, but I prefer having everything in one place
		Identity:    domainEntity.Identity.FirstName + " " + domainEntity.Identity.LastName,
		Birthday:    domainEntity.Birthday,
		Superpowers: domainEntity.Superpowers,
	}
}
