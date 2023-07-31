# perf-meas service

Build the project for the Advantech routers (arm64)
```
./build.sh min linux-amd64
git add .
git commit -m "update"
git push
```

Move the binary to the device and run it.
```
curl -LJO -k https://github.com/samiemostafavi/perfmeas/raw/main/pfm
chmod +x pfm
mv pfm /usr/bin/
```

To test it, once the server is running, use curl to send a POST request with a JSON payload containing a bash command. Replace <server_url> with the actual URL of your server (e.g., `http://localhost:50505/`):
```
curl -X POST -H "Content-Type: application/json" -d '{"cmd": "echo Hello from the server"}' <server_url>
```

