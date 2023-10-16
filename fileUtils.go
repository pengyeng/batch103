package main

import (
	"os"
)

type FileUtils struct {
}

func (f *FileUtils) OpenFixedWidthFile(filename string) (os.File, error) {
	fixedWidthFile, err := os.Open(filename)
	if err != nil {
		return *fixedWidthFile, err
	}
	//defer fixedWidthFile.Close()
	/*
		buf := make([]byte, 1024)
		for {
			n, err := fixedWidthFile.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				continue
			}
			if n > 0 {
				//log.Println(string(buf[:n]))
			}
		}
	*/
	return *fixedWidthFile, nil
}
