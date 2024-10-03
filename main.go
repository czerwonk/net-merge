package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	out := process(os.Stdin)
	fmt.Print(out)
}

func process(r io.Reader) string {
	cidrs := readToCIDRList(r)
	merged := Merge(cidrs)

	return toOutput(merged)
}

func readToCIDRList(r io.Reader) []net.IPNet {
	res := make([]net.IPNet, 0)

	s := bufio.NewScanner(r)
	for s.Scan() {
		ok, cidr := ParseCIDR(s.Text())
		if ok {
			res = append(res, *cidr)
		}
	}

	return res
}

func toOutput(cidrs []net.IPNet) string {
	sb := &strings.Builder{}

	for _, cidr := range cidrs {
		fmt.Fprintf(sb, "%s\n", cidr.String())
	}

	return sb.String()
}
