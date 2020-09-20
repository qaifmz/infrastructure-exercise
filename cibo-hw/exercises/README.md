# Overview

In this directory, you should find 5 exercise directories. Within each directory, you'll find
a README file with instructions and several Kubernetes manifests you need to deploy. Some of
these manifests will NOT apply cleanly, so you'll need to figure out what's wrong.

Each exercise will have you deploy a web app that can be verified by visiting some URL with
your browser. Each exercise's README will have more details about how to get this URL and
what you should expect to see.

Feel free to use any online resources to help you solve these problems. But you should know...
any commands you run on this host are being logged!

# Other Notes

If at any time you need to totally reset the exercise, just delete the corresponding namespace.
Each exercise has a manifest for recreating the namepsace. For example:
kubectl delete namespace exercise-1

If you need to reset the manifest files for an exercise, you can re-extract them from the tarball.
You don't have permissions to modify or delete that tarball, so the manifests within will be
pristine no matter what you accidentally mess up!

Take some notes as you progress. We're interested in your thought process, including dead-ends
or bad ideas, just as much as we're interested in a functional web page.
