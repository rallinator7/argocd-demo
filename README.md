# ArgoCD Demo


# What is this?
This repo demonstrates how argocd can be used to deploy microservices to environments in Kubernetes Clusters

# Setup

## Dependencies

This repo assumes you already have working knowledge on setting up and deploying to kubernetes clusters.  This tutorial uses tilt, which assumes you have a working cluster running that can be connected to from the machine you are running the demo on.

- [age](https://github.com/FiloSottile/age)
- [sops](https://github.com/mozilla/sops#encrypting-using-age)
- [yq](https://mikefarah.gitbook.io/yq/)
- tilt
- helm
- kubectl
- terraform


## Config File

There are several scripts in the repo that will require inputs from you, the end user.  To streamline this process, it's important that you fill out a config file so that the scripts will have all the resources they need.  Under the secrets repo, create a file called `config.yaml`.  This will hold information prior to starting the stack so that the startup can be automated.  Copy and paste this yaml structure into the config file before continuing on:

    github:
      email: # your github email
      token: # your github token
    argo:
      baseRepo: # Base of the repo you want to deploy. ex) if your app is in an org, use the ssh string for the repo all the way up until you hit the repo name
      fullRepo: # Repo that the demo app is in
      adminPassword: # The Password for the Argo UI
      publicKey: 
      privateKey: 
    sops:
      publicKey: 


## Creating a GitHub OAuth Token

This repo uses the GitHub Terraform provider to create and delete a temporary SSH key pair for the ArgoCD Operator. Before we can run the terraform code, we need to create an OAuth Token for GitHub. The token is used to create a temporary ssh key for the argocd operator to access the repos you want to deploy.  GitHub has a [tutorial](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) on how to do so.  Make sure the token has the ability to configure ssh keys for your profile.