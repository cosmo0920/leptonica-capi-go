sudo apt-get install zlib1g-dev
sudo apt-get install libpng12-dev libjpeg62-dev libtiff4-dev libsdl-gfx1.2-dev libsdl-image1.2-dev libsdl1.2-dev libavcodec-dev libavdevice-dev libavformat-dev libavutil-dev
sudo apt-get install python-enchant python-poppler

cd /tmp
wget http://www.leptonica.com/source/leptonica-1.70.tar.gz
tar -xvzof leptonica-1.70.tar.gz
cd leptonica-1.70
./configure
n_processors="$(grep '^processor' /proc/cpuinfo | wc -l)"
make -j${n_processors}
sudo make install
sudo ldconfig
