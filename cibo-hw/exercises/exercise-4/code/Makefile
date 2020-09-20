.PHONY: default
default: karma

karma: $(SOURCES)
	CGO_ENABLED=0 go build ${LDFLAGS} -o karma .

.PHONY: clean
clean:
	rm -f karma

.PHONY: release_binary
release_binary: clean
	GOOS=linux go build ${LDFLAGS} -a -o karma .

.PHONY: docker
docker: release_binary
	docker build --no-cache -t karma:latest .
	docker tag karma:latest 235156212133.dkr.ecr.us-east-1.amazonaws.com/karma:latest
	docker push 235156212133.dkr.ecr.us-east-1.amazonaws.com/karma:latest
