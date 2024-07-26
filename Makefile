# Makefile

include .env

.PHONY: run

postgres_init:
	./scripts/script.sh

api_run:
	docker build -t $(IMAGE_NAME) .
	docker run --rm -p 3000:3000  $(IMAGE_NAME)
