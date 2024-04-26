FROM tomyj/my-env:go
WORKDIR /app
ENV DB_HOSTNAME=database
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o api .
CMD [ "./api" ]