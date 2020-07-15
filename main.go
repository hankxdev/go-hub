package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	flag "github.com/ogier/pflag"
)

var (
	msg string
)

func main() {
	flag.Parse()

	// args := flag.Args()

	if len(flag.Args()) < 1 && flag.NFlag() == 0 {
		fmt.Printf("Usage:  %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	gitAdd := exec.Command("git", "add", ".")
	gitCommit := exec.Command("git", "commit", fmt.Sprintf("-m %s", msg))
	gitPush := exec.Command("git", "push")

	echo("Adding all Changes")
	err := gitAdd.Run()
	if err != nil {
		log.Fatal(err)
	}
	echo(fmt.Sprintf("Committing with message: %s", msg))
	err = gitCommit.Run()
	if err != nil {
		log.Fatal(err)
	}
	echo("Pushing")
	err = gitPush.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Push successfully")
}

func init() {
	flag.StringVarP(&msg, "msg", "m", "new code", "Commit Message")
}

func echo(msg string) {
	fmt.Println(msg)
}
