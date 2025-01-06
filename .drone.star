load("scripts/drone/pipelines/test_and_build.star", "pipeline_test_and_build")

def main(ctx):
    return [
        pipeline_test_and_build(ctx)
    ]
