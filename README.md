## Gin + MongoDB

### Create project

```js
go mod init trandung/server
```

### Insatll plugin: gin and mongo-driver

```js
go get -u github.com/gin-gonic/gin
go get -u go.mongodb.org/mongo-driver/mongo
```

### Remove all packages not used

```js
go mod tidy
```
