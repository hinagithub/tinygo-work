# RTL8720DNのファームウェアアップデート

## 1. ambd_flash_toolのセットアップ

```bash
git clone https://github.com/Seeed-Studio/ambd_flash_tool
cd ambd_flash_tool

# Python仮想環境を作成（macOS Homebrew環境でのエラー回避）
python3 -m venv venv
source venv/bin/activate

# RTL8720DNファームウェアをerase
python3 ambd_flash_tool.py erase
```

**注意：** erase処理中に「Sorry, the device you should have is not plugged in.」と表示されることがありますが、「Verify successful」が表示されていれば処理は成功しています。

## 2. TinyGoネットワーク機能のテスト

### 問題: initialize関数が使用できない

```bash
go get -u github.com/sago35/tinygo-examples/wioterminal/initialize
# エラー: tinygo.org/x/drivers/net: cannot find module providing package tinygo.org/x/drivers/net
```

**原因：** `github.com/sago35/tinygo-examples/wioterminal/initialize`は古いTinyGoサンプルで、現在の`tinygo.org/x/drivers`には`net`パッケージが含まれていません。

### 解決策

1. **TinyGo公式の最新ネットワークサンプルを使用**
2. **または、initialize.goをローカルにコピーして不要な依存関係を修正**

現在のTinyGoバージョンでは、このサンプルはそのままでは動作しません。

## トラブルシューティング

### Python環境エラー（解決済み）
macOSで`ambd_flash_tool.py`実行時に以下のエラーが出る場合：
```
× This environment is externally managed
```
→ **解決方法:** Python仮想環境を使用（上記手順1参照）
