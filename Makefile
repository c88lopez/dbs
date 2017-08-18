
install:
	### Clearing .dbs folder ###
	@rm -rf .dbs/
	
	### Replacing version ###
	@sed -i "s/~@DBS_VERSION@~/"$$(git rev-parse HEAD)"/g" src/help/help.go

	### Installing ###
	@go install
	
	### Reverting changes ###
	@git checkout -- src/help/help.go

	### Done! ###
.PHONY: install

clean:
	rm -rf .dbs/
