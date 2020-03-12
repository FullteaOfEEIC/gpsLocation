FROM python:3.7

RUN apt update


RUN pip install --upgrade pip setuptools
RUN pip install flask sqlalchemy flask-restful flask-sqlalchemy sqlalchemy_utils flask-migrate pymysql gunicorn flask_marshmallow marshmallow-sqlalchemy
RUN pip install folium pandas

WORKDIR /mnt

CMD python app.py
