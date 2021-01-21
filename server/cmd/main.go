package main

import (
	"fmt"
	hashes "server/internal/hash"
)

func main() {
	fmt.Printf("%s", hashes.GetHash("lalalala"))
}
