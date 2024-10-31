FROM docker.io/busybox:latest as downloader

RUN mkdir /tmp/world-1.0

WORKDIR /tmp/world-1.0

RUN wget https://ftp.postgresql.org/pub/projects/pgFoundry/dbsamples/world/world-1.0/world-1.0.tar.gz && tar xzvf world-1.0.tar.gz

FROM docker.io/library/postgres:16 

# dbsamples-0.1
COPY --from=downloader /tmp/world-1.0/dbsamples-0.1/world/world.sql /docker-entrypoint-initdb.d/01_world.sql

