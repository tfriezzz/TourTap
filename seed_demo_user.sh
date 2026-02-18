
#!/bin/sh

echo "waiting for backend"

until curl -s http://backend:8080/api/health > /dev/null; do
  sleep 1
done

echo "seeding demo user..."

curl -s -X POST http://backend:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "password": "password",
    "email": "test@email.com",
    "name": "Test McTestson"
  }'

echo "done"



