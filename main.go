package main

import (
	"io"
	"log"
)

func main() {
	log.Println("hello world")
	var fixedWidthFileReader = &FixedWidthFileReader{}
	var result, err1 = fixedWidthFileReader.Read("sample.txt")
	if err1 != nil {
		log.Println(err1)
	}
	log.Println(result)
	log.Println(len(result))

	var myFileUtils = &FileUtils{}
	var fileContent1, err = myFileUtils.OpenFixedWidthFile("sample.txt")
	if err != nil {
		log.Println("Error ", err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := fileContent1.Read(buf)
		log.Println(n)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		if n > 0 {
			log.Println(string(buf[:n]))
		}
	}

}
