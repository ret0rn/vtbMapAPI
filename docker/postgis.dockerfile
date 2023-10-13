FROM postgres:16

RUN apt-get update
RUN apt-get install --no-install-recommends --yes \
    postgresql-16-postgis-3 postgis

#COPY ../init.sql /docker-entrypoint-initdb.d/