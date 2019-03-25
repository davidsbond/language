################ STEP 1 #################
FROM golang:1.11-alpine as builder

# Copy source
COPY . /src
WORKDIR /src

ARG VERSION=latest

RUN adduser -D -g '' container
RUN apk update && apk add git
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s -X main.version=${VERSION}" -o /bin/language

################ STEP 2 #################
FROM scratch as main

# Copy binary from first step
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bin/language /bin/language

USER container
ENTRYPOINT [ "/bin/language" ]