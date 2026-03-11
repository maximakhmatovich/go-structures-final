package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var (
	errInvalidInput    = errors.New("некорректные входные данные")
	errInvalidActivity = errors.New("неизвестный тип тренировки")
	errInvalidSteps    = errors.New("некорректное количество шагов")
	errInvalidDuration = errors.New("некорректная продолжительность")
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parsedData := strings.Split(datastring, ",")
	if len(parsedData) != 3 {
		return errInvalidInput
	}

	steps, err := strconv.Atoi(parsedData[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return errInvalidSteps
	}

	t.Steps = steps
	t.TrainingType = parsedData[1]

	duration, err := time.ParseDuration(parsedData[2])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return errInvalidDuration
	}

	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var spentCalories float64
	var err error

	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpead := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}

	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}

	default:
		return "", errInvalidActivity
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpead, spentCalories)
	return result, nil
}
