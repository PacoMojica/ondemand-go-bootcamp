package database

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

type worker struct {
	id        int
	limit     int
	processed int
	filter    string
}

func (w *worker) run(input <-chan []string, output chan<- []string) {
	var filter int64
	if w.filter == "even" {
		filter = 0
	} else {
		filter = 1
	}

	for record := range input {
		if len(record) == 0 {
			log.Println("found empty record")
			continue
		}
		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			log.Printf("could not parse record ID: '%s'", record[0])
			continue
		}
		if w.processed < w.limit && id%2 == filter {
			output <- record
			w.processed += 1
		}
		if w.limit == w.processed {
			break
		}
	}
}

type pool struct {
	size           int
	itemsPerWorker int
	filter         string
	wg             *sync.WaitGroup
	input          chan []string
	output         chan []string
	items          [][]string
	count          int
	maxItems       int
}

type workerPool interface {
	run(*csv.Reader) [][]string
}

func (p *pool) spawn() {
	for id := 0; id < p.size; id++ {
		p.wg.Add(1)
		go func(id int) {
			defer p.wg.Done()
			w := worker{id: id, limit: p.itemsPerWorker, filter: p.filter}
			w.run(p.input, p.output)
		}(id)
	}
}

func (p *pool) read(r *csv.Reader) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		p.input <- record
	}
	close(p.input)
}

func (p *pool) close() {
	p.wg.Wait()
	close(p.output)
}

func (p *pool) run(r *csv.Reader) [][]string {
	p.spawn()
	go p.read(r)
	go p.close()

	for record := range p.output {
		if p.count < p.maxItems {
			p.items = append(p.items, record)
			p.count += 1
		} else {
			break
		}
	}

	return p.items
}

func create(filter string, maxItems, itemsPerWorker int) workerPool {
	input := make(chan []string, maxItems)
	output := make(chan []string, maxItems)
	wg := new(sync.WaitGroup)

	return &pool{
		size:           maxItems,
		itemsPerWorker: itemsPerWorker,
		filter:         filter,
		maxItems:       maxItems,
		wg:             wg,
		input:          input,
		output:         output,
	}
}

func (db *database) ConcurrentRead(filter string, maxItems, itemsPerWorker int) ([][]string, error) {
	f, err := os.OpenFile(db.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	p := create(filter, maxItems, itemsPerWorker)
	result := p.run(reader)

	return result, nil
}
