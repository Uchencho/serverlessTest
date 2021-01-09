OUTPUT = main # will be archived
ZIPFILE = lambda.zip

.PHONY: test
test:
	go test ./...

clean:
	rm -f $(OUTPUT)
	rm -f $(ZIPFILE)

main: main.go
	go build -o $(OUTPUT) main.go

build-local:
	go build -o $(OUTPUT) main.go

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) main

# create a lambda deployment package
$(ZIPFILE): clean lambda
	zip -9 -r $(ZIPFILE) $(OUTPUT)

run: build-local
	@echo ">> Running application ..."
	PORT=4000 \
	./$(OUTPUT)
	