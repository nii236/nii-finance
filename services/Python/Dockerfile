FROM python:2.7-alpine

MAINTAINER jacob.everist@gmail.com

RUN pip install werkzeug
RUN pip install requests
RUN pip install protobuf


EXPOSE 4000

ADD . /python_app

WORKDIR /python_app

CMD [ "python", "yahoo_ticker.py" ]



