FROM golang:1.12.0-alpine3.9
RUN apk add --no-cache git &&\ 
    go get "github.com/gorilla/mux" &&\
    go get "github.com/lib/pq"
ADD /src/ /app/ 
WORKDIR /app/
RUN go build -o main .
CMD ["/app/main"]



    
