package controller

import (
	"fmt"
	"testing"
)

func TestRemoveTopStruct(t *testing.T) {
	mapres := make(map[string]string, 5)
	mapreq := map[string]string{
		"c.sd.casd": "asdasd.aa",
		"c.sd.c":    "aasd.aa",
		"baidu.com": "andy",
	}
	mapres = RemoveTopStruct(mapreq)
	fmt.Println(mapres)

}
