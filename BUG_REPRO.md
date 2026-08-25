# BUG_REPRO

The following failures were observed while validating the initial project state.
Each section records what failed, how to reproduce it, and the complete command output.
They are preserved intentionally; only failing build gates are omitted from the generated Dockerfile.

## Failure 1: Go test (.)

- Observed problem: `Go test (.)` failed in the initial project state.
- Working directory: `.`
- Command: `cd /app && GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -count=1 ./...`
- Exit status: `1`

```text
?   	training-review/cmd/server	[no test files]
?   	training-review/internal/api	[no test files]
?   	training-review/internal/config	[no test files]
?   	training-review/internal/domain	[no test files]
?   	training-review/internal/importer	[no test files]
?   	training-review/internal/report	[no test files]
?   	training-review/internal/service	[no test files]
?   	training-review/internal/store	[no test files]
ok  	training-review/internal/flow001	0.025s
--- FAIL: Test670BusinessRegression (0.00s)
    regression_test.go:13: second number is A-001
FAIL
FAIL	training-review/internal/flow002	0.001s
ok  	training-review/internal/flow003	0.019s
FAIL
```

## Architecture reproduction

### linux/amd64
- Go toolchain version: exit `0`
- Node.js version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/server): exit `0`
- Frontend build (web): exit `0`
### linux/arm64
- Go toolchain version: exit `0`
- Node.js version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/server): exit `0`
- Frontend build (web): exit `0`
