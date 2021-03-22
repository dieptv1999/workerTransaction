package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	coinName           = "ETH"
	PAGE_SIZE    int64 = 100
	NUM_OF_TASKS       = 15
	m                  = sync.Mutex{}
	jobs         chan []string
	results      = make(chan int, 200)
	errors       chan DataAPIRessponse
	bc           = BlockchainAddress{
		Coin:      coinName,
		Available: false,
	}
	tasksChan chan *Task
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func main() {
	InitBigSetIf()
	//createDataTest()
	for alive := true; alive; {
		timer := time.NewTimer(5 * time.Minute)
		select {
		case <-timer.C:
			t := time.Now()
			////////////////////////////////////////////////////////
			if t.Hour() == 0 && t.Minute() < 5 {
				log.Println("worker check transaction start at ", t)
				//////////////////////////////////////////////////////////
				//TODO
				// check 5 phút một lần, đến khoảng thời gian từ 0h00 -> 0h05 sẽ thực hiện tác vụ (sẽ ko thực hiện đc 2 lần)
				go execute()
				/////////////////////////////////////////////////////////
			}
		}
	}
}

func execute() {
	jobs = make(chan []string, 100)
	errors = make(chan DataAPIRessponse, 200)
	///////////////////////////////////////
	// lấy tổng số address và tào pool task
	total, err := bc.GetTotalCount()
	var numOfPage int
	if total%PAGE_SIZE == 0 {
		numOfPage = int(total / PAGE_SIZE)
	} else {
		numOfPage = int(total/PAGE_SIZE) + 1
	}
	if err != nil {
		log.Println(err, "worker kiểm tra địa chỉ bị lỗi")
	}
	wg := sync.WaitGroup{}
	var numOfTask int
	if numOfPage > NUM_OF_TASKS {
		numOfTask = NUM_OF_TASKS
	} else {
		numOfTask = numOfPage
	}
	fmt.Println(numOfTask, "numOfTask")
	for w := 1; w <= numOfTask; w++ {
		go func(w int) {
			task := NewTask(w, &wg, jobs, errors, &m)
			task.Run()
		}(w)
	}
	////////////////////////////////////////
	fmt.Println(numOfPage, "numOfPage")
	wg.Add(numOfPage)
	for i := 0; i < numOfPage; i++ {
		start := i * int(PAGE_SIZE)
		bcAddress, _, err := bc.GetPaginate(int32(start), int32(PAGE_SIZE))
		if err != nil {
			log.Println(err, "worker kiểm tra địa chỉ bị lỗi")
		}
		var address []string
		for _, blockchainAddress := range bcAddress {
			address = append(address, blockchainAddress.BlockchainAddress)
		}
		jobs <- address
	}
	go func() {
		for k := true; k; {
			select {
			case err, ok := <-errors:
				if !ok {
					k = false
				} else {
					exeErr(err)
				}
			default:

			}
		}
	}()
	wg.Wait()
	close(jobs)
	close(errors)
	fmt.Println("end of execute")

}

func exeErr(err DataAPIRessponse) {
	fmt.Println("err: abcdefg")
	bTxError := mapDataApiToTxError(err)
	bTxError.Create()
}

func mapDataApiToTxError(data DataAPIRessponse) (err TxError) {
	err = TxError{
		BlockHash:         data.BlockHash,
		BlockNumber:       data.BlockNumber,
		Confirmations:     data.Confirmations,
		ContractAddress:   data.ContractAddress,
		CumulativeGasUsed: data.CumulativeGasUsed,
		From:              data.From,
		Gas:               data.Gas,
		GasPrice:          data.GasPrice,
		GasUsed:           data.GasUsed,
		Hash:              data.Hash,
		Input:             data.Input,
		IsError:           data.IsError,
		Nonce:             data.Nonce,
		TimeStamp:         data.TimeStamp,
		To:                data.To,
		TransactionIndex:  data.TransactionIndex,
		Txreceipt_status:  data.Txreceipt_status,
		Value:             data.Value,
		UpdatedAt:         time.Now().Unix(),
		CreatedAt:         time.Now().Unix(),
	}
	return
}

