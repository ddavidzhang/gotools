all: format run

file = gotcpspy

fmt_flags = -w=true 

format:
	gofmt $(fmt_flags) $(file).go

run:
	go run $(file).go $(args)

