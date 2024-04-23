FROM tomyj/my-env:go
WORKDIR /app
ENV DB_HOSTNAME=database
COPY . .
RUN go build -o api .
CMD [ "./api" ]