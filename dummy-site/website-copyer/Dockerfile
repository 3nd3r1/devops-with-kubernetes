FROM nginx:alpine

RUN apk add --no-cache wget && rm -rf /usr/share/nginx/html/*

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
