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
mutation CreateBook($input: BookInput!){
	CreateBook(input:$input){
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
    "title": "test2",
    "author": "me2",
    "publisher": "go2"
  }
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