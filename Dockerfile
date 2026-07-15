FROM golang:1.22.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION=0.1.0-dev
ARG GIT_COMMIT=unknown
ARG BUILD_TIME=unknown

RUN go build -ldflags "\
    -X github.com/matteoaricci/jot-api/version.Version=${VERSION} \
    -X github.com/matteoaricci/jot-api/version.GitCommit=${GIT_COMMIT} \
    -X github.com/matteoaricci/jot-api/version.BuildTime=${BUILD_TIME}" \
    -o main .

EXPOSE 8080
CMD ["./main", "-local"]