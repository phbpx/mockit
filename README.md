# mockit 

`mockit` is a dead simple yet powerful HTTP API mock CLI.

# Features

- Simple CLI.
- Simple configuration file (and you can split it into multiple files).
- Support response body template values.
- Buitin functions to generate random template values such as: `uuid`, `now`, `name` and `email`

## Install

### Pre-compiled executables

Get them [here](https://github.com/phbpx/mockit/releases)

### Source

```sh
git@github.com:phbpx/mockit.git
cd mockit
make mockit
sudo mv mockit ~/usr/local/bin # Or elsewhere, up to you.
```

## Run `mockit` using CLI

CLI usage:
```bash 
$ mockit -h
Usage of mockit:
  -config value
        config file path
  -port string
        http port to listen on (default "8080")
```

1. Define your mocks in the configuration file:
```yml
endpoints:
  - method: "GET"
    url: "/accounts/:id"
    response:
      code: 200
      headers:
        Content-Type: "application/json"
      body: |
        {
          "id": "{{ urlParam "id" }}",
          "name": "{{ name }}",
          "email": "{{ email }}",
          "username": "{{ username }}",
          "fixed": "abcd",
          "createdAt": "{{ now.Format "2006-01-02T15:04:05Z" }}"
        }
  - method: "POST"
    url: "/accounts"
    response:
      code: 201
      headers:
        Content-Type: "application/json"
      body: |
        {
          "id": "{{ uuid }}",
          "name": "{{ name }}",
          "email": "{{ email }}",
          "username": "{{ username }}",
          "fixed": "abcd",
          "createdAt": "{{ now.Format "2006-01-02T15:04:05Z" }}"
        }
```

2. Run CLI:
```bash
mockit -config conf.yml -port 8080
```

## Run `mockit` using docker/postman

1. Define your mocks in the configuration file:
```yml
endpoints:
  - method: "GET"
    url: "/accounts/:id"
    response:
      code: 200
      headers:
        Content-Type: "application/json"
      body: |
        {
          "id": "{{ urlParam "id" }}",
          "name": "{{ name }}",
          "email": "{{ email }}",
          "username": "{{ username }}",
          "fixed": "abcd",
          "createdAt": "{{ now.Format "2006-01-02T15:04:05Z" }}"
        }
```

2. Run docker:

```bash
docker run \
  -v "$(pwd):/src" \
  -e MOCKIT_CONFIG="/src/conf.yml" \
  -e MOCKIT_PORT="8080" \
  -p 8080:8080 \
  quay.io/phbpx/mockit:latest
```

## Template functions

- urlParam(paramName string): 
  - Ex.: For URL `/accounts/:id` use `{{ urlParam "id" }}` to print the `:id` parameter. 
- uuid: `{{ uuid }}`
- now: `{{ now }}`
- username: `{{ username }}`
- name: `{{ name }}`
- email: `{{ email }}`
- phone: `{{ phone }}`
- int: `{{ int }}`
- digit: `{{ digit }}`
- digitN(n int): `{{ digitN 4 }}`
- letter: `{{ letter }}`
- letterN(n int): `{{ letterN 10 }}`
- word: `{{ word }}`
- phrase: `{{ phrase }}`
- loremIpsum(n int): `{{ loremIpsum 10 }}`