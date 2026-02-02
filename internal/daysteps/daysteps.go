package daysteps

import (
	"fmt"
	
	"log"

	"time"

	"strings"

	"strconv"

	"errors"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

var (
conversionError = errors.New("ошибка преобразования типа")

invalidStepsError = errors.New("некорректное количество шагов")

parseError = errors.New("ошибка парсинга")

invalidDurationError = errors.New("некорректное время")
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

	if steps <= 0 {

		return 0, 0, invalidStepsError
		}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {

		return 0, 0, conversionError
		}
	if duration <= 0 {

		return 0, 0, invalidDurationError
	}	
	return steps, duration, nil

	// TODO: реализовать функцию
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {

		log.Println(err)

		return ""
	}
	if weight <= 0 {

		return ""
	}
	if height <= 0 {
		return ""
	}
	distance := (stepLength * float64(steps)) / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		
		return ""
	}
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)

	return info
}
