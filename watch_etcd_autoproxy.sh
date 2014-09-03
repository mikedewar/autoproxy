sudo service nginx start
./autoproxy
sudo service nginx restart

cat /etc/nginx/nginx.conf

while :;
do
  curl -L -s "http://172.17.42.1:4001/v2/keys/endpoints?wait=true&recursive=true";
  ./autoproxy
  sudo service nginx restart
  cat /etc/nginx/nginx.conf
done
