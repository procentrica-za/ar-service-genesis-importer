FROM golang:alpine
RUN apk add --no-cache git 
ADD /src/ /app/ 
WORKDIR /app/
RUN go build -o main .
CMD ["/app/main"]



    
