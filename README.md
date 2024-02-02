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

```rest
GET http://127.0.0.1:8080/user/details HTTP/1.1
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkpvaG4iLCJleHAiOjE3MDY5MTU5NzZ9.U_brDPjtRkDDQQQV8UFqW2hvUXk8GJHZps2kfZhb6rc


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
