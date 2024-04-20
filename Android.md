![](img/kuberneteslogo.png )

* [Kubernetes commands](#kubernetes-commands)  

# Kubernetes commands
Commands are based on OKD (_The Community Distribution of Kubernetes_)

*Deleted pod in project(also called namespace)*
```bash
oc delete pod kafka-brokers-0 -n kafka-dev-namespace
```
*Get pods in project/namespace(wider info)*
```bash
oc get pods -o wide
```
*Get endpoints for service kafka-brokers-headless*
```
oc get ep kafka-brokers-headless
```
*Describe service*
```bash
oc describe svc kafka-brokers-headless
```
*Get all services*
```bash
oc get service
```

