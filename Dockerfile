FROM golang:1.16-alpine as build
WORKDIR /app
ENV GOPROXY https://mirrors.aliyun.com/goproxy/
ADD go.mod /app
ADD go.sum /app
RUN go mod download
ADD main.go /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
COPY --from=build /app/app /app
EXPOSE 8080
CMD [ "/app" ]
