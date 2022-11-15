package infrastructure

import (
	"encoding/json"
	"github.com/OmgAbear/aoe/internal/superheroes/entity"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const heroFileEnv = "HERO_FILE_ENV"

var lock = &sync.Mutex{}

var instance *SuperheroRepo

type SuperheroReaderRepoI interface {
	LoadBySuperpowers(map[string]struct{}) []entity.Superhero
}

type SuperheroCreatorRepoI interface {
	SuperheroReaderRepoI
	Create(superhero entity.Superhero) entity.Superhero
}

//SuperheroRepo is the "repository pattern" implementation for superhero
type SuperheroRepo struct {
	//for the purpose of the exercise I will just load here to memory the superhero data
	//no size guards on this is intentional
	data []entity.Superhero
}

//getRepoInstance - initiates the repo "instance" while also loading the data
func getRepoInstance() *SuperheroRepo {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &SuperheroRepo{
				data: loadData(),
			}
		}
	}

	return instance
}

//NewSuperheroRepo - return a new "instance" of the superhero repo
func NewSuperheroRepo() *SuperheroRepo {
	return getRepoInstance()
}

//LoadBySuperpowers - returns a list of superheroes by superpowers
//the filtering logic is implemented here due to the fact that "normally" I'd probably do a "where in"
func (repository *SuperheroRepo) LoadBySuperpowers(superpowers map[string]struct{}) []entity.Superhero {
	if superpowers == nil {
		return repository.data
	}
	var filtered []entity.Superhero
	//O(N²) should be fine in this case
	for _, entry := range repository.data {
		for _, superpower := range entry.Superpowers {
			if _, ok := superpowers[superpower]; ok {
				filtered = append(filtered, entry)
				break
			}
		}
	}

	return filtered
}

//Create - create a new superhero entry
//normally would return some error but in this case unlikely to fail
func (repository *SuperheroRepo) Create(superhero entity.Superhero) entity.Superhero {
	repository.data = append(repository.data, superhero)

	return superhero
}

//loadData - dummy method to load the data from json file
func loadData() []entity.Superhero {
	location, found := os.LookupEnv(heroFileEnv)

	if !found {
		log.Fatal("set HERO_FILE_ENV with file location")
	}
	path, _ := filepath.Abs(location)
	bytes, _ := ioutil.ReadFile(path)

	var data []entity.Superhero
	_ = json.Unmarshal(bytes, &data)
	return data
}
