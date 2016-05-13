NAME=jsonptr
IMPORT=github.com/mickep76/${NAME}

all: test readme

format:
	gofmt -w=true .

test: format
	golint ${NAME}.go
	go vet ${NAME}.go
#	go test

readme:
	cat README.header >README.md
	godoc2md . >>README.md
	sed -i .bak 's!import "\."!import "${IMPORT}"!' README.md
	rm README.md.bak
