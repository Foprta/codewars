# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.134.0/containers/cpp/.devcontainer/base.Dockerfile
ARG VARIANT="buster"
FROM mcr.microsoft.com/vscode/devcontainers/cpp:0-${VARIANT}

# Installing GoogleTest
RUN cd /tmp googletest && git clone https://github.com/google/googletest.git \
    && cd googletest && mkdir build && cd build \
    && cmake .. -DBUILD_GMOCK=OFF -DBUILD_SHARED_LIBS=ON -DINSTALL_GTEST=ON -DCMAKE_INSTALL_PREFIX:PATH=/usr \
    && make && sudo make install && sudo ldconfig