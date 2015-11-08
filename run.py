import sys
import docker
from docker import Client
cli = Client(base_url='unix://var/run/docker.sock')

volumes = ["/var/ci/build","/var/ci/output"]
binds = ["/home/vagrant/jarvis/build:/var/ci/build","/home/vagrant/jarvis/output:/var/ci/output/"]
container = cli.create_container(image='golang:1.5', command='bash -c "bash /var/ci/build/build.sh"', volumes=volumes, host_config=docker.utils.create_host_config(binds=binds))
response = cli.start(container=container.get("Id"))
if response is not None:
    print(response)
else:
    for line in cli.logs(container=container.get("Id"), stderr=True, stdout=True, stream=True):
        sys.stdout.write(line)

