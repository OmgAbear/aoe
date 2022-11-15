package superheroes

import (
	"github.com/OmgAbear/aoe/internal/infrastructure"
	"github.com/OmgAbear/aoe/internal/superheroes/entity"
	"github.com/OmgAbear/aoe/internal/superheroes/value_object"
	"os"
	"strconv"
)

const CypherShiftBy = "ENV_SHIFT_BY"
const DefaultShiftBy = 3
const alphabet = "abcdefghijklmnopqrstuvwxyz"

var alphabetIndices = map[string]int{
	"a": 0, "b": 1, "c": 2, "d": 3, "e": 4, "f": 5, "g": 6, "h": 7, "i": 8, "j": 9, "k": 10, "l": 11,
	"m": 12, "n": 13, "o": 14, "p": 15, "q": 16, "r": 17, "s": 18, "t": 19, "u": 20, "v": 21, "w": 22,
	"x": 23, "y": 24, "z": 25,
}

//SuperheroService contains all the superheroes business logic required
type SuperheroService struct {
	validator     superheroValidator
	superheroRepo infrastructure.SuperheroCreatorRepoI
}

func NewSuperheroService() SuperheroService {
	return SuperheroService{
		validator:     newSuperheroValidator(),
		superheroRepo: infrastructure.NewSuperheroRepo(),
	}
}

//Encrypt accepts a string and returns the encrypted value
func (service SuperheroService) Encrypt(identity value_object.Identity) value_object.Identity {
	encryptor := func(toEncrypt string, shiftBy int64) string {
		returnVal := ""
		for i := 0; i < len(toEncrypt); i++ {
			//let's just add the space and move on
			if string(toEncrypt[i]) == " " {
				returnVal += " "
				continue
			}
			//this iterates through all the characters that need to be encrypted
			//gets the index of the letter
			//adds that index together with the shiftBy value
			//and uses that result to retrieve the "encrypted" value of the letter
			//takes out of range into account
			newPosition := alphabetIndices[string(toEncrypt[i])] + int(shiftBy)
			maxAlphabetIdx := len(alphabetIndices) - 1
			if newPosition > maxAlphabetIdx {
				newPosition = newPosition - maxAlphabetIdx
			}
			returnVal += string(alphabet[newPosition])
		}
		return returnVal
	}
	shiftBy := service.loadShiftByValue()
	return value_object.Identity{
		FirstName: encryptor(identity.FirstName, shiftBy),
		LastName:  encryptor(identity.LastName, shiftBy),
	}
}

func (service SuperheroService) ValidateSuperpowers(superhero entity.Superhero) error {
	return service.validator.ValidateSuperpowers(superhero)
}

//LoadHeroesByParams - loads heroes by provided params with encryption if required
func (service SuperheroService) LoadHeroesByParams(params map[string]struct{}, shouldEncrypt bool) []entity.Superhero {
	listOfHeroes := service.superheroRepo.LoadBySuperpowers(params)
	if shouldEncrypt {
		for idx, hero := range listOfHeroes {
			listOfHeroes[idx].Identity = service.Encrypt(hero.Identity)
		}
	}

	return listOfHeroes
}

//Create - Creates the superhero Record
func (service SuperheroService) Create(superhero entity.Superhero) entity.Superhero {
	service.ValidateSuperpowers(superhero)
	return service.superheroRepo.Create(superhero)
}

//loadShiftByValue - loads the desired shiftBy value for the cypher
func (service SuperheroService) loadShiftByValue() int64 {
	shiftByEnvVal, exists := os.LookupEnv(CypherShiftBy)
	shiftBy, _ := strconv.ParseInt(shiftByEnvVal, 10, 8)
	if shiftBy == 0 || !exists {
		//some failsafe
		shiftBy = DefaultShiftBy
	}

	return shiftBy
}
