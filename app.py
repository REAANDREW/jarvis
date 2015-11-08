import os
import random
import time
import sys
import docker
import pystache
import yaml
import HTMLParser
import uuid
import tempfile
import shutil
import time
import subprocess

h = HTMLParser.HTMLParser()

from flask import Flask, request, render_template, session, flash, redirect, url_for, jsonify
from celery import Celery
from celery.task.control import inspect
from docker import Client

cli = Client(base_url='unix://var/run/docker.sock')

app = Flask(__name__)
app.config['SECRET_KEY'] = 'top-secret!'

# Celery configuration
app.config['CELERY_BROKER_URL'] = 'redis://0.0.0.0:6379/0'
app.config['CELERY_RESULT_BACKEND'] = 'redis://0.0.0.0:6379/0'
#app.config['CELERY_ALWAYS_EAGER'] = True

# Initialize Celery
celery = Celery(app.name, broker=app.config['CELERY_BROKER_URL'])
celery.conf.update(app.config)

@celery.task(bind=True)
def long_task(self):
    p = subprocess.Popen(["docker", "--version"], stdout=subprocess.PIPE)
    output, err = p.communicate()
    print("DOCKER VERSION",output)
    suffix = str(uuid.uuid4())
    dirpath = tempfile.mkdtemp(suffix=suffix)
    output_path = os.path.join(dirpath,"output")
    build_path = os.path.join(dirpath,"build")
    os.mkdir(output_path)
    os.mkdir(build_path)
    with open('/code/template-build.sh','r+') as f:
        template = f.read()
        with open('/code/.travis.yml') as y:
            obj = yaml.load(y)
            obj["working_directory"] = "/go/src/github.com/guzzlerio/enanos"
            obj["repository"] = "git@github.com:guzzlerio/enanos.git"
            obj["branch"] = "develop"
            build_file_contents=h.unescape(pystache.render(template, obj))
            with open(os.path.join(build_path,"build.sh"),"w+") as build_file:
                build_file.write(build_file_contents)

    print("OUTPUT PATH", os.path.exists(output_path))
    print("BUILD PATH", os.path.exists(build_path))
    print("BUILD FILE PATH", os.path.exists(os.path.join(build_path,"build.sh")))

    volumes = ["/var/ci/build","/var/ci/output"]
    binds = ["{path}:/var/ci/build".format(path=output_path),"{path}:/var/ci/output/".format(path=build_path)]
    print("BINDS", binds)
    container = cli.create_container(image='golang:1.5', command='bash -c "bash /var/ci/build/build.sh"', volumes=volumes, host_config=docker.utils.create_host_config(binds=binds))
    response = cli.start(container=container.get("Id"))
    message = ""
    if response is not None:
        print(response)
    else:
        for line in cli.logs(container=container.get("Id"), stderr=True, stdout=True, stream=True):
            message.append(line)
            self.update_state(state='PROGRESS', meta={'status': message})
        return message
    #self.update_state(state='PROGRESS', meta={'status': "1"})
    #time.sleep(2)
    #self.update_state(state='PROGRESS', meta={'status': "2"})
    #time.sleep(2)
    #self.update_state(state='PROGRESS', meta={'status': "3"})
    #time.sleep(2)
    #self.update_state(state='PROGRESS', meta={'status': "4"})
    #time.sleep(2)
    #self.update_state(state='PROGRESS', meta={'status': "5"})
    #return "done"


@app.route('/', methods=['GET', 'POST'])
def index():
    msg = "Scheduled <br/>"
    report = inspect()
    for i in report.scheduled():
        msg.append(str(i)+"<br/>")

    msg = "Active <br/>"
    for i in report.active():
        msg.append(str(i)+"<br/>")
    return msg

@app.route('/longtask', methods=['POST'])
def longtask():

    task = long_task.apply_async()
    return jsonify({}), 202, {'Location': url_for('taskstatus', task_id=task.id)}

@app.route('/status/<task_id>')
def taskstatus(task_id):
    task = long_task.AsyncResult(task_id)
    info = task.info
    return str(task.state) +'\n' + str(task.info)

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True)
