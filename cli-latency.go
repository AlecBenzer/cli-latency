package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

var latencyA = flag.Duration("a", 10*time.Millisecond, "First latency to compare")
var latencyB = flag.Duration("b", 100*time.Millisecond, "Second latency to compare")

func readOrINT(r *bufio.Reader, sig <-chan os.Signal) []byte {
	ch := make(chan []byte)
	go func() {
		b, _ := r.ReadBytes('\n')
		ch <- b
	}()
	select {
	case b := <-ch:
		return b
	case <-sig:
		return nil
	}
}

func waitFor(d time.Duration, ch <-chan os.Signal) bool {
	fmt.Print("Start...")
	select {
	case <-time.After(d):
	case <-ch:
		return false
	}
	fmt.Print("finish!")
	return true
}

func main() {
	flag.Parse()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	var correct, wrong int

	latencies := []time.Duration{*latencyA, *latencyB}
	r := bufio.NewReader(os.Stdin)

	for _, l := range latencies {
		fmt.Printf("This is %v (Press Enter):", l)
		if readOrINT(r, ch) == nil {
			return
		}
		if !waitFor(l, ch) {
			return
		}
		time.Sleep(1 * time.Second)
		fmt.Println()
	}

	fmt.Println("\n")

	for {
		chosen := rand.Intn(len(latencies))

		fmt.Println("Press Enter when ready...")
		if readOrINT(r, ch) == nil {
			break
		}

		if !waitFor(latencies[chosen], ch) {
			break
		}

		time.Sleep(1 * time.Second)

		fmt.Println("\n\nHow long was that?")
		for i, l := range latencies {
			fmt.Printf("%c) %v\n", 'A'+i, l)
		}

		b := readOrINT(r, ch)
		if b == nil {
			break
		}
		guess := int(b[0] - 'A')
		if chosen == guess {
			correct += 1
		} else {
			wrong += 1
		}
		fmt.Println()
	}
	fmt.Printf("correct: %d\n", correct)
	fmt.Printf("wrong: %d\n", wrong)
	if correct+wrong != 0 {
		fmt.Printf("\n%d%% accuracy\n", 100.0*correct/(correct+wrong))
	}
}
