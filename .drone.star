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
        "image": "publicawesome/golang:1.22.7-devtooling",
        "commands": [
            "./scripts/go-test.sh"
        ],
        "environment": {
            "GOPROXY": "http://goproxy"
        }
    }
