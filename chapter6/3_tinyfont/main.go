package main

import (
	"image/color"

	"github.com/sago35/tinydisplay/examples/initdisplay"
	"tinygo.org/x/tinyfont"
)

func main() {
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
