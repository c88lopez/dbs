.PHONY: install
install:
	rm -rf .dbs/
	
	rpl "~@DBS_VERSION@~" "$$(git rev-parse HEAD)" help/help.go
	go install
	
	git checkout -- help/help.go
clean:
	rm -rf .dbs/
