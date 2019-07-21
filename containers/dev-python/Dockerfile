
FROM python:2.7-slim

RUN apt-get update && \
    apt-get install -y zip unzip

COPY run.sh /opt/run.sh

CMD ["/opt/run.sh"]
