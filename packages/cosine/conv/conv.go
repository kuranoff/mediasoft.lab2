package conv

import (
	"strconv"
)

// ConvertX ...
func ConvertX(v []string) []float64 {

	var vX = []float64{}

	// Корректируем массив (удаляем последний элемент со значением nil)
	vcorrect := v[:len(v)-1]

	for _, i := range vcorrect {

		j, err := strconv.ParseFloat(i, 64)

		if err != nil {
			panic(err)
		}

		vX = append(vX, j)

	}

	return vX

}
