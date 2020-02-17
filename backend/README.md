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

## migration

```
建立 Database
soda create -a

移除 Database
soda drop -a
```

```
創建模型
soda generate model user name:text email:text
```

```
執行
soda migrate up

返回
soda migrate down
```