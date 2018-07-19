FROM alpine:3.8

WORKDIR /bin/

COPY app .

EXPOSE 8080

# Uncomment to run the seed before running the server:
# CMD /bin/app seed && /bin/app daemon
CMD exec /bin/app daemon
