## やってること
画面チカチカ

## コマンド

- wioterminalへの書き込み
```
tinygo flash --target wioterminal
```
 
- シリアル通信の出力結果を見る

```
minicom -D /dev/tty.usbmodem14301
```

<img width="473" alt="Image" src="https://github.com/user-attachments/assets/57ac362f-f251-49df-a66f-db1dcc36bd9b" />

- monitorコマンドでもいい

```
tinygo monitor --target wioterminal
```

<img width="453" alt="Image" src="https://github.com/user-attachments/assets/0d5bd133-6523-4081-b808-38653056ef00" />