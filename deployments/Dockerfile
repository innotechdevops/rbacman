# build stage
FROM golang:alpine AS builder

ENV GO111MODULE=on

# install git.
RUN apk update && apk add --no-cache git

RUN mkdir -p /go/src/github.com/innotechdevops/rbacman
WORKDIR /go/src/github.com/innotechdevops/rbacman
COPY . .

# Download all dependencies
RUN go mod download

# With go ≥ 1.10
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/rbacman cmd/rbacman/main.go

# small image
FROM alpine:3.7

WORKDIR /app
COPY --from=builder /go/src/github.com/innotechdevops/rbacman/configuration/production.yml /app/configuration/production.yml
COPY --from=builder /go/src/github.com/innotechdevops/rbacman/policy /app/policy
COPY --from=builder /go/bin/rbacman .

ENV TZ=Asia/Bangkok
RUN echo "Asia/Bangkok" > /etc/timezone

# run binary.
ENTRYPOINT ["/app/rbacman", "-env", "production"]