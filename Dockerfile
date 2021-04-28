FROM scratch
COPY yoz-comment /usr/bin/yoz-comment
ENTRYPOINT ["/usr/bin/yoz-comment"]
