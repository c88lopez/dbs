.PHONY: install
install:
	rm -rf .dbs/
	go install

clean :
	rm -rf .dbs/