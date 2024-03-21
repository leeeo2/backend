#/bin/bash


cp Dockerfile ..
cp entrypoint.sh ..

cd ..
ls
docker build -t $1 .

# rm -f Dockerfile entrypoint.sh
