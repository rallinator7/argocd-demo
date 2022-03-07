# Your GitHub OAuth Token Goes Here
export GITHUB_TOKEN=""
# Replace With Your Own Email
EMAIL=""

ssh-keygen -t ed25519 -C "$EMAIL" -f ./secrets/key-pair -q -N ""

cd github
terraform init
terraform apply -auto-approve