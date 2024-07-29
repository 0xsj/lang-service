FROM golang:1.22-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sun ./
RUN go mod download 


COPY [^go.sum]* ./
COPY src src 
COPY proto proto 
COPY config.yml /app

WORKDIR /app/src
RUN go build =o user-service
RUN chmod a+x /user-service

EXPOSE 3002
CMD ["/user-service"]