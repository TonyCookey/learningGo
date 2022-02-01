
FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD /usr/games/fortune -a
#ENTRYPOINT /usr/games/fortune -a