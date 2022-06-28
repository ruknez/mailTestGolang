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

	resp, err := doGetRequest(ctx, url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("cannot ReadAll body %w", err)
	}
	return strings.Count(string(body), word), nil
}

// doGetRequest делает GET запрос, в случае успеха необходимо закрыть тело.
func doGetRequest(ctx context.Context, url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot make create request %s %w", url, err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot make do request %s %w", url, err)
	}
	return resp, nil
}
