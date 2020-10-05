# Go Plugins Example: Shipping Calculator

This is the source code for [How To Build Extensible Go Applications with Plugins
](https://blog.manimuridi.com/2020/10/05/how-to-build-extensible-go-applications-with-plugins/)

### Clone this project
```console
$ git clone https://github.com/ManiMuridi/go-plugins-shipping-calculator.git
```

### Plugin Interface
```go
// Your implemented interface must comply with this interface
type Shipper interface {
	Name() string
	Currency() string
	CalculateRate(weight float32) float32
}
```

### Add New Plugin Implementation

```go
package main

type shipper struct {}

func (s shipper) Name() string {
	return "United States Postal Service (USPS)"
}

func (s shipper) Currency() string {
	return "USD"
}

func (s shipper) CalculateRate(weight float32) float32 {
	cost := weight * 1.5
	tax := cost * .10

	return cost + tax
}

var Shipper shipper
```
### Build the Plugin Manually
```console
$ go build -buildmode=plugin -o ./plugins/usps.so usps/usps.go
```

### Build the Using the Makefile
```makefile
# Update the Makefile to look like this
.PHONY: usps fedex royalmail

usps:
	go build -buildmode=plugin -o ./plugins/usps.so usps/usps.go

fedex:
	go build -buildmode=plugin -o ./plugins/fedex.so fedex/fedex.go

royalmail:
	go build -buildmode=plugin -o ./plugins/royalmail.so royalmail/royalmail.go
```
Then call the make usps command in your terminal
```console
$ make usps
```

### Run the app
```console
# command: go run main.go {shipping method} {weight}
$ go run main.go usps 5
```