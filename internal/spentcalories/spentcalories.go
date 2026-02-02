package spentcalories

import (
	"fmt"

	"time"

	"strings"

	"strconv"

	"errors"

	"log"
)
var (
	conversionError = errors.New("ошибка преобразования типа")

	parseError = errors.New("ошибка парсинга")

	unknownActiviryError = errors.New("неизвестный тип тренировки")

	invalidStepsError = errors.New("некорректное количество шагов")

	invalidDurationError = errors.New("некорректное время")

	invalidHeight = errors.New("некорректный рост")

	invalidWeight = errors.New("некорректный вес")
)
// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	trainingParts := strings.Split(data, ",")

	if len(trainingParts) != 3 {
		return 0, "", 0, parseError
	}

	steps, err := strconv.Atoi(trainingParts[0])
	if err != nil {

		return 0,"", 0, conversionError
	}
	if steps <= 0 {

		return 0, "", 0, invalidStepsError
	}
	duration, err := time.ParseDuration(trainingParts[2])
	if err != nil {

		return 0,"", 0, conversionError
	}
	if duration <= 0 {

		return 0, "", 0, invalidDurationError
	}
	activityType:= trainingParts[1]
	
	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if height <= 0 {

		return 0, invalidHeight
	}
	if steps <= 0 {

		return 0, invalidStepsError
	}
	stepLength := height * stepLengthCoefficient 

	distanceKm := (float64(steps) * stepLength) / mInKm

	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if height <= 0 {

		return 0, invalidHeight
	}
	if steps <= 0 {

		return 0, invalidStepsError
	}
	if duration <= 0 {

		return 0
	}
	distanceKm := distance(steps, height)

	walkingHours := duration.Hours()

	return distanceKm / walkingHours
	// TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {

	steps, activityType, duration, err := parseTraining(data)
	if err != nil {

		log.Println("ошибка парсинга:", err)

		return "", parseError
	}
	if height <= 0 {

		return 0, invalidHeight
	}
	if weight <= 0 {

		return 0, invalidWeight
	}
	activityHours := duration.Hours()
	switch(activityType) {

	case "Ходьба":

		walkingDistance := distance(steps, height)
		walkingSpeed := meanSpeed(steps, height, duration)
		walkingCalories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		walkingInfo := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий:%.2f\n", 
		activityType, activityHours, walkingDistance, walkingSpeed, walkingCalories )

		return walkingInfo, nil

	case "Бег":

		runningDistance := distance(steps, height)
		runningSpeed := meanSpeed(steps, height, duration)
		runningCalories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		runningInfo := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий:%.2f\n", 
		activityType, activityHours, runningDistance, runningSpeed, runningCalories)

		return runningInfo, nil

	default:

		return "", unknownActiviryError

	}


	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {

		return 0, invalidDurationError

	}
	if height <= 0 {

		return 0, invalidHeight
	}
	if weight <= 0 {

		return 0, invalidWeight
	}
	if steps <= 0 {

		return 0, invalidStepsError
	}
	speed := meanSpeed(steps, height, duration)

	runningMinutes := duration.Minutes()

	return (weight * speed * runningMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {

		return 0, invalidDurationError
	}
		if height <= 0 {

		return 0, invalidHeight
	}
	if weight <= 0 {

		return 0, invalidWeight
	}
	if steps <= 0 {

		return 0, invalidStepsError
	}
	speed := meanSpeed(steps, height, duration)

	walkingMinutes := duration.Minutes()

	calories := (weight * speed * walkingMinutes) / minInH

	return calories * walkingCaloriesCoefficient, nil
}
