
print("""
-----------------------------------------------------------------
✨ Hello Tilt! This appears in the (Tiltfile) pane whenever Tilt
   evaluates this file.
-----------------------------------------------------------------
""".strip())

#Extensions
load('ext://helm_resource', 'helm_resource')

#Global
nspath = "./namespaces"

# Configuration Scripts
local_resource('setup-github', cmd='./scripts/setup-github.sh')
local_resource('setup-key', cmd='./scripts/setup-key.sh')

#SOPS
#Namespace
k8s_yaml("{}/{}.yaml".format(nspath, "sops"))

#Chart
helm_resource(
  name="sops-operator",
  namespace='sops',
  chart="./charts/sops-operator",
  resource_deps=['setup-key']
)


# #Argo
# #Namespace
# k8s_yaml("{}/{}.yaml".format(nspath, "argo"))

# #Charts
# helm_resource(
#   name="argo-cd",
#   namespace='argo',
#   chart="./charts/argo-cd",
#   resource_deps=['sops-operator', 'setup-github']
# )

# helm_resource(
#   name="argo-applicationset",
#   namespace='argo',
#   chart="./charts/argo-applicationset",
#   resource_deps=['sops-operator', 'argo-cd', 'setup-github' ]
# )

#Purl Microservice
#Namespaces
k8s_yaml("{}/{}.yaml".format(nspath, "production"))
k8s_yaml("{}/{}.yaml".format(nspath, "staging"))