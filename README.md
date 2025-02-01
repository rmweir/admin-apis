# License API

## Update Types

Run
```
just gen
```

## Add, remove, edit features
1. Add the features to `pkg/features.yaml`.
2. Run
```
just gen
```

## Test Stripe feature upload CI locally
1. Create token in Stripe sandbox
2. run
```
STRIPE_API_TOKEN=<sandbox-token> just upload-ci-local
```
