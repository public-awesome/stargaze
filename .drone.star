def main(ctx):
    return[
        test_and_build(ctx),
    ]

def test_and_build(ctx):
    return {
    "kind": "pipeline",
    "type": "docker",
    "name": "test_and_build",
    "steps": [
      {
        "name": "build",
        "image": "alpine",
        "commands": [
            "echo hello world"
        ]
      }
    ]
  }
