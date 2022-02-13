package main

import (
	"os"
    "os/exec"
	"path/filepath"
	"strings"

	"github.com/zserge/lorca"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func openWithLorca(url string) {
    ui, err := lorca.New(url, "", 800, 600)
    checkError(err)
    defer ui.Close()
    <-ui.Done()
}

func openWithCmd(name string, args ...string) {
    cmd := exec.Command(name, args...)
    err := cmd.Start()
    checkError(err)
}

func main() {
    executable, err := os.Executable()
    checkError(err)

    input := filepath.Base(executable)
    input = strings.TrimSuffix(input, ".sm")
    input = strings.TrimSpace(input)

    if input == "smartfile" {
        openWithLorca("https://www.bxtkezhan.xyz/post/project-smartfile/")
        return
    }

    cmds := strings.SplitN(input, ":", 2)
    switch cmds[0] {
    case "https": openWithLorca(strings.ReplaceAll(input, "\\", "/"))
    case "gedit", "firefox", "gvim": openWithCmd(cmds[0], cmds[1:]...)
    default:
    }
}
