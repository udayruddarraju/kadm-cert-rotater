# kadm-cert-rotater

`kadm-cert-rotate` is a wrapper on top of kubeadm, that:

- Renews certificates in /etc/kubernetes/pki
- Updates kubeconfigs on the control plane nodes
- Restarts the control plane components, kube-apiserver, kube-controller-manager, kube-scheduler and etc.

> kadm-cert-rotate needs to be deployed with hostIPC set to true for it to restart control plane components. This setting will help the binary discover control plane process IDs to issue a SIGHUP. If this is not set, it would only rotate/renew the certificates but not restart the components, which would mean that the new certificates are not yet in effect.

```
kadm-cert-rotate - rotates/renews control plane certificates if the expiration is lesser than the threshold.

Usage:
  kadm-cert-rotate [flags]

Flags:
  -h, --help                       help for kadm-cert-rotate
      --pki-dir string             path to the pki directory that is used during kubeadm init.
      --renewal-threshold string   duration with which the rotater is expected to renew expiring certificates.
```

## Development Guide

Clone the repository

```bash
cd ~/go/src
git clone https://github.com/udayruddarraju/kadm-cert-rotater
cd kadm-cert-rotater
```

Install locally
```bash
GOOS=linux GOARCH=amd64 go install cmd/kadm-cert-rotate
# on darwin run
cp ~/go/bin/linux_amd64/kadm-cert-rotate ./
# on linux run
cp ~/go/bin/kadm-cert-rotate ./
docker build . -t kadm-cert-rotate:latest 
```

To load the image in kind, run:
```bash
kind load docker-image kadm-cert-rotate:v4 --name=kadm-cert-rotate
```
