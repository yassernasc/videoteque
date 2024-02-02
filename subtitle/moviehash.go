package subtitle

import (
	"encoding/binary"
	"fmt"
	"io"
	"videoteque/movie"
)

func moviehash() (string, error) {
	const blockSize = 64 * 1024 // 64Kb
	const uint64SizeInBytes = 8

	var hash uint64

	v := movie.VideoRef
	videoSize := int64(v.Size())
	offsets := []int64{0, videoSize - blockSize} // first and last 64Kb's

	videoReader := v.Reader()
	defer videoReader.Close()

	for _, offset := range offsets {
		if _, err := videoReader.Seek(offset, io.SeekStart); err != nil {
			return "", err
		}

		partialReader := io.LimitReader(videoReader, blockSize)
		tmpBuf := make([]uint64, blockSize/uint64SizeInBytes)
		if err := binary.Read(partialReader, binary.LittleEndian, &tmpBuf); err != nil {
			return "", err
		}

		for _, n := range tmpBuf {
			hash += n
		}
	}
	hash += uint64(videoSize)

	return fmt.Sprintf("%016x", hash), nil
}
