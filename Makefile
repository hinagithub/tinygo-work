
build:
	tinygo build  -target wioterminal ./chapter5/button2/main.go

flash:
	tinygo flash -target wioterminal ./chapter5/button2/main.go

monitor:
	tinygo monitor --target wioterminal

minicom:
	minicom -D /dev/tty.usbmodem14301

# tinydisplay setup
tinydisplay-download:
	curl -OL https://github.com/sago35/tinydisplay/releases/download/0.3.0/tinydisplay_0.3.0_macos_amd64.tgz
	tar xvzf tinydisplay_0.3.0_macos_amd64.tgz
	cd tinydisplay_0.3.0_macos_amd64 && chmod 755 tinydisplay

tinydisplay-run:
	go run github.com/sago35/tinydisplay/examples/pyportal_boing@latest