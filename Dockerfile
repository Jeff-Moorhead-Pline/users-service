FROM golang:1.18-bullseye

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY api/ ./api/
COPY cmd/ ./cmd/
COPY data/ ./data/
RUN go mod download
RUN mkdir ./bin
RUN go build -o ./bin/users-service ./cmd/*.go
EXPOSE 8080
CMD [ "./bin/users-service" ]
