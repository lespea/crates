package crates

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/hashicorp/go-version"
)

type InstalledInfo struct {
	Name              string
	Version           *version.Version
	Url               *url.URL
	Bins              []string
	Features          []string
	AllFeatures       bool
	NoDefaultFeatures bool
	Profile           string
	Target            string
	Rustc             string
}

type InstallMap map[string]InstalledInfo

type Installs struct {
	Installs InstallMap
}

func (i Installs) ParseInfos() (InstallMap, error) {
	installed := make(InstallMap, len(i.Installs))

	for crateStr, infos := range i.Installs {
		newInfo := infos

		if name, ver, u, err := parseInstallName(crateStr); err != nil {
			return nil, err
		} else {
			newInfo.Name = name
			newInfo.Version = ver
			newInfo.Url = u
			installed[newInfo.Name] = newInfo
		}
	}

	return installed, nil
}

var inameRex = regexp.MustCompile(
	`^(\S+) (v?\d\S+) \((\S+)\)$`,
)

func parseInstallName(crateStr string) (string, *version.Version, *url.URL, error) {
	parts := inameRex.FindStringSubmatch(crateStr)
	if len(parts) != 4 {
		return "", nil, nil, fmt.Errorf("unknown crate entry: '%s'", crateStr)
	}

	name := parts[1]

	if name == "" {
		return "", nil, nil, fmt.Errorf("empty name: '%s'", crateStr)
	}

	ver, err := version.NewVersion(parts[2])
	if err != nil {
		return "", nil, nil, fmt.Errorf("bad version: '%s'; %w", crateStr, err)
	}

	u, err := url.Parse(parts[3])
	if err != nil {
		return "", nil, nil, fmt.Errorf("bad url: '%s'; %w", crateStr, err)
	}

	return name, ver, u, nil
}
