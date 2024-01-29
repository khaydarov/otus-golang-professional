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

	if fromFileInfo.Size()-offset < limit {
		barLimit = fromFileInfo.Size() - offset
	}

	if limit != 0 {
		bar = pb.Full.Start64(barLimit)
		barReader := bar.NewProxyReader(fromFile)
		io.CopyN(toFile, barReader, limit)
		bar.Finish()

		return nil
	}

	bar = pb.Full.Start64(fromFileInfo.Size() - offset)
	barReader := bar.NewProxyReader(fromFile)
	io.Copy(toFile, barReader)
	bar.Finish()

	return nil
}
