FROM golang:1.16.6-buster as build

WORKDIR /src

COPY . .

RUN go build .


FROM cypress/browsers:node14.17.0-chrome91-ff89

RUN apt update && apt install git -y

COPY --from=build /src/main /app

CMD /app

