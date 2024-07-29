FROM golang:1.22-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download 

COPY [^go.sum]* ./
COPY src src 
COPY proto proto 
COPY config.yml /app

WORKDIR /app/src
RUN go build =o conduit
RUN chmod a+x /conduit

EXPOSE 3001
CMD ["/conduit"]