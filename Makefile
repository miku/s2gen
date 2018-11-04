s2gen: cmd/s2gen/main.go schema.go
	go build -o $@ $<

test:
	go test -v ./...

clean:
	rm -f s2gen
