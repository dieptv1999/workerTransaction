package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"wokerTrans/helps"
)

type Task struct {
	id     int
	jobs   chan []string
	m      *sync.Mutex
	wg     *sync.WaitGroup
	errors chan DataAPIRessponse
}

func NewTask(id int, wg *sync.WaitGroup, jobs chan []string, errors chan DataAPIRessponse, m *sync.Mutex) *Task {
	return &Task{
		id:     id,
		jobs:   jobs,
		errors: errors,
		m:      m,
		wg:     wg,
	}
}

func (t *Task) Run() {
	for alive := true; alive; {
		m.Lock()
		address, ok := <-t.jobs
		m.Unlock()
		if !ok {
			fmt.Println("ok", t.id)
			break
		}
		fmt.Println(address, "address")
		for _, address := range address {
			url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&apikey=%s", address, API_KEY_TOKEN)
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(REQUEST_TIMEOUT)*time.Second)
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			fmt.Println(address)

			if err != nil {
				fmt.Println(err.Error(), "err.Error() ------- task address request")
				continue
			}
			var resResult = TransactionTraceAPIResponse{}
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil || resp.Body == nil {
				fmt.Println(err.Error(), "err.Error() ------- task address request err")
				continue
			}
			if err := json.NewDecoder(resp.Body).Decode(&resResult); err != nil {
				fmt.Println(err.Error(), "err.Error() ------- Decode resp api err")
				continue
			}
			if resResult.Status != "1" || resResult.Message != "OK" {
				continue
			}
			trans := MapTxHash2Trans{}
			for _, m := range resResult.Result {
				fmt.Println(m, "mmmmmmmmmm")
				_, err := trans.GetFromKey(m.BlockHash)
				if err != nil || helps.IsError(err) {
					log.Printf("transaction %s không hợp lệ", m.BlockHash)
					t.errors <- m
					fmt.Println("aaaaa")
				}
			}
		}
		fmt.Println("end", t.id)
		t.wg.Done()
	}
}
