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
$ curl -H 'Content-Type:application/json' -X POST -d 'query{userList{userId, userName, email}}' 'http://localhost:3000/graphql'
# => {"data":{"userList":[{"email":"user01@gmail.com","userId":"2ed35dac-ce2c-4790-9444-17ace556a3ce","userName":"user01"},{"email":"user02@gmail.com","userId":"6bc63079-8678-41b5-aa5d-9b2a7341164f","userName":"user02"},{"email":"user03@gmail.com","userId":"50bf9aff-1346-46e5-8a19-773275dac567","userName":"user03"}]}}
```

```
# Create new user
$ curl -H 'Content-Type:application/json' -X POST -d 'mutation{createUser(userName:"user04", email: "user04@gmail.com"){userId, userName, email}}' 'http://localhost:3000/graphql'
# => {"data":{"createUser":{"email":"user04@gmail.com","userId":"4212d3a2-a434-4e29-a313-c1b5232691a1","userName":"user04"}}}
```

- Case 2
```
# Get user list
http://localhost:3000/graphql?query={userList{userId,userName,email}}
# => {"data":{"userList":[{"email":"user01@gmail.com","userId":"2ed35dac-ce2c-4790-9444-17ace556a3ce","userName":"user01"},{"email":"user02@gmail.com","userId":"6bc63079-8678-41b5-aa5d-9b2a7341164f","userName":"user02"},{"email":"user03@gmail.com","userId":"50bf9aff-1346-46e5-8a19-773275dac567","userName":"user03"}]}}
```

```
# Create new user
http://localhost:3000/graphql?query=mutation+_{createUser(userName:%22user04%22,email:%22user04@gmail.com%22){userId,userName,email}}
# => {"data":{"createUser":{"email":"user04@gmail.com","userId":"4212d3a2-a434-4e29-a313-c1b5232691a1","userName":"user04"}}}
```
