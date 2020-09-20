# Instructions

There's something wrong with these manifests. Apply them, see what breaks, then fix it.

The code (config, really) for this web page is provided in the code directory. You don't
need to modify it, it's just there for reference.

When complete, get the hostname from the service object like this:

``` shell
kubectl get svc -n exercise-3 exercise-3-service -o jsonpath='{.status.loadBalancer.ingress[0].hostname}{"\n"}'
```

It takes a few minutes for the Elastic Load Balancer to spin up in AWS, during
which time you won't be able to hit it from your browser. You can use the host
command to check if it's ready. For example:
> host a32afe6e4c0fd11e8be070afdf0421c6-629025564.us-east-1.elb.amazonaws.com
Host a32afe6e4c0fd11e8be070afdf0421c6-629025564.us-east-1.elb.amazonaws.com not found: 3(NXDOMAIN)

When it's ready, you'll see output like this:
> host a32afe6e4c0fd11e8be070afdf0421c6-629025564.us-east-1.elb.amazonaws.com
a32afe6e4c0fd11e8be070afdf0421c6-629025564.us-east-1.elb.amazonaws.com has address 34.233.47.33

Then put that hostname into your browser and you should see a picture of CiBO's unofficial mascot!
