package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"hm8/fileio"

	"hm8/duplicates"
)

var fs fileio.FileIO

type Filehash struct {
	hash string
	path string
	size int64
}

func FindFilesInPath(rootDir string) <- chan Filehash {
	fileChannel := make(chan Filehash)

	go func() {
		defer close(fileChannel)
		fs.Walk(rootDir, func(path string, into files.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			fileChannel <- Filehash{"", path, info.Size()}
			return nil
		})

	}()

	return fileChannel
}

const fileHashSize = 4096

func emitDuplicates(fileChannel <- chan Filehash, dupesChannel chan<- duplicates.Duplicate) {
	hashMap := make(map[int64][]Filehash)
	for hf := range fileChannel {
		if fileslice, ok := hashMap[hf.size]; ok {
			hf.hash, _ = fs.MD5HashFile(hf.path, fileHashSize)
			for i := range fileslice {
				f := &fileslice[i]
				if f.hash == "" {
					f.hash, _ = fs.MD5HashFile(f.path, fileHashSize)
				}
				if hf.hash == f.hash && fs.FilesBytesAreEqual(hf.path, f.path) {
					dupesChannel <- duplicates.Duplicate{Value1: f.path, Value2: hf.hash}
					break
				}
			}
		}
		hashMap[hf.size] = append(hashMap[hf.size], hf)
	}
}

func FindDuplicates(fileChannel <-chan Filehash) <- chan duplicates.Duplicate {
	dupesChannel := make(chan duplicates.Duplicate)

	go func() {
		defer close(dupesChannel)
		emitDuplicates(fileChannel, dupesChannel)
	}()
	return dupesChannel
}

func GetDuplicateFileDeleter(writer io.Writer) duplicates.DuplicateHandler {
	return func (d duplicates.Duplicate) {
		path := d.Value2
		fmt.Fprintf(writer, "DELETING %s\n", path)
		fs.Delete(d.Value2)
	}
}

func ProcessDuplicateFiles(dir string, dupeHandler duplicates.DuplicateHandler) {
	duplicates.ApplyFuncToChan(
		FindDuplicates(
			FindFilesInPath(dir)), dupeHandler)
}
func SetFS(newFS fileio.FileIO) {
	fs = newFS
}

func main() {
	csv := flag.Bool("c", false, "Print duplicate value as a CSV to the console")
	del := flag.Bool("d", false, "Delete all duplicate values")

	flag.Parse()

	dupeHandler := duplicates.GetWriter(os.Stdout)
	if *csv {
		dupeHandler = duplicates.GetCsvWriter(os.Stdout)
	}else if *del {
		dupeHandler = GetDuplicateFileDeleter(os.Stdout)
	}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s [optons] directory\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "Calling without any options does a dry run and lists he files to be delete")
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		flag.Usage()
	}

	dir := flag.Args(0)

	SetFS(fileio.FS{})
	ProcessDuplicateFiles(dir, dupeHandler)
}
