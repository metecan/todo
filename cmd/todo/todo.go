package main

import (
	"flag"
	"fmt"
	"os"
	"todo-cli/services"
	"todo-cli/utils"
)

const (
	todoFile = "./data.json"
)

func main() {

	list := flag.Bool("list", false, "list all todos")
	create := flag.Bool("create", false, "create a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	remove := flag.Int("remove", 0, "remove a todo")
	flag.Parse()

	todos := services.NewTodo(todoFile)
	err := todos.Load(todoFile)
	utils.Error(err)

	switch {

	case *list:
		todos.List()

	case *create:
		label, err := utils.ReadInput(os.Stdin, flag.Args()...)
		utils.Error(err)

		err = todos.Create(label)
		utils.Error(err)

		todos.List()
	case *complete > 0:
		err := todos.Complete(*complete)
		utils.Error(err)

		todos.List()
	case *remove > 0:
		err := todos.Remove(*remove)
		utils.Error(err)

		todos.List()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
