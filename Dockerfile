FROM golang:1.26-alpine

WORKDIR /home/app

COPY ./Backend-Server /home/app

RUN go build -o job-grinder .

CMD ["./job-grinder"]