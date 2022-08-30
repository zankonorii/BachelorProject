FROM golang:1.18 AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /src

# Moduls
COPY go.sum go.mod /src/

WORKDIR /src

# Download modules
RUN go mod download

COPY . /src

# Build components.
# Put built binaries and runtime resources in /app dir ready to be copied over or used.
RUN go build -o usersService

RUN ls

CMD ["./usersService"]
