package util

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type DirProcessor interface {
	Tarzip(src string, f *os.File) error
}

type LocalProcessor struct{}

func NewLocalProcessor() *LocalProcessor {
	return &LocalProcessor{}
}

func (s *LocalProcessor) Tarzip(src string, f *os.File) error {

	fmt.Printf("Tarzipping file %s...\n", f.Name())

	if _, err := os.Stat(src); err != nil {
		return fmt.Errorf("fail to tar file: %v", err)
	}
	gzw := gzip.NewWriter(f)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		header.Name = file

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		f, err := os.Open(file)
		if err != nil {
			return err
		}

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		f.Close()

		return nil
	})
}
