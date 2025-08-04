## tinydisplayのダウンロードコマンド

Makefileで実行
```
make tinydisplay-download
make tinydisplay-run
```

もしくは手動で

```

# zipダウンロード
curl -OL https://github.com/sago35/tinydisplay/releases/download/0.3.0/tinydisplay_0.3.0_macos_amd64.tgz

# 展開
tar xvzf tinydisplay_0.3.0_macos_amd64.tgz

# 起動
cd tinydisplay_0.3.0_macos_amd64
chmod 755 tinydisplay
# finderからctrl + 右クリックで「開く」押下


# サンプル実行
go run github.com/sago35/tinydisplay/examples/pyportal_boing@latest

<img width="333" height="280" alt="Image" src="https://github.com/user-attachments/assets/6fba5dbb-eb03-4e0e-be61-b8cb2bf4adff" />

```