SOPS_DIR=""
SECRETS_DIR="./secrets"
PUBLIC_KEY_FILE=public-key.txt
KEYS_FILE=keys.txt
KEY_PAIR=$(age-keygen)
PUBLIC_KEY=$(echo $KEY_PAIR | grep -oh "\w*age1\w*")
PRIVATE_KEY=$(echo $KEY_PAIR | grep -oh "\w*AGE-SECRET-KEY-\w*")
CHART_SECRET_FILE="./charts/sops-operator/templates/age_secret.yaml"

if [[ $OSTYPE =~ ^linux ]]; then
    SOPS_DIR=$XDG_CONFIG_HOME/sops/age
else
    SOPS_DIR="$HOME/Library/Application Support/sops/age"
fi

if [[ ! -d $SOPS_DIR ]] ; then
    mkdir "$SOPS_DIR"
fi

if [[ ! -d $SECRETS_DIR ]] ; then
    mkdir "$SECRETS_DIR"
fi

if [[ ! -f "$SOPS_DIR/$KEYS_FILE" ]] ; then
    touch "$SOPS_DIR/$KEYS_FILE"
fi

echo "# public key for argocd demo: $PUBLIC_KEY" >> "$SOPS_DIR/$KEYS_FILE"
echo $PUBLIC_KEY > "$SECRETS_DIR/$PUBLIC_KEY_FILE"
echo $PRIVATE_KEY >> "$SOPS_DIR/$KEYS_FILE"

AGE_KUBE_SECRET=$(kubectl create secret generic --from-file "$SOPS_DIR/$KEYS_FILE" --dry-run=client -o=yaml age-keys)
echo "$AGE_KUBE_SECRET" > $CHART_SECRET_FILE