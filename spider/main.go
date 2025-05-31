package main

import (
	"runtime"
	"sync"

	"github.com/scott-mescudi/eximius/spider/internal/parser"
	"github.com/scott-mescudi/eximius/spider/pkg/models"
)



func worker(tasks <-chan models.Document, wg *sync.WaitGroup) {
    defer wg.Done()
    for task := range tasks {
		parser.ParseDocument(&task)
    }
}

func main() {
    tasks := make(chan models.Document, 100_000)
    var wg sync.WaitGroup


    for range runtime.NumCPU() {
        wg.Add(1)
        go worker(tasks, &wg)
    }


    close(tasks) 
    wg.Wait()    
}
