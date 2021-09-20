# Authentication using sessions and middleware

```
$ go run main.go
```

## Access protected page without login
```
curl http://127.0.0.1:8080/home
```
Will get unauthorized and redirect back to "/"

## Simulate login
```
curl -s -I http://127.0.0.1:8080/login
```
Copy the session cookie

## Access protected page with session cookie
```
curl -s --cookie "session=MTYzMjA0ODA1NHxEd"  http://127.0.0.1:8080/home

Protected page accessible ...!!!


