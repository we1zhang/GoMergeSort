package main

import (
	"GoMergeSort2/pipeline"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//inFileName := "small.in"
	inFileName := "large.in"
	//p := createPipeLine(inFileName, 512, 4)
	p := createPipeLine(inFileName, 800000000, 4)

	//outFileName := "small.out"
	outFileName := "large.out"
	WriteToFile(p, outFileName)

	PrintFile(outFileName)
}

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)

	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}

func WriteToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer, p)
}

func createPipeLine(filename string, fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount
	pipeline.Init()
	var sortResults []<-chan int

	for i := 0; i < chunkCount; i++ {

		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		//0表示从头开始
		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))

	}
	return pipeline.MergeN(sortResults...)
}
