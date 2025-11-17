package main

import (
	"encoding/json"
	"os"
)

// Storage представляет собой хранилище данных в файле.
type Storage[T any] struct {
	FileName string
}

// NewStorage инициализирует и возвращает новый экземпляр Storage.
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

// Save сохраняет данные в файл в формате JSON.
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

// Load загружает данные из файла в формате JSON.
func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}
