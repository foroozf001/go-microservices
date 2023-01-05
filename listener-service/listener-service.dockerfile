FROM golang:1.19-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o listenerApp .

FROM alpine:3.17.0
LABEL maintainer="farazforoozan@gmail.com"
EXPOSE 80/tcp
RUN mkdir /app
COPY --from=builder /app/listenerApp /app
RUN adduser -HD -u 1500 -g user user
WORKDIR /app
USER user
CMD ["/app/listenerApp"]