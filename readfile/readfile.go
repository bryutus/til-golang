package readfile

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func readfile(filepath string) {
	start := time.Now()

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true
	var line []string

	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
		fmt.Println(line)
	}

	end := time.Now()
	fmt.Printf("%fs\n", (end.Sub(start)).Seconds())
}
