package q4

import "errors"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	outgoingCities := make(map[string]bool)

	// Marca todas as cidades de partida
	for _, path := range caminhos {
		outgoingCities[path[0]] = true
	}

	// Encontra a cidade de destino
	for _, path := range caminhos {
		destinationCity := path[1]
		if !outgoingCities[destinationCity] {
			return destinationCity, nil
		}
	}
	return "", errors.New("cidade de destino não encontrada")
}
