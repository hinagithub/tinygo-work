#### VSCodeの設定
- .settings.jsonを編集
https://tinygo.org/docs/guides/ide-integration/vscode/

・・・デフォルトではmachine.LEDが見つからないなど上手く動かなかったのでGOROOTを書き換え

TINYGOROOTを調べる
```
tinygo-work >>>tinygo env TINYGOROOT
/usr/local/Cellar/tinygo/0.37.0
```

settings.jsonに設定
```
"GOROOT": "/usr/local/Cellar/tinygo/0.37.0",
```

- コマンドパレットでTinyGo targetを選択 > Ardinoを選択
