FROM golang:1.18.2-bullseye AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download  
COPY . ./  
RUN cd /build/test1/mainprogram &&\
 go run mainprogram.go &&\
 go build -o mainprogram mainprogram.go

# TODO - can't get version in contaier
FROM build AS test 
WORKDIR /test
COPY --from=build /build ./
RUN cd /test/test2 &&\
	go test

# TODO - can't get version in contaier
FROM alpine:latest AS production
WORKDIR /app/
COPY --from=build /build/test1/mainprogram/mainprogram ./
RUN ls -al /app/mainprogram
COPY --from=build /build/entrypoint.sh ./
RUN chmod 755 /app/entrypoint.sh
CMD ["/app/entrypoint.sh"]

