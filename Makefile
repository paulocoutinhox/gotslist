GOFMT=gofmt

GOFILES=\
	gotslist.go\

format:
	${GOFMT} -w gotslist.go
	${GOFMT} -w gotslist_test.go

test:
	go test -v