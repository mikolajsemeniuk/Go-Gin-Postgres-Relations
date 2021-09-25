# Go-Gin-Postgres-Relations
commands:
```sh
docker-compose up -d
docker container ls
docker cp init.sql <container_id>:/home/init.sql
docker exec -it <container_id> /bin/bash
psql postgres://root:P%40ssw0rd@localhost
\l
CREATE DATABASE db;
\c db
\dt
\i /home.init.sql
\dt
select * from users;
\q
exit
```