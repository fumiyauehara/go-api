package handler

import (
	"fmt"
	"net/http"
	"time"
)

func EventIndexHandler(w http.ResponseWriter, r *http.Request) {
	// クライアントとの接続を維持するループ
	for {
		// イベントデータを生成
		eventData := fmt.Sprintf("data: %s\n\n", time.Now().Format(time.RFC3339))

		// データを送信
		if _, err := fmt.Fprint(w, eventData); err != nil {
			panic(err)
		}

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		// 一定時間待機（例：1秒）
		time.Sleep(1 * time.Second)
	}
}
