package gofinderr

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// getOS return os type
func getOS() (os string) {
	os = runtime.GOOS
	return
}

// Google search error on www.google.com
func Google(e error) {
	os := getOS()

	var (
		cmd *exec.Cmd
		err error
	)

	cmd, err = getCommand(os, e, "google")

	if err != nil {
		log.Println(err.Error())
		return
	}

	if err = cmd.Run(); err != nil {
		log.Println(err.Error())
	}
	return
}

// StackOverflow search error on www.stackoverflow.com
func StackOverflow(e error) {
	os := getOS()

	var (
		cmd *exec.Cmd
		err error
	)

	cmd, err = getCommand(os, e, "stackoverflow")

	if err != nil {
		log.Println(err.Error())
		return
	}

	if err = cmd.Run(); err != nil {
		log.Println(err.Error())
	}
	return
}

// Generate a command line to open browser based on os type
func getCommand(os string, e error, site string) (cmd *exec.Cmd, err error) {

	url := fmt.Sprintf("https://www.%s.com/search?q=%v", site, e.Error())

	switch os {
	case "darwin":
		cmd = exec.Command("open", url)
	case "window":
		cmd = exec.Command("start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		err = fmt.Errorf("unsupport os")
	}
	return
}
