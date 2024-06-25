# OpenFeature Demo

This demo app shows how to use [OpenFeature](https://openfeature.dev/) with GoLang
to standardize feature flag evaluation.

The demo app uses the vendors [Statsig](https://statsig.com/) and [Flagsmith](https://www.flagsmith.com/).

## Running the application

```bash
$ STATSIG_KEY=secret-123 FLAGSMITH_KEY=ser.123 \
go run .
```