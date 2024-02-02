# services
Services


## Sign in

```rest
POST http://127.0.0.1:8080/user HTTP/1.1
content-type: application/json

{
    "username": "John",
    "password": "secret"
}
```

## Login

```rest
POST http://127.0.0.1:8080/access HTTP/1.1
content-type: application/json

{
    "username": "John",
    "password": "secret"
}
```
