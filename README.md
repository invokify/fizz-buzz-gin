# fizz-buzz-gin

## The game rules

The original fizz-buzz consists in writing all numbers from `1` to `100`, and replacing:

- all multiples of `3` by `fizz`,
- all multiples of `5` by `buzz`,
- and all multiples of `15` by `fizzbuzz`.

The output would look like this:

```
1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,fizz,...
```

> More specifically :
> `Expose a REST API endpoint` that accepts five parameters : two strings (say, string1 and string2), and three integers (say, int1, int2 and limit), and returns a JSON
> It must return a list of strings with numbers from 1 to limit, where:
>
> - all multiples of int1 are replaced by string1,
> - all multiples of int2 are replaced by string2,
> - all multiples of int1 and int2 are replaced by string1string2

## Solution

I have used `gin-gonic` framework to handle the web server part, and build/test my endpoints. This part is located in the `pkg/server` folder.

The core of the game rules are located in the `pkg/business` folder.

## Getting started

For both following cases, you can test the endpoint:

- either in the [browser](http://localhost/api/v1/fizz-buzz?string1=Fizz&string2=Buzz&int1=3&int2=5&limit=10)
- either in [swagger](http://localhost/swagger/index.html)

### Launch locally

> The project has been set up on Windows, following commands might change on other OS

Launch all tests and check coverage:

```
go test -race -covermode atomic -coverprofile .\coverage.txt ./...
```

Build:

```
go build -o=fizzBuzzApp .\cmd\api\main.go
```

Run:

```
go run .\cmd\api\main.go
```

### Using Docker Compose

#### **_Requirements_**

_Install Docker Engine_

The image is build in the Gitlab CI according to the `.gitlab-ci.yml` file and the `Dockerfile`. It it then pushed to the registry. The `docker-compose.yaml` file pulls it from the registry and runs a container within Docker Engine.

```
docker-compose up -d
```
