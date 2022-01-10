build:
	go build -o bin/sort src/sort.go

clean:
	rm -f temp/*

test:
	bin/sort example1.dat output

check:
	bin/valsort output
