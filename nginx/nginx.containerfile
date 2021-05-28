FROM nginx:1.16

COPY site.conf /etc/nginx/sites-enabled/000-default
COPY site.conf /etc/nginx/sites-available/000-default