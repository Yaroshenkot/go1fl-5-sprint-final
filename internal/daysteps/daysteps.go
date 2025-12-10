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
	//trimmed := strings.TrimSpace(datastring)
	//if trimmed != datastring {
	//	return errors.New("неверный формат строки, ожидается: `шаги, длительность`")
	//}

	//if strings.HasPrefix(datastring, " ") || strings.HasSuffix(datastring, " ") {
	//	return errors.New("неверный формат строки: пробелы в начале или конце строки")
	//}
	//if strings.Contains(datastring, " ,") || strings.Contains(datastring, ", ") {
	//	return errors.New("неверный формат строки, ожидается: `шаги, длительность`")
	//}
	//datastring = strings.TrimSpace(datastring)

	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат строки, ожидается: шаги, длительность")
	}
	stepsStr := strings.TrimSpace(parts[0])
	if stepsStr != parts[0] {
		return errors.New("неверный формат строки, ожидается: шаги, длительность")
	}
	//if stepsStr == "" {
	//return errors.New("количество шагов не может быть пустым")
	//}
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("ошибка преобразования количества шагов: %v", err)
	}

	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	ds.Steps = steps

	durationStr := strings.TrimSpace(parts[1])
	if durationStr != parts[1] {
		return errors.New("неверный формат строки, ожидается: шаги, длительность")
	}
	if strings.Contains(durationStr, " ") {
		return errors.New("ошибка преобразования длительности")
	}
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

	info := fmt.Sprintf("Количество шагов: %d.\n"+"Дистанция составила %.2f км.\n"+"Вы сожгли %.2f ккал.\n", ds.Steps, distance, calories)
	return info, nil
}
