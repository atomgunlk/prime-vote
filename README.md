# Prime vote
## "Prime vote", The demo project
Included : 
- Unit test, Use Mockery for mock data of dependencies.
- Request validator
- Graceful shutdown

# Run
```
docker-compose up
```

# API
POST /vote

POST /voteitem

GET /voteitem

GET /voteitem/{{id}}

PUT /voteitem/{{id}}

PUT /voteitem/{{id}}/clear

DELETE /voteitem/{{id}}

GET /voteresult

GET /voteresult/export

# POSTMAN
[postman-collection](./test/Prime-vote.postman_collection.json)

# Run Test
```
godotenv -f ./test/.env go test ./cmd/prime-vote/handler  -run "TestUnitHandler" -cover
```

