<h1>Run E2E test

## Prerequisites

- npm
- npx
- hoverfly / hoverctl
- playwright

## Procedure

### 1. Install Tools

### 2. Install hoverfly Cert in your PC

```bash
wget https://raw.githubusercontent.com/SpectoLabs/hoverfly/master/core/cert.pem -O /tmp/cert.pem
```

after add it.

### 3. Run App

```bash
make db-run

HTTP_PROXY=http://localhost:8500 \
HTTPS_PROXY=http://localhost:8500 \
NO_PROXY=localhost,127.0.0.1 \
IS_E2E_MODE=true make run
```

### 4. Run test

```bash
make test-e2e
```

if you want to open ui,

```bash
make test-e2e-ui
```
