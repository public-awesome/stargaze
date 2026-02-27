
go_dev_image = "publicawesome/golang:1.23.8-devtooling"
go_image = "golang:1.23.8-alpine3.20"
wasmvm_version = "v2.3.2"
wasmvm_x86_84_hash = "4d03a4bf508c89a303e8d7d0236feac44a40f6b6e221df4076968abe9d1e49c6"
docker_image = "docker:24"
docker_dind_image = "docker:dind"

def pipeline_test_and_build(ctx):
    return {
    "kind": "pipeline",
    "type": "docker",
    "name": "test_and_build",
    "steps": [
      step_fetch(ctx),
      step_debug_dind(ctx),
      step_test(ctx),
      step_build(ctx),
      step_build_docker(ctx),

    ],
    "volumes": [
      create_volume_dockersock(ctx)
    ],
    "services": [
      service_dind(ctx)
    ]
  }

# Fetch the latest tags from the repository
def step_fetch(ctx):
    return {
        "name": "fetch",
        "image": "alpine/git",
        "commands": [
            "git fetch --tags"
        ]
    }

def step_test(ctx):
    return {
        "name": "test",
        "image": go_dev_image,
        "commands": [
            "./scripts/go-test.sh"
        ],
        "environment": {
            "GOPROXY": "http://goproxy"
        }
    }

def step_build(ctx):
    return {
        "name": "build",
        "image": go_image,
        "commands": [
            "apk add --no-cache ca-certificates build-base git",
            "wget https://github.com/CosmWasm/wasmvm/releases/download/{}/libwasmvm_muslc.x86_64.a -O /lib/libwasmvm_muslc.x86_64.a".format(wasmvm_version),
            "echo '{} /lib/libwasmvm_muslc.x86_64.a' | sha256sum -c".format(wasmvm_x86_84_hash),
            "LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true  make build",
            "echo 'Ensuring binary is statically linked ...' && (file $PWD/bin/starsd | grep 'statically linked')"
        ],
        "environment": {
            "GOPROXY": "http://goproxy"
        }
    }


def step_build_docker(ctx):
    return {
        "name": "build_docker",
        "image": docker_image,
        "commands": [
            "docker build -t publicawesome/stargaze:latest ."
        ],
        "volumes": [
            mount_volume(ctx, "dockersock", "/var/run")
        ]
    }

def step_debug_dind(ctx):
    return {
        "name": "debug_dind",
        "image": "alpine",
        "commands": [
            "sleep 10",
            "ls -l /var/run/docker.sock",
            "test -S /var/run/docker.sock && echo 'Docker socket found' || echo 'Docker socket missing'"
        ],
        "volumes": [
            mount_volume(ctx, "dockersock", "/var/run")
        ]
    }

def service_dind(ctx):
    return {
        "name": "dind",
        "image": docker_dind_image,
        "privileged": True,
        "volumes": [
          mount_volume(ctx, "dockersock", "/var/run")
        ]
    }

def mount_volume(ctx, name, path):
    return {
        "name": name,
        "path": path
    }

def create_volume_dockersock(ctx):
    return {
        "name": "dockersock",
        "temp": dict()
    }

def volume_docker_export(ctx):
    return {
        "name": "docker_export",
        "path": "/containers/export"
    }
