.PHONY: fedex royalmail

fedex:
	go build -buildmode=plugin -o ./plugins/fedex.so fedex/fedex.go

royalmail:
	go build -buildmode=plugin -o ./plugins/royalmail.so royalmail/royalmail.go