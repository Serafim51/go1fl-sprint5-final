package trainings

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	var (
		errInvalidFormat = errors.New("invalid data format")
		errNumberParsing = errors.New("failed to parse step count")
		errZeroSteps     = errors.New("step count must be greater than 0")
		errTimeParsing   = errors.New("failed to parse time value")
		errDuration      = errors.New("invalid format or duration <= 0")
	)
	units := strings.Split(datastring, ",") // разделение строки на слайс строк
	if len(units) != 3 {                    // проверка длины слайса
		return fmt.Errorf("%w: expected format - 'number,string,duration'", errInvalidFormat) //проверка соответствия формату
	}
	number, err := strconv.Atoi(units[0])
	if err != nil { //проверка преобразования строки в число (количество шагов)
		return fmt.Errorf("%w: %v", errNumberParsing, err)
	}
	if number <= 0 { //проверка шагов на > 0
		return fmt.Errorf("%w", errZeroSteps)
	}
	duration, err := time.ParseDuration(units[2])
	if err != nil { //проверка формата времени
		return fmt.Errorf("%w: %v", errTimeParsing, err)
	}
	if duration <= 0 { //проверка формата времени
		return fmt.Errorf("%w", errDuration)
	}
	t.Steps = number
	t.TrainingType = units[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distant := spentenergy.Distance(t.Steps, float64(t.Personal.Height))

	meanSpeed := spentenergy.MeanSpeed(t.Steps, float64(t.Personal.Height), t.Duration)
	var (
		errInvalidTriningType = errors.New("неизвестный тип тренировки")
	)
	switch t.TrainingType {
	case "Бег":
		runCal, err := spentenergy.RunningSpentCalories(t.Steps, float64(t.Personal.Weight), float64(t.Personal.Height), t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distant, meanSpeed, runCal), nil

	case "Ходьба":
		walkCal, err := spentenergy.WalkingSpentCalories(t.Steps, float64(t.Personal.Weight), float64(t.Personal.Height), t.Duration)
		if err != nil {
			log.Println(err) // При ошибке логируем и прерываем выполнение, чтобы не обрабатывать некорректные данные
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distant, meanSpeed, walkCal), nil
	default:
		return "", fmt.Errorf("%w: ", errInvalidTriningType)
	}

}
