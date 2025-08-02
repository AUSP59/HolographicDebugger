
#!/bin/bash
echo "Installing HoloDebug..."
mkdir -p /opt/holodebug
cp -r * /opt/holodebug/
ln -sf /opt/holodebug/bin/holodebug /usr/local/bin/holodebug
echo "Installed to /opt/holodebug"
