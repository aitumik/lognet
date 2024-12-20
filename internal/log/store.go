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

	w,err := s.buf.Write(p)
	if err != nil {
	       return 0,0,err
	}

	w += lenWidth

	s.size += uint64(w)
	
	return uint64(w),pos,nil
}


func (s *store) Read(pos uint64) ([]byte,err) {
     	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
	       return nil,err
	}

	size := make([]byte,lenWidth)

	// are we reading from s.File to size or what. What is
	// the documentation for this crap

	if _,err := s.File.ReadAt(size,int64(pos)); err != nil {
	   	 return nil,err
	}

	b := make([]byte,lenWidth)
	if _,err := s.File.ReadAt(b,enc.Uint64(s.pos + lenWidth)); err != nil {
	   	 return nil,err
	}

	return b
}

func (s *store) ReadAt(p []byte,pos uint64) (int,error) {
     	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.File.Flush(); err != nil {
	   return 0,err
	}

	
	return s.File.ReadAt(p,pos)
}

func (s *store) Close() error {
     	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.buf.Flush()
	if err != nil {
	       return err
	}

	return s.File.Close()
}



