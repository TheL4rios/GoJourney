FROM postgres:16.2

COPY ./init-db.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]