FROM postgres:15.1

ENV PG_CRON_VERSION "1.4.2"

RUN apt-get update && apt-get install -y --no-install-recommends \
    postgresql-server-dev-15 postgresql-contrib-15 \
    libcurl4-openssl-dev \
    wget jq cmake build-essential ca-certificates && \
    mkdir /build && \
    cd /build && \
    wget https://github.com/citusdata/pg_cron/archive/v$PG_CRON_VERSION.tar.gz && \
    tar xzvf v$PG_CRON_VERSION.tar.gz && \
    cd pg_cron-$PG_CRON_VERSION && \
    make && \
    make install && \
    cd / && \
    rm -rf /build && \
    apt-get remove -y wget jq cmake build-essential ca-certificates && \
    apt-get autoremove --purge -y && \
    apt-get clean && \
    apt-get purge && \
    rm -rf /var/lib/apt/lists/*

RUN echo "shared_preload_libraries = 'pg_cron'" >> /var/lib/postgresql/data/postgresql.conf
RUN echo "cron.database_name = '${PG_CRON_DB:-pg_cron}'" >> /var/lib/postgresql/data/postgresql.conf

COPY ./docker-entrypoint.sh /usr/local/bin/

RUN chmod a+x /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 5432
CMD ["postgres"]

FROM golang:latest

RUN go version

COPY . /avito-tech-2023/
WORKDIR /avito-tech-2023/

# build go app
RUN go mod download
RUN GOOS=linux go build -o app ./cmd/main.go

RUN chmod +x docker-entrypoint.sh

CMD ["./app"]