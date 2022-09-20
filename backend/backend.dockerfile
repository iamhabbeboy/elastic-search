FROM golang:1.18-alpine as builder

WORKDIR /app

COPY . ./

RUN go build -v -o /web cmd/web/main.go

CMD ["/web"]

#-----------------HOT-RELOAD-----------------
# WORKDIR /app
# ENV CGO_ENABLED 0 
# ENV GOOS linux 
# COPY . .
# RUN go get github.com/githubnemo/CompileDaemon
# RUN chmod +x ./hot-reload.sh
# ENTRYPOINT [ "./hot-reload.sh" ]