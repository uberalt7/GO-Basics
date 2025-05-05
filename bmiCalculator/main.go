package main

import (
	"fmt"
	"math"
)

func main() {
	// const bmiPower = 2
	var userHeight, userWeight float64
	fmt.Print("Введите рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите вес в килограммах: ")
	fmt.Scan(&userWeight)
	bmi := userWeight / math.Pow(userHeight/100, 2)
	result := fmt.Sprintf("Ваш индекс массы тела: %.f", bmi)
	fmt.Print("Ура! ", result)
}

/*
if <condition1> {
	...
} else if <condition2.1> && <condition2.2> || <condition2.3> {
    ...
} else {
	...
}
*/
