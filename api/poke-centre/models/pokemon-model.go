package models

import (
	"encoding/json"

	"github.com/bal3000/PokeCentre/proto/pokemon"
)

type PokemonModel struct {
	Id             int32     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	IsDefault      bool      `json:"isDefault,omitempty"`
	Number         int32     `json:"number,omitempty"`
	Types          []Type    `json:"types,omitempty"`
	Description    string    `json:"description,omitempty"`
	Moves          []Move    `json:"moves,omitempty"`
	Abilities      []Ability `json:"abilities,omitempty"`
	Stats          []Stat    `json:"stats,omitempty"`
	Height         int32     `json:"height,omitempty"`
	Weight         int32     `json:"weight,omitempty"`
	BaseExperience int32     `json:"baseExperience,omitempty"`
	Order          int32     `json:"order,omitempty"`
	Sprites        Sprite    `json:"sprites,omitempty"`
	Species        Species   `json:"species,omitempty"`
}

type Type struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slot int32  `json:"slot,omitempty"`
}

type Stat struct {
	Name     string `json:"name,omitempty"`
	Effort   int32  `json:"effort,omitempty"`
	BaseStat int32  `json:"baseStat,omitempty"`
}

type Ability struct {
	Name     string `json:"name,omitempty"`
	Slot     int32  `json:"slot,omitempty"`
	IsHidden bool   `json:"isHidden,omitempty"`
}

type VersionGroupDetails struct {
	VersionName     string `json:"versionName,omitempty"`
	MoveLearnMethod string `json:"moveLearnMethod,omitempty"`
	LevelLearnedAt  int32  `json:"levelLearnedAt,omitempty"`
}

type Move struct {
	Name                string                `json:"name,omitempty"`
	VersionGroupDetails []VersionGroupDetails `json:"versionGroupDetails,omitempty"`
}

type Sprite struct {
	FrontDefault     string `json:"frontDefault,omitempty"`
	BackDefault      string `json:"backDefault,omitempty"`
	FrontShiny       string `json:"frontShiny,omitempty"`
	BackShiny        string `json:"backShiny,omitempty"`
	FrontFemale      string `json:"frontFemale,omitempty"`
	BackFemale       string `json:"backFemale,omitempty"`
	BackShinyFemale  string `json:"backShinyFemale,omitempty"`
	FrontShinyFemale string `json:"frontShinyFemale,omitempty"`
}

type SpeciesData struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Species struct {
	BaseHappiness      int32         `json:"baseHappiness,omitempty"`
	CaptureRate        int32         `json:"captureRate,omitempty"`
	Color              SpeciesData   `json:"color,omitempty"`
	EggGroups          []SpeciesData `json:"eggGroups,omitempty"`
	EvolutionChainId   int32         `json:"evolutionChainId,omitempty"`
	EvolvesFromSpecies SpeciesData   `json:"evolvesFromSpecies,omitempty"`
	FlavorText         string        `json:"flavorText,omitempty"`
	Genus              string        `json:"genus,omitempty"`
	Shape              SpeciesData   `json:"shape,omitempty"`
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
	pm, err := MapFromProto[*pokemon.Pokemon, PokemonModel](p)
	if err != nil {
		panic(err)
	}

	return pm
	// return PokemonModel{
	// 	Id:             p.Id,
	// 	Name:           p.Name,
	// 	Description:    p.Description,
	// 	IsDefault:      p.IsDefault,
	// 	Number:         p.Number,
	// 	Height:         p.Height,
	// 	Weight:         p.Weight,
	// 	BaseExperience: p.BaseExperience,
	// 	Order:          p.Order,
	// 	Types:          MapTypes(p.Types),
	// }
}

func MapCollectionToModel(p []*pokemon.Pokemon) []PokemonModel {
	list := make([]PokemonModel, 0)
	for _, r := range p {
		list = append(list, MapProtoToModel(r))
	}
	return list
}

func MapTypes(t []*pokemon.Type) []Type {
	types := make([]Type, 0)
	for _, r := range t {
		types = append(types, Type{
			Id:   r.Id,
			Name: r.Name,
			Slot: r.Slot,
		})
	}
	return types
}

func MapFromProto[T any, U any](from T) (U, error) {
	var result U
	js, err := json.Marshal(from)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal(js, &result); err != nil {
		return result, err
	}

	return result, nil
}
