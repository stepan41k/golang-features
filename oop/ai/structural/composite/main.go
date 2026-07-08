package main

import "fmt"

type Component interface {
	Search(keyword string)
}

type File struct {
	name string
}

func (f *File) Search(k string) {
	fmt.Println("Ищем в файле", f.name)
}

type Folder struct {
	components []Component
	name string
}

func (f *Folder) Search(k string) {
	fmt.Println("Заходим в папку", f.name)
	for _, composite := range f.components {
		composite.Search(k)
	}
}