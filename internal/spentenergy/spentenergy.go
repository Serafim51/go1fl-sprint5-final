package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	var errInvalidInput = errors.New("invalid input data")
	if steps <= 0 {
		return 0, fmt.Errorf("%w: step count must be greater than 0", errInvalidInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w: weight must be greater than 0", errInvalidInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w: height must be greater than 0", errInvalidInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w: training duration must be greater than 0", errInvalidInput)
	}
	// реализация функции ** для себя** - требует теста
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	var errInvalidInput = errors.New("invalid input data")
	if steps <= 0 {
		return 0, fmt.Errorf("%w: step count must be greater than 0", errInvalidInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w: weight must be greater than 0", errInvalidInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w: height must be greater than 0", errInvalidInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w: training duration must be greater than 0", errInvalidInput)
	}
	// реализация функции ** для себя** - требует теста
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {

		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient) * float64(steps) / mInKm
}
