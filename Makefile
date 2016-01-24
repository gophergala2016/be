build: clean
	go build -o bin/be

clean:
	rm -rf bin

container:
	docker build -t be .
