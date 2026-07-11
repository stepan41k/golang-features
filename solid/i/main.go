package main

// BAD
// type SmartDevice interface {
// 	Print()
// 	Scan()
// 	Fax()
// }

// type SimplePrinter struct {}

// func (s SimplePrinter) Print() {}
// func (s SimplePrinter) Scan() {}
// func (s SimplePrinter) Fax() {}

// GOOD

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

func PrintDocument(p Printer) {
	p.Print()
}