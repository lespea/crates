package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/lespea/crates"
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
	jo := json.NewEncoder(os.Stdout)
	jo.SetEscapeHTML(false)
	jo.SetIndent("", "  ")

	if false {
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
			jo.Encode(v)
			fmt.Println()
		}
	}

	if true {
		resp, err := http.Get(crates.CrateUrl("atuin").String())
		if err != nil {
			panic(err)
		}

		scn := bufio.NewScanner(resp.Body)

		infos := make(crates.CrateInfos, 0, 30)

		for scn.Scan() {
			var info crates.CrateInfo
			e(json.Unmarshal(scn.Bytes(), &info))
			infos = append(infos, &info)
		}

		infos.ParseVerStrs()

		for _, v := range infos {
			v.Deps = nil
			jo.Encode(v)
			fmt.Println()
		}
	}
}
