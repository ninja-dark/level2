package fileio

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"hm8/bytecompare"
)

type FileInfo interface {
	IsDir() bool
	Size() int64
}


type WalkFn = func(path string, info FileInfo, err error) error

type FileIO interface {
	Delete(string) error
	FilesBytesAreEqual(string, string) bool
	MD5HashFile(string, int64) (string, error)
	Walk(string, WalkFn) error
}
type FS struct{}

func readfile(path string) []byte {
	if data, err := ioutil.ReadFile(path); err == nil {
		return data
	}
	return nil
}

func (FS) Delete(path string) error {
	return os.Remove(path)
}

func (FS) FilesBytesAreEqual(path1 string, path2 string) bool {
	b1 := readfile(path1)
	b2 := readfile(path2)
	return bytecompare.BytesAreEqual(b1, b2)
}

func (FS) MD5HashFile(path string, hashSize int64) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return bytecompare.MD5Hash(file, hashSize), nil
}

func (FS) Walk(root string, walkFn WalkFn) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error{
		return walkFn(path, info, err)
	})
}