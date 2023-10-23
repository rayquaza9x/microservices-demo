# FROM public.ecr.aws/docker/library/kong:3.0.0
FROM kong:2.8.0

USER root

# ARG DEBIAN_FRONTEND="noninteractive"
# RUN apt-get update
# RUN apt-get install openssh-client -y

# RUN apk add openssh-client

# RUN echo "    IdentityFile ~/.ssh/id_rsa" >> /etc/ssh/ssh_config

RUN CURRENT_DIR=$(pwd)
COPY ./containers/kong/kong-middleman-plugin /custom-plugin/middleman
# RUN luarocks install /tmp/kong-plugin/middleman-0.1.1-1.rockspec
WORKDIR /custom-plugin/middleman
RUN luarocks make *.rockspec
WORKDIR $CURRENT_DIR

# RUN cd /tmp/kong-plugin && luarocks install middleman-0.1.1-1.rockspec

# RUN mkdir /root/.ssh/
# RUN touch /root/.ssh/known_hosts
# RUN eval "$(ssh-agent -s)"
# RUN ssh-add -l -E sha256
# RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
# RUN ssh-keyscan -t rsa,dsa,ed25519,ecdsa -H github.com >> ~/.ssh/known_hosts

# USER kong

# RUN git clone git@github.com:pantsel/kong-middleman-plugin.git /tmp/middleman

# RUN luarocks install https://github.com/pantsel/kong-middleman-plugin --check-lua-versions

USER kong