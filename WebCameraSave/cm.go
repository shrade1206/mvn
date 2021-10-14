package main

import (
	"log"
	"os"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	// go run ./cm.go filename.jpg

	//檔案名稱.副檔名
	saveFile := os.Args[1]

	// 設定視訊設備
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Println(err)
	}

	defer webcam.Close()
	// 延長視訊鏡頭時間 1s，以免曝光太短
	time.Sleep(time.Second)

	// 設定一個框，拿來儲存畫面
	img := gocv.NewMat()

	defer img.Close()
	// 設備讀到畫面
	webcam.Read(&img)

	// 保存設定的名字與擷取的照片
	gocv.IMWrite(saveFile, img)
}
