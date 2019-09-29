.PHONY: images
images:

	docker build --no-cache -t build/general ./images/general

.PHONY: debug
debug:

	docker run --rm -it build/general bash

.PHONY: python3
python3:

	go run cmd/build.go -file "$(CURDIR)/examples/$@/config.json" > $(CURDIR)/examples/$@/build.sh && \
	docker run --rm -it -v "$(CURDIR)/examples/$@:/home/app" -w /home/app build/general bash build.sh

.PHONY: bash
bash:

	go run cmd/build.go -file "$(CURDIR)/examples/$@/config.json" > $(CURDIR)/examples/$@/build.sh && \
	docker run --rm -it -v "$(CURDIR)/examples/$@:/home/app" -w /home/app build/general bash build.sh

.PHONY: test
test:

	go test ./... -coverprofile=coverage.out

.PHONY: cover
cover:

	go tool cover -html=coverage.out
