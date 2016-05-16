// Pointers project main.go
package main

import (
	"fmt"
)

var shippingDays = 30
var shippingDaysPtr = new(int)

func main() {
	shippingDaysAdjustments(shippingDays)
	fmt.Println("after shippingDaysAdjustments call", shippingDays)

	shippingDaysAdjustmentsPtr(&shippingDays)
	fmt.Println("after shippingDaysAdjustmentsPtr call", shippingDays)
	fmt.Println("after shippingDaysAdjustmentsPtr call", &shippingDays)

	shipper := shipper{}
	shipper.id = 400
	shipper.name = "Pacific Shippers"
	shipperUpdates(&shipper)
	fmt.Println("shipper.id after shipperUpdates call", shipper.id)
	fmt.Println("shipper.name after shipperUpdates call", shipper.name)

	*shippingDaysPtr = 55
	shippingDaysAdjustmentsPtr(shippingDaysPtr)
	fmt.Println("shippingDaysPtr after shippingDaysAdjustmentsPtr call", *shippingDaysPtr)
}
func shippingDaysAdjustments(shippingDays int) {
	shippingDays += 10
}

func shippingDaysAdjustmentsPtr(shippingDays *int) {
	*shippingDays += 10
}

func shipperUpdates(s *shipper) {
	s.id += 10
	s.name += " Inc."
}

type shipper struct {
	name string
	id   int
}
