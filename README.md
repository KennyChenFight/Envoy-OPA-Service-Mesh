# Envoy-OPA-Service-Mesh
Integrate Istio with Open-Agent-Policy to demo service-mesh

## How to run
you can use minikube to run k8s cluster
1. `minikube start`
2. use istioctl to install Istio in k8s cluser: `istioctl install`<br>
reference: https://istio.io/latest/docs/setup/install/istioctl/
3. apply sidecar and OPA policy etc.: `kubectl apply -f sidecar.yaml` 
4. apply service a: `kubectl apply -f service-a.yaml`
5. apply service b: `kubectl apply -f service-b.yaml`
6. access service-a or service-b:
   ```bash
   minikube service service-a --url
   http://127.0.0.1:52783 # it is port-forward docker container to localhost
   minikube service service-b --url
   http://127.0.0.1:52992
   ```
7. testing service-a can access service-b: `curl http://127.0.0.1:52783/calling-service-b?name=serviceb` <br>
   testing service-b can access service-a: `curl http://127.0.0.1:52992/calling-service-a?name=servicea` <br>
   each response code should be 200.

## Explain OPA policy
you can look at service-a main.go and service-b main.go code. <br>
policy rule: when one service want to access the other service it should provide correct path and query string. <br>
These are simple rules, just for demo. <br>
```rego
package istio.authz

import input.attributes.request.http as http_request

default allow = false

allow {
  http_request.method == "GET"
  input.parsed_path == ["service-b-hello"]
  input.parsed_query.name == ["serviceb"]
}

allow {
  http_request.method == "GET"
  input.parsed_path == ["service-a-hello"]
  input.parsed_query.name == ["servicea"]
}
```
