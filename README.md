### Numbers service

![Numbers logo](/github/logo.jpg)

## Description

The service collects numbers from the various URLs passed in the GET request by url `/numbers` into `u` parameters.
 
These numbers are sorted by value and duplicates are deleted. If the service responds more than 500 ms, then its response will not be taken.

By default server start on port 8080

### Example

GET request:

You can pass multiple URL in `u` parameters
```
http://localhost:8080/numbers?u=http://localhost:8090/primes&u=http://localhost:8090/fibo
```
Response:
```
{
    "numbers": [
        1,
        2,
        3,
        5,
        7,
        8,
        11,
        13,
        21
    ]
}
```

### How to run?

```
go get -u github.com/pcherednichenko/numbers
```
Then from folder
```
go run main.go
```

Created by [Pavel Cherednichenko](https://www.linkedin.com/in/pavel-cherednichenko-0a2a0b118/)

#### [Test server](/testserver)