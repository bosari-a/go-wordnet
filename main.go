package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/bosari-a/array-utils/search"
)

func main() {
	// the category/type of words defined in array "pos"
	pos := [4]string{"noun", "verb", "adj", "adv"}
	// regexp to match a word in a line read from index.pos file
	wordRegexp := regexp.MustCompile("[a-zA-Z]+")
	// regexp to match buffer offsets to definitions in data file
	offsetRegexp := regexp.MustCompile(`\d{8}?`)

	v, err := os.ReadFile("./dict/index.verb")
	if err != nil {
		log.Fatal(err)
	}

	content := string(v)
	lines := strings.Split(content, "\n")
	sort.Slice(lines, func(i2, j int) bool {
		return lines[i2] < lines[j]
	})

	var words []string
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		sep := strings.Split(line, " ")
		words = append(words, sep[0])
	}

	bs := search.BinaryStrFactory(words, true)
	i := bs.BinarySearchStr("act")
	// fmt.Printf("found %v at %v\n", lines[i], i)
	regWord := regexp.MustCompile("[a-zA-Z]+")
	found := regWord.Find([]byte(lines[i]))

	regOffsets := regexp.MustCompile(`\d{8}?`)
	offsetsb := regOffsets.FindAll([]byte(lines[i]), -1)
	offsets := make([]string, len(offsetsb))
	for i3, v := range offsetsb {
		offsets[i3] = string(v)
	}
	fmt.Printf("word: %v\noffsets: %v\n", string(found), offsets)
}
