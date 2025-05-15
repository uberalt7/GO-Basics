package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var userHeight, userWeight float64
	checker := true

	for checker {
		userHeight, userWeight = 0, 0
		fmt.Print("Введите рост в сантиметрах: ")
		fmt.Scan(&userHeight)
		fmt.Print("Введите вес в килограммах: ")
		fmt.Scan(&userWeight)
		bmi, err := calculateBMI(userWeight, userHeight)
		if err != nil {
			fmt.Println(err)
			fmt.Println("==================")
			continue
		}
		printBMI(bmi)
		checker = checkRepeat()
	}
}

func calculateBMI(userWeight float64, userHeight float64) (float64, error) {
	const bmiPower = 2
	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("Неверное значение роста или веса")
	}
	return userWeight / math.Pow(userHeight/100, bmiPower), nil
}

func printBMI(bmi float64) {
	var result string
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

func checkRepeat() bool {
	var choise string
	fmt.Print("\nХотите воспользоваться калькулятором ещё раз? y/n: ")

	fmt.Scan(&choise)
	for choise != "y" && choise != "yes" && choise != "ja" && choise != "да" && choise != "д" && choise != "n" && choise != "no" && choise != "nein" && choise != "нет" && choise != "н" {
		fmt.Print("Введите y или n: ")
		fmt.Scan(&choise)
	}

	switch choise {
	case "n", "no", "nein", "нет", "н":
		fmt.Print("До свидания!")
		return false
	case "y", "yes", "ja", "да", "д":
		fmt.Println("==================")
		return true
	}
	return true
}

/*
If
if <condition1> {
	<...>
} else if <condition2.1> && <condition2.2> || <condition2.3> {
    <...>
} else {
	<...>
}

Switch — case
switch variable {
case 1:
	<...>
case 2, 3, 4:
	<...>
default:
	<...>
}

Цикл (только for)
for i := 0; i < 10; i++ {
	<...>
}

Бесконечный цикл, break и continue
for {
	<...>
	if <condition1> {
		break
	} else if <condition2> {
		continue
	}
}
*/
