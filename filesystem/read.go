package filesystem

import (
	"io"
	"os"
	"strings"
)

func Read(filepath string) ([]string, error) {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0755); if err != nil {
		return nil, err
	}

	var data string
	for {
		chunk := make([]byte, 1042)
		n, err := f.Read(chunk)

		if err == io.EOF {
			break
		}

		data += string(chunk[:n])
	}

	data = strings.TrimSpace(data)
	return strings.Split(data, "\n"), nil
}
