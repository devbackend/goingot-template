#!/bin/bash

IFS=', ' read -r -a addr_array <<<"$PROD_SERVER_ADDR"

for element in "${addr_array[@]}"; do
  ssh gitlab-runner@"$element" docker login "$CI_REGISTRY" -u "$REGISTRY_USER" -p "$REGISTRY_TOKEN"
  ssh gitlab-runner@"$element" docker pull "$CI_REGISTRY_IMAGE":"$CI_COMMIT_REF_SLUG" || exit 1
  ssh gitlab-runner@"$element" docker stop "$CI_PROJECT_NAME" || echo nothing to stop
  ssh gitlab-runner@"$element" docker rm "$CI_PROJECT_NAME" || echo nothing to remove
  ssh gitlab-runner@"$element" docker run -d \
    --name "$CI_PROJECT_NAME" --net host \
    -h "\$HOSTNAME" \
    -v /opt/corplimits/config:/opt/corplimits/config:ro \
    --log-opt max-size=2000m --log-opt max-file=3 \
    --restart=always \
    "$CI_REGISTRY_IMAGE":"$CI_COMMIT_REF_SLUG"
  ssh gitlab-runner@"$element" docker ps
  ssh gitlab-runner@"$element" docker logs "$CI_PROJECT_NAME"
done
