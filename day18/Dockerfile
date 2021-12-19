FROM golang:1.17-alpine as BASE
WORKDIR /src
COPY ./ .
RUN go build -o /bin/aoc

FROM alpine:latest
COPY --from=BASE /bin/aoc /bin/aoc
COPY entrypoint.sh . 
RUN chmod +x ./entrypoint.sh
RUN apk add coreutils
COPY input.txt .
ENTRYPOINT ["./entrypoint.sh", "/bin/aoc"]