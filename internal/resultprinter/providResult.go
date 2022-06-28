package resultprinter

import (
	"context"
	"fmt"

	"github.com/mailTestGolang/internal/manager"
)

// resultPrinter реализует интерфейс ProvidingResults.
type resultPrinter struct {
	ctx context.Context
}

// NewResultPrinter возвращает resultPrinter.
func NewResultPrinter(ctx context.Context) manager.ProvidingResults {
	return &resultPrinter{ctx: ctx}
}

// WorkerResult выводит результаты каждого подсчета в консоль.
func (r *resultPrinter) WorkerResult(path string, count int) {
	fmt.Println("Count for ", path, ": ", count)
}

// TotalResult выводит общие результаты.
func (r *resultPrinter) TotalResult(count int) {
	fmt.Println("Total: ", count)
}
