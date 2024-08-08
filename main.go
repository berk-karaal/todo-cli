package main

import (
	"github.com/berk-karaal/todo-cli/cmd"
	_ "github.com/berk-karaal/todo-cli/cmd/database"
)

func main() {
	cmd.Execute()
}
