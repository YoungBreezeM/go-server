#! /bin/bash
echo "====== START ====="

basepath=$(pwd)
pkgname="pkg"
mkdir -p $pkgname/dist
#
cd $basepath/server && ./build.sh
cp -r $basepath/server/build/* $basepath/$pkgname
cd $basepath/front-ui&& npm run build
cp -r $basepath/front-ui/dist $basepath/$pkgname
tar -cvf $pkgname.tar $pkgname/
#
echo "====== DONE ======"