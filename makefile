
run:
	@go run pull/main.go

hotrun:
	@gowatch -o ./out/pull -p ./pull

run-push:
	@go run push/main.go

hotrun-push:
	@gowatch -o ./out/push -p ./push -args='dumps/kb'

build:
	@go build -o out/pull pull/main.go

build-push:
	@go build -o out/push push/main.go
	
clean:
	@rm -rf out
