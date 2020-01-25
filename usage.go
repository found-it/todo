
package usage

import (
    "fmt"
)

func All() {
    var usg string = `

Usage:

    todo <command> [arguments]

The commands are:

    list        list tasks
    add         add a task
    rm          remove a task
    done        mark a task as done
    tag         tag a task
    clean       clean out finished tasks

`
    fmt.Println(usg)
}

func Add() {
    var usg string = `

Usage: todo add [--tag <tag>] "<task>"

The --tag flag adds tags to new task.

`
    fmt.Println(usg)
}

func Rm() {
    var usg string = `

Usage: todo rm

`
    fmt.Println(usg)
}

func Done() {
    var usg string = `

Usage: todo done

`
    fmt.Println(usg)
}

func Tag() {
    var usg string = `

Usage: todo tag

`
    fmt.Println(usg)
}

func Clean() {
    var usg string = `

Usage: todo clean

`
    fmt.Println(usg)
}

