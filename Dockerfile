FROM golang:1.13

RUN go get github.com/oxequa/realize
ENTRYPOINT ["realize"]
CMD ["start"]
