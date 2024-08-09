# todo-cli

A cli todo application designed for daily todos.

[![asciicast](https://asciinema.org/a/MyHsZgQfVaW6eBgIoQkvBHVBD.svg)](https://asciinema.org/a/MyHsZgQfVaW6eBgIoQkvBHVBD)

## Install

```
$ go install github.com/berk-karaal/todo-cli@latest
```

This will add `todo-cli` binary to your GOBIN directory, renaming this binary to
`todo` could be efficient for use. (You can find the binary location with
`which todo-cli` command on Linux.)

## Basic Usage

`todo help` to display help text contains commands, description and usages.

`todo new "Todo text"` to create new todo.

`todo list` to list todos created today.<br>
`todo list -d 1` to list todos created yesterday.

`todo set TODO_ID STATUS` to update status of a todo. See available status
values with `todo help set`.

`todo remove TODO_ID` to remove a todo.

`todo rename TODO_ID NEW_NAME` rename a todo.

You can check other commands and details of above commands with `todo help`.  
