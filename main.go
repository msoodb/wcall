package main

import (
    "os"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "bytes"
    "bufio"
)

func main() {
    //help()
    files, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        /*cmd := exec.Command("wc", "-l", f.Name())
        stdout, err := cmd.Output()

        if err != nil {
            continue
        }*/
	file, err := os.Open(f.Name())
	if err != nil{
            continue
	}
	line, err := lineCounter(bufio.NewReader(file))
	if err != nil{
	    continue
	}
	fmt.Println(" ", f.Name(), " ",  line)
    }
}

func help() {
	fmt.Println("Usage:")
	fmt.Println(" ./wcall [flags]")
}


func lineCounter(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
}

