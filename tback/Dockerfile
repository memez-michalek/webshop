FROM golang:1.16-buster

WORKDIR /app

#RUN apt get git gcc libc-dev

RUN go get -u -v github.com/cosmtrek/air

COPY . .

CMD ["air", "-d"]