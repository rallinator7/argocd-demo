
print("""
-----------------------------------------------------------------
✨ Hello Tilt! This appears in the (Tiltfile) pane whenever Tilt
   evaluates this file.
-----------------------------------------------------------------
""".strip())

#Extensions
load('ext://helm_resource', 'helm_resource')

# Settings
update_settings ( max_parallel_updates = 1 , k8s_upsert_timeout_secs = 120 , suppress_unused_image_warnings = None )

#Global
nspath = "./demo/namespaces"
argo_app_path= "./demo/argo-app"
demo_deps_path="./demo/charts"
chart_path="./"

# Configuration Scripts
local_resource('setup-github', cmd='./scripts/setup-github.sh')
local_resource('setup-argo', cmd='./scripts/setup-argo.sh', resource_deps=['setup-github'])
local_resource('setup-key', cmd='./scripts/setup-key.sh')

#SOPS
#Namespace
k8s_yaml("{}/{}.yaml".format(nspath, "sops"))

#Chart
helm_resource(
  name="sops-operator",
  namespace='sops',
  chart="{}/{}".format(demo_deps_path, "sops-operator"),
  resource_deps=['setup-key']
)


# Argo
# Namespace
k8s_yaml("{}/{}.yaml".format(nspath, "argo"))

# #Charts
helm_resource(
  name="argo-cd",
  namespace='argo',
  chart="{}/{}".format(demo_deps_path, "argo-cd"),
  resource_deps=['sops-operator', 'setup-github'],
  port_forwards=["8080"]
)

helm_resource(
  name="argo-applicationset",
  namespace='argo',
  chart="{}/{}".format(demo_deps_path, "argocd-applicationset"),
  resource_deps=['sops-operator', 'argo-cd', 'setup-github' ]
)

k8s_yaml("{}/{}.yaml".format(argo_app_path, "greeter"))
k8s_resource(
  objects=['greeter:applicationset'],
  resource_deps=['argo-applicationset'],
  new_name="greeter-application-set"
)

# Microservice
# Namespaces
k8s_yaml("{}/{}.yaml".format(nspath, "production"))
k8s_yaml("{}/{}.yaml".format(nspath, "staging"))
k8s_yaml("{}/{}.yaml".format(nspath, "dev"))

#Dev Deployment
docker_build('greeter',
            context='.',
            dockerfile='./app/docker/Dockerfile',
            entrypoint='/main',
            only=[
                './app/cmd/main.go',
                './app/go.mod',
                './app/go.sum'
            ],
             live_update=[
                sync('./app/go.mod', '/service/go.mod'),
                sync('./app/go.sum', '/service/go.sum'),
                sync('./app/cmd', '/service/cmd'),
             ]
)

helm_resource(
  name="greeter",
  namespace='dev',
  chart="./chart",
  image_deps= ["greeter"],
  image_keys= [('deployment.image', 'deployment.tag')],
  port_forwards=["5555"]
)