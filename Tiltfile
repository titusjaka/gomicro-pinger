load('ext://restart_process', 'docker_build_with_restart')

local_resource(
    'gomicro-pinger-compile',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/gomicro-pinger ./',
    deps=['./main.go', './internal/', './commands/'],
    labels=['build']
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
    sync('./build', '/app'),
  ],
)

k8s_yaml('deploy/ponger.micro-grpc.k8s.tilt.yaml')
k8s_resource(
    'ponger-micro-grpc',
    labels=['server']
)

k8s_yaml('deploy/pinger.micro-grpc.k8s.tilt.yaml')
k8s_resource(
    'pinger-micro-grpc',
    labels=['client']
)

k8s_yaml('deploy/dashboard.k8s.tilt.yaml')
k8s_resource(
    'micro-dashboard',
    labels=['dashboard']
)

k8s_resource(
    'micro-dashboard', 
    port_forwards='8081',
    # resource_deps=['micro-dashboard'],
  )
