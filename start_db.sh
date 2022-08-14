MONGO_ADMIN_USER="admin"
MONGO_ADMIN_PASS="bantengmerah"

docker run -d \
    -- name jadwalin-user-server \
    --p 27017:27017 \
    -e MONGO_USER=$MONGO_ADMIN_USER \
    -e MONGO_PASS=$MONGO_ADMIN_PASS \

docker run \
    --name golang-redis \
    -p 7001:6379 \
    -d redis