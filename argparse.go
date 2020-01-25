
package argparse

import (
    "flag"
    "os"
    "fmt"
)

import (
    "usage"
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
 */
func Parse() (string, []string) {

    listCmd  := flag.NewFlagSet("list",  flag.ExitOnError)

    addCmd   := flag.NewFlagSet("add",   flag.ExitOnError)
    addCmdTag := addCmd.String("tag", "", "tag")

    rmCmd    := flag.NewFlagSet("rm",    flag.ExitOnError)

    doneCmd  := flag.NewFlagSet("done",  flag.ExitOnError)

    tagCmd   := flag.NewFlagSet("tag",   flag.ExitOnError)

    cleanCmd := flag.NewFlagSet("clean", flag.ExitOnError)


    argCount := len(os.Args[1:])

    if argCount < 1 {
        // print usage
        usage.All()
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
            usage.Add()
            os.Exit(1)
        }
        cmd = "add"
        addCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'add'")
        fmt.Println("  tag:", *addCmdTag)
        arg = addCmd.Args()
        for _,v := range arg {
            fmt.Println(v)
        }

    case "rm":
        if argCount < 2 {
            usage.Rm()
            os.Exit(1)
        }
        cmd = "rm"
        rmCmd.Parse(os.Args[2:])
        arg = rmCmd.Args()

    case "done":
        if argCount < 2 {
            usage.Done()
            os.Exit(1)
        }
        cmd = "done"
        doneCmd.Parse(os.Args[2:])
        arg = doneCmd.Args()

    case "tag":
        if argCount < 2 {
            usage.Tag()
            os.Exit(1)
        }
        cmd = "tag"
        tagCmd.Parse(os.Args[2:])
        arg = tagCmd.Args()

    case "clean":
        if argCount < 2 {
            usage.Clean()
            os.Exit(1)
        }
        cmd = "clean"
        cleanCmd.Parse(os.Args[2:])
        arg = cleanCmd.Args()

    default:
        fmt.Println("Expected different subcommand than [", os.Args[1], "]")
        // print usage
        usage.All()
        os.Exit(1)
    }
    return cmd, arg
}

// func parse_add(
