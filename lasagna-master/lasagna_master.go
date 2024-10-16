package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, count int) (total int) {
	levelsLayers := len(layers)
	if count == 0 {
		count = 2
	}
	total = levelsLayers * count
	return
}

// TODO: define the 'Quantities()' function
func Quantities(ingred []string) (noodles int, sauce float64) {
	for i := range ingred {
		switch ingred[i] {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
	secretIngred := ""
	for _, fv := range friendsList {

		isExist := false
		for _, mv := range myList {
			if fv == mv {
				isExist = true
				break
			}
		}
		if isExist == false {
			secretIngred = fv
		}
	}

	lastSymbol := len(myList) - 1
	myList[lastSymbol] = secretIngred
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(originalList []float64, countOfportions int) []float64 {
	neededSlice := make([]float64, 0, len(originalList))
	for _, v := range originalList {
		needIngredValue := v / 2 * float64(countOfportions)
		neededSlice = append(neededSlice, needIngredValue)
	}

	return neededSlice
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
