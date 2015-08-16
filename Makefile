# Makefile for resource program

PROG = resource
VERSION = 0.1.0
TARDIR = $(HOME)/Desktop/TarPit/
DATE = 	`date "+%Y-%m-%d.%H_%M_%S"`
DOCOUT = README-$(PROG)-godoc.md

all:
	go build -v

# change cp to echo if you really don't want to install the program

install:
	go build -v
	go tool vet .
	go tool vet -shadow .
	gofmt -w *.go
#	go install
	cp $(PROG) $(HOME)/bin

docs:
	godoc2md . > $(DOCOUT)
	godepgraph -md -p . >> $(DOCOUT)
	deadcode -md >> $(DOCOUT)
	cp README-$(PROG).md README.md
	cat $(DOCOUT) >> README.md
	cp README.md README2.md

neat:
	go fmt ./...

dead:
	deadcode > problems.dead

index:
	cindex .

clean:
	go clean ./...
	rm -f *~ problems.dead count.out loki.go
#	rm -f $(DOCOUT)


tar:
	echo $(TARDIR)$(PROG)_$(VERSION)_$(DATE).tar
	tar -ncvf $(TARDIR)$(PROG)_$(VERSION)_$(DATE).tar .

test:
# remove possibly broken version of loki.go first (only needed during debug)
	rm -f loki.go
	go build -v
	./resource -source="loki.jpg" -rc="loki.go" -var="lokiJpgBites"
	gofmt -w *.go

# Coverage test maker
#cover:
#	go test -covermode=count -coverprofile=count.out
#	cover -html=count.out
