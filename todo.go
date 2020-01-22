
package main

import (
    "fmt"
    "os"
    "flag"
    "reflect"
)


func usage() {
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


/*

Commands:
    list
    add
    rm
    done
    tag

*/
func parse_args() (string, []string) {

    list_cmd := flag.NewFlagSet("list", flag.ExitOnError)
    add_cmd  := flag.NewFlagSet("add",  flag.ExitOnError)
    rm_cmd   := flag.NewFlagSet("rm",   flag.ExitOnError)
    done_cmd := flag.NewFlagSet("done", flag.ExitOnError)
    tag_cmd  := flag.NewFlagSet("tag",  flag.ExitOnError)


    if len(os.Args) < 2 {
        // print usage
        usage()
        os.Exit(1)
    }

    switch os.Args[1] {

    case "list":
        fmt.Println("list")
        list_cmd.Parse(os.Args[2:])
        fmt.Println(reflect.TypeOf(list_cmd.Args()))
        return "list", list_cmd.Args()
    case "add":
        fmt.Println("add")
        add_cmd.Parse(os.Args[2:])
        fmt.Println(add_cmd.Args())
    case "rm":
        fmt.Println("rm")
        rm_cmd.Parse(os.Args[2:])
        fmt.Println(rm_cmd.Args())
    case "done":
        fmt.Println("done")
        done_cmd.Parse(os.Args[2:])
        fmt.Println(done_cmd.Args())
    case "tag":
        fmt.Println("tag")
        tag_cmd.Parse(os.Args[2:])
        fmt.Println(tag_cmd.Args())
    default:
        fmt.Println("Expected different subcommand")
        // print usage
        usage()
        os.Exit(1)
    }
    return "none", []string{}
}

// This is the main function
func main() {


    cmd, arr := parse_args()


    fmt.Println("cmd: ", cmd)
    fmt.Println("arr: ", arr)



}
