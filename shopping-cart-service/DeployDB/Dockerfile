# Dockerfile-mysql
FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=shopping_cart
ENV MYSQL_USER=user123
ENV MYSQL_PASSWORD=pass123

COPY database.sql /docker-entrypoint-initdb.d/database.sql