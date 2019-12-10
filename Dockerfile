FROM golang:1.13.1 AS builder

ARG app_name=user

WORKDIR /go/src/${app_name}

COPY . /go/src/${app_name}

RUN go get . 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  

ARG app_name=user

RUN apk --no-cache add ca-certificates

WORKDIR /root/

ENV RUN_MODE=production

COPY --from=builder /go/src/${app_name}/app .

COPY --from=builder /go/src/${app_name}/prod.yml .

CMD ["./app"]  
