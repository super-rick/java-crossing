# 多阶段构建:golang 构建 → scratch 运行(静态二进制,≈10MB)
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM scratch
COPY --from=builder /app/app /app
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app"]
