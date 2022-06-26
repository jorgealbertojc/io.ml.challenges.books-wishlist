


.PHONY: bootstrap
bootstrap:
	@$(shell which go) mod tidy &&\
	$(shell which go) mod vendor
