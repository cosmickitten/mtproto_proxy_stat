[![Build Status](https://travis-ci.com/trigun117/mtproto_proxy_stat.svg?branch=master)](https://travis-ci.com/trigun117/mtproto_proxy_stat) [![Go Report Card](https://goreportcard.com/badge/github.com/trigun117/mtproto_proxy_stat)](https://goreportcard.com/report/github.com/trigun117/mtproto_proxy_stat)

# mtproto_proxy_stat

Telegram MTProto Proxy Docker image with statistics output

# Examples

MTProto Proxy browser output

![example of browser output](https://github.com/trigun117/mtproto_proxy_stat/blob/master/image.JPG)

MTProto Proxy Datadog graph

![datadog graph example](https://github.com/trigun117/mtproto_proxy_stat/blob/master/image1.JPG)

# Get Started

For start, build docker image from Dockerfile and run with this command
```
docker run -d \
-p 443:443 \
-p 80:80 \
-e WORKERS=16 \
-v proxy-config:/data \
--restart always \
image_name
```
and visit http://your_mtproto_server_ip

If you want send metrics to Datadog, start Datadog Docker agent and use this command
```
docker run -d \
-p 443:443 \
-p 80:80 \
-e WORKERS=16 \
-e DDGIP=datadog_ip \
-e TGN=datadog_tag_name \
-v proxy-config:/data \
--link=dd-agent \
--name=mtproto \
--restart always \
image_name
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details