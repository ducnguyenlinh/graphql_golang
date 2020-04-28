## graphql_golang
このリポジトリはgraphQLの練習のための簡単なgolang Webアプリケーションです。

## 設定
`go mod`を設定する

```
$ go mod init
$ go mod vendor
```

## 起動
```cmd
$ go run main
```

- Case 1:
```
# Get user list
$ curl -H 'Content-Type:application/json' -X POST -d 'query{userList{id, name, email}}' 'http://localhost:3000/graphql'
# => {"data":{"userList":[{"email":"user01@gmail.com","id":"2ed35dac-ce2c-4790-9444-17ace556a3ce","name":"user01"},{"email":"user02@gmail.com","id":"6bc63079-8678-41b5-aa5d-9b2a7341164f","name":"user02"},{"email":"user03@gmail.com","id":"50bf9aff-1346-46e5-8a19-773275dac567","name":"user03"}]}}
```

```
# Create new user
$ curl -H 'Content-Type:application/json' -X POST -d 'mutation{createUser(name:"user04", email: "user04@gmail.com"){id, name, email}}' 'http://localhost:3000/graphql'
# => {"data":{"createUser":{"email":"user04@gmail.com","id":"4212d3a2-a434-4e29-a313-c1b5232691a1","name":"user04"}}}
```

```
# Find user by id
$ curl -H 'Content-Type:application/json' -X POST -d '{user(id:"849b9ad1-847d-48b6-98e9-169b51fbb4c3"){id, name, email}}' 'http://localhost:3000/graphql'
# => {"data":{"user":{"email":"user02@gmail.com","id":"849b9ad1-847d-48b6-98e9-169b51fbb4c3","name":"user02"}}}
```

- Case 2
```
# Get user list
http://localhost:3000/graphql?query={userList{id,name,email}}
# => {"data":{"userList":[{"email":"user01@gmail.com","id":"2ed35dac-ce2c-4790-9444-17ace556a3ce","name":"user01"},{"email":"user02@gmail.com","id":"6bc63079-8678-41b5-aa5d-9b2a7341164f","name":"user02"},{"email":"user03@gmail.com","id":"50bf9aff-1346-46e5-8a19-773275dac567","name":"user03"}]}}
```

```
# Create new user
http://localhost:3000/graphql?query=mutation+_{createUser(name:%22user04%22,email:%22user04@gmail.com%22){id,name,email}}
# => {"data":{"createUser":{"email":"user04@gmail.com","id":"4212d3a2-a434-4e29-a313-c1b5232691a1","name":"user04"}}}
```

```
# Find user by id
http://localhost:3000/graphql?query={user(id:"759c4add-5984-4ace-8824-5a16b5add106"){id,name,email}}
# => {"data":{"user":{"email":"user02@gmail.com","id":"759c4add-5984-4ace-8824-5a16b5add106","name":"user02"}}}
```
