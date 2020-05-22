# Rbitrage

![meow](docs/media/pirate-pixel-meowth.png)

Save 💵 on stateless, fault tolerant, distributed workloads.

## Setup

These instructions will go through setting up a simple dataplane cluster that leverages rbitrage  
*Note: These instructions use Amazon EKS for the control cluster*

### Prereqs

1. AWS CLI installed
2. EKS cluster provisioned
3. Helm installed

### Install and configure Crossplane

![colorful-icecream](docs/media/crossplane.png)

Install Crossplane and provider-aws into the control cluster:  

```bash
kubectl create namespace crossplane-system
helm repo add crossplane-alpha https://charts.crossplane.io/alpha
helm install crossplane \
    --namespace crossplane-system crossplane-alpha/crossplane \
    --version 0.8.0 \
    --set clusterStacks.aws.deploy=true \
    --set clusterStacks.aws.version=v0.6.0 \
    --disable-openapi-validation
```

Load AWS credentials into the control cluster so that Crossplane is able to provision infrastructure on your behalf.
In this example we will only deploy resources into `us-west-2`.
Here are [additional Crossplane docs](https://crossplane.io/docs/master/cloud-providers/aws/aws-provider.html) on how
to add AWS credentials.

```bash
BASE64ENCODED_AWS_ACCOUNT_CREDS=$(echo -e "[default]\naws_access_key_id = $(aws configure get aws_access_key_id --profile default)\naws_secret_access_key = $(aws configure get aws_secret_access_key --profile default)" | base64  | tr -d "\n")

cat > aws-credentials.yaml <<EOF
---
apiVersion: v1
kind: Secret
metadata:
  name: aws-account-creds
  namespace: crossplane-system
type: Opaque
data:
  credentials: ${BASE64ENCODED_AWS_ACCOUNT_CREDS}
---
apiVersion: aws.crossplane.io/v1alpha3
kind: Provider
metadata:
  name: aws-provider-west
spec:
  credentialsSecretRef:
    name: aws-account-creds
    namespace: crossplane-system
    key: credentials
  region: us-west-2

kubectl apply -f "aws-credentials.yaml"
```

Create a namespace for your application

```bash
kubectl create namespace rbitrage-dataplane-example
```

### Install Argo CD

![space-worm](docs/media/argo.png)

Use the following commands to install Argo

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

Use the following command and navigate to `localhost:8080` in order to view the Argo UI on your local machine

```bash
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

For initial login, the username is admin and the password is the pod name of the Argo CD API server. To find your generated pod name, run the following command

```bash
kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2
```

### Deploy application infrastructure

In Argo CD, the term Application is used to refer to a set of configuration files that should be deployed as a single unit. An application allows you to specify a source repository for your configuration files, then it watches for updates and creates or updates objects in your Kubernetes cluster based on observed changes.

1. Launch the Argo UI, log in, go to `Settings > Projects > New Project`. For the sake of this example
use the `default` project.

2. Click `Edit as YAML` and use the following config:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: rbitrage
spec:
  destination:
    namespace: rbitrage-dataplane-example
    server: 'https://kubernetes.default.svc'
  source:
    path: examples/simple/app-cluster-infra
    repoURL: 'https://github.com/cmdallas/rbitrage.git'
    targetRevision: HEAD
    directory:
      recurse: true
  project: default
  syncPolicy:
    automated:
      automated:
        prune: false
        selfHeal: false
      prune: true
```

### Deploy the application in us-west-2

⚠️ Only continue to this step once the `rbitrage-dataplane-cluster` is ready. This may take some time.

1. In the Argo UI, create a new `Application`.

2. Click `Edit as YAML` and use the following config:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: rbitrage-app
spec:
  destination:
    namespace: rbitrage-dataplane-example
    server: 'https://kubernetes.default.svc'
  source:
    path: examples/simple/app
    repoURL: 'https://github.com/cmdallas/rbitrage.git'
    targetRevision: HEAD
    directory:
      recurse: true
  project: default
  syncPolicy:
    automated:
      automated:
        prune: false
        selfHeal: false
      prune: true
```
