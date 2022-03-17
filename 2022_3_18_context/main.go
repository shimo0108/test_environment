package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("開始")
	time.Sleep(2 * time.Second)
	fmt.Println("終了")

	ch <- "実行結果"
}

func main() {
	// チャネルをmake関数を使って生成
	ch := make(chan string)

	// contextを作成
	ctx := context.Background()

	//ctxに1秒間の時間制限をつけて再定義している
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	// リソースの解放処理を行う。
	defer cancel()

	go longProcess(ctx, ch)
L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("#####Error#####")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}

	}
	fmt.Println("ループ抜けた")
}
