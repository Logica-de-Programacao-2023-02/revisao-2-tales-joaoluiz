package bonus

import (
	"fmt"
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
	var finalSmallPrice, finalLargePrice []float64
	var finalSmallSize, finalLargeSize []string
	var divisorSmall, divisorLarge float64
	var mediaSmall, mediaLarge float64
	var sumSmall, sumLarge float64
	var xCountS, xCountL int

	//countXMap := make(map[string]int)

	var (
		smallPrice  []float64
		largePrice  []float64
		slicePriceS []float64
		slicePriceM []float64
		slicePriceL []float64
		smallSize   []string
		largeSize   []string
		sliceSizeS  []string
		sliceSizeM  []string
		sliceSizeL  []string
	)

	if len(shirts) == 0 {
		return 0, 0, fmt.Errorf("empty")
	}

	for _, shirt := range shirts {
		if strings.Contains(shirt.Size, "S") {
			slicePriceS = append(slicePriceS, shirt.Price)
			sliceSizeS = append(sliceSizeS, shirt.Size)

		} else if strings.Contains(shirt.Size, "M") {
			slicePriceM = append(slicePriceM, shirt.Price)
			sliceSizeM = append(sliceSizeM, shirt.Size)
		} else {
			slicePriceL = append(slicePriceL, shirt.Price)
			sliceSizeL = append(sliceSizeL, shirt.Size)
		}
	}

	//smallSize / Price, largeSize / Price
	//tem o com M e S ------ len(sliceSizeS) > 0 && len(sliceSizeM) > 0 && len(sliceSizeL) == 0
	//tem 2 com S, M e L --- len(sliceSizeS) > 0 && len(sliceSizeM) > 0 && len(sliceSizeL) > 0
	//tem o que só tem L --- len(sliceSizeS) == 0 && len(sliceSizeM) == 0 && len(sliceSizeL) > 0
	//tem o que só tem S --- len(sliceSizeS) > 0 && len(sliceSizeM) == 0 && len(sliceSizeL) == 0

	if len(sliceSizeS) > 0 && len(sliceSizeM) > 0 && len(sliceSizeL) == 0 {
		smallSize = sliceSizeS
		smallPrice = slicePriceS
		largeSize = sliceSizeM
		largePrice = slicePriceM

		for _, size := range smallSize {
			if len(size) > xCountS {
				xCountS = len(size)
			}
		}

		for _, size := range largeSize {
			if len(size) > xCountL {
				xCountL = len(size)
			}
		}

		for i, size := range smallSize {
			if len(size) == xCountS {
				finalSmallPrice = append(finalSmallPrice, smallPrice[i])
				finalSmallSize = append(finalSmallSize, size)
			}
		}

		for i, size := range largeSize {
			if len(size) == xCountL {
				finalLargePrice = append(finalLargePrice, largePrice[i])
				finalLargeSize = append(finalLargeSize, size)
			}
		}

		divisorSmall = float64(len(finalSmallSize))
		divisorLarge = float64(len(finalLargeSize))

		for _, price := range finalSmallPrice {
			sumSmall += price
		}

		for _, price := range finalLargePrice {
			sumLarge += price
		}
	}

	if len(sliceSizeS) > 0 && len(sliceSizeM) > 0 && len(sliceSizeL) > 0 {
		smallSize = sliceSizeS
		smallPrice = slicePriceS
		largeSize = sliceSizeL
		largePrice = slicePriceL

		for _, size := range smallSize {
			if len(size) > xCountS {
				xCountS = len(size)
			}
		}

		for _, size := range largeSize {
			if len(size) > xCountL {
				xCountL = len(size)
			}
		}

		for i, size := range smallSize {
			if len(size) == xCountS {
				finalSmallPrice = append(finalSmallPrice, smallPrice[i])
				finalSmallSize = append(finalSmallSize, size)
			}
		}

		for i, size := range largeSize {
			if len(size) == xCountL {
				finalLargePrice = append(finalLargePrice, largePrice[i])
				finalLargeSize = append(finalLargeSize, size)
			}
		}

		divisorSmall = float64(len(finalSmallSize))
		divisorLarge = float64(len(finalLargeSize))

		for _, price := range finalSmallPrice {
			sumSmall += price
		}

		for _, price := range finalLargePrice {
			sumLarge += price
		}
	}

	if len(sliceSizeS) == 0 && len(sliceSizeM) == 0 && len(sliceSizeL) > 0 {
		smallSize = sliceSizeL
		smallPrice = slicePriceL
		largeSize = sliceSizeL
		largePrice = slicePriceL

		//para o menor ter um tamanho inicial
		xCountS = len(largeSize[0])

		for _, size := range largeSize {
			if len(size) < xCountS {
				xCountS = len(size)
			}

			if len(size) > xCountL {
				xCountL = len(size)
			}
		}

		for i, size := range largeSize {
			if len(size) == xCountS {
				finalSmallPrice = append(finalSmallPrice, smallPrice[i])
				finalSmallSize = append(finalSmallSize, size)
			}

			if len(size) == xCountL {
				finalLargePrice = append(finalLargePrice, largePrice[i])
				finalLargeSize = append(finalLargeSize, size)
			}
		}

		divisorSmall = float64(len(finalSmallSize))
		divisorLarge = float64(len(finalLargeSize))

		for _, price := range finalSmallPrice {
			sumSmall += price
		}

		for _, price := range finalLargePrice {
			sumLarge += price
		}
	}

	if len(sliceSizeS) > 0 && len(sliceSizeM) == 0 && len(sliceSizeL) == 0 {
		smallSize = sliceSizeS
		smallPrice = slicePriceS
		largeSize = sliceSizeS
		largePrice = slicePriceS

		//para o menor ter um tamanho inicial
		xCountS = len(smallSize[0])

		for _, size := range smallSize {
			if len(size) < xCountS {
				xCountS = len(size)
			}

			if len(size) > xCountL {
				xCountL = len(size)
			}
		}

		for i, size := range smallSize {
			//quanto menor a quantidade de X, maior a camisa
			if len(size) == xCountS {
				finalLargePrice = append(finalLargePrice, largePrice[i])
				finalLargeSize = append(finalLargeSize, size)
			}

			//quanto maior a quantidade de X, menor a camisa
			if len(size) == xCountL {
				finalSmallPrice = append(finalSmallPrice, smallPrice[i])
				finalSmallSize = append(finalSmallSize, size)
			}
		}

		divisorSmall = float64(len(finalSmallSize))
		divisorLarge = float64(len(finalLargeSize))

		for _, price := range finalSmallPrice {
			sumSmall += price
		}

		for _, price := range finalLargePrice {
			sumLarge += price
		}
	}

	mediaSmall = sumSmall / divisorSmall
	mediaLarge = sumLarge / divisorLarge

	return mediaLarge, mediaSmall, nil
}
