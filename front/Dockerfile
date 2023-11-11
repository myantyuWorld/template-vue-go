FROM node:16.20
WORKDIR /app
COPY front/vue_app/ ./

RUN apt-get update -qq \
    && yarn 

# TODO : ここでこける、手動でコンテナに入ると成功する(６行目のyarnが走ってない？)
# CMD ["yarn", "run", "dev", "--host"]