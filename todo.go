package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Структура задачи
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	UpdatedAt   *time.Time
}

// Список задач
type Todos []Todo

// Добавляет новую задачу в список
func (todos *Todos) add(Title string) {
	todo := Todo{
		Title:       Title,
		Completed:   false,
		CompletedAt: nil,
		UpdatedAt:   nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}

// Проверяет, что индекс задачи валиден
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("неверный индекс задачи")
		fmt.Println(err)
		return err
	}
	return nil
}

// Удаляет задачу по индексу
func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

// Редактирует заголовок задачи по индексу
func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	now := time.Now()
	t[index].UpdatedAt = &now

	return nil
}

// Выводит список задач в табличном формате
func (todos *Todos) print() {
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)
	writer.AppendHeader(table.Row{"#", "Задача", "Статус", "Создана", "Выполнена", "Обновлена"})

	for index, t := range *todos {
		completed := "❌"
		completedAt := "-"
		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		updatedAt := "-"
		if t.UpdatedAt != nil {
			updatedAt = t.UpdatedAt.Format(time.RFC1123)
		}

		writer.AppendRow(table.Row{
			strconv.Itoa(index),
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
			updatedAt,
		})
	}
	writer.SetStyle(table.StyleColoredYellowWhiteOnBlack)
	writer.Render()
}

// Переключает статус задачи (выполнена/не выполнена)
func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	now := time.Now()
	t[index].UpdatedAt = &now
	t[index].Completed = !t[index].Completed

	if t[index].Completed {
		t[index].CompletedAt = &now
	} else {
		t[index].CompletedAt = nil
	}

	return nil
}
