FROM golang:alpine as build-stage

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o api ./cmd/main.go


FROM alpine as production-stage

WORKDIR /app

COPY --from=build-stage /app/api .

EXPOSE 8080

CMD ./api
