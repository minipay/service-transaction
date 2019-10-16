FROM golang:1.13.0-alpine
# RUN apk add git
# RUN git clone https://github.com/vishnubob/wait-for-it.git
RUN apk add tzdata && \
        cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
        echo "Asia/Jakarta" >  /etc/timezone && \
        date && \
        apk del tzdata
COPY . /home/app/
WORKDIR /home/app/
# CMD ./wait-for-it/wait-for-it.sh --host=mysql --port=3307 --timeout=60 -- "go run main.go"
CMD go run main.go
EXPOSE 5050