# `docker build` or `az acr build`?

One of the great things about [Cloud Native](https://www.cncf.io/) systems is that they're centered around containers. You can build your app, package it up into a container image, and then you can run the image in all sorts of places: your machine, your CI/CD system, your staging setup, and even production.

**`bld` is a little tool to make building container images a little easier.**

More background here: https://arschles.com/blog/az-acr-build-or-docker-build/

# How to Build this Binary

Make sure you're running with [Go modules](https://github.com/golang/go/wiki/Modules) turned on. You can also optionally set `GOPROXY` to a modules proxy (`https://microsoftgoproxy.azurewebsites.net` is one :smile:)

When you're ready, run:

```
go build
```

You'll have a `./bld` binary that you can use. Read on for how.

# How to Use `bld`

I try to make it simple to use the binary. You always run something like this:

```
bld -t my/image -f my/Dockerfile my/build/context
```

The `-f` flag is optional, the other stuff is required. If `docker` is available in the executable `PATH`, then `bld` will run `docker build -t my/image -f my/Dockerfile my/build/context`. Otherwise, it'll run `az acr build -t my/image my/Dockerfile my/build/context`.

In the former case, make sure your Docker daemon is up and running, and your `docker` CLI can talk to it.

In the latter case, make sure you're logged into your Azure account from the `az` CLI. Also remember that you have to pay for stuff! For example:

- You have to pay for builds
- Builds automatically store images in [ACR](https://cda.ms/HC), and you also have to pay for that

