/**
 * Concrete implementation
 */
package main

import "fmt"

type Epson struct {
}

func (p *Epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}
