# Instructions

These manifests deploy a karma tracking app that has a frontend/web service and a Redis backend for
storing karma. The web service has the following endpoints:

  - GET  /api/healthz - Simple health check that checks for Redis connectivity
  - GET  /api/karma - Endpoint that retrieves the current karma value
  - POST /api/karma/increment - Endpoint to increment karma
  - POST /api/karma/decrement - Endpoint to decrement karma

The source code for the frontend and web service can be found in the "code" directory.

Fix the manifests. Note that the source code does not need to be modified and is only available for reference.

When complete, get the hostname from the service object like this:

``` shell
kubectl get svc -n exercise-4 karma-app -o jsonpath='{.status.loadBalancer.ingress[0].hostname}{"\n"}'
```


It takes a few minutes for the Elastic Load Balancer to spin up in AWS, during
which time you won't be able to hit it from your browser. You can use the host
command to check if it's ready. For example:
> host a32afe6e4c0fd11e8be070afdf0421c6-629025564.us-east-1.elb.amazonaws.com
Host a32afe6e4c0fd11e8be070afdf0421c6-629025564.us-east-1.elb.amazonaws.com not found: 3(NXDOMAIN)

When it's ready, you'll see output like this:
> host a32afe6e4c0fd11e8be070afdf0421c6-629025564.us-east-1.elb.amazonaws.com
a32afe6e4c0fd11e8be070afdf0421c6-629025564.us-east-1.elb.amazonaws.com has address 34.233.47.33

Then put that hostname into your browser and you should see a page that lets you raise and
lower your karma!
