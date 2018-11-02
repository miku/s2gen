solrstructgen: cmd/solrstructgen/main.go
	go build -o $@ $<

clean:
	rm -f solrstructgen
