# TinyGo組み込み開発学習リポジトリ

書籍「基礎から学ぶ TinyGoの組込み開発」 (C&R研究所) を読みながら手を動かしてみた学習記録のリポジトリ

## 参考書籍
- [基礎から学ぶ TinyGoの組込み開発](https://github.com/sago35/tinygobook) (C&R研究所)

## メモ

### VSCodeでコード補完が効かない(machine.LEDなど)
`.vscode/settings.json`でGOROOTをTinyGoのパスに設定して解決


## tinydisplay/examples/initdisplay のエラー解決方法

Wio Terminalで `github.com/sago35/tinydisplay/examples/initdisplay` を使う際、

```
cannot use machine.SPI3 (variable of type *machine.SPI) as machine.SPI value in argument to ili9341.NewSPI
```
のようなエラーが出る場合は、tinydisplayライブラリのバージョンが古い可能性があります。

### 解決手順

1. プロジェクトディレクトリで以下を実行し、tinydisplayを最新版にアップデートします。

```sh
cd chapter6/5_ili9341
# または該当ディレクトリへ移動

go get -u github.com/sago35/tinydisplay

go mod tidy
```

2. その後、通常通り `tinygo build` や `tinygo flash` でビルド・書き込みが可能になります。

---

この方法で、initdisplayを使ったまま最新のTinyGo環境で動作させることができます。


# 参考

[Woman Who Goのリポジトリ](https://github.com/WomenWhoGoTokyo/book-reading-party/tree/master/learn-embedded-development-with-tinygo)