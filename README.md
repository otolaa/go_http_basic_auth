## HTTP Basic Auth in Go(Golang)

```
curl -u username:password http://localhost:8080/
```

If your username or password contains a special character, such as white-space, then you might want to surround credentials with single quotes:

```
curl -u 'username:password' http://localhost:8080/
```

If you want to have a full control over your HTTP request, you might want to Base64 encode your username:password and place it into Authorization header.
Curl command should look like this:

```
curl -H 'Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=' http://localhost:8080/
```
and this https://devpal.co/base64-encode/ => username:password the is dXNlcm5hbWU6cGFzc3dvcmQ=


```
curl -H "Authorization: Basic $(base64 <<<"joeuser:secretpass")" http://localhost:8080/
```

```
curl -i 'Accept:application/json' -u username:password http://localhost:8080/
```


## Kill Port

```
killall -9 go
```