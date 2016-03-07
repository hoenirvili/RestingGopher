all :
	go build
	./RestingGopher
reset:
	rm RestingGopher
	echo "" > server.log
	cat server.log
e:
	cat server.log
