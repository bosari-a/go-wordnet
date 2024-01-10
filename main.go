package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/bosari-a/array-utils/search"
)

// Word struct for what a word should contain.
// Offsets are stored for the purpose of finding definitions in data.pos files
type Word struct {
	Word        string
	Definitions map[string]string
	offsets     map[string][]int
}

// the category/type of words defined in array "pos"
var posExt = map[string]string{
	"n": "noun",
	"v": "verb",
	"r": "adv",
	"a": "adj",
}

func (word *Word) FindOffsets(w string, dictPath string) error {
	word.offsets = make(map[string][]int)
	for pos, ext := range posExt {
		file := fmt.Sprintf("index.%v", ext)
		filepath := path.Join(dictPath, file)
		data, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}
		lines := dataToContent(data)
		bs := search.BinaryStrFactory(lines, true)
		// index of the word (-1 if it isn't found)
		i := bs.BinarySearchStr(w, func(s string) string {
			return strings.Split(s, " ")[0]
		})
		if i != -1 {
			line := lines[i]
			// regexp to match buffer offsets to definitions in data file
			offsetRegexp := regexp.MustCompile(`\d{8}?`)
			offsetsBytes := offsetRegexp.FindAll([]byte(line), -1)
			offsets := make([]int, len(offsetsBytes))
			for k, v := range offsetsBytes {
				offsets[k], _ = strconv.Atoi(string(v))
			}
			word.offsets[pos] = offsets
			word.Word = w
		}
	}
	return nil
}

// takes in data from os.ReadFile("index.pos")
// returns lines, words arrays (sorted)
func dataToContent(d []byte) []string {
	content := string(d)
	lines := strings.Split(content, "\n")
	sort.Slice(lines, func(i, j int) bool {
		return strings.Split(lines[i], " ")[0] < strings.Split(lines[j], " ")[0]
	})
	return lines
}

func main() {
	var word Word
	err := word.FindOffsets("pleasantly", "./dict")
	if err != nil {
		panic(err)
	}
	fmt.Printf("offsets: %v\n", word.offsets)
}
