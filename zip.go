// Package funks implements basic functions that I use frequently to save time on replicating code across different projects.
package funks

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ZipAFolder zips a folder and its contents into an archive named after the folder.
func ZipAFolder(f string) error {
	newZipArchive, err := os.Create(f + ".zip")
	if err != nil {
        return fmt.Errorf("unable to create new archive %s - %v", f, err)
	}
	defer newZipArchive.Close()

	zipWriter := zip.NewWriter(newZipArchive)
    defer zipWriter.Close()

    files, err := ioutil.ReadDir(f)
    if err != nil {
        return err
    }
    for _, j := range files {
        if err := AddFileToZip(filepath.Join(f, j.Name()), zipWriter); err != nil {
            return fmt.Errorf("unable to add %s to archive %s", j.Name(), newZipArchive.Name())
        }
    }
    return nil
}

// AddFileToZip adds a file f to a given *zip.Writer z. Remember to create a new archive file with os.Create and call zip.NewWriter on the *os.File.
func AddFileToZip(f string, z *zip.Writer) error {
	fileToZip, err := os.Open(f)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(fileToZip.Name())
	header.Method = zip.Deflate

	writer, err := z.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}
