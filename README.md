## Redeam Code Challenge

### Running the API

```
/*> docker-compose up
```

### Running the tests

```
/tests> go get github.com/stretchr/testify/assert
/tests> go test
```

### Notes

- It appears that the `mysql` container needs 20-30 seconds to initialize prior to running tests.
- `go get` dependencies are included in the docker container for the API, but the dependency for the tests must be installed manually.