# jenkins_golang_test

## run unit test
```sh
go test -v github.com/jenkins_golang_test/operator
```

## generate code coverage
```sh
go test -v -coverprofile=cover.out github.com/jenkins_golang_test/operator
```

## review the code coverage
```sh
# Open a web browser displaying annotated source code:
go tool cover -html=cover.out

# Write out an HTML file instead of launching a web browser:
# go tool cover -html=cover.out -o coverage.html

# Display coverage percentages to stdout for each function:
# go tool cover -func=cover.out

```