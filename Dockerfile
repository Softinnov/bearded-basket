
FROM jpetazzo/dind

#- custom from google/golang
RUN apt-get update -y && apt-get install --no-install-recommends -y -q curl build-essential ca-certificates git mercurial bzr
RUN mkdir /goroot && curl https://storage.googleapis.com/golang/go1.2.2.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1
RUN mkdir -p /gopath/bin /gopath/pkg /gopath/src

ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
#-

WORKDIR /gopath/src/github.com/softinnov/bearded-basket/

CMD ["bash"]
