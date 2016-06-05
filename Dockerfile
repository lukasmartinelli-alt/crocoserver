FROM python:3.5
WORKDIR /usr/src/app
COPY requirements.txt /usr/src/app/
RUN pip install --no-cache-dir -r requirements.txt

COPY . /usr/src/app
RUN pip install -e .
CMD ["croco"]
EXPOSE 5000
