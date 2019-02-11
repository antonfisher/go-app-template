package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/antonfisher/go-app-template/pkg/counter" //TODO replace repository path
	"github.com/antonfisher/go-app-template/pkg/params"  //TODO replace repository path
)

func main() {
	var (
		fVersion = flag.Bool("version", false, "print application version")
		// other application params...
	)

	flag.Parse()

	if *fVersion {
		fmt.Printf("%s@%s-%s\n", params.Name, params.Version, params.Commit)
		os.Exit(0)
	}

	// main function
	c := counter.New()
	c.Start()

	// user interrupts program by pressing Ctrl+C
	var wg sync.WaitGroup
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	wg.Add(1)
	go func() {
		<-signalCh
		doneCh := c.Stop()
		<-doneCh

		fmt.Println("Interrupted")
		wg.Done()
		os.Exit(0)
	}()
	wg.Wait()
}
