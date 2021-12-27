package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func GetTotalGas(filepath string) (int64, error) {

	csvFile, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}

	defer csvFile.Close()


	r := csv.NewReader(bufio.NewReader(csvFile))
	r.Comma = ','
	var res int64


	index := 0
	for {
		index++
		line, error := r.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("err", err)
			return 0, err
		}
		if index == 1 {
			continue
		}
		gas := line[2]

		val, err  := strconv.ParseInt(gas, 10, 32)
		if err == nil {

			println("gas vl", val)
			res += val;
		}


	}

	return res, nil
}
