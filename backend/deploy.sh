echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

docker push sjtujj/api-auth
docker push sjtujj/api-user
docker push sjtujj/api-sellinfo
sudo docker push sjtujj/api-file
# docker api insert before this
docker push sjtujj/srv-auth
docker push sjtujj/srv-user
docker push sjtujj/srv-sellinfo
sudo docker push sjtujj/srv-file
# docker srv insert before this

openssl aes-256-cbc -K $encrypted_3b095c6852fb_key -iv $encrypted_3b095c6852fb_iv -in private.key.enc -out private.key -d
chmod 400 private.key
ssh -i private.key -p 30710 centos@202.120.40.8 bash update.sh