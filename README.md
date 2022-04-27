# Learning Go With Tests

This is code written as part of [Learning Go With Tests](https://quii.gitbook.io/learn-go-with-tests).

## Running said tests

You can run all the tests using the handy make target:

```shell
make test
```

The CI will also run all tests on every push.

### Other make targets

There are a couple of additional make targets:

- help => Print some help info
- test_coverage => Will run all tests with coverage
- doc => Will run the godoc on [localhost](http://localhost:6060)

## Principles

To keep things simple, there will be a file ( and associated test ) for each chapter.
