FROM nvidia/cuda:9.0-cudnn7-devel-ubuntu16.04

# 更新apt源
COPY sources.list /etc/apt

# 更新apt并安装基础工具
RUN apt-get update --fix-missing
RUN  apt-get install -y wget bzip2 ca-certificates \
    libglib2.0-0 libxext6 libsm6 libxrender1 \
    git mercurial subversion vim 
    && rm -rf /var/lib/apt/lists/*

# 安装编译工具cmake
RUN  apt-get install -y build-essential cmake  && rm -rf /var/lib/apt/lists/*

# 安装anaconda
ENV LANG=C.UTF-8 LC_ALL=C.UTF-8
ENV PATH /opt/conda/bin:$PATH
RUN wget https://mirrors.tuna.tsinghua.edu.cn/anaconda/archive/Anaconda3-2020.02-Linux-x86_64.sh -O ~/anaconda.sh && \
    /bin/bash ~/anaconda.sh -b -p /opt/conda && \
    rm ~/anaconda.sh && \
    ln -s /opt/conda/etc/profile.d/conda.sh /etc/profile.d/conda.sh && \
    echo ". /opt/conda/etc/profile.d/conda.sh" >> ~/.bashrc && \
    echo "conda activate base" >> ~/.bashrc

# 安装tensorflow
RUN pip install  -i https://pypi.tuna.tsinghua.edu.cn/simple  tensorflow

# 安装pytorch
RUN pip install torch --user

# 安装opencv3依赖库
RUN apt-get install libgtk2.0-dev libavcodec-dev libavformat-dev libjpeg.dev libtiff4.dev libswscale-dev libjasper-dev  && rm -rf /var/lib/apt/lists/*


# 安装npm和nodejs
RUN pip install npm && conda install -c conda-forge nodejs

# 安装caffe
RUN  apt-get install -y --no-install-recommends build-essential libatlas-base-dev libboost-all-dev libgflags-dev  libgoogle-glog-dev libhdf5-serial-dev libleveldb-dev liblmdb-dev libopencv-dev  libprotobuf-dev libsnappy-dev protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*
ENV CAFFE_ROOT=/opt/caffe
RUN cd $CAFFE_ROOT && git clone https://github.com/BVLC/caffe.git .

# 安装MNXnet
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ mxnet-cu90

