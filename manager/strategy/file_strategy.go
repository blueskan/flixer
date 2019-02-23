package strategy

import (
	"bufio"
	"fmt"
	"os"
)

type FileStrategy struct {
	filename string
}

func (fs FileStrategy) Process(values string) {
	f, _ := os.Create(fs.filename)
	defer f.Close()

	w := bufio.NewWriter(f)

	fmt.Fprintf(w, fmt.Sprintf("%s\n", values))
	w.Flush()
}

func NewFileStrategy(filename string) OutputStrategy {
	return &FileStrategy{
		filename: filename,
	}
}
