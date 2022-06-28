package fileparser

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mailTestGolang/internal/workmanager"
)

// fileParser реализует интерфейс WordCalculator для файла.
type fileParser struct {
	ctx context.Context
}

// NewFileParser создает fileParser.
func NewFileParser(ctx context.Context) workmanager.WordCalculator {
	return &fileParser{ctx: ctx}
}

// Calculate открывает файл, построчно считает вхождения слово.
func (r *fileParser) Calculate(path, word string) (int, error) {
	wordsCount := 0
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("error when opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordsCount += strings.Count(scanner.Text(), word)
	}

	if err = scanner.Err(); err != nil {
		return 0, fmt.Errorf("error when scanning file: %w", err)
	}
	return wordsCount, nil
}
