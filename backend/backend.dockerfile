FROM golang:1.18-buster as deps

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

# # #-----------------BUILD-----------------
FROM deps AS app-build

RUN go build -v -o ./web cmd/web/main.go

CMD ["./web"]

#-----------------HOT-RELOAD-----------------

# FROM deps AS hot-reload
# WORKDIR /app
# ENV CGO_ENABLED 0 
# ENV GOOS linux 
# RUN go install github.com/githubnemo/CompileDaemon
# COPY ./hot-reload.sh hot-reload.sh

# RUN sh hot-reload.sh
# ENTRYPOINT [ "hot-reload.sh" ]
# RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
#     && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# CMD air

# go get github.com/pilu/fresh


# FROM golang:1.18-alpine as builder

# WORKDIR /app

# COPY . .

# RUN go build -v -o /web cmd/web/main.go
# # RUN go build -o /cmd/web

# CMD ["/web"]

# FROM golang:1.18-alpine as builder
# COPY . /app
# WORKDIR /app
# RUN CGO_ENABLED=0 go build -o backendApp ./cmd/web
# RUN chmod +x ./backendApp
# #build a tiny docker image 

# # FROM alpine:latest
# # RUN mkdir /app
# COPY --from=builder ./backendApp /.
# CMD ["./backendApp"]

#-----------------HOT-RELOAD-----------------
# WORKDIR /app
# ENV CGO_ENABLED 0 
# ENV GOOS linux 
# COPY . .
# RUN go get github.com/githubnemo/CompileDaemon
# RUN chmod +x ./hot-reload.sh
# ENTRYPOINT [ "./hot-reload.sh" ]