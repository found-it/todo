//
package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    // "io/ioutil"
    "github.com/fatih/color"
    "log"
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

    const filename string = ".todo.txt"

    yellow := color.New(color.FgYellow).SprintFunc()
    cyan   := color.New(color.FgCyan).SprintFunc()

    cmd, arg := Parse()

    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    switch cmd {

    case "list":
        scanner := bufio.NewScanner(file)
        success := true
        for success {
            success = scanner.Scan()
            if success == false {
                if scanner.Err() != nil {
                    log.Fatal(scanner.Err())
                } else {
                    break
                }
            }

            ll := strings.Split(scanner.Text(), ":")
            fmt.Printf("%s: %s\n", yellow(ll[0]), cyan(ll[1]))
        }


    case "add":
        line := "TODO:" + arg[0] + "\n"
        _, err := file.WriteString(line)
        if err != nil {
            panic(err)
        }

    // case "done":
    //     txt, err := ioutil.ReadFile(


    default:
        fmt.Println("Expected different subcommand than [", os.Args[1], "]")
        // print usage
        All()
        os.Exit(1)
    }

}
