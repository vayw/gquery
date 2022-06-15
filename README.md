# Stupid query builder for GraphQL

## How-to

```
// create a query to resource "hero"
query := GQuery{QFrom: "hero"}

// add condition; note the quotes in the string - it's must if you testing against the string and not the type
query.Where("name", "_similar", "\"%% Skywalker\"")
// search in relation to resource "friends"
query.Where("friends.name", "_in", "[\"Han Solo\"]")

// specify property you need
query.Get("name")
query.Get("birthday")

// build the query string
query.Build()
```
