.PHONY: test
test:
		go run main.go -d _example
cfgtest:
		go run main.go -c _example/te.config.json
