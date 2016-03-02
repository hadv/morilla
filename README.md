# Simple JSON REST API using [Gorilla Mux][1]

[![Build Status](https://travis-ci.org/hadv/morilla.svg?branch=master)](https://travis-ci.org/hadv/morilla)

## Installation
```
$ go get github.com/hadv/morilla
```
## Dependencies Installation
```
$ go get github.com/gorilla/mux
$ go get github.com/onsi/ginkgo
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
```

## Compile & Running the Server
```
$ go install github.com/hadv/morilla
$ morilla
```

Then the server will start at this local path: http://localhost:8080/

## Checking the REST API by using curl
```
$ curl -X GET http://localhost:8080/api/v1/movies -v
```

## Running Test (TDD/BDD with ginkgo/gomega/gory)
```
$ cd src/github.com/hadv/morilla/api
$ ginkgo -r --randomizeAllSpecs -cover
Running Suite: Morilla Suite
============================
Random Seed: 1456914522 - Will randomize all specs
Will run 3 of 3 specs

•••
Ran 3 of 3 Specs in 0.004 seconds
SUCCESS! -- 3 Passed | 0 Failed | 0 Pending | 0 Skipped PASS
coverage: 20.9% of statements

Ginkgo ran 1 suite in 2.116968979s
Test Suite Passed
```

[1]: http://www.gorillatoolkit.org/pkg/mux
