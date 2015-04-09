release:
	go get
	go build
	tar -zcvf build.tar.gz nexii README.md
