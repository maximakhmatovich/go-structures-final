package spentenergy

import (
	"errors"
	"log"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	errInvalidSteps    = errors.New("некорректное количество шагов")
	errInvalidHeight   = errors.New("некорректный рост")
	errInvalidWeight   = errors.New("некорректный вес")
	errInvalidDuration = errors.New("некорректная продолжительность")
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errInvalidSteps
	}

	if height <= 0 {
		return 0, errInvalidHeight
	}

	if weight <= 0 {
		return 0, errInvalidWeight
	}

	if duration <= 0 {
		return 0, errInvalidDuration
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	spentCalories := (weight * meanSpeed * duration.Minutes()) / minInH
	return spentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errInvalidSteps
	}

	if height <= 0 {
		return 0, errInvalidHeight
	}

	if weight <= 0 {
		return 0, errInvalidWeight
	}

	if duration <= 0 {
		return 0, errInvalidDuration
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	spentCalories := (weight * meanSpeed * duration.Minutes()) / minInH
	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		log.Println(errInvalidSteps)
		return 0
	}

	if height <= 0 {
		log.Println(errInvalidHeight)
		return 0
	}

	if duration <= 0 {
		log.Println(errInvalidDuration)
		return 0
	}

	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		log.Println(errInvalidSteps)
		return 0
	}

	if height <= 0 {
		log.Println(errInvalidHeight)
		return 0
	}

	stepLength := height * stepLengthCoefficient
	distance := (stepLength * float64(steps)) / mInKm
	return distance
}
