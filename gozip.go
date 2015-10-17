package gozip

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	patch "github.com/jniltinho/gozip/patchzip"
)

func IsZip(path string) bool {
	r, err := zip.OpenReader(path)
	if err == nil {
		r.Close()
		return true
	}
	return false
}

func Zip(path string, dirs []string) (err error) {
	if IsZip(path) {
		return errors.New(path + " is already a zip file")
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	startoffset, err := f.Seek(0, os.SEEK_END)
	if err != nil {
		return
	}

	w := patch.NewWriterAt(f, startoffset)

	for _, dir := range dirs {
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			fh, err := patch.FileInfoHeader(info)
			if err != nil {
				return err
			}
			fh.Name = path

			p, err := w.CreateHeader(fh)
			if err != nil {
				return err
			}
			if !info.IsDir() {
				content, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				_, err = p.Write(content)
				if err != nil {
					return err
				}
			}
			return err
		})
	}
	err = w.Close()
	return
}

func Unzip(zippath string, destination string) (err error) {
	r, err := zip.OpenReader(zippath)
	if err != nil {
		return err
	}
	for _, f := range r.File {
		fullname := path.Join(destination, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fullname, f.FileInfo().Mode().Perm())
		} else {
			os.MkdirAll(filepath.Dir(fullname), 0755)
			perms := f.FileInfo().Mode().Perm()
			out, err := os.OpenFile(fullname, os.O_CREATE|os.O_RDWR, perms)
			if err != nil {
				return err
			}
			rc, err := f.Open()
			if err != nil {
				return err
			}
			_, err = io.CopyN(out, rc, f.FileInfo().Size())
			if err != nil {
				return err
			}
			rc.Close()
			out.Close()

			mtime := f.FileInfo().ModTime()
			err = os.Chtimes(fullname, mtime, mtime)
			if err != nil {
				return err
			}
		}
	}
	return
}

func UnzipList(path string) (list []string, err error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return
	}
	for _, f := range r.File {
		list = append(list, f.Name)
	}
	return
}
