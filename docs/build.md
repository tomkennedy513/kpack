# Builds

A Build is a resource that schedules and run a single [Cloud Native Buildpacks](http://buildpacks.io) build.

Corresponding `kp` cli command docs [here](https://github.com/vmware-tanzu/kpack-cli/blob/main/docs/kp_build.md).

Unlike with the [Image resource](image.md), using Builds directly allows granular control of when builds execute. Each build resource is immutable and corresponds to a single build execution. You will need to create a new build for every build execution as builds will not rebuild on source code and buildpack updates. Additionally, you will need to manually specify the source, and the cache volume. 

A Build resource is comparable to `pack build`. Are you familiar with pack? if so, you can check the [comparison section](#kpack-vs-pack) 

### Configuration

```yaml
apiVersion: kpack.io/v1alpha2
kind: Build
metadata:
  name: sample-build
spec:
  tags:
  - sample/image
  serviceAccountName: service-account
  builder:
    image: gcr.io/paketo-buildpacks/builder:base
    imagePullSecrets:
    - name: builder-secret
  cache:
    volume:
      persistentVolumeClaimName: persisent-volume-claim-name
  projectDescriptorPath: path/to/project.toml
  source:
    git:
      url: https://github.com/buildpack/sample-java-app.git
      revision: main
  activeDeadlineSeconds: 1800
  env:
  - name: "JAVA_BP_ENV"
    value: "value"
  resources:
    requests:
      cpu: "0.25"
      memory: "128M"
    limits:
      cpu: "0.5"
      memory: "256M"
  tolerations:
    - key: "key1"
      operator: "Exists"
      effect: "NoSchedule"
  nodeSelector:
    disktype: ssd
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
          - matchExpressions:
              - key: kubernetes.io/e2e-az-name
                operator: In
                values:
                  - e2e-az1
                  - e2e-az2
```

- `tags`: A list of docker tags to build. At least one tag is required.
- `serviceAccount`: The Service Account name that will be used for credential lookup. Check out the [secrets documentation](secrets.md) for more information. 
- `builder.image`: This is the tag to the [Cloud Native Buildpacks builder image](https://buildpacks.io/docs/using-pack/working-with-builders/) to use in the build. Unlike on the Image resource, this is an image not a reference to a Builder resource.    
- `builder.imagePullSecrets`: An optional list of pull secrets if the builder is in a private registry. [To create this secret please reference this link](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/#registry-secret-existing-credentials)
- `source`: The source location that will be the input to the build. See the [Source Configuration](#source-config) section below.
- `activeDeadlineSeconds`: Optional configurable max active time that the pod can run for
- `cache`: Caching configuration, two variants are available:
  - `volume.persistentVolumeClaimName`: Optional name of a persistent volume claim used for a build cache across builds.
  - `registry.tag`: Optional name of a tag used for a build cache across builds.
- `env`: Optional list of build time environment variables.
- `defaultProcess`: The [default process type](https://buildpacks.io/docs/app-developer-guide/run-an-app/) for the built OCI image
- `projectDescriptorPath`: Path to the [project descriptor file](https://buildpacks.io/docs/reference/config/project-descriptor/) relative to source root dir or `subPath` if set. If unset, kpack will look for `project.toml` at the root dir or `subPath` if set.
- `resources`: Optional configurable resource limits on `CPU` and `memory`.
- `tolerations`: Optional configurable pod spec tolerations
- `nodeSelector`: Optional configurable pod spec nodeSelector
- `affinity`: Optional configurabl pod spec affinity

> Note: All fields on a build are immutable. Instead of updating a build, create a new one.
 
##### <a id='source-config'></a>Source Configuration

The `source` field is a composition of a source code location and a `subpath`. It can be configured in exactly one of the following ways:

* Git

    ```yaml
    source:
      git:
        url: ""
        revision: ""
      subPath: ""
    ```
    - `git`: (Source Code is a git repository)
        - `url`: The git repository url. Both https and ssh formats are supported; with ssh format requiring a [ssh secret](secrets.md#git-secrets).
        - `revision`: The git revision to use. This value may be a commit sha, branch name, or tag.
    - `subPath`: A subdirectory within the source folder where application code resides. Can be ignored if the source code resides at the `root` level.

* Blob

    ```yaml
    source:
      blob:
        url: ""
        stripComponents: 0
      subPath: ""
    ```
    - `blob`: (Source Code is a blob/jar in a blobstore)
        - `url`: The URL of the source code blob. This blob needs to either be publicly accessible or have the access token in the URL
        - `stripComponents`: Optional number of directory components to strip from the blobs content when extracting.
    - `subPath`: A subdirectory within the source folder where application code resides. Can be ignored if the source code resides at the `root` level.

* Registry

    ```yaml
    source:
      registry:
        image: ""
        imagePullSecrets:
        - name: ""
      subPath: ""
    ```
    - `registry` ( Source code is an OCI image in a registry that contains application source)
        - `image`: Location of the source image
        - `imagePullSecrets`: A list of `dockercfg` or `dockerconfigjson` secret names required if the source image is private
    - `subPath`: A subdirectory within the source folder where application code resides. Can be ignored if the source code resides at the `root` level.



#### Status

When a build complete successfully its status will report the fully qualified built image reference.

If you are using `kubectl` this information is available with `kubectl get <build-name>` or `kubectl describe <build-name>`. 

```yaml
status:
  conditions:
  - lastTransitionTime: "2020-01-17T16:16:36Z"
    status: "True"
    type: Succeeded
  latestImage: index.docker.io/sample/image@sha256:d3eb15a6fd25cb79039594294419de2328f14b443fa0546fa9e16f5214d61686
  ...
``` 

When a build fails its status will report the condition Succeeded=False. 

```yaml
status:
  conditions:
  - lastTransitionTime: "2020-01-17T16:13:48Z"
    status: "False"
    type: Succeeded
  ...
``` 

### kpack vs pack 

So, you are a [pack][_pack] user trying to learn about [kpack][_kpack] and get your [Cloud Native Buildpacks][_cnb] journey to the next level? then you are in the right place, on the next sections we are going to explain the similarities between [pack][_pack] and [kpack][_kpack].

First of all, both [kpack][_kpack] and [pack][_pack] implement the [platform interface](https://github.com/buildpacks/spec/blob/main/platform.md) [specification](https://github.com/buildpacks/spec/blob/main/platform.md), but they do it for two non-overlapping contexts: while [pack][_pack] targets developers and local builds, [kpack][_kpack] manages containerization on day-2 and at scale and is a [Kubernetes](https://kubernetes.io/) native implementation.

We will define some basic use case scenarios and see how we can get the output from both tools.

#### Assumptions

In order to make our comparison very simple, lets make some assumptions:
1. Our application source code is one of the [samples](https://github.com/buildpacks/samples/tree/main/apps) application
2. We are going to use [Cloud Native Buildpacks](_cnb) [sample builder](https://hub.docker.com/r/cnbs/sample-builder) to build our application source code
3. We need **write** access to a remote registry to publish our application image

#### Build scenario

Let's define our most basic use case as follows: 

`As a [pack|kpack] user, I want to convert my application source code into an image and publish it into a remote registry`

##### Pack Implementation

In order to build our application source code using [pack][_pack] we need to run a command similar to this:

`pack build --publish --path apps/<APP> --builder cnbs/sample-builder:<bionic OR alpine> <app-image-name>`

After building your '<app-image-name>' must be written into your remote registry.

##### Kpack Implementation

How do we get a similar functionality to a `pack build` command using [kpack][_kpack]? the answer is the Build resource!

Once you have [kpack][_kpack] up and running on a kubernetes cluster, you need to create a Build resource and apply it to your cluster. for our scenario it looks like this:

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

Once you create yaml file, the next step is just to apply the resource into your kubernetes cluster, for example using 

```bash
kubectl apply -f <your-build-resource.yaml>
```

After building, your '<app-image-name>' must be also written into your remote registry.

**Note** Probably you will need to create some [secrets](secrets.md) to give [kpack][_kpack] access to your remote registry, but this is also required on [pack][_pack], so please check the documentation depending on your registry provider

#### Re-build scenario

`As a [pack|kpack] user, I want to rebuild my application source code after some change and publish a new image into a remote registry`

##### Pack Implementation

In [pack][_pack], in order to re-build your application image, you just need to run the `pack build` command after saving your application source code changes

##### Kpack Implementation

As we mentioned above, All fields on a build are immutable, this mean that every time we want to run a build we must create a new `Build` resource. One way to do this is using `generateName` field in our resource definition.

From our previous resource definition, let's remove the `metadata.name` field, and replace it with a `metadata.generateName` this value will be used by the server, to generate a unique name ONLY IF the Name field has not been provided.

```yaml
apiVersion: kpack.io/v1alpha2
kind: Build
metadata:
  generateName: sample-build- # this value will be a prefix 
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
Once you create yaml file, any time you want to create a build, just run

```bash
kubectl create -f <your-build-resource.yaml>
```

A new build resource will be created and a unique suffix will be added to the value provided, for example: `sample-build-2vsz5`

Note: use `create` instead of `apply` when using `generateName`

#### Rebase scenario

`As a [pack|kpack] user, I want to rebase my application image with a new run-image from the stack`

##### Pack Implementation

[pack][_pack] offers the `pack rebase` command to accomplish this goal, for example:

```bash
pack rebase --publish <app-image-name>
```

##### Kpack Implementation

A standalone build can be triggered to be rebase if [kpack][_kpack] detects it is a "rebase-able" build. 

A build is considered "rebase-able" if the following conditions are met: 

- the field `spec.lastBuild.stackId` is equal to <same-stack-id-as-the-builder> 
- An annotation key `image.kpack.io/reason` is equal to `STACK`

An example resource that also is configured to use the `generateName` field could be as follows:

```yaml
apiVersion: kpack.io/v1alpha2
kind: Build
metadata:
  generateName: sample-build- # this value will be a prefix 
  annotations:
    image.kpack.io/reason: STACK
spec:
  lastBuild:
    stackId: <same-stack-id-as-the-builder>
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
Once you create yaml file, just run

```bash
kubectl create -f <your-build-resource.yaml>
```

[kpack][_kpack] will create a pod execution the rebase operation


[_pack]:https://github.com/buildpacks/pack
[_kpack]:https://github.com/pivotal/kpack
[_cnb]:https://buildpacks.io
