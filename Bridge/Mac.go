/**
 * Refined abstract
 */
package main

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *Mac) setPrinter(p Printer) {
	m.printer = p
}