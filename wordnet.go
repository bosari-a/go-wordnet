package gowordnet

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

// Word struct. Contains the actual word and a map of definitions with their
// corresponding pos (where pos is noun, adjective, adverb, verb).
type Word struct {
	Word        string
	Definitions map[string]string
}

// The category/type of words mapped to the file extension
// of index.pos and data.pos files
var posExt = map[string]string{
	"n": "noun",
	"v": "verb",
	"r": "adv",
	"a": "adj",
}

// This function takes the word to be searched and the path to the wordnet dict folder.
// It mutates a word struct by adding definitions to the Definitions hash map.
// Returns error (nil if no error occurs).
func (word *Word) GetDefinitions(w string, dictPath string) error {
	word.Definitions = make(map[string]string)
	for _, ext := range posExt {
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
			for k, v := range offsetsBytes {
				offset, err := strconv.Atoi(string(v))
				if err != nil {
					return err
				}
				key := fmt.Sprintf("%v(%v)", ext, k+1)
				datafile := fmt.Sprintf("data.%v", ext)
				def, err := ParseDataFile(path.Join(dictPath, datafile), offset)
				if err != nil {
					return err
				}
				word.Definitions[key] = def
			}
			word.Word = w
		}
	}
	return nil
}

// This function parses definition files (data.pos files).
// It takes path to data.pos and the offset to the definition.
// Returns the definition,nil or [empty string],err.
func ParseDataFile(dataPath string, offset int) (string, error) {
	// open data.pos file and read the definition for current offset
	// add definition to word.Definition
	fd, err := os.Open(dataPath)
	if err != nil {
		return "", err
	}
	fd.Seek(int64(offset), io.SeekStart)
	// use bufio to make a new reader so we can directly read a line
	r := bufio.NewReader(fd)
	defLine, _, err := r.ReadLine()
	if err != nil {
		return "", err
	}
	fd.Close()
	def := strings.Split(string(defLine), "|")[1]
	return strings.TrimSpace(def), nil
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
