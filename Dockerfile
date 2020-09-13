FROM golang:alpine3.11 AS builder
WORKDIR /github.com/alimy/tl-explorer
RUN apk --no-cache --no-progress add --virtual \
  build-deps \
  build-base \
  git
COPY . .
RUN git submodule update --init
RUN make build

FROM golang:alpine3.11
WORKDIR /app
COPY --from=builder /github.com/alimy/tl-explorer/tl-explorer .
EXPOSE 8080
ENTRYPOINT ["/app/tl-explorer"]
