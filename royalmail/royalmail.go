package main

type shipper struct {}

func (s shipper) Name() string {
	return "Royal Mail (RM)"
}

func (s shipper) Currency() string {
	return "GBP"
}

func (s shipper) CalculateRate(weight float32) float32 {
	cost := weight * .9
	tax := cost * .5

	return cost + tax
}

var Shipper shipper