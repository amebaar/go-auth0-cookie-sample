FROM node:16.11-alpine3.11
WORKDIR /usr/src/app

RUN mkdir ~/.npm-global && \
    npm config set prefix '~/.npm-global' && \
    echo ' export PATH=~/.npm-global/bin:$PATH' >> ~/.bash_profile && \
    source ~/.bash_profile