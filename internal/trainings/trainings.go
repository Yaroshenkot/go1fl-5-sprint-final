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

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат строки, ожидается: шаги, тип_тренеровки, длительность")
	}
	StepsStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(StepsStr)
	if err != nil {
		return fmt.Errorf("ошибка преобразования количества шагов: %v", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	t.Steps = steps

	trainingType := strings.TrimSpace(parts[1])
	t.TrainingType = trainingType
	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %v", err)
	}
	if duration <= 0 {
		return errors.New("длительность должна быть положительной")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch strings.ToLower(t.TrainingType) {
	case "бег", "run", "running":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "ходьба", "walk", "walking":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %v", err)
	}
	info := fmt.Sprintf("Тип тренировки: %s\n"+"Длительность: %.2f ч.\n"+"Дистанция: %.2f км.\n"+"Скорость: %.2f км/ч\n"+"Сожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, calories)
	return info, nil
}
