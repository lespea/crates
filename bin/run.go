package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/lespea/crates"
	"os"
)

func e(err error) {
	if err != nil {
		panic(err)
	}
}

func fe(f func() error) {
	e(f())
}

func main() {
	fh, err := os.Open(`/home/adam/.cargo/.crates2.json`)
	e(err)
	defer fe(fh.Close)

	buf := bufio.NewReader(fh)
	j := json.NewDecoder(buf)

	var cr crates.Installs
	e(j.Decode(&cr))

	installs, err := cr.ParseInfos()
	e(err)

	for _, v := range installs {
		fmt.Printf("%+v\n", v)
	}
}
