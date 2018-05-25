package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	csvf, err := os.Open("zxc.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(csvf))

	header, err := r.Read()
	col := 0
	for i, h := range header {
		if "lineItem/UnblendedCost" == h {
			col = i
			break
		}
	}

	sum := 0.0
	for i := 0; ; i++ {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		itemCost, err := strconv.ParseFloat(line[col], 64)
		if err != nil {
			log.Println("error paring item cost at line", i, ":", line[col])
		}
		sum += itemCost
	}

	fmt.Println("Total sum:", sum)
}
