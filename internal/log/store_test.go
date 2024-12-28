package log

import (
       "io/ioutil"
       "os"
       "testing"
       "github.com/stretchr/testify/requre"
)

var (
	write = []byte("hello world")
	width = uint64(len(write)) + lenWidth
)

func TestStoreAppendRead(t *testing.T) {
     	f,err := ioutil.TempFile("",
}

