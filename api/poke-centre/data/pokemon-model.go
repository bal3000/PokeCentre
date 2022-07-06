package data

import (
	"encoding/json"

	"github.com/bal3000/PokeCentre/proto/pokemon"
)

type PokemonModel struct {
	Id          int32  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Number      int32  `json:"number,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

func (p PokemonModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

type PokemonCollection struct {
	Pokemon []PokemonModel `json:"pokemon"`
}

func (p PokemonCollection) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func MapProtoToModel(p *pokemon.Pokemon) PokemonModel {
	return PokemonModel{
		Id:          p.Id,
		Name:        p.Name,
		Number:      p.Number,
		Type:        p.Types[0].Name,
		Description: p.Description,
	}
}

func MapCollectionToModel(p []*pokemon.Pokemon) []PokemonModel {
	list := make([]PokemonModel, 0)
	for _, r := range p {
		list = append(list, MapProtoToModel(r))
	}
	return list
}
