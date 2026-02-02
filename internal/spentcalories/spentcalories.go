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

	divisorZero = errors.New("ошибка деления на 0")

	parseError = errors.New("Ошибка парсинга")
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

	duration, err := time.ParseDuration(trainingParts[2])
	if err != nil {

		return 0,"", 0, conversionError
	}
	activityType:= trainingParts[1]
	
	return steps, duration, activityType, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient 

	distanceKm := (float64(steps) * stepLength) / mInKm

	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

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
	activityHours := duration.Hours()
	switch(activityType) {

	case "Ходьба":

		walkingDistance := distance(steps, height)
		walkingSpeed := meanSpeed(steps, height, duration)
		walkingCalories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		walkingInfo := fmt.Sprintf("Тип тренировки: %s.\nДлительность: %f ч.\nДистанция: %f км.\nСкорость: %f км/ч\nСожгли калорий:%f\n", 
		activityType, activityHours, walkingDistance, walkingSpeed, walkingCalories )

		return walkingInfo, nil

	case "Бег":

		runningDistance := distance(steps, height)
		runningSpeed := meanSpeed(steps, height, duration)
		runningCalories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		runningInfo := fmt.Sprintf("Тип тренировки: %s.\nДлительность: %f ч.\nДистанция: %f км.\nСкорость: %f км/ч\nСожгли калорий:%f\n", 
		activityType, activityHours, runningDistance, runningSpeed, runningCalories)

		return runningInfo, nil

	default:

		return "Неизвестный тип тренировки", nil

	}


	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {

		return 0, divisorZero
	}
	speed := meanSpeed(steps, height, duration)

	runningMinutes := duration.Minutes()

	return (weight * speed * runningMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {

		return 0, divisorZero
	}
	speed := meanSpeed(steps, height, duration)

	walkingMinutes := duration.Minutes()

	calories := (weight * speed * walkingMinutes) / minInH

	return calories * walkingCaloriesCoefficient, nil
}
