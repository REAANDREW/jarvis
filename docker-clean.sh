#! /bin/bash
for i in `docker ps -a | tr -s " " | grep -v Up | cut -f 1 -d " " | tail -n +2`
do 
	docker rm $i
done
for i in `docker images | tr -s " " | grep none | cut -f 3 -d " "`
do
	docker rmi $i
done

