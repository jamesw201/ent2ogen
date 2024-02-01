# ent2ogen

Quite often there's a need to expose some database entities through service api.\
If you are using [ent](https://github.com/ent/ent) as a database abstraction layer, and [ogen](https://github.com/ogen-go/ogen) for an API, then you might need to do type conversion between ent and ogen types, like this:

```go
func (s *Server) GetUser(ctx context.Context, params openapi.GetUserParams) (openapi.User, error) {
	u, err := s.db.Users.Get(params.UserID)
	if err != nil {
		return openapi.User{}, fmt.Errorf("query user: %w", err)
	}

	return openapi.User{
		ID:       u.ID,
		Username: u.Username,
		Age:      u.Age,
		// and so on...
	}
}
```

Writing such conversion by hand is annoying and error-prone (especially when the types are big with deep nesting).\
Also these conversions may become out of sync over time because of database or api schema updates.\
Ent2ogen solves this problem by generating mapping functions automatically.

## How to use

1. Create openapi schema
2. Create ent schema
3. Create ```entc.go``` file ([sample](example/ent/entc.go))
4. Create ```generate.go``` file ([sample](example/ent/generate.go))
5. Use following ent schema annotations:

* ```ent2ogen.BindTo("")``` - generate mapping function to specified openapi schema component.
* ```ent2ogen.Bind()``` - similar to BindTo but uses ent schema name by default.

6. Run ```go generate```

### Running
6. run `docker-compose up`
7. run `go run -mod=mod main.go`

### Testing
8. put http://localhost:8082/products into a Rest client (postman/insomnia) and POST the json body of the request from step 5.
![alt text](https://github.com/jamesw201/goent-otel-cockroach/blob/main/images/create-product.png?raw=true)
9. run `cockroach sql --execute "SELECT * FROM products" --insecure` to see your schema object created in CockroachDB
10. follow steps 8-9 for a Customer that has the newly created Product
![alt text](https://github.com/jamesw201/goent-otel-cockroach/blob/main/images/create-customer.png?raw=true)
11. try getting the endpoint: customers/{id}/products
![alt text](https://github.com/jamesw201/goent-otel-cockroach/blob/main/images/get-customer-products.png?raw=true)
12. open `http://localhost:16686/search` and search traces to find OpenTelemetry of your requests
We should see something like this: 
![alt text](https://github.com/jamesw201/goent-otel-cockroach/blob/main/images/opentelemetry.png?raw=true)

![alt text](https://github.com/jamesw201/goent-otel-cockroach/blob/main/images/otel-span.png?raw=true)
