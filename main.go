package main

import (
	"fmt"
	"os"
	"os/exec"

	flag "github.com/ogier/pflag"
)

var (
	msg string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage:  %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// gitAdd := exec.Command("hub", "add", ".")
	// gitCommit := exec.Command("hub", "commit", fmt.Sprintf("-m '%s'", msg))
	// gitPush := exec.Command("git", "push")
	// var out bytes.Buffer
	// if err := pipe.Command(
	// 	&out,
	// 	gitAdd,
	// 	gitCommit,
	// 	gitPush,
	// 	exec.Command("git", "status"),
	// 	exec.Command("git", "add ."),
	// 	exec.Command("ls"),
	// ); err != nil {
	// 	log.Fatal(err)
	// }

	// if _, err := io.Copy(os.Stdout, &out); err != nil {
	// 	log.Fatal(err)
	// }

	gitAdd := exec.Command("hub", "add", ".")
	gitCommit := exec.Command("hub", "commit", fmt.Sprintf("-m '%s'", msg))

	err := gitAdd.Run()
	err1 := gitCommit.Run()
	if err != nil {
		fmt.Println(err)
	}
	if err1 != nil {
		fmt.Println(err1)
	}
}

func init() {
	flag.StringVarP(&msg, "msg", "m", "new code", "Commit Message")
}
