package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mailTestGolang/internal/manager"
	"github.com/mailTestGolang/internal/resultprinter"
	"github.com/mailTestGolang/internal/workmanager"
	"github.com/mailTestGolang/internal/workmanager/fileparser"
	"github.com/mailTestGolang/internal/workmanager/siteparser"
)

func main() {
	ctx := context.Background()
	k := 5
	word := "Go"
	getTimeOut := 2 * time.Second

	ch := goScanStdInput()

	work := workmanager.NewWorkManager(ctx, word, fileparser.NewFileParser(ctx), siteparser.NewSiteParser(ctx, getTimeOut))

	manage := manager.NewManager(ctx, k, ch, work, resultprinter.NewResultPrinter(ctx))
	manage.StartManage()
}

// goScanStdInput возвращает канал по которому идут строки из stdin.
func goScanStdInput() chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			ch <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Println(fmt.Errorf("scanStdInput error detected: %w", err))
		}
	}()

	return ch
}
