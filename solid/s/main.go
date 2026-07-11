package main

import "os"

// BAD
// type Report struct {
// 	Content string
// }

// func (r *Report) GenerateData() {
// 	r.Content = "Data"
// }

// func (r *Report) SaveToFile(filename string) {
// 	os.WriteFile(filename, []byte(r.Content), 0644)
// }

// GOOD
type Report struct {
	Content string
}

func (r *Report) GenerateData() {
	r.Content = "Data"
}

type ReportSaver struct {}

func (r *ReportSaver) Save(filename string, data string) {
	os.WriteFile(filename, []byte(data), 0644)
}