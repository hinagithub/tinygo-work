
build:
	tinygo build  -target wioterminal ./chapter5/button2/main.go

flash:
	tinygo flash -target wioterminal ./chapter5/button2/main.go

monitor:
	tinygo monitor --target wioterminal

minicom:
	minicom -D /dev/tty.usbmodem14301