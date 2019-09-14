.PHONY: images
images:

	docker build --no-cache -t build/general ./images/general

.PHONY: debug
debug:

	docker run --rm -it build/general bash

.PHONY: python3
python3:

	go run bash.go python3 && \
	docker run --rm -it -v "$(CURDIR)/examples/python3:/home/app" -w /home/app build/general bash build.sh

.PHONY: bash
bash:

	go run bash.go bash && \
	docker run --rm -it -v "$(CURDIR)/examples/bash:/home/app" -w /home/app build/general bash build.sh
