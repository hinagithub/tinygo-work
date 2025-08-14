# TinyFont フォント生成手順

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