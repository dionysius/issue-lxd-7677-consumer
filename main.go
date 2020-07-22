package main

import (
	"fmt"

	lxd "github.com/dionysius/issue-lxd-7677-lib/lxd"
	// uncomment to see what each of these do in the vendor/ folder
	// _ "github.com/dionysius/issue-lxd-7677-lib/lxd/include1"
	// _ "github.com/dionysius/issue-lxd-7677-lib/lxd/include2"
	// _ "github.com/dionysius/issue-lxd-7677-lib/lxd/include3"
	_ "github.com/dionysius/issue-lxd-7677-lib/lxd/implementation"
)

func main() {
	fmt.Printf("Hello %s\n", lxd.Something)
}
