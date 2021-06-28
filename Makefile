default:
	cd ./src && go build -o ../bifocals

.PHONY: install
install:
	cp ./bifocals /usr/local/bin/bifocals
