FROM python:2.7

ENV C_FORCE_ROOT 1
ENV APP_USER ci
ENV APP_ROOT /code
ENV PYTHONUNBUFFERED 1

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get -y install libffi-dev libssl-dev python-dev python-setuptools

RUN curl https://get.docker.com | bash

RUN easy_install pip

RUN groupadd -r ${APP_USER} \
    && useradd -r -m \
    --home-dir ${APP_ROOT} \
    -s /usr/sbin/nologin \
    -g ${APP_USER} ${APP_USER}

WORKDIR ${APP_ROOT}
ADD requirements.txt ${APP_ROOT}/
RUN pip install -r requirements.txt
RUN pip install git+https://github.com/mbr/flask-bootstrap --upgrade
#USER ${APP_USER}
ADD . ${APP_ROOT}/
RUN echo `ls ${APP_ROOT}`
