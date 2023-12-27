FROM python:3.10

WORKDIR /app

COPY requirements.txt /app/requirements.txt

COPY entrypoint.sh /app/entrypoint.sh

RUN pip install -r requirements.txt

CMD [ "sh", "entrypoint.sh" ]