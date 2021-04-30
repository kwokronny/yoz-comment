FROM scratch

WORKDIR /app

COPY yoz-comment /app/yoz-comment
EXPOSE 9975
VOLUME ["/app"]
ENTRYPOINT ["/app/yoz-comment"]
