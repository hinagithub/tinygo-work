# TinyGo組み込み開発学習リポジトリ

書籍「基礎から学ぶ TinyGoの組込み開発」 (C&R研究所) を読みながら手を動かしてみた学習記録のリポジトリ

## 参考書籍
- [基礎から学ぶ TinyGoの組込み開発](https://github.com/sago35/tinygobook) (C&R研究所)

## よく使うコマンド

### ビルド
```bash
tinygo build -o hello.uf2 -target wioterminal -size short main.go
```

### フラッシュ書き込み
```bash
tinygo flash --target wioterminal
```

### シリアル通信監視
```bash
# minicomを使用
minicom -D /dev/tty.usbmodem14301

# TinyGo monitorを使用（推奨）
tinygo monitor --target wioterminal
```

## メモ

### VSCodeでコード補完が効かない(machine.LEDなど)
`.vscode/settings.json`でGOROOTをTinyGoのパスに設定して解決
