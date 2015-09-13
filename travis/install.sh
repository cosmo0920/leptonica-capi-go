#!/bin/bash
sudo apt-get install zlib1g-dev
sudo apt-get install libpng12-dev libjpeg62-dev libtiff4-dev libsdl-gfx1.2-dev libsdl-image1.2-dev libsdl1.2-dev libavcodec-dev libavdevice-dev libavformat-dev libavutil-dev
sudo apt-get install python-enchant python-poppler

export LEPTONICA_VERSION="1.71"
cd /tmp
wget http://www.leptonica.com/source/leptonica-${LEPTONICA_VERSION}.tar.gz
tar -xvzof leptonica-${LEPTONICA_VERSION}.tar.gz
cd leptonica-${LEPTONICA_VERSION}
./configure
n_processors="$(grep '^processor' /proc/cpuinfo | wc -l)"
make -j${n_processors}
sudo make install
sudo ldconfig
