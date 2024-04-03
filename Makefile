.PHONY: docker-build

docker-build:
	@echo building
	docker build -t stori .
	@echo Done
