package crates

import (
	"net/url"

	"github.com/hashicorp/go-version"
)

type Dep struct {
	Name            string
	Req             string
	Features        []string
	Optional        bool
	DefaultFeatures bool
	Kind            string
	Target          *string
	Registry        *string
	Package         *string
}

type CrateInfo struct {
	Name          string
	Vers          string
	ParsedVersion *version.Version
	Cksum         string
	Deps          []Dep
	Features      map[string][]string
	Yanked        bool
	Links         *string
	V             uint32
	Features2     map[string][]string
	RustVersion   string
}

func (ci *CrateInfo) ParseVerStr() {
	if v, err := version.NewSemver(ci.Vers); err != nil {
		panic(err)
	} else {
		ci.ParsedVersion = v
	}
}

type CrateInfos []*CrateInfo

func (cis CrateInfos) ParseVerStrs() {
	for _, ci := range cis {
		ci.ParseVerStr()
	}
}

func CrateUrl(crate string) *url.URL {
	u, err := url.Parse(`https://index.crates.io/`)
	if err != nil {
		panic(err)
	}

	parts := make([]string, 0, 3)

	switch len(crate) {
	case 1:
		parts = append(parts, "1")

	case 2:
		parts = append(parts, "2")

	default:
		last := len(crate)
		if last > 4 {
			last = 4
		}

		parts = append(parts, crate[:2], crate[2:last])
	}

	parts = append(parts, crate)

	return u.JoinPath(parts...)
}
