# crossflux

Crossflux is an example of managing an entire organization from a single
repository using GitOps on [Crossplane](https://crossplane.io/),
[Flux](https://fluxcd.io/), and [Kubernetes](https://kubernetes.io/).

## Repository Layout

- `clusters`: sources and reconciliations for Flux
    - Controls how Flux will reconcile changes in this repo
- `packages`: source for Crossplane configuration packages
- `platform`: manifests for building platform on Crossplane
    - Controls what platform APIs are offered to developers in the cluster
    - Installs configurations from `packages`
- `teams`: team-specific manifests for applications and infrastructure
    - See the [teams documentation](./docs/teams.md) for more information

## Setup

The following steps assume a running Kubernetes cluster and a `kubeconfig`
context that points to it.

1. Install Flux CLI

```
curl -s https://toolkit.fluxcd.io/install.sh | sudo bash
```

2. Install Latest Flux Version

```
flux install --version=latest --namespace=flux-system
```

3. Create Github Deploy Key Secret

```
flux create secret git git-deploy-key \
    --url=ssh://git@github.com/hasheddan/crossflux \
    --namespace=flux-system \
    --export > git-deploy-key.yaml
```

Add public key as a deploy key in your repo by following the steps in [the
GitHub
documentation](https://docs.github.com/en/developers/overview/managing-deploy-keys#deploy-keys),
then create the `Secret` in the cluster.

```
kubectl apply -f credentials/git-deploy-key.yaml
```

4. Create Provider Credential Secrets

```
kubectl create secret generic aws-creds -n crossplane-system --from-file=creds=./credentials/aws-creds.conf
```

5. Bootstrap Crossflux Cluster

```
kubectl apply -f clusters/controlplane
```
