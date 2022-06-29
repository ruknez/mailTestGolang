package manager

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

// Worker интерфейс определяющий воркера.
type Worker interface {
	Work(st string) (int, error)
}

// ProvidingResults определяет то что обрабатывает результат воркера.
type ProvidingResults interface {
	WorkerResult(path string, count int)
	TotalResult(count int)
}

// Manager структура с основной бизнес логико. Запускает воркеров и отправляет их результат в обработчик.
type Manager struct {
	ctx            context.Context
	maxWorkers     int
	worker         Worker
	resultProvider ProvidingResults
	inputData      <-chan string
}

// NewManager принимает канал которому идут данные и реализацию интерфейсов Worker и ProvidingResults.
func NewManager(
	ctx context.Context,
	k int, inputData <-chan string,
	worker Worker,
	resultProvider ProvidingResults) *Manager {
	if k <= 0 {
		panic(fmt.Sprint("workers count is not right ", k))
	}
	return &Manager{
		ctx:            ctx,
		maxWorkers:     k,
		worker:         worker,
		resultProvider: resultProvider,
		inputData:      inputData,
	}
}

// StartManage функция с основной бизнес логикой.
func (m *Manager) StartManage() {
	var totalResult int32
	countWorkers := 0
	wg := sync.WaitGroup{}
	condVariable := sync.NewCond(&sync.Mutex{})

	for inputLine := range m.inputData {
		condVariable.L.Lock()
		for countWorkers >= m.maxWorkers {
			condVariable.Wait()
		}
		countWorkers++
		condVariable.L.Unlock()

		wg.Add(1)
		go func(input string) {
			defer func() {
				wg.Done()
				condVariable.L.Lock()
				countWorkers--
				condVariable.Signal()
				condVariable.L.Unlock()
			}()

			res, err := m.worker.Work(input)
			if err != nil {
				log.Println(fmt.Errorf("error from worker %w", err))
				return
			}
			m.resultProvider.WorkerResult(input, res)
			atomic.AddInt32(&totalResult, int32(res))
		}(inputLine)
	}
	wg.Wait()
	m.resultProvider.TotalResult(int(totalResult))
}
