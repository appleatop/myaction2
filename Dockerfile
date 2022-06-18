FROM golang:1.18.2-bullseye AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download  
COPY . ./  
RUN go test ./test2 -c -o test2_program
RUN go test ./test3 -c -o test3_program
RUN cd /build/test1/mainprogram &&\
 go run mainprogram.go &&\
 go build -o mainprogram mainprogram.go 
# TODO - can't get version in container
FROM build AS test 
WORKDIR /test
COPY --from=build /build ./
RUN cd /test/test2 &&\
 go test 
RUN ./test2_program
RUN ./test3_program
# TODO - can't get version in contaieet i
FROM alpine:latest AS integration_test 
WORKDIR /app/
COPY --from=test /test ./
COPY --from=build /build/test1/mainprogram ./
RUN cp mainprogram mainprogram2 && chmod 755 ./mainprogram2
COPY --from=build /build/entrypoint.sh ./
COPY --from=build /build/testintegration.sh ./
RUN chmod 755 /app/entrypoint.sh
RUN chmod 755 /app/testintegration.sh
RUN ls -al /app/
CMD ["/app/entrypoint.sh"]

