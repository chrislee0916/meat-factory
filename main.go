package main

import (
	"fmt"
	"sync"
	"time"
)

type meat struct {
	Type string // 肉品類型
	Cost int    // 處理該肉品所需的時間
}

type factory struct {
	Meats   []int    // 肉品的列表
	Workers []string // 員工列表
}

func NewFactory(meats []int, workers []string) *factory {
	return &factory{
		Meats:   meats,
		Workers: workers,
	}
}

func (f *factory) Run() {
	meatChan := make(chan meat, len(f.Meats)/2) // 建立一個用來存放所有肉品的channel
	meatType := map[int]string{                 // 定義肉品類型的映射關係
		1: "牛肉",
		2: "豬肉",
		3: "雞肉",
	}

	var wg sync.WaitGroup // 創建一個 WaitGroup 用於等待所有的 goroutine 完成工作
	wg.Add(len(f.Meats))  // 設置 WaitGroup 的計數器為肉品數量

	go func() {
		defer close(meatChan)
		for _, v := range f.Meats {
			meatChan <- meat{Type: meatType[v], Cost: v} // 將肉品加入 channel 中
		}
	}()

	for i := 0; i < len(f.Workers); i++ {
		go func(i int) {
			for meat := range meatChan {
				// 員工取得肉品
				fmt.Printf("%s 在 %s 取得%s\n", f.Workers[i], time.Now().Format("2006-02-01 15:04:05"), meat.Type)
				// 員工處理肉品所需的時間
				time.Sleep(time.Duration(meat.Cost) * time.Second)
				// 員工處理完肉品
				fmt.Printf("%s 在 %s 處理完%s\n", f.Workers[i], time.Now().Format("2006-02-01 15:04:05"), meat.Type)
				// 一個肉品處理完成 WaitGroup 計數器 -1
				wg.Done()
			}
		}(i)
	}

	wg.Wait() // 等待所有的肉品處理完成

}

func main() {
	// 肉品的列表
	meats := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3}
	// 員工的列表
	workers := []string{"A", "B", "C", "D", "E"}

	factory := NewFactory(meats, workers)
	factory.Run() // 執行工廠
}
