package crates

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/hashicorp/go-version"
)

func Test_parseInstallName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str     string
		name    string
		verStr  string
		urlStr  string
		wantErr bool
	}{
		{
			"ast-grep 0.12.5 (registry+https://github.com/rust-lang/crates.io-index)",
			"ast-grep",
			"0.12.5",
			"registry+https://github.com/rust-lang/crates.io-index",
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			wantVer, err := version.NewVersion(tt.verStr)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			wantUrl, err := url.Parse(tt.urlStr)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			name, ver, u, err := parseInstallName(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInstallName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if name != tt.name {
				t.Errorf("parseInstallName() got = %v, want %v", name, tt.name)
			}
			if !reflect.DeepEqual(wantVer, ver) {
				t.Errorf("parseInstallName() got1 = %v, want %v", ver, wantVer)
			}
			if !reflect.DeepEqual(wantUrl, u) {
				t.Errorf("parseInstallName() got2 = %v, want %v", u, wantUrl)
			}
		})
	}
}
