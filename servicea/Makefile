REPOSITORY := kennychenfight/service-a
VERSION := v1

.PHONY: all
all: image

.PHONY: image
image: build-image push-image

.PHONY: build-image
build-image:
	@echo "Building the docker image: $(REPOSITORY)..."
	docker build --no-cache -t $(REPOSITORY):latest -t $(REPOSITORY):$(VERSION) .

.PHONY: push-image
push-image: build-image
	@echo "Pushing the docker image for $(REPOSITORY)/$(VERSION) and $(REPOSITORY):latest..."
	docker push $(REPOSITORY):$(VERSION)
    docker push $(REPOSITORY):latest