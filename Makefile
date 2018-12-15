pp:
	go run ./pipe/pipe.go

fork:
	go build ./forkexec/script.go
	go run ./forkexec/forkexec.go
