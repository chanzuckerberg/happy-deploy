# happy
Happy Path Deployment Tool

### Security

Please note: If you believe you have found a security issue, please responsibly disclose by contacting us at security@chanzuckerberg.com

Visit the Happy Path documentation for more details: https://chanzuckerberg.github.io/happy/

Happy Path is an open source project intended to simplify web application, microservice, and cron jobs deployments, in adherence to CZI security practices. 

### Features
* Manages short-lives infrastructure (we deploy into your compute)
* Groups services together (we call it a `stack` for co-deployment), each `stack` is isolated, and you can have multiple `stacks` created for the same application.
* Easily promote changes from lower to higher environments
* Has an extensive set of Github workflows
* Supports both AWS ECS and EKS compute, and allows for an easy migration between the two
* Abstracts out IaC code with the intent that developers should only focus on the application code
* Supports Linkerd service mesh, mTLS and service-to-service authorization when deployed on EKS with Linkerd installed
* Plays nicely with `external-dns`, `karpenter`, `cluster-autoscaler`
* Integrates with Datadog for dashboarding (assuming you have a `datadog` agent deployed into your EKS)
* Provides service discovery

### Prerequisites

You will need to have Docker desktop, AWS CLI, and `terraform` installed to use Happy.

### Install

Install `happy`:

#### MacOS
```sh
brew tap chanzuckerberg/tap
brew install happy
```

#### Linux

Binaries are available on the releases page. Download one for your architecture, put it in your path and make it executable.

Instructions on downloading the binary:

1. Go here: <https://github.com/chanzuckerberg/happy/releases> to find which version of happy you want.
2. Run `curl -s https://raw.githubusercontent.com/chanzuckerberg/happy/master/download.sh | bash -s -- -b HAPPY_PATH VERSION`
   1. HAPPY_PATH is the directory where you want to install happy
   2. VERSION is the release you want
3. To verify you installed the desired version, you can run `happy version`.


### Getting started

Clone this repo: 
```sh
git clone https://github.com/chanzuckerberg/happy.git
```

Navigate to an example app and try happy out:
```sh
cd examples/typical_app
happy list
happy create mystack
happy update mystack
happy delete mystack
```

### Contributing

This project adheres to the Contributor Covenant code of conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to opensource@chanzuckerberg.com.
//
