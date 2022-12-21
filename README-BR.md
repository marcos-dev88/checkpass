# checkpass

<br>

## Sumário
- [Sobre o projeto](#sobre-o-projeto)
- [Necessário](#necessario)
- [Instruções](#instrucoes)
- [Iniciar o projeto](#iniciar-o-projeto)
- [Requisições](#requisicoes)

<br>

## Sobre o projeto
Esse projeto tem como objetivo checar e verificar cada regra definida pelo usuário.

-----

## Necessário
* GNU Make >= v4.3
* Docker >= v20.10.17
* Docker Compose >= v2.6.1
* Golang >= v1.18

-----

## Instructions

1. Renomeie seu **.env.example** para **.env**;
2. Defina suas configurações padrões em seu arquivo **.env**;
3. Inicie o projeto; (Siga as instruções de como inicar em: [Iniciar o projeto](#iniciar-o-projeto) )
4. Verifique suas senhas; (Siga as instruções de como verificar suas senhas: [Requisições](#requisicoes) )

-----

## Iniciar o projeto

### Iniciando com docker (Linux)
```shell
$ make run-docker
```

### Iniciando com docker (Linux/Windows)
```shell
$ docker-compose up
```

### Iniciando com arquivo binário (Linux)
```shell
$ make run
```

### Iniciando com arquivo binário (Windows)
```shell
$ ./bin/checkpass.exe
```

-----

## Requisições

Faça uma requisição **POST** com esse corpo:

```
{
    "password": "My4maz1ngStr0ngPa$$W0Rd", # as senhas que desejas verificar    
    "rules": [                             # lista de regras que desejas verificar em sua senha
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
            "value": 0                   # 0 é igual à "false" e 1 é igual a "true"
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

### Exemplos de requisições:

*sucesso:*
```shell
# sucesso

$ curl -XPOST 'http://127.0.0.1:8000/verify' -H 'Content-Type: application/json' -d '{ "password": "YADFgsba$*(50)88some2n3a$hfqp*(mxnv!32!@2SDFGsOaSFVjh", "rules": [ { "rule": "minSize", "value": 8 }, { "rule": "minSpecialChars", "value": 8 }, { "rule": "noRepeted", "value": 0 }, { "rule": "minDigit", "value": 8 }, { "rule": "minUppercase", "value": 8 }, { "rule": "minLowercase", "value": 8 } ] }'

```

*Sem repetição:*
```shell
# erro de verificação da regra de sem repetição

$ curl -XPOST 'http://127.0.0.1:8000/verify' -H 'Content-Type: application/json' -d '{ "password": "YADFgsba$*(50)88some2n3a$hfqp*(mxnv!32!@2SDFGsOaSFVjh", "rules": [ { "rule": "minSize", "value": 8 }, { "rule": "minSpecialChars", "value": 8 }, { "rule": "noRepeted", "value": 1 }, { "rule": "minDigit", "value": 8 }, { "rule": "minUppercase", "value": 8 }, { "rule": "minLowercase", "value": 8 } ] }'
```
