package application

import (
	"github.com/OmgAbear/aoe/internal/application/dto"
	"github.com/OmgAbear/aoe/internal/superheroes"
	"net/url"
	"strconv"
)

const UrlQuerySuperpowerKey = "superpowers"
const UrlQueryEncryptionKey = "encrypted"

//Service is the glue/facade between different layers (superheroes & infra in this case)
type Service struct {
	superheroService superheroes.SuperheroService
}

func NewService() Service {
	return Service{
		superheroService: superheroes.NewSuperheroService(),
	}
}

func (service Service) Create(inputDto dto.SuperheroInputDto) (dto.SuperheroOutputDto, error) {
	domainEntity := SuperHeroDtoToEntity(inputDto)
	err := service.superheroService.ValidateSuperpowers(domainEntity)

	if err != nil {
		return dto.SuperheroOutputDto{}, err
	}
	entity := service.superheroService.Create(SuperHeroDtoToEntity(inputDto))
	return SuperheroEntityToDto(entity), nil
}

//List - lists the heroes
//parses input and transforms to required internal structures
func (service Service) List(query url.Values) []dto.SuperheroOutputDto {
	var queriedSuperpowers map[string]struct{}
	returnList := make([]dto.SuperheroOutputDto, 0)
	if _, exists := query[UrlQuerySuperpowerKey]; exists {
		queriedSuperpowers = map[string]struct{}{}
		for _, queriedPower := range query[UrlQuerySuperpowerKey] {
			queriedSuperpowers[queriedPower] = struct{}{}
		}
	}

	shouldEncrypt, _ := strconv.ParseBool(query.Get(UrlQueryEncryptionKey))
	heroes := service.superheroService.LoadHeroesByParams(queriedSuperpowers, shouldEncrypt)

	for _, hero := range heroes {
		returnList = append(returnList, SuperheroEntityToDto(hero))
	}

	return returnList
}
