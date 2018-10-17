# Build Docker Images Locally or in the Cloud with `bld`

One of the great things about [Cloud Native](https://www.cncf.io/) systems is that they're centered around containers. You can build your app, package it up into a container image, and then you can run the image in all sorts of places: your machine, your CI/CD system, your staging setup, and even production.

After the image is built, you're golden! But in order to build that image, you have to have a very specific setup. For example, to build Docker images on my Mac, I have to use virtualization to run a Docker server locally, and then get a Mac OS X CLI build to talk to that server. Docker provides installers to make this setup easy for me to do, and it's worth it because I use Docker a lot.

But what about systems where it's not worth it? It might be a friend's system who just wants to try building your Docker-based app once. Or it might be your CI/CD system where you literally can't install Docker. Or it might be on production where you want to build a Docker image from inside of a [Kubernetes](https://kubernetes.io) pod (inception, right!?!?!?)

Enter `bld`. This CLI 
