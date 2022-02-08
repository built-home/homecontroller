


### CRDs


`.envrc`
```sh
export IPADDRESS="10.0.0.1"
```

`file.yaml`
```yaml
kind: Device
metadata:
  name: laptop
spec:
  name: PersonalDesktop
  type: Macbook
  os: darwin
  ipAddress: ${IPADDRESS}
```

```sh
envsubst < path/to/file.yaml | kubectl apply -f -
EOF<<<
kind: Device
metadata:
  name: laptop
spec:
  name: PersonalDesktop
  type: Macbook
  os: darwin
  ipAddress: 10.0.0.1
>>>
```
