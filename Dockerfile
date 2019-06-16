FROM bigrocs/alpine:ca-data

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY ./bin/service /usr/local/bin/
CMD ["service"]