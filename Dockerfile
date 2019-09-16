FROM golang:1.13.0-alpine
RUN apk add tzdata && \
        cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
        echo "Asia/Jakarta" >  /etc/timezone && \
        date && \
        apk del tzdata
COPY . /home/app/
WORKDIR /home/app/
CMD ["go","run","main.go"]
EXPOSE 5050