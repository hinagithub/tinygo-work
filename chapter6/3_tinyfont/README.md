# TinyFont フォント生成手順

<img width="329" height="277" alt="Image" src="https://github.com/user-attachments/assets/985ea149-4dcd-4e9f-8a80-170db49e4c5a" />

## 1. tinyfontgen-ttfツールのインストール

```bash
go install tinygo.org/x/tinyfont/cmd/tinyfontgen-ttf@dev
```

## 2. Noto Sans JP フォントのダウンロード

```bash
curl -OL https://github.com/google/fonts/raw/main/ofl/notosansjp/NotoSansJP%5Bwght%5D.ttf
```

## 3. フォントファイル生成

```bash
~/go/bin/tinyfontgen-ttf --output font.go --fontname NotoSans32pt --size 32 "NotoSansJP%5Bwght%5D.ttf" --string "こんにちは"
```

## 生成されるファイル

- `font.go`: 「こんにちは」の文字データを含むGoフォントファイル
- フォント名: `NotoSans32pt`
- サイズ: 32pt

## 使用方法

生成された `font.go` ファイルをTinyGoプロジェクトで使用して、日本語文字を表示できます。

## 実行手順

### tinydisplayエミュレータで実行

1. **tinydisplayエミュレータを起動**

```bash
cd ./tinydisplay_0.3.0_macos_amd64
./tinydisplay &
```

2. **プログラム実行**

```bash
cd ./chapter6/3_tinyfont
go run .
```



### Wio Terminalで実行

```bash
cd ./chapter6/3_tinyfont
tinygo flash --target wioterminal .
```

## トラブルシューティング

### `undefined: NotoSans32pt` エラー

**原因**: TinyGoはパッケージレベルの変数初期化順序に敏感で、`font.go`と`main.go`の初期化順序が不定のため、`NotoSans32pt`が未定義になる場合があります。
Goでは、パッケージレベルの変数は宣言された順序で初期化されますが、TinyGoの場合は特に・・・
- font.goとmain.goの初期化順序が不定
- main.goが先に処理されると、NotoSans32ptがまだ未定義
- コンパイラが依存関係を正しく解決できない
となることがある

**解決法**: `main()` 関数内で変数を宣言

```go
// NG: パッケージレベルの変数宣言
var font = &NotoSans32pt

// OK: 関数内で宣言
func main() {
    font := &NotoSans32pt  // 実行時に初期化されるため安全
    // ...
}
```

### `connection refused` エラー

tinydisplayエミュレータが起動してない。先にエミュレータを起動してからプログラムを実行