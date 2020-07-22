package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
	msg     string
	branch  string
	cBranch string
)

var blankRegex = regexp.MustCompile(`\s+`)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 && flag.NFlag() == 0 {
		fmt.Printf("Usage:  %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	flagActions(os.Args[1])
}

func flagActions(arg string) {
	switch arg {
	case "-m", "msg":
		quickPush()
	case "-b", "branch":
		quickBranch()
	case "-c", "checkout":
		quickCheckout()
	}

}

func quickBranch() {
	branch = strings.TrimSpace(branch)
	if branch == "" {
		echo("Empty branch name is not allowed")
		os.Exit(1)
	}
	branch = blankRegex.ReplaceAllString(branch, "-")
	gitBranch := exec.Command("git", "branch", branch)
	gitCheckoutBranch := exec.Command("git", "checkout", branch)
	err := gitBranch.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = gitCheckoutBranch.Run()
	if err != nil {
		log.Fatal(err)
	}
	echo(fmt.Sprintf("create and switched to branch %s", branch))
}

func quickCheckout() {
	gitCheckout := exec.Command("git", "checkout", cBranch)

	err := gitCheckout.Run()
	if err != nil {
		log.Fatal(err)
	}
	echo(fmt.Sprintf("Switched to branch %s", cBranch))
}

func quickPush() {
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
	flag.StringVarP(&msg, "msg", "m", "new code", "add all your changed files, commit with the message, then push")
	flag.StringVarP(&branch, "branch", "b", "new-branch", "create a new branch and switch to it")
	flag.StringVarP(&cBranch, "checkout", "c", "master", "checkout a branch")

}

func echo(msg string) {
	fmt.Println(msg)
}
