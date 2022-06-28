package workmanager

import (
	"context"
	"net/url"

	"github.com/mailTestGolang/internal/manager"
)

// WordCalculator основной интерфейс компонентов воркера.
type WordCalculator interface {
	Calculate(path, word string) (int, error)
}

// workManager отвечает за выбор компонента который должен заниматься подсчетом слов.
type workManager struct {
	ctx        context.Context
	wordToFind string
	fileReader WordCalculator
	siteParser WordCalculator
}

// NewWorkManager принимает слово которое надо искать и возможные компоненты.
func NewWorkManager(ctx context.Context, word string, fileReader, siteParser WordCalculator) manager.Worker {
	return &workManager{
		ctx:        ctx,
		wordToFind: word,
		fileReader: fileReader,
		siteParser: siteParser,
	}
}

// Work определяет по входной строке какой компонент должен считать слова.
func (w *workManager) Work(st string) (int, error) {
	if isItURL(st) {
		return w.siteParser.Calculate(st, w.wordToFind)
	}
	return w.fileReader.Calculate(st, w.wordToFind)
}

// isItURL проверяет - является ли строка полноценным урлом.
func isItURL(st string) bool {
	parsedURL, err := url.ParseRequestURI(st)
	if err != nil {
		return false
	}
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}
