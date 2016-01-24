build: clean
	go build

clean:
	rm -f be

container:
	docker build -t be .
