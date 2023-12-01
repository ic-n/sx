FROM golang:1.21

WORKDIR /opt/commitreveal

COPY go.mod /opt/commitreveal/
COPY go.sum /opt/commitreveal/
RUN go mod download

COPY cmd/ /opt/commitreveal/cmd/
COPY pkg/ /opt/commitreveal/pkg/
COPY contracts/ /opt/commitreveal/contracts/
RUN go build -o /opt/commitreveal/bin/commitreveal cmd/server/main.go

CMD [ "/opt/commitreveal/bin/commitreveal" ]
