package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"plugin"
	"strconv"
)

type Shipper interface {
	Name() string
	Currency() string
	CalculateRate(weight float32) float32
}

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		pluginName := args[0]
		weight, _ := strconv.ParseFloat(args[1], 32)

		// Load the plugin
		// 1. Search the plugins directory for a file with the same name as the pluginName
		// that was passed in as an argument and attempt to load the shared object file.
		plug, err := plugin.Open(fmt.Sprintf("plugins/%s.so", pluginName))
		if err != nil {
			log.Fatal(err)
		}

		// 2. Look for an exported symbol such as a function or variable
		// in our case we expect that every plugin will have exported a single struct
		// that implements the Shipper interface with the name "Shipper"
		shipperSymbol, err := plug.Lookup("Shipper")

		if err != nil {
			log.Fatal(err)
		}

		// 3. Attempt to cast the symbol to the Shipper
		// this will allow us to call the methods on the plugins if the plugin
		// implemented the required methods or fail if it does not implement it.
		var shipper Shipper
		shipper, ok := shipperSymbol.(Shipper)

		if !ok {
			log.Fatal("Invalid shipper type")
		}

		// 4. If everything is ok from the previous assertions, then we can proceed
		// with calling the methods on our shipper interface object
		rate := shipper.CalculateRate(float32(weight))
		rate1Day := fmt.Sprintf("%.2f %s", rate, shipper.Currency())

		rate2Days := fmt.Sprintf("%.2f %s",
			rate - (rate * .20),
			shipper.Currency())

		rate7Days := fmt.Sprintf("%.2f %s",
			rate - (rate * .70),
			shipper.Currency())

		table := tablewriter.NewWriter(os.Stdout)

		fmt.Println(shipper.Name())
		table.SetHeader([]string{"Number of Days", "Rate"})

		table.Append([]string{"1 Day Express", rate1Day})
		table.Append([]string{"2 Days Shipping", rate2Days})
		table.Append([]string{"7 Days Shipping", rate7Days})
		table.Render()
	}
}