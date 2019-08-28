
FROM maven:3.3-jdk-8

RUN mkdir /home/container && \
    chmod 0777 /home/container && \
    mkdir /home/container/.m2 && \
    chmod 0777 /home/container/.m2 && \
    mkdir /data && \
    chmod 0777 /data

RUN groupadd --gid 1000 -r localgrp -o && \
    useradd --system --uid=1000 --gid=1000 --home-dir /home/container local1000 -o && \
    useradd --system --uid=1001 --gid=1000 --home-dir /home/container local1001 -o

ENV MAVEN_CONFIG /home/container/.m2

CMD ["mvn", "spring-boot:run"]
