//
package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "io/ioutil"
    "github.com/fatih/color"
    "log"
    "os/user"
    "path/filepath"
    "os/exec"
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

    // Get the home directory
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }

    // TODO: Figure out how to make this const
    var filename string = filepath.Join(usr.HomeDir, ".todo.txt")

    yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
    blue   := color.New(color.FgBlue).SprintFunc()

    cmd, arg := Parse()

    // TODO: Open file in each function
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


    // Operate on commands
    switch cmd {

        case "list":
            fmt.Println("\n")
            scanner := bufio.NewScanner(file)
            var dones strings.Builder
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
                if ll[0] == "DONE" {
                    dones.WriteString(scanner.Text() + "\n")
                } else {
                    fmt.Printf("%s: %s\n", yellow(ll[0]), blue(ll[1]))
                }
            }
            if len(dones.String()) > 0 {
                fmt.Println("\n")
                fmt.Printf("%s", yellow("Finished Tasks:"))
                fmt.Println("\n")
                ll := strings.Split(dones.String(), "\n")
                for _, done := range ll {
                    if len(done) > 0 {
                        dd := strings.Split(done, ":")
                        fmt.Printf("%s: %s\n", yellow(dd[0]), blue(dd[1]))
                    }
                }

            }
            fmt.Println("\n")


        case "add":
            fmt.Println("Adding:", arg[0])
            line := "TODO:" + arg[0] + "\n"
            _, err := file.WriteString(line)
            if err != nil {
                panic(err)
            }


        // TODO: Alert when task is already marked as done
        // TODO: Use hashes to ID tasks
        case "done":
            fmt.Println("Marking", arg[0], "as done")
            file, err := ioutil.ReadFile(filename)
            if err != nil {
                log.Fatal(err)
            }
            lines := strings.Split(string(file), "\n")

            for i, line := range lines {
                if strings.Contains(line, arg[0]) {
                    fmt.Println("Replacing")
                    strings.Replace(lines[i], "TODO", "DONE", 1)
                }
                fmt.Println(lines[i])
            }
            output := strings.Join(lines, "\n")
            err = ioutil.WriteFile(filename, []byte(output), 0644)
            if err != nil {
                log.Fatal(err)
            }

        default:
            fmt.Println("Expected different subcommand than [", os.Args[1], "]")
            // print usage
            All()
            os.Exit(1)
    }

}
