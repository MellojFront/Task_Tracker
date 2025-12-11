package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func printUsage() {
	fmt.Println("Трекер задач - простое CLI-приложение для управления вашими задачами.")
	fmt.Println("\nПРИМЕРЫ:")
	fmt.Println(`  go run . -add="Новая задача"`)
	fmt.Println(`  go run . -edit="0:Обновленный заголовок"`)
	fmt.Println("  go run . -toggle=0")
	fmt.Println("  go run . -del=0")
	fmt.Println("  go run . -list")
}

func NewCmdFlags() *CmdFlags {
	cf := &CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Добавить новую задачу")
	flag.StringVar(&cf.Edit, "edit", "", "Редактировать задачу, формат: 'id:Новый заголовок'")
	flag.IntVar(&cf.Del, "del", -1, "Удалить задачу по ID")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Переключить статус задачи по ID")
	flag.BoolVar(&cf.List, "list", false, "Вывести список всех задач")

	flag.Usage = printUsage

	flag.Parse()
	return cf

}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Ошибка, проверьте формат команды. Пожалуйста используйте id: new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Ошибка: неверный формат индекса задачи")
			os.Exit(1)
		}

		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		// Если не указано ни одной команды, по умолчанию выводим список задач.
		// flag.NFlag() возвращает количество флагов, установленных в командной строке.
		if flag.NFlag() == 0 {
			todos.print()
		}
	}
}
