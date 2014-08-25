sudo service nginx start
./autoproxy
nginx -s reload -c autoproxy.conf

cat /usr/share/nginx/autoproxy.conf

while :;
do
  curl -L -s "http://172.17.42.1:4001/v2/keys/endpoints?wait=true&recursive=true";
  ./autoproxy
  nginx -s reload -c autoproxy.conf

  cat /usr/share/nginx/autoproxy.conf
done
