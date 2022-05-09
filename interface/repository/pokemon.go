package repository

import (
	"fmt"
	"go-bootcamp/domain/model"
	"go-bootcamp/infrastructure/database"
	"go-bootcamp/usecase/repository"
	"reflect"
	"strconv"
	"strings"
)

type pokemonRepository struct {
	db database.DB
}

func New(db database.DB) repository.PokemonRepository {
	return &pokemonRepository{db}
}

func recordToPokemon(r []string, p *model.Pokemon, ri *int, v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		// fmt.Println(v.Type().Field(i).Name)
		switch f.Kind() {
		case reflect.String:
			f.SetString(r[*ri])
			*ri++
		case reflect.Int:
			parsedInt, err := strconv.ParseInt(r[*ri], 10, 32)
			if err != nil {
				return err
			}
			f.SetInt(parsedInt)
			*ri++
		case reflect.Uint:
			parsedUint, err := strconv.ParseUint(r[*ri], 10, 32)
			if err != nil {
				return err
			}
			f.SetUint(parsedUint)
			*ri++
		case reflect.Bool:
			parsedBool, err := strconv.ParseBool(r[*ri])
			if err != nil {
				return err
			}
			f.SetBool(parsedBool)
			*ri++
		case reflect.Struct:
			si := reflect.ValueOf(f.Addr().Interface())
			s := si.Elem()
			err := recordToPokemon(r, p, ri, s)
			if err != nil {
				return err
			}
		case reflect.Slice:
			sliceElem := v.Type().Field(i).Type.Elem()
			kind := sliceElem.Kind()
			values := strings.Split(r[*ri], ",")

			switch kind {
			case reflect.String:
				newSlice := reflect.AppendSlice(f, reflect.ValueOf(values))
				f.Set(newSlice)
			case reflect.Struct:
				newSlice := reflect.MakeSlice(f.Type(), 0, len(values))
				for _, value := range values {
					newStruct := reflect.New(sliceElem).Elem()
					attributes := strings.Split(value, "$")
					for j := 0; j < newStruct.NumField(); j++ {
						sf := newStruct.Field(j)
						switch sf.Kind() {
						case reflect.String:
							sf.SetString(attributes[j])
						case reflect.Bool:
							b, err := strconv.ParseBool(attributes[j])
							if err != nil {
								return err
							}
							sf.SetBool(b)
						}
					}
					newSlice = reflect.Append(newSlice, newStruct)
				}
				f.Set(newSlice)
			}
			*ri++
		}
	}
	return nil
}

func parsePokemon(r [][]string) ([]model.Pokemon, error) {
	var ps []model.Pokemon

	for _, record := range r {
		p := model.Pokemon{}
		pv := reflect.ValueOf(&p)
		v := pv.Elem()
		recordIndex := 0

		err := recordToPokemon(record, &p, &recordIndex, v)
		if err != nil {
			return nil, fmt.Errorf("Parsing pokemon data: %w", err)
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func parseRecord(p *model.Pokemon) []string {
	aSlice := []string{}
	for _, a := range p.Abilities {
		aSlice = append(aSlice, fmt.Sprintf("%v$%v", a.Name, a.IsHidden))
	}
	abilities := fmt.Sprintf("%v", strings.Join(aSlice, ","))
	moves := fmt.Sprintf("%v", strings.Join(p.Moves, ","))
	types := fmt.Sprintf("%v", strings.Join(p.Types, ","))

	return []string{
		strconv.FormatUint(uint64(p.ID), 10),
		p.Name,
		p.Image,
		strconv.FormatUint(uint64(p.Weight), 10),
		strconv.FormatUint(uint64(p.Height), 10),
		strconv.FormatUint(uint64(p.BaseExperience), 10),
		p.Species.Name,
		strconv.FormatInt(int64(p.Species.GenderRate), 10),
		strconv.FormatBool(bool(p.Species.IsBaby)),
		strconv.FormatBool(bool(p.Species.IsLegendary)),
		strconv.FormatBool(bool(p.Species.IsMythical)),
		p.Species.Habitat,
		abilities,
		moves,
		types,
	}
}

func (pr *pokemonRepository) FindAll() ([]model.Pokemon, error) {
	r, err := pr.db.Read()
	if err != nil {
		return nil, err
	}
	p, err := parsePokemon(r)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pr *pokemonRepository) FindById(ID uint) (p model.Pokemon, err error) {
	r, err := pr.db.Read()
	if err != nil {
		return
	}
	ps, err := parsePokemon(r)
	if err != nil {
		return
	}

	for _, item := range ps {
		if item.ID == ID {
			p = item
			break
		}
	}
	return
}

func (pr *pokemonRepository) Create(p *model.Pokemon) (*model.Pokemon, error) {
	r := parseRecord(p)
	if err := pr.db.Write(r); err != nil {
		return nil, err
	}

	return p, nil
}
