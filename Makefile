.PHONY: install
install:
	### Clear ###
	rm -rf .dbs/
	
	### Replace version before install ###
	sed -i "s/~@DBS_VERSION@~/"$$(git rev-parse HEAD)"/g" src/help/help.go
	go install
	
	### Revert changes ###
	git checkout -- src/help/help.go
clean:
	rm -rf .dbs/
