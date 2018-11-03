solrstructgen: cmd/solrstructgen/main.go schema.go
	go build -o $@ $<

test:
	go test -v ./...

clean:
	rm -f solrstructgen
