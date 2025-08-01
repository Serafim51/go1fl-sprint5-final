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

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	var (
		errInvalidFormat = errors.New("data format error")
		errNumberParsing = errors.New("step count retrieval error")
		errTimeParsing   = errors.New("time value retrieval error")
		errZeroSteps     = errors.New("step count <= 0")
		errTimeZero      = errors.New("time value <= 0")
	)
	units := strings.Split(datastring, ",") // разделение строки data на слайс строк
	if len(units) != 2 {                    // проверка длины слайса
		return fmt.Errorf("%w: expected format - 'number,duration'", errInvalidFormat)
	}
	number, err := strconv.Atoi(units[0])
	if err != nil {
		return fmt.Errorf("%w: %v", errNumberParsing, err)
	}
	if number <= 0 { //Проверка количества шагов
		return fmt.Errorf("%w: step count must be greater than 0", errZeroSteps)
	}
	duration, err := time.ParseDuration(units[1])
	if err != nil {
		return fmt.Errorf("%w: %v", errTimeParsing, err)
	}
	if duration <= 0 {
		return fmt.Errorf("%w", errTimeZero)
	}
	ds.Steps = number
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, float64(ds.Personal.Height))
	if ds.Steps <= 0 {
		return "", errors.New("step count must be greater than 0")
	}

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, float64(ds.Personal.Weight), float64(ds.Personal.Height), ds.Duration)
	if err != nil { //определение ошибки

		return "", fmt.Errorf("Spent calories error: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
