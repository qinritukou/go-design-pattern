/**
 * Abstraction
 */
package main

type Computer interface {
	print()
	setPrinter(printer Printer)
}
