package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

)

func main()  {

}

func NewCsvFile(date [][]string)  {
	f,err := os.Create("test.csv")
	if err !=nil {
		println(err)
	}

	defer f.Close()
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	w.WriteAll(date)
	w.Flush()
}

func ReadCsv(fileName string)  {
	fs,err :=os.Open(fileName)
	if err != nil {
		println(err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			println("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
	}
	fmt.Println("\n---------------------------\n")
	// 针对小文件，也可以一次性读取所有的文件。注意，r要重新赋值，因为readall是读取剩下的
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		println("can not readall, err is %+v", err)
	}
	for _, row := range content {
		fmt.Println(row)
	}

}