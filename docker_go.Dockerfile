FROM golang:1.24 AS build

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/cmd ./cmd
COPY backend/pgk ./pgk

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o server ./cmd/server

FROM alpine:3.22

RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Europe/Oslo

COPY --from=build /app/server /app/server

EXPOSE 8182

CMD ["/app/server"]
