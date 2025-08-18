package main

import (
	"fmt"
	"image/color"

	"github.com/sago35/tinydisplay/examples/initdisplay"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freeserif"
)

func main() {

	// P208〜
	// 文字列の基本の描画
	// base()

	// P211〜
	// 文字列を何度も描画する
	writeMultipleTimes()
}

func base() {
	// フォントと色の設定
	// 注意: TinyGoでは初期化順序の問題を避けるため、関数内で変数を宣言
	// 本ではmain関数外に定義しているがうまく動かなかった。README参照。
	font := &NotoSans32pt                       // font.goで生成された日本語フォント
	black := color.RGBA{0x00, 0x00, 0x00, 0xFF} // 黒色

	// ディスプレイの初期化
	display := initdisplay.InitDisplay()

	// 文字列を表示
	tinyfont.WriteLine(display, font, 5, 50, "tinyfont", black) // 英語
	tinyfont.WriteLine(display, font, 5, 100, font.Name, black) // フォント名
	tinyfont.WriteLine(display, font, 5, 150, "こんにちは", black)   // 日本語
	tinyfont.WriteLine(display, font, 5, 200, "おはよう", black)    // 日本語

	// プログラム終了を防ぐ
	select {}
}

func writeMultipleTimes() {
	// 重なる例
	black := color.RGBA{0x00, 0x00, 0x00, 0xFF}
	white := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	display := initdisplay.InitDisplay()
	display.FillScreen(white)

	count := 0
	label := NewLabel(320, 50)
	for {
		count++
		str := fmt.Sprintf("count: %d", count)

		// ① ただ単に書き込む。文字列が重なる
		tinyfont.WriteLine(display, &freeserif.Regular18pt7b, 10, 50, str, black)

		// ② 画面を白く塗ってから書き込む。チラつく
		display.FillRectangle(0x50, 320, 50, white)
		tinyfont.WriteLine(display, &freeserif.Regular18pt7b, 10, 50, str, black)

		// ③ Labelを使う
		label.FillScreen(white)
		tinyfont.WriteLine(label, &freeserif.Regular18pt7b, 10, 30, str, black)
		display.DrawRGBBitmap(0x100, label.Buf, label.W, label.H)

	}
}

type Label struct {
	Buf  []uint16
	W, H int16
}

func NewLabel(w, h, int16) *Label {
	return &Label{
		Buf: make([]uint16, w*h),
		W:   w,
		H:   h,
	}
}

func (l *Label) Size() (x, y, int16) {
	return l.W, l.H
}
