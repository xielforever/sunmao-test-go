FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/app .

FROM alpine:3.19
RUN adduser -D -H appuser
COPY --from=build /out/app /app
USER appuser
EXPOSE 8080
ENV PORT=8080
ENTRYPOINT ["/app"]
