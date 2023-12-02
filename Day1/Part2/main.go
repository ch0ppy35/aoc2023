package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Processor struct {
	Scanner *bufio.Scanner
	File    *os.File
}

func NewProcessor(filePath string) (*Processor, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(file)
	if err := s.Err(); err != nil {
		return nil, err
	}

	return &Processor{
		File:    file,
		Scanner: s,
	}, nil
}

func (p *Processor) ProcessCodes() int {
	var total int
	for p.Scanner.Scan() {
		line := p.Scanner.Text()
		var l int

		for i := 0; i < len(line); i++ {
			if ll, ok := p.StringNumberToInt(line[i:]); ok {
				l = ll * 10
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if ll, ok := p.StringNumberToInt(line[i:]); ok {
				l += ll
				break
			}
		}
		total += l
	}

	return total
}

func (p *Processor) StringNumberToInt(line string) (int, bool) {
	switch {
	case len(line) > 0 && line[0] >= '0' && line[0] <= '9':
		return int(line[0] - '0'), true
	default:
		return p.stringNumberLookup(line)
	}
}

func (p *Processor) stringNumberLookup(line string) (int, bool) {
	switch {
	case strings.HasPrefix(line, "zero"):
		return 0, true
	case strings.HasPrefix(line, "one"):
		return 1, true
	case strings.HasPrefix(line, "two"):
		return 2, true
	case strings.HasPrefix(line, "three"):
		return 3, true
	case strings.HasPrefix(line, "four"):
		return 4, true
	case strings.HasPrefix(line, "five"):
		return 5, true
	case strings.HasPrefix(line, "six"):
		return 6, true
	case strings.HasPrefix(line, "seven"):
		return 7, true
	case strings.HasPrefix(line, "eight"):
		return 8, true
	case strings.HasPrefix(line, "nine"):
		return 9, true
	default:
		return 0, false
	}
}

func main() {
	processor, err := NewProcessor("input.txt")
	if err != nil {
		log.Fatal("Error creating code processor: ", err)
	}
	defer processor.File.Close()

	total := processor.ProcessCodes()
	log.Printf("The total is %d", total)
}
