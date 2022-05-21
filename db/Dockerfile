FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD admin
ENV MYSQL_USER dev
ENV MYSQL_PASSWORD dev
ENV MYSQL_DATABASE book-service

ADD init.sql /docker-entrypoint-initdb.d

EXPOSE 3306