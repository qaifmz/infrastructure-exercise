# Instructions

For this exercise, you need to write a shell script that can deploy the manifests in this directory
to any namespace that's passed as an argument. The shell script should:
* Check that a namespace name was passed and, if not, print an error message and exit 1.
* Print an error message and exit 1 if the provided namespace already exists.
* Deploy the manifests in this directory into the namespace. You'll need to use kubectl create commands
  with the -n flag.
* After deploying those manifests, wait until the pods are ready. These pods may take some time to start.
* Print a friendly success message to the screen.
* Exit 0.

The web app being deployed by these manifests is the Jake web page from before. And like before, you
can check that deployment was successful by getting the hostname from the service (substitute
$NAMESPACE with the namespace name you passed to your script) like this:

``` shell
kubectl get svc -n $NAMESPACE exercise-5-service -o jsonpath='{.status.loadBalancer.ingress[0].hostname}{"\n"}'
```

Like in previous exercises, you'll have to wait a couple of minutes for the new ELB to come up. Then
you should see another Jake page.
