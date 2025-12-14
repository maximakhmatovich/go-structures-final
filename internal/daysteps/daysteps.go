package daysteps

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
	errInvalidSteps    = errors.New("некорректное количество шагов")
	errInvalidDuration = errors.New("некорректная продолжительность")
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parsedData := strings.Split(datastring, ",")
	if len(parsedData) != 2 {
		return errInvalidInput
	}

	steps, err := strconv.Atoi(parsedData[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return errInvalidSteps
	}

	duration, err := time.ParseDuration(parsedData[1])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return errInvalidDuration
	}

	ds.Duration = duration
	ds.Steps = steps
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories)
	return result, nil
}
