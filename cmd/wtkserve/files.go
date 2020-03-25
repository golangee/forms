package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// A file wraps an absolute path name and the already read file info. This avoids redundant file i/o in our context.
type file struct {
	AbsoluteFileName string
	Info             os.FileInfo
}

// listFiles returns all visible files (not folders) in lexical order.
func listFiles(root string) ([]file, error) {
	var res []file
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Mode().IsRegular() && !strings.HasPrefix(info.Name(), ".") {
				res = append(res, file{path, info})
			}

			if strings.HasPrefix(info.Name(), ".") && info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		})
	if err != nil {
		return res, fmt.Errorf("failed to list files from %s:%w", root, err)
	}
	return res, nil
}

// calculateHash takes all input files and calculates a common hash
func calculateHash(files []file) (string, error) {
	h := sha256.New()
	for _, file := range files {
		f, err := os.Open(file.AbsoluteFileName)
		if err != nil {
			return "", fmt.Errorf("failed to open file %s: %w", file.AbsoluteFileName, err)
		}
		if _, err := io.Copy(h, f); err != nil {
			_ = f.Close()
			return "", fmt.Errorf("failed to calculate hash of %s: %w", file.AbsoluteFileName, err)
		}
		err = f.Close()
		if err != nil {
			return "", fmt.Errorf("failed to close file %s: %w", file.AbsoluteFileName, err)
		}
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// copy copies all bytes from source to dest
func copy(fromName string, toName string) (int64, error) {
	from, err := os.Open(fromName)
	if err != nil {
		return 0, fmt.Errorf("failed to open source file %s: %w", fromName, err)
	}
	defer from.Close()

	to, err := os.OpenFile(toName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, fmt.Errorf("failed to open destination file %s: %w", fromName, err)
	}
	defer to.Close()

	n, err := io.Copy(to, from)
	if err != nil {
		return n, fmt.Errorf("cannot copy from %s -> %s: %w", fromName, toName, err)
	}
	return n, nil
}

// copyDir copies a whole directory recursively
func copyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = copyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

// copyFile copies a single file from src to dst
func copyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
