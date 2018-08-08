package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"syscall"
	"unsafe"
)

const (
	maxBytes = 1<<50 - 1
)

type ReaderAt struct {
	data []byte // io.ReaderAt
	size int64
}

func (r *ReaderAt) ReadAt(readBuf []byte, offset int64) (readBytes int, err error) {
	if r.data == nil {
		return 0, fmt.Errorf("closed reader?")
	}

	if offset < 0 || offset > int64(len(r.data)) {
		return 0, fmt.Errorf("offset=%d is out of range", offset)
	}

	readBytes = copy(readBuf, r.data[offset:])
	if readBytes < len(readBuf) {
		return readBytes, io.EOF
	}

	return readBytes, nil
}

func OpenFileAsMemMapper(filePath string) (*ReaderAt, error) {

	if !exists(filePath) {
		return nil, errors.New("file does not exits")
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("cannot open file")
	}
	defer f.Close()

	// get some info
	fi, err := f.Stat()
	if err != nil {
		return nil, errors.New("cannot Stat() on file")
	}

	size := fi.Size()
	if size == 0 {
		return &ReaderAt{}, nil
	}

	if size < 0 {
		return nil, fmt.Errorf("file size is less than zero")
	}

	// low, high -> should be passed in 8 byte word
	low, high := uint32(size), uint32(size>>32)

	hMap, err := syscall.CreateFileMapping(syscall.Handle(f.Fd()), nil, syscall.PAGE_READONLY, high, low, nil)
	if err != nil {
		return nil, err //fmt.Errorf("call of CreateFileMapping() failed")
	}
	defer syscall.CloseHandle(hMap)

	ptr, err := syscall.MapViewOfFile(hMap, syscall.FILE_MAP_READ, 0 /* h */, 0 /* low */, uintptr(size))
	if err != nil {
		return nil, err
	}

	data := (*[maxBytes]byte)(unsafe.Pointer(ptr))[:size]

	return &ReaderAt{data: data, size: size}, nil
}
