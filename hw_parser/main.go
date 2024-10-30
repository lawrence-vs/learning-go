package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
)

type Parser struct {
	Status   string
	Blocks   []string
	Filename string
}

func (p *Parser) FromHwFile(raw, filename string) {
	p.Blocks = regexp.MustCompile(`\n{2,}`).Split(raw, -1)
	p.Filename = filename
}

func (p *Parser) ParseBlock(block string) bool {
	// Sample regex pattern for the ADDRESS extraction
	match, _ := regexp.MatchString(`\sADDRESS\s+(\d+),\s(\d+),\s(\d+)`, block)
	if !match {
		p.Status = "No symbols found"
		return false
	}
	p.Status = "Symbols found and processed"
	return true
}

func (p *Parser) Run() bool {
	for _, block := range p.Blocks {
		if p.ParseBlock(block) {
			fmt.Println(p.Status)
		}
	}
	p.Export()
	return true
}

func (p *Parser) Export() bool {
	file, err := os.Create(p.Filename)
	if err != nil {
		fmt.Println("Could not create file:", err)
		return false
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Header1", "Header2", "Header3"})
	writer.Write([]string{"Row1Col1", "Row1Col2", "Row1Col3"})

	fmt.Println("Exported to", p.Filename)
	return true
}

func main() {
	parser := Parser{}
	parser.FromHwFile("raw data here", "output.csv")
	parser.Run()
}
