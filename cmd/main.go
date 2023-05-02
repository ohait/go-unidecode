package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	convert("./unidecode/src/main/resources/data/")
}

func convert(dir string) {
	d, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, e := range d {
		if e.IsDir() {
			continue
		}
		name := filepath.Join(dir, e.Name())
		y, err := strconv.ParseInt(e.Name()[1:], 16, 32)
		if err != nil {
			log.Printf("ignoring %q", name)
			continue
		}
		code := int(y)
		log.Printf("name: %q", name)
		f, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		data, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		f, err = os.Create(e.Name() + ".go")
		if err != nil {
			panic(err)
		}
		f.WriteString("package unidecode\n\n")
		f.WriteString("func init() {")
		for i, s := range strings.Split(string(data), "\n") {
			x := code*256 + i
			ln := fmt.Sprintf("    m[%d] = %q\n", x, s)
			//log.Printf("%04x[%q]: `%s`", x, rune(x), s)
			f.WriteString(ln)
		}
		f.WriteString("}")
	}
}
