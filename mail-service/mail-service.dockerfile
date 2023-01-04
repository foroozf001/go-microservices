FROM golang:1.19-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o mailApp ./cmd/api

FROM alpine:3.17.0
LABEL maintainer="farazforoozan@gmail.com"
EXPOSE 80/tcp
RUN mkdir /app
COPY --from=builder /app/mailApp /app
COPY --from=builder /app/templates /app/templates
RUN adduser -HD -u 1500 -g user user
WORKDIR /app
USER user
CMD ["/app/mailApp"]