//
//
//
//
//
package main

//
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "github.com/fatih/color"
    "log"
    "os/user"
    "path/filepath"
    "hash/fnv"
    "strconv"
)


//
//
//
type TodoLine struct {

    Id      uint32
    Status  string
    Task    string
    Tags    string

}


//
//
//
//
//
func hash (s string) uint32 {

    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()

}




//
//
// Grabs lines from the text file and returns an array of structs
//
//
func fillarray (filename string) []TodoLine {

    file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Instantiate a new scanner
    scanner := bufio.NewScanner(file)
    todos := []TodoLine{}

    // Loop through file and handle each line
    // TODO: Fill array of todo structs - define common function.
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
        // Append new struct
        line := strings.Split(scanner.Text(), ":")
        i64, err := strconv.ParseUint(line[0], 10, 32)
        if err != nil {
            log.Fatal(err)
        }
        u32 := uint32(i64)
        t := TodoLine {
            Id     : u32,
            Status : line[1],
            Task   : line[2]}
        todos = append(todos, t)
    }

    return todos
}



//
//
//
//
//
func list (tl []TodoLine) {
    fmt.Println("\n")

    // Declare the color printers
    yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
    blue   := color.New(color.FgBlue).SprintFunc()

    // Loop through file and handle each line
    // TODO: Fill array of todo structs - define common function.
    for _,ll := range tl {
        fmt.Printf("%d  %s: %s\n", ll.Id, yellow(ll.Status), blue(ll.Task))
    }
    fmt.Println("\n")
}



//
//
// This is the main function
//
//
func main() {

    // Get the home directory
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }

    // TODO: Figure out how to make this const
    var filename string = filepath.Join(usr.HomeDir, ".todo.txt")

    cmd, arg := Parse()

    // TODO: Open file in each function


    // Operate on commands
    switch cmd {

        case "list":
            list(fillarray(filename))



        case "add":
            file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                log.Fatal(err)
            }
            defer file.Close()
            fmt.Println("Adding:", arg[0])
            line := fmt.Sprintf("%d:TODO:%s\n", hash(arg[0]), arg[0])
            _, err = file.WriteString(line)
            if err != nil {
                panic(err)
            }


        // TODO: Alert when task is already marked as done
        // TODO: Write out to the file
        case "done":
            fmt.Println("Marking", arg[0], "as done")
            ll := fillarray(filename)

            i64, err := strconv.ParseUint(arg[0], 10, 32)
            if err != nil {
                log.Fatal(err)
            }
            id := uint32(i64)
            for i,l := range ll {
                if l.Id == id {
                    fmt.Println("Marking", l.Task, "as done")
                    ll[i].Status = "DONE"
                }
            }
            file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                log.Fatal(err)
            }
            defer file.Close()

            var sb strings.Builder
            for _,l := range ll {
                line := fmt.Sprintf("%d:%s:%s\n", l.Id, l.Status, l.Task)
                sb.WriteString(line)
            }
            _, err = file.WriteString(sb.String())
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
