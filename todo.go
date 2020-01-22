
package main

import (
    "fmt"
    "os"
    "flag"
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
    list <tag>
    add "task" <tag>
    rm <tag>
    done "task"
    tag "task" <tag>

*/
func parse_args() (string, []string) {

    list_cmd  := flag.NewFlagSet("list",  flag.ExitOnError)
    add_cmd   := flag.NewFlagSet("add",   flag.ExitOnError)
    rm_cmd    := flag.NewFlagSet("rm",    flag.ExitOnError)
    done_cmd  := flag.NewFlagSet("done",  flag.ExitOnError)
    tag_cmd   := flag.NewFlagSet("tag",   flag.ExitOnError)
    clean_cmd := flag.NewFlagSet("clean", flag.ExitOnError)


    if len(os.Args) < 2 {
        // print usage
        usage()
        os.Exit(1)
    }

    var cmd string = "none"
    arg := []string{}

    switch os.Args[1] {

    case "list":
        cmd = "list"
        list_cmd.Parse(os.Args[2:])
        arg = list_cmd.Args()
    case "add":
        cmd = "add"
        add_cmd.Parse(os.Args[2:])
        arg = add_cmd.Args()
    case "rm":
        cmd = "rm"
        rm_cmd.Parse(os.Args[2:])
        arg = rm_cmd.Args()
    case "done":
        cmd = "done"
        done_cmd.Parse(os.Args[2:])
        arg = done_cmd.Args()
    case "tag":
        cmd = "tag"
        tag_cmd.Parse(os.Args[2:])
        arg = tag_cmd.Args()
    case "clean":
        cmd = "clean"
        clean_cmd.Parse(os.Args[2:])
        arg = clean_cmd.Args()
    default:
        fmt.Println("Expected different subcommand than [", os.Args[1], "]")
        // print usage
        usage()
        os.Exit(1)
    }
    return cmd, arg
}

// This is the main function
func main() {


    cmd, arg := parse_args()


    fmt.Println("cmd: ", cmd)
    fmt.Println("arg: ", arg)



}
