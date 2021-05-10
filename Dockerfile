FROM golang:alpine AS builder

RUN apk update && apk add --no-cache make bash gcc musl-dev libc-dev ca-certificates curl
RUN adduser -D -g '' appuser

WORKDIR /app
COPY . .

RUN make build


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/bin/auth-svc /app/bin/auth-svc

USER appuser
EXPOSE 3000

ENTRYPOINT ["/app/bin/auth-svc"]
