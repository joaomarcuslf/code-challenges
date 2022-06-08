package challenge03

import (
	"sort"
	"strings"
)

type Pokemon struct {
	Dex     int
	Name    string
	PrevEvo []int
	NextEvo []int
}

func getNormalSearch(pokemon []Pokemon, searchTerm string) []Pokemon {
	var result []Pokemon

	for _, p := range pokemon {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(searchTerm)) {
			result = append(result, p)
		}
	}

	return result
}

func mapSliceToDex(pokemon []Pokemon) map[int]Pokemon {
	var result = make(map[int]Pokemon)

	for _, p := range pokemon {
		result[p.Dex] = p
	}

	return result
}

func mapDexToSlice(dex map[int]Pokemon) []Pokemon {
	var result []Pokemon

	for _, p := range dex {
		result = append(result, p)
	}

	return result
}

func getFirstOnEvoLine(pokemon Pokemon, dex map[int]Pokemon) Pokemon {
	aux := pokemon

	for len(aux.PrevEvo) > 0 {

		aux = dex[aux.PrevEvo[0]]
	}

	return aux
}

func getNextEvoLine(pokemon Pokemon, dex map[int]Pokemon) []Pokemon {
	results := []Pokemon{pokemon}

	var pokemonsToMap = []int{pokemon.Dex}

	// Get all next evolutions until the end

	for len(pokemonsToMap) > 0 {
		var nextPokemons []Pokemon

		for _, p := range pokemonsToMap {
			nextPokemons = append(nextPokemons, dex[p])
		}

		results = append(results, nextPokemons...)

		pokemonsToMap = []int{}

		for _, p := range nextPokemons {
			pokemonsToMap = append(pokemonsToMap, p.NextEvo...)
		}

	}

	return results
}

// https://www.freecodecamp.org/news/graph-algorithms-and-data-structures-explained-with-java-and-c-examples/
// https://www.geeksforgeeks.org/graph-data-structure-and-algorithms/

func SearchPokemon(searchType string, pokemon []Pokemon, searchTerm string) []Pokemon {
	dex := mapSliceToDex(pokemon)

	search := getNormalSearch(pokemon, searchTerm)

	switch searchType {
	case "evolution":
		resultDex := mapSliceToDex([]Pokemon{})

		for _, p := range search {
			fe := getFirstOnEvoLine(p, dex)

			resultDex[fe.Dex] = fe
		}

		for _, p := range resultDex {
			evos := getNextEvoLine(p, dex)

			for _, e := range evos {
				resultDex[e.Dex] = e
			}
		}

		search = mapDexToSlice(resultDex)

	default:
		break
	}

	sort.Slice(search, func(i, j int) bool {
		return search[i].Dex < search[j].Dex
	})

	return search
}
