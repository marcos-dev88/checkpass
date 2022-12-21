# checkpass

<br>

## Summary
- [About the project](#about-the-project)
- [Requirements](#requirements)
- [Instructions](#instructions)
- [Run the project](#run-the-project)
- [Request to API](#request-to-api)

<br>

## About the project
This project has the responsibility of checking and verifying each rule that you define.

-----

## Requirements
* GNU Make >= v4.3
* Docker >= v20.10.17
* Docker Compose >= v2.6.1
* Golang >= v1.18

-----

## Instructions

1. Rename your **.env.example** to **.env**;
2. Set your defaults in your **.env** file;
3. Run the project; (Follow the guide on to run it: [Run the project](#run-the-project) )
4. Check your passwords; (Follow the guide on to check your passwords: [Request to api](#request-to-api) )

-----

## Run the project:

### Running with docker (Linux)
```shell
$ make run-docker
```

### Running with docker (Linux/Windows)
```shell
$ docker-compose up
```

### Running with binary file (Linux)
```shell
$ make run
```

### Running with binary file (Windows)
```shell
$ ./bin/checkpass.exe
```

-----

## Request to API

Make a **POST** Request with this body:

```
{
    "password": "My4maz1ngStr0ngPa$$W0Rd", # the password that you want to check
    "rules": [                             # list of rules that you want to verify in your password
        {
            "rule": "minSize",
            "value": 8
        },
        {
            "rule": "minSpecialChars",
            "value": 8
        },
        {
            "rule": "noRepeted",
            "value": 0                   # 0 is equivalent to false and 1 to true
        },
        {
            "rule": "minDigit",
            "value": 8
        },
        {
            "rule": "minUppercase",
            "value": 8
        },
        {
            "rule": "minLowercase",
            "value": 8
        }
    ]
}
```

### Examples of requests:

*success:*
```shell
# success

$ curl -XPOST 'http://127.0.0.1:8000/verify' -H 'Content-Type: application/json' -d '{ "password": "YADFgsba$*(50)88some2n3a$hfqp*(mxnv!32!@2SDFGsOaSFVjh", "rules": [ { "rule": "minSize", "value": 8 }, { "rule": "minSpecialChars", "value": 8 }, { "rule": "noRepeted", "value": 0 }, { "rule": "minDigit", "value": 8 }, { "rule": "minUppercase", "value": 8 }, { "rule": "minLowercase", "value": 8 } ] }'

```

*noRepeat issue:*
```shell
# error of verify in noRepeted

$ curl -XPOST 'http://127.0.0.1:8000/verify' -H 'Content-Type: application/json' -d '{ "password": "YADFgsba$*(50)88some2n3a$hfqp*(mxnv!32!@2SDFGsOaSFVjh", "rules": [ { "rule": "minSize", "value": 8 }, { "rule": "minSpecialChars", "value": 8 }, { "rule": "noRepeted", "value": 1 }, { "rule": "minDigit", "value": 8 }, { "rule": "minUppercase", "value": 8 }, { "rule": "minLowercase", "value": 8 } ] }'
```
