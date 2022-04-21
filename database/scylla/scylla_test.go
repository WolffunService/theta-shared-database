package scylla

import (
	"fmt"
	"testing"
)

func TestScylla(t *testing.T) {
	ses, err := NewSession("thetan", "localhost")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("connect success")
	defer func() {
		//fmt.Println(ses.Closed())
		ses.Close()
		//fmt.Println(ses.Closed())
	}()
	//wait
}
