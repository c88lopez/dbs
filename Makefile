install: applyversionhash goinstall revertversionchanges
	@echo "### Done! ###"
.PHONY: install

removefolder:
	@echo "### Clearing .dbs folder ###"
	@rm -rf .dbs/

applyversionhash:
	@echo "### Replacing version ###"
	@sed -i "s/~@DBS_VERSION@~/"$$(git rev-parse HEAD)"/g" cmd/version.go

goinstall:
	@echo "### Installing ###"
	@go install -ldflags "-w"

revertversionchanges:
	@echo "### Reverting changes ###"
	@git checkout -- cmd/version.go
