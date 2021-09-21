package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	s "strings"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

func plotGraph(values plotter.XYs){
	p := plot.New()
	p.Add(plotter.NewGrid())
	s, err := plotter.NewScatter(values)
	if err != nil{
		panic(err)
	}
	p.Add(s)
	if err := p.Save(3*vg.Inch, 3*vg.Inch, "graph.png"); err != nil{
		panic(err)
	}
}

//parallelizes ping command using go routine
func main() {
	urlList := os.Args[1:]//list of urls defined by user inputted args

	out := make(chan string, len(urlList))

	fmt.Println("GOMAXPROCS max value", maxCPU())

	//this tests runtime of the program with default value of GOMAXPROCS up to max value of GOMAXPROCS
	gmpMax := maxCPU()
	var durationList = make(map[int]int64)//list to store and display the duration of the pings according to the value of GOMAXPROCS
	for i := 0; i <= gmpMax; i++ {
		fmt.Println("\nTesting GOMAXPROCS value = ", i)
		fmt.Println("url/min/avg/max/stddev") //to clarify what the output means
		runtime.GOMAXPROCS(i)
		start := time.Now()
		for range urlList {
			go ping(out, urlList)
		}
		duration := time.Since(start)
		durationList[i] = duration.Milliseconds()
		fmt.Println("time: ", duration)
	}
	fmt.Println("\nFinal list of durations in milliseconds for the code according to the value of GOMAXPROCS")
	fmt.Println(durationList)
	values := make(plotter.XYs, len(durationList))
	var i = 0
	for k, v := range durationList {
		values[i].X = float64(k)
		values[i].Y = float64(v)
		fmt.Println(values[i])
		i++
	}
	plotGraph(values)
}
