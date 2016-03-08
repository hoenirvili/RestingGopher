all :
	go build
	./RestingGopher
clean:
	rm RestingGopher
	echo "" > server.log
	cat server.log
e:
	cat server.log
