#!/bin/bash

if [ -f .env ]; then
  source .env
else
  echo "error: .env file not found"
  exit 1
fi


if [ -z "$POSTGRES_USER"  ] || [ -z "$POSTGRES_PASSWORD" ] || [ -z "$POSTGRES_DB" ]; then
  echo "error: missing required env vars in .env"
  exit 1
fi


CONTAINER_NAME="backend-db-1"

case "$1" in 
  start)
  docker run -d \
    --name "$CONTAINER_NAME" \
    -e POSTGRES_USER="$POSTGRES_USER" \
    -e POSTGRES_PASSWORD="$POSTGRES_PASSWORD" \
    -e POSTGRES_DB="$POSTGRES_DB" \
    -p 5431:5432 \
    -v "$(pwd)/pgdata:/var/lib/postgresql/data" \
    postgres:15

if [ $? -eq 0 ]; then
  echo "DB started, connect with:"
  echo "psql -h localhost -p 5431 -U $POSTGRES_USER -d $POSTGRES_DB"
  echo "$POSTGRES_PASSWORD"
fi
;;

stop)
  docker stop "$CONTAINER_NAME"
  docker rm "$CONTAINER_NAME"
  echo "DB stopped."
  ;;

logs)
  docker logs -f $CONTAINER_NAME
  ;;

delete)
  docker stop "$CONTAINER_NAME"
  docker rm "$CONTAINER_NAME"
  sudo rm -rf pgdata
  echo "pgdata deleted"
  ;;

connect)
  psql -h localhost -p 5431 -U $POSTGRES_USER -d $POSTGRES_DB
  ;;

migrate)
  cd sql/schema
  goose postgres postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5431/$POSTGRES_DB up
  cd ..
  cd ..
  ;;

revert)
  cd sql/schema
  goose postgres postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5431/$POSTGRES_DB down
  cd ..
  cd ..
  ;;

reset-groups)
  curl -X POST http://localhost:8080/admin/reset-groups
  ;;


*)
  echo "usage: ./db.sh start|stop|delete|logs|connect|migrate|revert|reset-groups"
  exit 1
  ;;
esac
