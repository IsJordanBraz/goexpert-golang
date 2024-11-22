- install golang-migrate

- smigrate create -ext=sql -dir=sql/migrations -seq init
- create databases in 000001_init.up.sql
- migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

- docker-compose exec mysql bash
- mysql -uroot -p courses

### mysql> show tables;
+-------------------+
| Tables_in_courses |
+-------------------+
| categories        |
| courses           |
| schema_migrations |
+-------------------+
3 rows in set (0.00 sec)

### mysql> desc courses;
+-------------+---------------+------+-----+---------+-------+
| Field       | Type          | Null | Key | Default | Extra |
+-------------+---------------+------+-----+---------+-------+
| id          | varchar(36)   | NO   | PRI | NULL    |       |
| category_id | varchar(36)   | NO   | MUL | NULL    |       |
| name        | text          | NO   |     | NULL    |       |
| description | text          | YES  |     | NULL    |       |
| price       | decimal(10,2) | NO   |     | NULL    |       |
+-------------+---------------+------+-----+---------+-------+
5 rows in set (0.00 sec)

- migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

- create queries

- sqlc generate