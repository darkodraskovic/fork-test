package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "helloworld",
		Short: "A fancy Hello World in Go",
		Run: func(cmd *cobra.Command, args []string) {
			printHello()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func printHello() {
	// Spinner animation in a goroutine
	spinDone := make(chan struct{})
	go spinner(spinDone)

	// Simulate some work
	time.Sleep(2 * time.Second)
	close(spinDone)

	// Colored output
	green := color.New(color.FgGreen).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	// Random greetings
	greetings := []string{"Hello", "Howdy", "Ahoy", "Salut", "Ciao", "Hola"}
	rand.Seed(time.Now().UnixNano())
	greet := greetings[rand.Intn(len(greetings))]

	fmt.Printf("\n%s %s, %s!\n", green(greet), green("World"), bold("Gopher"))

	// System info
	fmt.Println("Running on:", runtime.GOOS, runtime.GOARCH)
	fmt.Println("Go version:", runtime.Version())
}

func spinner(done <-chan struct{}) {
	chars := []rune{'|', '/', '-', '\\'}
	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\r%c Thinking...", chars[i%len(chars)])
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}
}
