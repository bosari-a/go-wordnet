package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/bosari-a/array-utils/search"
)

// the category/type of words defined in array "pos"
var pos [4]string = [4]string{"adv", "adj", "noun", "verb"}

type Word struct {
	Word        string
	Definitions map[string]string
	offsets     map[string][]uint64
}

func (word *Word) FindOffsets(w string, dictPath string) error {
	word.offsets = make(map[string][]uint64)
	for _, ext := range pos {
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
			// regexp to match corresponding pos
			posRegexp := regexp.MustCompile(`\s[a-z]{1}?\s`)
			// corresponding pos is:
			corrPos := string(posRegexp.Find([]byte(line)))
			println(corrPos)
			// regexp to match buffer offsets to definitions in data file
			offsetRegexp := regexp.MustCompile(`\d{8}?`)
			offsetsBytes := offsetRegexp.FindAll([]byte(line), -1)
			offsets := make([]uint64, len(offsetsBytes))
			for k, v := range offsetsBytes {
				offsets[k] = binary.BigEndian.Uint64(v)
			}
			word.offsets[corrPos] = offsets
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
	err := word.FindOffsets("act", "./dict")
	if err != nil {
		panic(err)
	}
	fmt.Printf("offsets: %v\n", word.offsets)
}
