package lib

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Zip(srcDir string, destZip string) error {
	zf, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zf.Close()

	archive := zip.NewWriter(zf)
	defer archive.Close()

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == srcDir {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, srcDir+string(os.PathSeparator))
		header.Method = zip.Deflate

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return err
}

func Unzip(zipFile string, destDir string) error {
	zf, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zf.Close()

	for _, f := range zf.File {
		filename := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(filename, os.ModePerm)
		} else {
			_ = os.MkdirAll(filepath.Dir(filename), os.ModePerm)

			reader, err := f.Open()
			if err != nil {
				return err
			}
			defer reader.Close()

			file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, reader)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
