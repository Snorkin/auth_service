FROM golang:1.22

#ENV TZ=Europe/Moscow
#RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 5000

ENTRYPOINT CompileDaemon --build="go build cmd/auth/main.go" --command=./main
