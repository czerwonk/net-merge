package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	out := process(os.Stdin)
	fmt.Print(out)
}

func process(r io.Reader) string {
	cidrs := readToCIDRList(r)
	merged := merge(cidrs)

	return toOutput(merged)
}

func readToCIDRList(r io.Reader) []net.IPNet {
	res := make([]net.IPNet, 0)

	s := bufio.NewScanner(r)
	for s.Scan() {
		ok, cidr := parseCIDR(s.Text())
		if ok {
			res = append(res, *cidr)
		}
	}

	return res
}

func toOutput(cidrs []net.IPNet) string {
	out := ""

	for _, cidr := range cidrs {
		out += fmt.Sprintf("%s\n", cidr.String())
	}

	return out
}
