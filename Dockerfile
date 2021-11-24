FROM golang:1.16.9

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY=https://goproxy.cn
	
WORKDIR /build

COPY . .

RUN go mod tidy

RUN go build -o bing_pic .

WORKDIR /app
RUN mv /build/bing_pic .
RUN rm -rf /build
EXPOSE 1000

CMD ["/app/bing_pic"]
