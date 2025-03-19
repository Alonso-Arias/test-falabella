# test-cleverit

## Requisitos

* Go 1.20
* MySQL 5.7.x ( docker pull mysql:5.7.33 )
## Ambiente Local ( BD basado en Docker )

* BD: `docker run --name test-db -e MYSQL_ROOT_PASSWORD=123456 -d -p 3306:3306 mysql:5.7.33`

## Creación Esquema y Tablas - Carga datos iniciales ( Basado en Docker)

* Copiar scripts dentro del contenedor : `docker cp ./db/scripts/ test-db:/tmp/`
* Eliminación y creación de esquema y tablas : `docker exec -t test-db /bin/sh -c 'mysql -u root -p123456 </tmp/scripts/create-db.sql'`
* Carga de datos iniciales : `docker exec -t test-db /bin/sh -c 'mysql -u root -p TEST -p123456 </tmp/scripts/load-data.sql'`


## Compilación y Ejecución

* `make run`


## Ejecución de Tests

* `make test`

## Docker

Comandos para generación de contenedor de API. No es necesario para ambiente local.

* `docker build -t exam-test:1.0 .`
* `docker run -p 1323:1323 --name exam-test exam-test:1.0`

O para levantar todo

* `docker-compose up`
