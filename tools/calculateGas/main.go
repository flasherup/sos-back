package main

import . "github.com/flasherup/sos-back/tools/calculateGas/csv"

func main() {
	gas, err := GetTotalGas("export-GasUsed.csv")
	if err != nil {
		println("error", err)
		return
	}
	println("gas", gas)
}
