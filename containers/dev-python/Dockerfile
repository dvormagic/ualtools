
FROM python:2.7-slim

RUN apt-get update && \
    apt-get install -y zip unzip

RUN mkdir /home/container && \
    chmod 0777 /home/container && \
    mkdir /data && \
    chmod 0777 /data

COPY run.sh /opt/run.sh

CMD ["/opt/run.sh"]
