package main

import (
	"os"

	"gocv.io/x/gocv"
)

func main() {
	// go run ./cm.go filename.jpg

	//檔案名稱.副檔名
	saveFile := os.Args[1]

	// 設定視訊設備
	webcam, _ := gocv.VideoCaptureDevice(0)

	defer webcam.Close()
	// 設定視窗跟名稱
	window := gocv.NewWindow("Hello")
	// 設定一個框，拿來儲存畫面
	img := gocv.NewMat()

	defer img.Close()
	// 設備讀到畫面
	webcam.Read(&img)
	// 在指定的地方顯示圖片
	window.IMShow(img)

	defer window.Close()
	// 保存設定的名字與擷取的照片
	gocv.IMWrite(saveFile, img)
}
