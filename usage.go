
package main

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
