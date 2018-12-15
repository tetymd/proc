pp:
	go run ./pipe/pipe.go

fork:
	go build -o ./forkexec/script ./forkexec/script.go
	go run ./forkexec/forkexec.go
	rm ./forkexec/script
