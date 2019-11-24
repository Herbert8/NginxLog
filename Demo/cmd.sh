docker run -d --rm --name httpbin -p 12345:80 kennethreitz/httpbin



docker run -d --rm \
        --name intranet_nginx \
        -p 7000:80 \
        --link=httpbin:httpbin \
        -v /NginxLog/Intranet/conf/:/etc/nginx/conf.d/ \
        -v /NginxLog/Intranet/logs/:/etc/nginx/logs/ \
        openresty/openresty:centos


docker run -d --rm \
        --name dmz_nginx \
        -p 6000:80 \
        --link=intranet_nginx:intranet_nginx \
        -v /NginxLog/DMZ/conf/:/etc/nginx/conf.d/ \
        -v /NginxLog/DMZ/logs/:/etc/nginx/logs/ \
        openresty/openresty:centos


curl -i -H 'a:1' -H 'b:2' -d 'f=123' 'http://localhost:6000/dmz/anything?c=3'