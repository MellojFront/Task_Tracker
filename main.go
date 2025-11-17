package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.toogle(1)
	todos.print()
	storage.Save(todos)

}
