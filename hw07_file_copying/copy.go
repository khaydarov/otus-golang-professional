package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fromFile.Close()

	fromFileInfo, err := fromFile.Stat()
	if err != nil {
		return err
	}

	if fromFileInfo.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer toFile.Close()

	if offset != 0 {
		fromFile.Seek(offset, io.SeekStart)
	}

	var bar *pb.ProgressBar
	barLimit := limit
	if limit > fromFileInfo.Size() {
		barLimit = fromFileInfo.Size()
	}

	if fromFileInfo.Size()-offset < limit || limit == 0 {
		barLimit = fromFileInfo.Size() - offset
	}

	bar = pb.Full.Start64(barLimit)
	barReader := bar.NewProxyReader(fromFile)

	copyLimit := fromFileInfo.Size()
	if limit != 0 {
		copyLimit = limit
	}

	_, err = io.CopyN(toFile, barReader, copyLimit)
	bar.Finish()

	if err != nil && !errors.Is(err, io.EOF) {
		err = os.Remove(toPath)
		return err
	}

	return nil
}
