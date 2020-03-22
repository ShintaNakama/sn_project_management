FROM golang:1.13
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo


RUN go get github.com/oxequa/realize
ENTRYPOINT ["realize"]
CMD ["start"]
