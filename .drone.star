
go_dev_image = "publicawesome/golang:1.22.7-devtooling"
go_image = "golang:1.22.7-alpine3.19"
def main(ctx):
    return[
        pipeline_test_and_build(ctx),
    ]


def pipeline_test_and_build(ctx):
    return {
    "kind": "pipeline",
    "type": "docker",
    "name": "test_and_build",
    "steps": [
      step_fetch(ctx),
      step_test(ctx),
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
            "wget https://github.com/CosmWasm/wasmvm/releases/download/v2.2.1/libwasmvm_muslc.x86_64.a -O /lib/libwasmvm_muslc.x86_64.a",
            "echo 'b3bd755efac0ff39c01b59b8110f961c48aa3eb93588071d7a628270cc1f2326  /lib/libwasmvm_muslc.x86_64.a' | sha256sum -c",
            "LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true  make build",
            "echo 'Ensuring binary is statically linked ...' && (file $PWD/bin/starsd | grep 'statically linked')"
        ],
        "environment": {
            "GOPROXY": "http://goproxy"
        }
    }
