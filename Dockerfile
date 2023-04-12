FROM golang

WORKDIR /app

COPY . .

RUN go build

CMD ["api_sensor"]