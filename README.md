# Prime vote
## "Prime vote", The demo project
Included : 
- Unit test, Use Mockery for mock data of dependencies.
- Request validator
- Graceful shutdown

# Build & Run
```
cd deployments
./dockerbuild.sh
docker-compose up
```

## Test users
```
Username : test0
Password : test0

Username : test1
Password : test1
```

# API

<details>
<summary><code>POST</code> <code><b>/login</b></code> (User login)</summary>

#### Request [`application/json`]
> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `username`        |  required | string         | Username   |
> | `password`        |  required | string         | Password   |
#### Response
> | name              | data type                         | response                                                            |
> |-------------------|-----------------------------------|---------------------------------------------------------------------|
> | `responseStatus`  | Object                            | `{"code":"00000","message":"Success"}`                              |
> | `token`           | string                            |                               |
</details>

<details>
<summary><code>POST</code> <code><b>/vote</b></code> (Give a vote)</summary>

#### Request [`application/json`]
> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `id`              |  required | int            | Vote item ID   |
#### Response
> | name              | data type                         | response                                                            |
> |-------------------|-----------------------------------|---------------------------------------------------------------------|
> | `responseStatus`  | Object                            | `{"code":"00000","message":"Success"}`                              |
</details>

<details>
<summary><code>POST</code> <code><b>/voteitem</b></code> (Create a vote item)</summary>

#### Request [`application/json`]
> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `name`            |  required | string         | Vote item name   |
> | `description`     |  required | string         | Vote item description   |
#### Response
> | name              | data type                         | response                                                            |
> |-------------------|-----------------------------------|---------------------------------------------------------------------|
> | `responseStatus`  | Object                            | `{"code":"00000","message":"Success"}`                              |
</details>

<details>
<summary><code>GET</code> <code><b>/voteitem</b></code> (List vote items)</summary>

#### Response
> | name              | data type                         | response                                                            |
> |-------------------|-----------------------------------|---------------------------------------------------------------------|
> | `responseStatus`  | Object                            | `{"code":"00000","message":"Success"}`                              |
> | `items`           | Array of Item                     | `[{"id": 1, "name": "item1", "description": "", "vote_count": 4 }, {"id": 2, ...}]`                              |
</details>

<details>
<summary><code>PUT</code> <code><b>/voteitem/{{id}}</b></code> (Update vote item)</summary>

#### Request [`application/json`]
> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `name`            |  required | string         | Vote item name   |
> | `description`     |  required | string         | Vote item description   |
#### Response
> | name              | data type                         | response                                                            |
> |-------------------|-----------------------------------|---------------------------------------------------------------------|
> | `responseStatus`  | Object                            | `{"code":"00000","message":"Success"}`                              |
</details>

<details>
<summary><code>PUT</code> <code><b>/voteitem/{{id}}/clear</b></code> (Clear vote item count)</summary>

#### Response
> | name              | data type                         | response                                                            |
> |-------------------|-----------------------------------|---------------------------------------------------------------------|
> | `responseStatus`  | Object                            | `{"code":"00000","message":"Success"}`                              |
</details>

<details>
<summary><code>DELETE</code> <code><b>/voteitem/{{id}}</b></code> (Delete vote item)</summary>

#### Response
> | name              | data type                         | response                                                            |
> |-------------------|-----------------------------------|---------------------------------------------------------------------|
> | `responseStatus`  | Object                            | `{"code":"00000","message":"Success"}`                              |
</details>


<details>
<summary><code>GET</code> <code><b>/voteresult</b></code> (Get vote result)</summary>

#### Response
> | name              | data type                         | response                                                            |
> |-------------------|-----------------------------------|---------------------------------------------------------------------|
> | `responseStatus`  | Object                            | `{"code":"00000","message":"Success"}`                              |
</details>

<details>
<summary><code>GET</code> <code><b>/voteresult/export</b></code> (Export vote result as XLSX file)</summary>
</details>

# POSTMAN Collection
[postman-collection](./test/Prime-vote.postman_collection.json)

# Run Test
```
godotenv -f ./test/.env go test ./cmd/prime-vote/handler  -run "TestUnitHandler" -cover
```

