package q4

import (
	"fmt"
	"sort"
)

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	var destinationCities []string
	var originCities []string
	var destinationCity string
	var infiniteTravel bool

	if len(caminhos) == 0 {
		return "", fmt.Errorf("empty")
	}

	for _, travel := range caminhos {
		for i, city := range travel {
			if i == 0 {
				originCities = append(originCities, city)
			}

			if i == 1 {
				destinationCities = append(destinationCities, city)
			}
		}
	}

	sort.Strings(originCities)
	sort.Strings(destinationCities)

	copyDestination := destinationCities

	for i := 0; i < len(originCities); i++ {
		infiniteTravel = true

		if originCities[i] != destinationCities[i] {
			infiniteTravel = false
			break
		}
	}

	if infiniteTravel {
		return "", fmt.Errorf("infinite travel")
	}

	for i := 0; i < len(destinationCities); i++ {
		for j := 0; j < len(originCities); j++ {
			if originCities[j] == destinationCities[i] {
				copyDestination = append(copyDestination[:i], copyDestination[i+1:]...)
			}
		}
	}

	destinationCity = copyDestination[0]

	return destinationCity, nil
}
