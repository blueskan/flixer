package strategy

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

type FileStrategy struct {
	filename string
}

func (fs FileStrategy) Process(values url.Values) {
	f, _ := os.Create(fs.filename)
	defer f.Close()

	w := bufio.NewWriter(f)

	for key, value := range values {
		fmt.Fprintf(w, "%s:%s\n", key, value)
	}
}

func NewFileStrategy(filename string) OutputStrategy {
	return &FileStrategy{
		filename: filename,
	}
}
