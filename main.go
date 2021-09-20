package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	s "strings"
)

//maxCPU is a functions that finds the highest possible gomaxprocs value by comparing the number of cpus.
func maxCPU() int {
	maxProcs := runtime.GOMAXPROCS(0) //default value
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

//does the pinging as well as formats the result so that only the url and stats come out as the result
func ping(out chan string, urlList []string) {

	listLength := len(urlList)

	//This loop iterates through the list of urls which were passed in by the user, and running ping each time.
	//Note that this works for the Mac OS shell, and is untested for Windows and Linux.
	for i := 0; i < listLength; i++ {
		go func(i int) {
			tempUrl := urlList[i]
			result, err := exec.Command("ping", "-c 5", tempUrl).Output() //uses shell to ping a url 5 times
			if err != nil {
				fmt.Println(err.Error())
			} else {
				indexOfStats := s.Index(string(result), " = ") //this piece of code finds where the stats are within the output of using ping
				stats := tempUrl + "/" + string(result[indexOfStats+3:])
				out <- stats
			}
		}(i)
	}
}

//parallelizes ping command using go routine
func main() {
	maxProcs := maxCPU()
	runtime.GOMAXPROCS(maxProcs)
	urlList := os.Args[1:]//list of urls defined by user inputted args

	out := make(chan string, len(urlList))
	start := time.Now()

	go ping(out, urlList)

	fmt.Println("url/min/avg/max/stddev") //to clarify what the output means
	for range urlList {
		fmt.Println(<-out)
	}

	duration := time.Since(start)
	fmt.Println("time: ", duration)
}
