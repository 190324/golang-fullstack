## graphql schema
go run artisan.go gql api
go run artisan.go gql admin

## database sql
go run artisan.go migration [create_table_name] [db_name]

go run artisan.go migration create_products solomo

## database migrate
go run artisan.go migrate -e [db_name]

UP
go run artisan.go migrate -e solomo

DOWN
go run artisan.go migrate down -e solomo

## project layout
https://github.com/golang-standards/project-layout

email
動態參數


error if

golang no optional