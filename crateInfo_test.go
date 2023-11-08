package crates

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_crateUrl(t *testing.T) {
	t.Parallel()

	for _, ti := range [][2]string{
		{"o", "1/o"},
		{"db", "2/db"},
		{"thr", "th/r/thr"},
		{"four", "fo/ur/four"},
		{"fiver", "fi/ve/fiver"},
	} {
		want, err := url.Parse(`https://index.crates.io/` + ti[1])
		assert.NoError(t, err)

		assert.Equal(t, CrateUrl(ti[0]).String(), want.String())
	}
}
