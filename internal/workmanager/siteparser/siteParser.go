package siteparser

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mailTestGolang/internal/workmanager"
)

// siteParser реализует интерфейс WordCalculator для парсинга урла.
type siteParser struct {
	ctx     context.Context
	timeOut time.Duration
}

// NewSiteParser создает siteParser.
func NewSiteParser(ctx context.Context, timeOut time.Duration) workmanager.WordCalculator {
	return &siteParser{ctx: ctx, timeOut: timeOut}
}

// Calculate делает Get запрос по урлу и считает количество совпадений слова.
func (s *siteParser) Calculate(url, word string) (int, error) {
	ctx, cancel := context.WithTimeout(s.ctx, s.timeOut)
	defer cancel()

	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("cannot get %s %w", url, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("cannot ReadAll body %w", err)
	}
	return strings.Count(string(body), word), nil
}
