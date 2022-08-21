package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		// 判断字体文件是否存在
		if strings.Contains(path, "simsun.ttc") {
			fmt.Println(path)
			// 设置环境变量
			_ = os.Setenv("FYNE_FONT", path)
		}
	}
}

func main() {
	MyApp := app.New()
	c := MyApp.NewWindow("acrd 卡片")

	acrds(c)
	c.Resize(fyne.NewSize(600, 600))
	c.ShowAndRun()

	//a := app.New()
	//window := a.NewWindow("局域网共享空间")
	//
	//window.Resize(fyne.NewSize(600, 600))
	//
	//text1 := widget.NewLabel("文件1")
	//text2 := widget.NewLabel("文件2")
	//text3 := widget.NewLabel("文件3")
	//grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	//// 3 列
	////content := container.NewWithoutLayout(widget.NewLabel("文件1"), widget.NewLabel("文件2"))
	////content := container.New(layout.NewGridLayoutWithRows(10), widget.NewLabel("文件1"), widget.NewLabel("文件1"))
	//
	//window.SetContent(grid)
	//window.ShowAndRun()
}

func acrds(w fyne.Window) {

	acrd1 := widget.NewCard("acrd", "12123", container.NewVBox(
		widget.NewButton("clikc", nil),
		widget.NewLabel("123")))

	acrd1.Image = canvas.NewImageFromFile("1.jpg")
	w.SetContent(acrd1)
}
