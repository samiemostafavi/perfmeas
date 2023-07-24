# adv-mobile-info

Build the project for the Advantech routers (arm64)
```
./build.sh min linux-arm64
git add .
git commit -m "update"
git push
```

Move the binary to the device and run it.
```
curl -LJO -k https://github.com/samiemostafavi/advmobileinfo/raw/main/ami
chmod +x ami
mv ami /usr/bin/
```

Add the following to the scripts tab of the router
```
ami > /root/ami.log 2>&1 &
```
