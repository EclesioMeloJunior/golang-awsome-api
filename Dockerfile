FROM golang:1.15 as build

WORKDIR /app
COPY . ./

RUN go build -o app.so ./main.go

FROM debian

EXPOSE 8080

WORKDIR /
COPY --from=build /app/app.so .

RUN apt-get update
RUN apt-get install ca-certificates -y
RUN ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

ENTRYPOINT [ "./app.so" ]