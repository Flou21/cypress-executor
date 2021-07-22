FROM golang:1.16.6-buster as build

WORKDIR /src

COPY . .

RUN go build .


FROM registry.flou21.de/firefox-cypress-browser:firefox86

RUN apt update && apt install git -y

COPY --from=build /src/main /app

CMD /app

