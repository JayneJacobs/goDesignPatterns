package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Scan(d Document)
	Fax(d Document)
}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedMachine interface {
	Print(d Document)
}

func (m OldFashionedPrinter) Print(d Document) {

}

type Printer interface {
	Print(d Document)
}
type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {

}

type MyPrinter interface {
	func (m MyPrinter)Print() 
}

func (m MyPrinter)Print()  {
	
}

type Photocopier struct {

}

func (p Photocopier) Print(d Document)  {
	
}


type MultiFunctionDevice interface {
	Printer
	Scanner
	Faxer
}
// ISP
func main() {

}
