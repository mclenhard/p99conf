FROM ubuntu:20.04 as build
ENV  DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
apt install -y bison python3-pip build-essential cmake flex git libedit-dev  \
libllvm7 llvm-7-dev libclang-7-dev python zlib1g-dev libelf-dev libfl-dev bpfcc-tools curl



# Let's people find our Go binaries


RUN git clone --branch  v0.23.0 https://github.com/iovisor/bcc.git
RUN mkdir bcc/build
WORKDIR  bcc/build
RUN cmake ..
RUN make
RUN  make install
RUN cmake -DPYTHON_CMD=python3 .. # build python3 binding
WORKDIR /bcc/build/src/python/
RUN make
RUN make install

WORKDIR /usr/sbin/
RUN curl -s https://storage.googleapis.com/golang/go1.19.1.linux-amd64.tar.gz | tar -v -C /usr/local -xz
ENV PATH $PATH:/usr/local/go/bin

COPY . .
RUN go mod download
RUN go build -o pod2pid





CMD [ "python3", "./main.py" ]