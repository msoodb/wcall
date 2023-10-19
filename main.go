package main

import (
    "fmt"
    "os/exec"
    "io/ioutil"
    "log"
)

func main() {
    help()
    files, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        cmd := exec.Command("wc", "-l", f.Name())
        stdout, err := cmd.Output()

        if err != nil {
            //fmt.Println(err.Error())
            continue
        }
        fmt.Println(string(stdout))
    }
}

func help() {
	fmt.Println("Usage:")
	fmt.Println(" ./wcall [flags]")
}

