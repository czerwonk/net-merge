package main

import (
	"bufio"
	"fmt"
	"sort"
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
172.24.0.1
192.168.2.0/24
192.168.0.0/16
`
	expected := `172.24.0.1/32
192.168.0.0/16
2001:678:1e0:100::/56
2001:678:1e0:200::2/128
2001:678:1e0::/64
`

	r := strings.NewReader(input)
	out := process(r)
	out = sortLines(out)

	assert.Equal(t, expected, out)
}

func sortLines(s string) string {
	r := strings.NewReader(s)
	sc := bufio.NewScanner(r)

	lines := []string{}
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	sort.Strings(lines)

	out := ""
	for _, l := range lines {
		out += fmt.Sprintf("%s\n", l)
	}

	return out
}
