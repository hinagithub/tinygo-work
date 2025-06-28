
build:
	tinygo build  -target wioterminal ./chapter5/button/main.go

flash:
	tinygo flash -target wioterminal ./chapter5/button/main.go