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
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат строки, ожидается: `шаги, длительность`")
	}
	stepsStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("ошибка преобразования количества шагов: %v", err)
	}

	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	ds.Steps = steps

	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)

	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %v", err)
	}

	if duration <= 0 {
		return errors.New("длительность должна быть положительной")
	}

	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %v", err)
	}

	info := fmt.Sprintf("количество шагов: %d\n"+"Дистанция составила %.2f км.\n"+"Вы сожгли %.2f кал.", ds.Steps, distance, calories)
	return info, nil
}
