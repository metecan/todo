package services

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Todo struct {
	File  string
	Todos []item
}

type item struct {
	Label       string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

func NewTodo(filename string) Todo {
	return Todo{
		File:  filename,
		Todos: []item{},
	}
}

func (t *Todo) List() {
	headerFmt := color.New(color.FgCyan, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgGreen).SprintfFunc()

	tbl := table.New("#", "Task", "Done", "Created At", "Completed At")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for index, todo := range t.Todos {
		isDone := "❌"
		completed := "Not Yet"

		if todo.Done {
			isDone = "✅"
			completed = todo.CreatedAt.String()
		}
		tbl.AddRow(index+1, todo.Label, isDone, todo.CreatedAt, completed)
	}

	tbl.Print()
}

func (t *Todo) Create(label string) error {
	todo := item{
		Label:       label,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	t.Todos = append(t.Todos, todo)

	return t.save(t.File)
}

func (t *Todo) Complete(index int) error {
	listOfTodos := t.Todos

	if index <= 0 || index > len(listOfTodos) {
		return errors.New("complete: invalid todo number")
	}

	theTask := &listOfTodos[index-1]

	theTask.CompletedAt = time.Now()
	theTask.Done = true

	return t.save(t.File)
}

func (t *Todo) Remove(index int) error {
	listOfTodos := t.Todos

	if index <= 0 || index > len(listOfTodos) {
		return errors.New("complete: invalid todo number")
	}

	t.Todos = append(listOfTodos[:index-1], listOfTodos[index:]...)

	return t.save(t.File)
}

func (t *Todo) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, &t.Todos)
	if err != nil {
		return err
	}

	return t.save(t.File)
}

func (t *Todo) save(filename string) error {
	data, err := json.Marshal(t.Todos)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
