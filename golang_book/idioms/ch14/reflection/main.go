package main

import (
	"encoding/csv"
	"fmt"
	csvmarshaller "mymodule/golang_book/idioms/ch14/reflection/csv_marshaller"
	"strings"
)

func main() {
	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var entries []MyData
	err = csvmarshaller.Unmarshall(allData, &entries)
	if err != nil {
		panic(err)
	}
	fmt.Println(entries)
	out, err := csvmarshaller.Marshall(entries)
	if err != nil {
		panic(err)
	}
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	w.WriteAll(out)
	fmt.Println(sb)
}
