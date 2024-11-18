# Dagger Go Utilities

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sagikazarmark/dagx/ci.yaml?style=flat-square)](https://github.com/sagikazarmark/dagx/actions/workflows/ci.yaml)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/sagikazarmark/dagx)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.20-61CFDD.svg?style=flat-square)
[![built with nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

**Simple utilities for working with the Dagger Go SDK.**

## Usage

### Pipeline

Use the `pipeline` package to run a bunch of jobs in parallel.
It's especially useful to wrap a number of objects implementing a `Sync` function (like containers, directories, files).

<details><summary>Before</summary><br><pre>
p := pool.New().WithErrors().WithContext(ctx)

p.Go(func(ctx context.Context) error {
    _, err := dag.Container.Sync(ctx)

    return err
})

p.Wait()
</pre></details>

After:

```go
p := pipeline.New(ctx)

pipeline.AddSyncStep(p, dag.Container())

pipeline.Run(p)
```

## Development

**For an optimal developer experience, it is recommended to install [Nix](https://nixos.org/download.html) and [direnv](https://direnv.net/docs/installation.html).**

## License

The project is licensed under the [MIT License](LICENSE).
