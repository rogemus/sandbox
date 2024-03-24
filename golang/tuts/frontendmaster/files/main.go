package main

import (
	"files/data"
	sampledata "files/sample_data"
)

var courses []data.Printable

func main() {
	courses = sampledata.GetData()
	for _, c := range courses {
		println(c.Print())
	}
}

