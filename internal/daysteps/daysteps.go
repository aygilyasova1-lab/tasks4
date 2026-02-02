package daysteps

import (
	"fmt"

	"time"

	"strings"

	"strconv"

	"errors"
)

var (
conversionError := errors.New("ошибка преобразования типа")

zeroSteps := errors.New("количество шагов равно 0")

parseError := errors.New("ошибка парсинга")
)
const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	parts := strings.Split(data, ",")

	if len(parts) != 2 {

		return 0, 0, parseError
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {

		return 0, 0, conversionError
		}

	if steps == 0 {

		return 0, 0, zeroSteps
		}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {

		return 0, 0, conversionError
		}	
	return steps, duration, nil

	// TODO: реализовать функцию
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {

		fmt.Println(err)

		return ""
	}
	if steps < 0 {

		return ""
	}
	distance := (stepLength * steps) / mInKm

	calories, err := WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		
		return ""
	}
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %f км.\nВы сожгли %f ккал.\n", steps, distance, calories)

	return info
}
