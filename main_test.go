package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	input := `2001:678:1e0:0::/64
2001:678:1e0::1
2001:678:1e0:100::/56
2001:678:1e0:110::1/128
2001:678:1e0:200::2/128
`
	expected := `2001:678:1e0:100::/56
2001:678:1e0::/64
2001:678:1e0:200::2/128
`

	r := strings.NewReader(input)
	out := process(r)

	assert.Equal(t, expected, out)
}
