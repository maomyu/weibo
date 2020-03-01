FROM debian:latest

RUN mkdir -p /weibo
WORKDIR /weibo
COPY .  /weibo/
ENV PORT 9090
EXPOSE 9090
CMD ["./main"]