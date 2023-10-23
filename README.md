# TemporalIO Hello World

## Running locally

> The compose file and dynamicconfig was copy n paste by the official repository: https://github.com/temporalio/docker-compose

- start temporal cluster and it's dependencies: `docker compose up -d`
- execute worker in background to run your code and send results to temporal cluster with: `make run-worker`
- to trigger workflow any time you want: `make run-workflow`
- and then you can check results on UI: http://localhost:8080/namespaces/default/workflows

## Ref

- https://learn.temporal.io/getting_started/go/hello_world_in_go/
- https://docs.temporal.io/kb/all-the-ways-to-run-a-cluster#docker--docker-compose