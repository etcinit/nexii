release:
	go get
	go get -u -v github.com/laher/goxc
	goxc -tasks='xc archive' -bc 'linux windows darwin' -d .
	#tar -zcvf build-win32.tar.gz release/nexii.exe README.md
clean:
	rm -rf snapshot
