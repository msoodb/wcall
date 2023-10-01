package main

import (
	"fmt"
    "os/exec"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		//fmt.Println(f.Name())
        cmd := exec.Command("wc", "-l", f.Name())
        stdout, err := cmd.Output()

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        // Print the output
        fmt.Println(string(stdout))
	}
}
