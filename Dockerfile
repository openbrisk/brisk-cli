# Build stage
FROM golang:1.9.2 as builder

WORKDIR /go/src/github.com/openbrisk/brisk-cli
COPY . .

# Run a gofmt and exclude all vendored code.
RUN test -z "$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*"))" || { echo "Run \"gofmt -s -w\" on your Golang code"; exit 1; }

# ldflags "-s -w" strips binary
# ldflags -X injects commit version into binary

RUN go test $(go list ./... | grep -v /vendor/ | grep -v /template/|grep -v /build/) -cover \
 && VERSION=$(git describe --all --exact-match `git rev-parse HEAD` | grep tags | sed 's/tags\///') \
 && GIT_COMMIT=$(git rev-list -1 HEAD) \
 && CGO_ENABLED=0 GOOS=linux go build --ldflags "-s -w \
    -X github.com/openbrisk/brisk-cli/version.GitCommitHash=${GIT_COMMIT} \
    -X github.com/openbrisk/brisk-cli/version.Version=${VERSION}" \
    -a -installsuffix cgo -o brisk-cli

# Release stage
FROM alpine:3.6

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/openbrisk/brisk-cli/brisk-cli .

CMD ["./brisk"]