## tinygoのビルド



buildというサブコマンドを使う

```
tinygo build -o hello.uf2 -target wioterminal -size short examples/blinky1
```

#### 何してる？
Go言語のコードを小さなバイナリにコンパイル

#### オプション
- -o: 出力先
- -target: ターゲット
- -size: サイズ

ターゲットに指定できるのは他だとたとえば
- pico (Raspberry Pi Pico)

など。

#### 用語解説
- uf2: USB Flashing Format
Microsoftが開発したファイルフォーマットで、特にマイクロコントローラへのファームウェア書き込みを容易にするために設計されました。Raspberry Pi Picoや一部のAdafruitボードなどで広く採用されています。

UF2が使われる背景:
従来のマイコンへの書き込みは、特定の書き込みツールやドライバが必要だったり、コマンドライン操作が必要だったりと、初心者には敷居が高いものでした。MicrosoftがPXT (現在のMakeCode) のためにUF2を開発した目的は、子供を含むより多くのユーザーが簡単にマイコンにプログラムを書き込めるようにすることでした。USBメモリのようにファイルをコピーするだけで書き込みが完了する手軽さが、その目的を達成しています。
