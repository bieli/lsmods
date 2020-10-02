FROM ubuntu:20.04

RUN apt-get update
RUN apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
RUN apt-get install -y golang kmod

WORKDIR /go/src/app
COPY . .

RUN go build -i main.go
RUN chmod a+x main.go

RUN ls -la /go/src/app/main
RUN uname -a

CMD ["./main"]
