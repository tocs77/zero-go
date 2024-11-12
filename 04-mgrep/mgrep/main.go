package main

import (
	"coursecontent/04-mgrep/worker"
	"coursecontent/04-mgrep/worklist"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexflint/go-arg"
)

func discoverDirs(wl *worklist.Worklist, searchDir string) {
	entries, err := os.ReadDir(searchDir)
	if err != nil {
		fmt.Println("Error reading", searchDir, err)
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			nextSearchDir := filepath.Join(searchDir, entry.Name())
			discoverDirs(wl, nextSearchDir)
			continue
		}
		wl.Add(worklist.NewJob(filepath.Join(searchDir, entry.Name())))
	}
}

var args struct {
	SearchTerm string `arg:"positional,required"`
	SearchDir  string `arg:"positional" default:"."`
}

func main() {
	arg.MustParse(&args)
	var workersWg sync.WaitGroup
	wl := worklist.New(100)
	results := make(chan worker.Result, 100)
	numWorkers := 10

	workersWg.Add(1)
	go func() {
		defer workersWg.Done()
		discoverDirs(&wl, args.SearchDir)
		wl.Finalize(numWorkers)
	}()

	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()

			for {
				workEntry := wl.Next()
				if workEntry.Path == "" {
					return
				}

				workerResult := worker.FindInFile(workEntry.Path, args.SearchTerm)
				if workerResult == nil {
					continue
				}

				for _, r := range workerResult.Inner {
					results <- r
				}
			}
		}()
	}
	blockWorkersWg := make(chan struct{})
	go func() {
		workersWg.Wait()
		close(blockWorkersWg)
	}()

	var displayWg sync.WaitGroup
	displayWg.Add(1)
	go func() {
		for {
			select {
			case result := <-results:
				fmt.Printf("%v[%v]:%v\n", result.Path, result.LineNum, result.Line)
			case <-blockWorkersWg:
				if len(results) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()

	displayWg.Wait()
}
