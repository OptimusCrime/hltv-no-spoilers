FROM golang:1.23

ENV TZ=Europe/Oslo

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/cmd ./cmd
COPY backend/pgk ./pgk

RUN go build -o server ./cmd/server

EXPOSE 8190

CMD ["/app/server"]
