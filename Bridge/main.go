/**
 * Client code
 */
package main

import "fmt"

func main() {
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()
}
