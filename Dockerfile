FROM golang:alpine 
RUN apk update && apk add --no-cache git
RUN go get github.com/go-sql-driver/mysql
EXPOSE 8081
ENV mySQLIPAddress 192.168.67.241
ENV mySQLIPPort 3306
ENV mySQLPassword tushar321
RUN echo $mySQLIPAddress
RUN echo $mySQLIPPort
RUN echo $mySQLPassword
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]