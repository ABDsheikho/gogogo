exe=./bin/gogogo
test_dir=~/Desktop/projects/test_gogogo
mod_path=testMod

build:
	@go build -o $(exe) main.go

test:
	@make build
	$(exe) $(test_dir) $(mod_path) $(flags)

clean:
	@rm -rf $(test_dir)

help:
	@$(exe) help
