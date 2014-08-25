FROM ubuntu
RUN sudo apt-get update
RUN sudo apt-get install -y nginx curl
ADD . .
EXPOSE 80
RUN chmod +x watch_etcd_autoproxy.sh
ENTRYPOINT /watch_etcd_autoproxy.sh
