package helpers

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// 生成节点实例
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println(node.GetId())
	}
}
