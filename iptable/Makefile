REPOSITORY := kennychenfight/proxy_init
VERSION := v1

.PHONY: all
all: build push

.PHONY: build
build: image

.PHONY: image
image:
	docker build -t $(REPOSITORY):latest -t $(REPOSITORY):$(VERSION) .

.PHONY: push
push: build
	docker push $(REPOSITORY):$(VERSION)
	docker push $(REPOSITORY):latest
