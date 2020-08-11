#!/bin/sh
sudo docker-compose exec db bash -c "chmod 0775 db/init-db.sh"
sudo docker-compose exec db bash -c "/db/init-db.sh"
