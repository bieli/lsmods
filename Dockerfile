FROM golang:1.21.6-bullseye

RUN apt-get update
RUN apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
RUN #apt-get install -y golang kmod
RUN apt-get install -y kmod

WORKDIR /go/src/app
COPY . .

RUN go build -o lsmods cmd/main.go
RUN chmod a+x lsmods

RUN ls -la /go/src/app/lsmods
RUN uname -a

CMD ["./lsmods"]