func createDataTest() {
	//listaddr := []string{
	//	"0xA6cC9e28Bf3F08467aa2d0761Ff84b74a7aBB12B",
	//	"0x1Cdb00D07b721B98Da52532DB9a7D82D2A4bF2e0",
	//	"0xc7218B8D2716efd5A5F69CcAAe754521E02d5E9F",
	//	"0x033eF6db9FBd0ee60E2931906b987Fe0280471a0",
	//	"0x4004C9E76b6C07Aa7Ef473764cDd3eF9E552b25c",
	//	"0x1f403C619f4EeC77a6Bfe269a08a02232F21711B",
	//	"0x7a250d5630b4cf539739df2c5dacb4c659f2488d",
	//	"0x2NFxPkWUMsNGoVvwLuLVWTi2bfptYPYdNq5",
	//	"0x2NCFxez61h34gqYn4hPxHMvtX7hmbyHuxpb",
	//	"0x2N7RHjfnn38TJZaBnYmoFcuXUE7NeGowEMn",
	//	"0x2N6DwwXvgA4Brcv6erHybuwAtEM4xwnsWRC",
	//	"0x2N451wqiJUbwHckZWFkEVXUvqGZhjYZkYwx",
	//	"0x2MzFTdQu33XjepHvctpGfxvDhg27iFmcsoM",
	//	"0x2MyW37aoHuJ4g5NmNs5SRzQTqsY1b1LxdsJ",
	//	"0x2My2xjRGhsnV4MeknLnuLsfCZ41eS8AjKxE",
	//	"0x2MxgueAZ4Nd9sudX54VKbwD5W27hPgmRaYP",
	//	"0x2Mx7wptxNu42xHtoV9ZAEnP1M2GCDKbvbz6",
	//	"0x2MvN2mHopmxspHwrXKLL2oWguFLKnHq8q4T",
	//	"0x2Mvg6LH5sw5WgjkVGtg9wDjYxoxHaErPxqK",
	//	"0xAFa152428B7255348f1e1D338f5118b572462bA8",
	//	"0xaF0F6fF04c51F4BF4c1EDe8339F3FAcA2A3C4935",
	//	"0x9eD9483Aa441cf1c1ade4e76e45E2915840B4b4b",
	//	"0x9C167cb5398f76721feA1dDb1C5F702C29410DAa",
	//	"0x7d720D28b7fD86758FA34b806cC13078D4A41909",
	//	"0x6BA068Ac716C577924f6D045f0Eb6758F585A02b",
	//	"0x38BD04629AE713D88d9ef543065A34999AA62856",
	//	"0x33703765534d1e1D2da85748c2Ef5565824E3579",
	//	"0x335406D216aA8eb929aedaa65165746AcC258420",
	//	"0x1EC4128E10Dc2eAb93068c8Ab4f8912E48672E95",
	//	"0x1E9675c94C1F37930e771EB2Fe9C43a1A9f5E55B",
	//	"0x0966f5c20c6165E72C70906Bc29b5B1B4a76Ae2F",
	//	"0x81b7E08F65Bdf5648606c89998A9CC8164397647",
	//	"0xa172De44903C8ba7bE568559eDD700b96fC2e246",
	//	"0xb74F51b3F27E06Cbd8Bb36b968F0062EfAa79884",
	//	"0xC8b793F2FEFf5614b3eC523Bc0A2736089b91543",
	//	"0x11f941c34A719Ba341f307d094bBA8bcAba3CA4E",
	//}

	listaddr := []string{
		"0xcA46F486A36594DAbBA36e6E71A07db87dD53171",
		"0x69030eB8915350E3537e2b97Eb9586F29F1519Fd",
		"0xdbaE96e97D00256a09322E1C49F062CB8beBb0F0",
		"0x0d85008A1282253C96CC655eD76b754D6913b17C",
		"0x9cBb00a772fC7d9A736Bdc177F0f6152f0975905",
		"0x4FBcc823E6E8D98cABf5c2271074898EA65b9292",
		"0x23049A681B570a75d20500974eb066137bfCf154",
		"0x834dfBf8b87b8df7a09d700621Dc6ecC32197912",
	}
	rand.Seed(time.Now().UnixNano())
	for _, s := range listaddr {
		bcs := BlockchainAddress{
			Coin:              coinName,
			Available:         false,
			BlockchainAddress: s,
			Type:              "1",
		}
		err := bcs.Create()
		if err != nil {
			fmt.Println(err)
		}
	}
	//for i := 0; i < 100; i++ {
	//	//0xD57aEE8Bf32d0E4c0a7E64Ab4d2cB81e96a96f2c
	//	b, _ := GenerateRandomBytes(i)
	//	bcs := BlockchainAddress{
	//		Coin:              coinName,
	//		Available:         false,
	//		BlockchainAddress: fmt.Sprintf("0x%s", helps.Hash256(b)),
	//		Type:              "1",
	//	}
	//	err := bcs.Create()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
}

func readDataError() {
	bTxError := TxError{}
	result, total, _ := bTxError.GetAll()
	fmt.Println(result, total, "aaaaaaaaaa")
}
