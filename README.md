![Logo of the project](https://github.com/gopher-lang/gopher/blob/master/doc/gopher/gophercolor.png)

# Code Challenge: HTTP key/value store (take-home exercise)
## Requirements
HTTP endpoints:

PUT /<key>  — Set the value of a key

GET /<key>  — Fetch the value of a key

DELETE /<key> — Delete a key

The server's database state is persisted to the file system (without relying in an external process, or using an existing key-value store library).

Restarting the server should not lose writes that have already been acknowledged with a 2XX response status code.

## Features

The server is listening on Port 5000.

to test the server, you can try out some of the example curl commands

Example:
1. Write a value to the key ```foo```:
```bash
$ curl -i -X PUT 'http://localhost:5000/foo' -H 'Content-Type: application/octet-stream' --data-binary 'hello world!'
HTTP/1.1 204 No Content
```

2. Fetch the value of ```foo```:
```bash
$ curl -i 'http://localhost:5000/foo'
HTTP/1.1 200 OK
Content-Type: application/octet-stream
Content-Length: 12

hello world!
```

3. Fetch the value of ```qux```(which does not exist):
```bash
$ curl -i -X PUT 'http://localhost:5000/foo' -H 'Content-Type: application/octet-stream' --data-binary 'hello world!'
HTTP/1.1 204 No Content
```

4. Delete the key ```foo```:
```bash
curl -i -X DELETE 'http://localhost:3000/foo'
HTTP/1.1 204 No Content
```
## Licensing

"The code in this project is licensed under MIT license."
