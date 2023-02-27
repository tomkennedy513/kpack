# Are you familiar with pack?

So, you are a  user trying to learn about [kpack](_kpack) and get your [Cloud Native Buildpacks](_cnb) journey to the next level? then you are in the right place, on the next sections we are going to explain the similarities between [pack](_pack) and [kpack](_kpack).

First of all, both [kpack](_kpack) and [pack](_pack) implement the [platform interface](https://github.com/buildpacks/spec/blob/main/platform.md) [specification](https://github.com/buildpacks/spec/blob/main/platform.md), but they do it for two non-overlapping contexts: while [pack](_pack) targets developers and local builds, [kpack](_kpack) manages containerization on day-2 and at scale and is a [Kubernetes](https://kubernetes.io/) native implementation. 

## kpack build vs pack build

Let's start by defining a simple use case scenario and see how we can get the output from both tools. 

### Scenario

Let's define our most basic use case as follows: `As a [pack|kpack] user, I want to convert my application source code into an image and publish it into a remote registry`

#### Implementation

In order to make our implementation very simple, lets make some assumptions:
1. Our application source code is one of the [samples](https://github.com/buildpacks/samples/tree/main/apps) application 
2. We are going to use [Cloud Native Buildpacks](_cnb) [sample builder](https://hub.docker.com/r/cnbs/sample-builder) to build our application source code
3. We need **write** access to a remote registry to publish our application image

##### Pack Implementation

In order to build our application source code using [pack](_pack) we need to run a command similar to this:

`pack build --publish --path apps/<APP> --builder cnbs/sample-builder:<bionic OR alpine> <app-image-name>`

After building your '<app-image-name>' must be written into your remote registry.

##### Kpack Implementation

How do we get a similar functionality to a pack build command using kpack? the answer is the [Build](build.md) resource!

[kpack](_kpack) offers a [Build](build.md) resource which is comparable to `pack build` and provides a simplified interface with more control over the build lifecycle including interoperability with [builders](builders.md) built by [pack](_pack).

Once you have [kpack](_kpack) up and running on a kubernetes cluster, you need to create a [Build](build.md) resource and apply it to your cluster. for our scenario it looks like this:

```yaml
apiVersion: kpack.io/v1alpha2
kind: Build
metadata:
  name: sample-build # This can be any name
spec:
  tags:
    - <app-image-name>
  builder:
    image: cnbs/sample-builder:<bionic OR alpine>
  source:
    git:
      url: https://github.com/buildpacks/samples.git
      revision: main
    subPath: "apps/<APP>"
```

Once you create yaml file, the next step is just to apply the resource into your kubernetes cluster, for example using `kubectl apply -f <your-build-resource.yaml>`. After building, your '<app-image-name>' must be also written into your remote registry.

**Note** Probably you will need to create some [secrets](secrets.md) to give [kpack](_kpack) access to your remote registry, but this is also required on [pack](_pack), so please check the documentation depending on your registry provider


[_pack]: https://github.com/buildpacks/pack
[_kpack]: https://github.com/pivotal/kpack
[_cnb]: https://buildpacks.io
