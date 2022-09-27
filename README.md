# gql-go
Example mutation
```
mutation CreateToDo ($input: NewTodo!){
  createTodo(input:$input){
    id
    text
    done
    user{
      id
      name
    }
  }
}
```
```
{
  "input": {
    "text": "hello world",
    "userId": "01"
  }
}
```