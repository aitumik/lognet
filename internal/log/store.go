package log

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	enc = binary.BigEndian
)

const (
	lenWidth = 8
)

type store struct {
	*os.File
	mu     sync.Mutex
	buf *bufio.Writer
	size   uint64
}

func newStore(file *os.File) (*store,error) {
     	// get the file info
     	fi,err := os.Stat(file.Name())
	if err != nil {
	   return nil,err
	}

	size := uint64(fi.Size())
	if err != nil {
	   return nil,err
	}
	
	return &store{
	       File: file,
	       size: size,
	       buf: bufio.NewWriter(f),
	},nil
}

func (s *store) Append(p []byte) (n uint64,pos uint64,err error) {
     	s.mu.Lock()
	defer s.mu.Unlock()

	pos := s.size

	if err := binary.Write(s.buf,enc,uint64(len(p))); err != nil {
	       return 0,0,err
	}
}
