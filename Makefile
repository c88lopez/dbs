.PHONY: install
install:
	rm -rf .dbs/

	rpl "~@DBS_VERSION@~" "$(git rev-parse HEAD)" help/help.go
	go install
	rpl "$(git rev-parse HEAD)" "~@DBS_VERSION@~" help/help.go

clean :
	rm -rf .dbs/