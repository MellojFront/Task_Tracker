package main

func main() {
	todos := Todos{}
	todos.add("Купить продукты")
	todos.add("Прочитать книгу")
	todos.toogle(0)
	todos.print()

}
