package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"flag"
)


type selpg struct {
	start int
	end int
	length int
	file string
	f bool
	printDest string
}

//init function
func initSelpg(Selpg *selpg) {
	Selpg.start = -1
	Selpg.end = -1
	Selpg.length = 72
	Selpg.f = false
	Selpg.file = ""
	Selpg.printDest = ""
}

type error_type int

const (
	invalid_s	error_type = iota
	invalid_e
	invalid_e_s
	invalid_length
)

//error message
var err_mesg = map[error_type]string {
	invalid_s: "start should be a positive number\n",
	invalid_e: "end should be a positive number\n",
	invalid_e_s: "start should be less than end\n",
	invalid_length: "length should be a positive number\n",
}

//choose what to do
func process(Selpg *selpg) {
	if Selpg.start < 1 {
		fmt.Printf("%s", err_mesg[invalid_s])
		return
	}
	if Selpg.end < 1 {
		fmt.Printf("%s", err_mesg[invalid_e])
		return
	}
	if Selpg.length < 1 {
		fmt.Printf("%s", err_mesg[invalid_length])
		return
	}
	if Selpg.start > 0 && Selpg.end > 0 && Selpg.start > Selpg.end {
		fmt.Printf("%s", err_mesg[invalid_e_s])
		return
	}
	if Selpg.printDest != "" {
		outputFile(Selpg)
	} else {
		readFile(Selpg)
	}
}

//read datas from a file and print
func readFile(Selpg *selpg) {
	inputFile, inputError := os.Open(Selpg.file)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the file\n")
		return
	}
	defer inputFile.Close()

	i := 1
	min := (Selpg.start - 1) * Selpg.length + 1
	max := Selpg.end * Selpg.length
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if i >= min && i <= max {
			fmt.Printf("%s", inputString)
		}
		i++
		if readerError == io.EOF {
			return
		}
	}
}

//copy datas from one file to another file
func outputFile(Selpg *selpg) {
	//check the destionation if exists
	outputFile, outputError := os.OpenFile(Selpg.printDest, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creating\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	//open the file
	inputFile, inputError := os.Open(Selpg.file)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the file\n")
		return
	}
	defer inputFile.Close()

	//if the data was included in the range
	//write it to the destionation file
	i := 1
	min := (Selpg.start - 1) * Selpg.length + 1
	max := Selpg.end * Selpg.length
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if i >= min && i <= max {
			outputWriter.WriteString(inputString)
		}
		i++
		if readerError == io.EOF {
			fmt.Printf("successfully write to " + Selpg.printDest)
			break
		}
	}

	outputWriter.Flush()
}

func main()  {
	var Selpg selpg
	initSelpg(&Selpg)

	//selpg -s 1 -e 2 -l 3 file
	flag.IntVar(&Selpg.start, "s", -1, "start line")
	flag.IntVar(&Selpg.end, "e", -1, "end line")
	flag.IntVar(&Selpg.length, "l", 72, "page length")
	flag.BoolVar(&Selpg.f, "f", false, "seperate pages")
	flag.StringVar(&Selpg.printDest, "d", "", "print name")
	flag.Parse()

	if len(flag.Args()) > 0 {
		Selpg.file = flag.Args()[0]
	}

	process(&Selpg)

	os.Exit(0)
}
