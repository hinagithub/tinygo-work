# Wi-Fi Access Point Test

Wio TerminalでWi-Fi接続をテストします。

## 使い方

```bash
# Wifi設定
export WIFI_SSID="tullys_Wi-Fi"
export WIFI_PASSWORD=""
tinygo flash -target=wioterminal main.go
```

## トラブルシューティング

環境変数エラーが出る場合は上記のexportコマンドを実行してください。
