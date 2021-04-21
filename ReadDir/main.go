package main

import (
	"fmt"
	"strings"
)

func main() {
	vnfdYamlLocalPath := "/tmp/vnfd/simple_scalable.vnfd.scalable.tosca.yaml"

	path := strings.SplitAfterN(vnfdYamlLocalPath, "/", 0)
	fmt.Println(path)

	return

}
