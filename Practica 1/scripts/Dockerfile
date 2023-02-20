FROM ubuntu:latest

RUN apt-get update && apt-get install  -y bash

COPY . . /

RUN chmod +x ./script.sh

CMD [ "/bin/bash", "./script.sh" ]