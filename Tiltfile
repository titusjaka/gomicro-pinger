load('ext://restart_process', 'docker_build_with_restart')

local_resource(
    'gomicro-pinger-compile',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/gomicro-pinger ./',
    deps=['./main.go', './internal/', './commands/'],
  )

docker_build_with_restart(
  'gomicro-pinger-image',
  context='.',
  entrypoint=['/app/gomicro-pinger'],
  dockerfile='deploy/Dockerfile.tilt',
  only=[
    './build',
  ],
  live_update=[
    sync('./build', '/app/build'),
  ],
)

k8s_yaml('deploy/ponger.micro-grpc.k8s.tilt.yaml')

k8s_resource(
    'ponger-micro-grpc',
    labels=['backend']
)

k8s_yaml('deploy/pinger.micro-grpc.k8s.tilt.yaml')

k8s_resource(
    'pinger-micro-grpc',
    labels=['backend']
)

# k8s_resource(
#     'example-go', 
#     port_forwards=6066,
#     resource_deps='example-go-compile'],
#   )
