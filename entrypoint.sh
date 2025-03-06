#!/bin/sh

# Aguarda o PostgreSQL estar pronto
echo "Waiting for PostgreSQL..."
while ! nc -z $DB_HOST $DB_PORT; do
  sleep 0.1
done
echo "PostgreSQL started"

# Executa as migrações
echo "Running migrations..."
migrate -path /app/migrations -database "postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up

# Verifica se as migrações foram bem-sucedidas
if [ $? -ne 0 ]; then
    echo "Migration failed!"
    exit 1
fi
echo "Migrations completed successfully!"

# Inicia a aplicação
echo "Starting application..."
exec ./main