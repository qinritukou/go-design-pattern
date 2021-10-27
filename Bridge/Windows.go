/**
 * Refined abstraction
 */
package main

import "fmt"

type Windows struct {
	printer Printer
}

func (w *Windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *Windows) setPrinter(p Printer) {
	w.printer = p
}
