# ArgoCD Demo


# What is this?
This repo demonstrates how argocd can be used to deploy microservices to environments in Kubernetes Clusters

# Dependencies

This repo assumes you already have working knowledge on setting up and deploying to kubernetes clusters.  This tutorial uses tilt, which assumes you have a working cluster running that can be connected to from the machine you are running the demo on.

- [age](https://github.com/FiloSottile/age)
- [sops](https://github.com/mozilla/sops#encrypting-using-age)
- tilt
- helm
- kubectl
- terraform

# Creating a GitHub OAuth Token

This repo uses the GitHub Terraform provider to create and delete a temporary SSH key pair for the ArgoCD Operator. If you already have an admin OAuth Token setup for your account, you can add it to the setup-github.sh file.  This is used to create a temporary ssh key for the argocd operator to access the repos you want to deploy.  If you do not, GitHub has a [tutorial](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) on how to do so.  Make sure to copy and paste the token into the setup-github.sh file so that you don't lose it.