# gql-go
- Run postgres in docker locally
```bash
docker run --name postgres-gql -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:14
```
- GQL migrate
```bash
go run github.com/99designs/gqlgen migrate
```
### Queries & Mutations example
Mutation 1
```
mutation CreateBook($input: BookInput!) {
  CreateBook(input: $input) {
    id
    title
    author
    publisher
  }
}
```
```
{
  "input": {
    "title": "test",
    "author": "me",
    "publisher": "go"
  }
}
```
Mutation 2
```
mutation UpdateBook {
  UpdateBook(
    id: 1
    input: {title: "test-new1", author: "me-new1", publisher: "go-new1"}
  ) {
    id
    title
    author
    publisher
  }
}
```
Mutation 3
```
mutation DeleteBook {
  DeleteBook(id: 1)
}
```
Query 1
```
query GetAllBooks {
  GetAllBooks {
    id
    title
    author
    publisher
  }
}
```
Query 2
```
query GetOneBook {
  GetOneBook(id: 1) {
    id
    title
    author
    publisher
  }
}
```