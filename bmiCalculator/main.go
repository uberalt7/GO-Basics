package main

import (
	"fmt"
	"math"
)

func main() {
	// const bmiPower = 2
	var userHeight, userWeight float64
	var result string
	fmt.Print("Введите рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите вес в килограммах: ")
	fmt.Scan(&userWeight)
	bmi := userWeight / math.Pow(userHeight/100, 2)
	if bmi < 16 {
		result = fmt.Sprintf("Ваш индекс массы тела: %.f.\nУ Вас выраженный дефицит массы тела.", bmi)
	} else if bmi >= 16 && bmi < 18 {
		result = fmt.Sprintf("Ваш индекс массы тела: %.f.\nУ Вас недостаточная масса тела.", bmi)
	} else if bmi >= 18 && bmi < 25 {
		result = fmt.Sprintf("Ваш индекс массы тела: %.f.\nУ Вас нормальный вес.", bmi)
	} else if bmi >= 25 && bmi < 30 {
		result = fmt.Sprintf("Ваш индекс массы тела: %.f.\nУ Вас избыточная масса тела.", bmi)
	} else if bmi >= 30 && bmi < 35 {
		result = fmt.Sprintf("Ваш индекс массы тела: %.f.\nУ Вас ожирение I степени.", bmi)
	} else if bmi >= 35 && bmi < 40 {
		result = fmt.Sprintf("Ваш индекс массы тела: %.f.\nУ Вас ожирение I степени.", bmi)
	} else if bmi >= 40 {
		result = fmt.Sprintf("Ваш индекс массы тела: %.f.\nУ Вас ожирение III степени.", bmi)
	}
	fmt.Print(result)
}

/*
if <condition1> {
	...
} else if <condition2.1> && <condition2.2> || <condition2.3> {
    ...
} else {
	...
}

switch variable {
case 1:
	<...>
case 2, 3, 4:
	<...>
default:
	<...>
}
*/
