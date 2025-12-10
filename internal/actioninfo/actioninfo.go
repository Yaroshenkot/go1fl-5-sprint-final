package actioninfo

import "fmt"

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Printf("Ошибка парсинга данных `%s`: %v\n", data, err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка получения информации: %v\n", err)
			continue
		}
		fmt.Println(info)
	}
}
