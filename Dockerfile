FROM alpine:3.6

EXPOSE 8080

COPY goneup /opt/
COPY static /opt/static

CMD ["./opt/goneup"]
