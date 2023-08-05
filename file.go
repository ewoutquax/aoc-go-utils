package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFileAsNumbers(baseDir string) (numbers []int) {
	var lines []string = ReadFileAsLines(baseDir)

	for _, string := range lines {
		numbers = append(numbers, ConvStrToI(string))
	}

	return
}

func ReadFileAsBlocks(baseDir string) (blocks [][]string) {
	var block_inputs []string = strings.Split(readFile(baseDir), "\n\n")

	for _, block_input := range block_inputs {
		blocks = append(blocks, strings.Split(block_input, "\n"))
	}
	return
}

func ReadFileAsLines(baseDir string) []string {
	return strings.Split(readFile(baseDir), "\n")
}

func ReadFileAsLine(baseDir string) string {
	return readFile(baseDir)
}

func readFile(baseDir string) string {
	fullPath := fmt.Sprintf("%s/%s", baseDir, "input.txt")
	absPath, _ := filepath.Abs(fullPath)

	raw, err := os.ReadFile(absPath)
	check(err)

	return strings.TrimSuffix(string(raw), "\n")
}
