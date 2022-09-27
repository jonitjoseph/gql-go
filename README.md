# gql-go
- Run postgres in docker locally
```bash
docker run --name postgres-gql -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:14
```