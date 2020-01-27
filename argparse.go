
package main

import (
    "flag"
    "os"
    "fmt"
)

/*
 *  Ensure that all the command line arguments are formatted in the expected
 *  format. Error out with explanation if the arguments do not match.
 *
 *
 *  Commands:
 *      list <tag>
 *      add "task" <tag>
 *      rm <tag>
 *      done "task"
 *      tag "task" <tag>
 *      clean <tag>
 *
 *  TODO: Figure out why go install does not pick up changes in this file??????
 *
 */
func Parse() (string, []string) {

    listCmd  := flag.NewFlagSet("list",  flag.ExitOnError)

    addCmd   := flag.NewFlagSet("add",   flag.ExitOnError)
    // addCmdTag := addCmd.String("tag", "", "tag")

    rmCmd    := flag.NewFlagSet("rm",    flag.ExitOnError)

    doneCmd  := flag.NewFlagSet("done",  flag.ExitOnError)


    argCount := len(os.Args[1:])

    if argCount < 1 {
        // print usage
        All()
        os.Exit(1)
    }

    var cmd string = "none"
    arg := []string{}

    switch os.Args[1] {

    case "list":
        cmd = "list"
        listCmd.Parse(os.Args[2:])
        arg = listCmd.Args()

    case "add":
        if argCount < 2 {
            Add()
            os.Exit(1)
        }
        cmd = "add"
        addCmd.Parse(os.Args[2:])
        arg = addCmd.Args()

    case "rm":
        if argCount < 2 {
            Rm()
            os.Exit(1)
        }
        cmd = "rm"
        rmCmd.Parse(os.Args[2:])
        arg = rmCmd.Args()

    case "done":
        if argCount < 2 {
            Done()
            os.Exit(1)
        }
        cmd = "done"
        doneCmd.Parse(os.Args[2:])
        arg = doneCmd.Args()

    default:
        fmt.Println("Expected different subcommand than [", os.Args[1], "]")
        // print usage
        All()
        os.Exit(1)
    }
    return cmd, arg
}

// func parse_add(
