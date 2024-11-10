FOLDER ?=

.PHONY: generate

generate:
ifeq ($(FOLDER),)
	@echo "Set folder name FOLDER=promlem-name-with-number"
	@exit 1
endif
	@mkdir -p $(FOLDER)
	@echo "Generate Go-files $(FOLDER).go and $(FOLDER)_test.go"
	
	@echo "package leetcode\n" > $(FOLDER)/$(FOLDER).go
	@echo "package leetcode\n" > $(FOLDER)/$(FOLDER)_test.go
