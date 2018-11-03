solrstructgen: cmd/solrstructgen/main.go schema.go
	go build -o $@ $<

clean:
	rm -f solrstructgen
