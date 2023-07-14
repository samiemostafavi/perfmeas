FROM golang:1.19-alpine
RUN go install github.com/samiemostafavi/advmobileinfo/cmd/ami@latest
EXPOSE 50500

ENTRYPOINT ["ami"]
