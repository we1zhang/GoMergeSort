package main

import (
	"GoMergeSort2/pipeline"
	"bufio"
	"fmt"
	"os"
)

func main() {

	//mergeDemo()

	//const filename = "small.in"
	const filename = "large.in"
	const n = 100000000

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	source := pipeline.RandomSource(n)

	//写到文件中
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, source)
	writer.Flush() //因为用到了bufio，所以需要flush

	file, err = os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	count := 0
	p := pipeline.ReaderSource(bufio.NewReader(file), -1)
	for v := range p {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}

func mergeDemo() {
	//p := pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4))
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))
	//这样输出也行
	//for {
	//	if num, ok := <-p; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}
	for v := range p {
		fmt.Println(v)
	}
}
