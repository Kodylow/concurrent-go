package main

import (
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
    stdOut := os.Stdout

    r, w, + := os.Pipe()
    os.Stdout = w

    var wg sync.WaitGroup
    wg.Add(1)

    go printSomething("alpha", &wg)

    wg.Wait()

    _ = w.Close()

    result, _ := io.ReadAll(r)
    output := string(result)
    
}