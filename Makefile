.PHONY: install
install:
	### Clear ###
	rm -rf .dbs/
	
	### Replace version before install ###
	rpl -q "~@DBS_VERSION@~" "$$(git rev-parse HEAD)" help/help.go
	go install
	
	### Revert changes ###
	git checkout -- help/help.go
clean:
	rm -rf .dbs/
