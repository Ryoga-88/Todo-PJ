FROM golang:1.23.0

WORKDIR /go/src/github.com/Ryoga-88/Todo-PJ/backend

COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@v1.29.0

COPY backend .

CMD [ "air" ]