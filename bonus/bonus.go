package bonus

import (
	"errors"
	"math"
	"strings"
)

//Receba uma lista de camisas, cada uma contendo o preço e o tamanho. O tamanho da camisa é representado por uma string,
//que pode ser "M" ou uma combinação de caracteres "X" seguida por "S" ou "L".
//
//Por exemplo, as strings "M", "XXL", "S" e "XXXXXXXS" podem representar tamanhos de algumas camisas. Já as strings "
//XM", "LL" e "SX" não representam tamanhos válidos.
//
//O objetivo é calcular a média entre o preço da maior camisa e o preço da menor camisa da lista.
//
//A comparação entre os tamanhos das camisas deve seguir as seguintes regras:
//
//- Qualquer tamanho pequeno (independentemente da quantidade de caracteres "X") é menor que o tamanho médio e qualquer
//  tamanho grande.
//- Qualquer tamanho grande (independentemente da quantidade de caracteres "X") é maior que o tamanho médio e qualquer
//  tamanho pequeno.
//- Quanto mais caracteres "X" antes de "S", menor o tamanho.
//- Quanto mais caracteres "X" antes de "L", maior o tamanho.
//
//Por exemplo:
//
//1. "XXXS" < "XS"
//2. "XXXL" > "XL"
//3. "XL" > "M"
//4. "XXL" = "XXL"
//5. "XXXXXS" < "M"
//6. "XL" > "XXXS"
//
//Dessa forma, ao receber a lista de camisas com seus respectivos preços e tamanhos, você deve calcular a média de preços
//da maior e da menor camisa.
//
//Caso não seja possível calcular a média, retorne um erro.

type Shirt struct {
	Size  string
	Price float64
}

func CalculateAveragePrice(shirts []Shirt) (max float64, min float64, err error) {
	if len(shirts) == 0 {
		return 0, 0, errors.New("no shirts")
	}

	shirtsMap := make(map[int][]float64)

	var (
		minSize = math.MaxInt
		maxSize = math.MinInt
	)

	for _, shirt := range shirts {
		var result int
		switch {
		case strings.Contains(shirt.Size, "S"):
			result = -1 - (10 * strings.Count(shirt.Size, "X"))
		case strings.Contains(shirt.Size, "L"):
			result = 1 + (10 * strings.Count(shirt.Size, "X"))
		}
		shirtsMap[result] = append(shirtsMap[result], shirt.Price)
		if result < minSize {
			minSize = result
		}
		if result > maxSize {
			maxSize = result
		}
	}

	return calculateAverage(shirtsMap[maxSize]), calculateAverage(shirtsMap[minSize]), nil
}

func calculateAverage(prices []float64) float64 {
	var sum float64
	for _, price := range prices {
		sum += price
	}
	return sum / float64(len(prices))
}
