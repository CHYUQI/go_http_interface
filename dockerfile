FROM golang AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLE=0 GOOS=linux go build -o app ./main.go


FROM alpine AS runner

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD [ "./app" ]