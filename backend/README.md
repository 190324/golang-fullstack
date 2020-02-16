## 更新 gqlgen 腳本

```
API 
go run scripts/gqlgen.go -c gqlgen/api/gqlgen.yml

ADMIN 
go run scripts/gqlgen.go -c gqlgen/admin/gqlgen.yml
```

## 啟動 API Server

```
go run api.go
```