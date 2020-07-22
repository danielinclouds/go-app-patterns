
# Small Tools for CI
This pattern describes how to include small tools (bash replacement) written in Go into Docker container used by CI pipeline.

## Pattern 1
Using builder pattern build apps and copy them to output container.

```sh
mv Dockerfile_pattern1 Dockerfile
docker build --no-cache -t ciimages/tooling:0.0.1 .
```

## Pattern 2
Build Docker containers for each app and using builder pattern copy binaries to output container.

```sh
cd image_splitter
docker build --no-cache -t image_splitter .
cd ..

cd printer
docker build --no-cache -t printer .

mv Dockerfile_pattern2 Dockerfile
docker build --no-cache -t ciimages/tooling:0.0.1 .
```