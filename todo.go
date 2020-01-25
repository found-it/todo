
package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "github.com/fatih/color"
)

import (
    "usage"
    "argparse"
)

type TodoLine struct {

    Id      int
    Status  string
    Task    string
    Tags    string

}


// func printTodoLine(line TodoLine

// This is the main function
func main() {

    yellow := color.New(color.FgYellow).SprintFunc()
    cyan   := color.New(color.FgCyan).SprintFunc()

    cmd, arg := argparse.Parse()

    switch cmd {

    case "list":
        content, err := ioutil.ReadFile(".todo.txt")
        if err != nil {
            panic(err)
        }
        ll := strings.Split(string(content), ":")
        fmt.Printf("%s: %s\n", yellow(ll[0]), cyan(ll[1]))


    case "add":
        line := "TODO: " + arg[0] + "\n"
        err := ioutil.WriteFile(".todo.txt", []byte(line), 0644)
        if err != nil {
            panic(err)
        }

    default:
        fmt.Println("Expected different subcommand than [", os.Args[1], "]")
        // print usage
        usage.All()
        os.Exit(1)
    }

}
