FROM golang:latest

WORKDIR /go/src/hrdwr
COPY . .

RUN apt-get install -y lm-sensors

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["hrdwr"]