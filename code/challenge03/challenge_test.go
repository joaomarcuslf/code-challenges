package challenge03

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name       string
	searchType string
	dex        []Pokemon
	search     string
	expected   []Pokemon
}

func MiniDex() []Pokemon {
	return []Pokemon{
		Pokemon{
			Dex:     35,
			Name:    "Clefairy",
			PrevEvo: []int{173},
			NextEvo: []int{36},
		},
		Pokemon{
			Dex:     36,
			Name:    "Clefable",
			PrevEvo: []int{35},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     133,
			Name:    "Eevee",
			PrevEvo: []int{},
			NextEvo: []int{134, 135, 136, 196, 197, 470, 471, 700},
		},
		Pokemon{
			Dex:     134,
			Name:    "Vaporeon",
			PrevEvo: []int{133},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     135,
			Name:    "Jolteon",
			PrevEvo: []int{133},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     136,
			Name:    "Flareon",
			PrevEvo: []int{133},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     173,
			Name:    "Cleffa",
			PrevEvo: []int{},
			NextEvo: []int{35},
		},
		Pokemon{
			Dex:     183,
			Name:    "Marill",
			PrevEvo: []int{298},
			NextEvo: []int{184},
		},
		Pokemon{
			Dex:     184,
			Name:    "Azumarill",
			PrevEvo: []int{183},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     196,
			Name:    "Espeon",
			PrevEvo: []int{133},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     197,
			Name:    "Umbreon",
			PrevEvo: []int{133},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     280,
			Name:    "Ralts",
			PrevEvo: []int{},
			NextEvo: []int{281},
		},
		Pokemon{
			Dex:     281,
			Name:    "Kirlia",
			PrevEvo: []int{280},
			NextEvo: []int{282, 475},
		},
		Pokemon{
			Dex:     282,
			Name:    "Gardevoir",
			PrevEvo: []int{281},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     298,
			Name:    "Azurill",
			PrevEvo: []int{},
			NextEvo: []int{183},
		},
		Pokemon{
			Dex:     470,
			Name:    "Leafeon",
			PrevEvo: []int{133},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     471,
			Name:    "Glaceon",
			PrevEvo: []int{133},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     475,
			Name:    "Gallade",
			PrevEvo: []int{281},
			NextEvo: []int{},
		},
		Pokemon{
			Dex:     700,
			Name:    "Sylveon",
			PrevEvo: []int{133},
			NextEvo: []int{},
		},
	}
}

func TestSearchPokemon(t *testing.T) {
	// Cases
	// fmt.Println("Exemplo 01: ", challenge03.SearchPokemon("evolution", minidex, "Ga"))
	// fmt.Println("Exemplo 02: ", challenge03.SearchPokemon("evolution", minidex, "al"))
	// fmt.Println("Exemplo 03: ", challenge03.SearchPokemon("evolution", minidex, "sylveon"))
	// fmt.Println("Exemplo 04: ", challenge03.SearchPokemon("evolution", minidex, "cleffa"))
	// fmt.Println("Exemplo 05: ", challenge03.SearchPokemon("evolution", minidex, "a"))
	testCases := []TestCase{
		{
			name:       "Search for a pokemon",
			searchType: "pokemon",
			dex:        MiniDex(),
			search:     "cleffa",
			expected: []Pokemon{
				{
					Dex:     173,
					Name:    "Cleffa",
					PrevEvo: []int{},
					NextEvo: []int{35},
				},
			},
		},
		{
			name:       "Search for a Evo line",
			searchType: "evolution",
			dex:        MiniDex(),
			search:     "cleffa",
			expected: []Pokemon{
				{
					Dex:     35,
					Name:    "Clefairy",
					PrevEvo: []int{173},
					NextEvo: []int{36},
				},
				{
					Dex:     36,
					Name:    "Clefable",
					PrevEvo: []int{35},
					NextEvo: []int{},
				},
				{
					Dex:     173,
					Name:    "Cleffa",
					PrevEvo: []int{},
					NextEvo: []int{35},
				},
			},
		},
		{
			name:       "Search for no result",
			searchType: "evolution",
			dex:        MiniDex(),
			search:     "Furfles",
			expected:   []Pokemon{},
		},
	}

	for _, tc := range testCases {
		e1 := SearchPokemon(tc.searchType, tc.dex, tc.search)

		if (len(e1) > 0 && len(tc.expected) > 0) && !reflect.DeepEqual(e1, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, e1, tc.expected, "[]Pokemon")
		}
	}
}
