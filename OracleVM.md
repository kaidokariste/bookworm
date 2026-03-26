# Connect PuTTY to Ubuntu 22.04 (VirtualBox) via SSH Port Forwarding  
This guide explains how to connect from your host machine to an Ubuntu 22.04 virtual machine running in VirtualBox using PuTTY and SSH.



## Requirements

- VirtualBox installed
- Ubuntu 22.04 VM
- PuTTY installed (Windows)

## 1. Install SSH Server in Ubuntu VM

Open terminal inside your Ubuntu VM and run:

```bash
sudo apt update
sudo apt install openssh-server
```

# Check SSH status:

```bash
sudo systemctl status ssh
```

If not running:

```bash
sudo systemctl enable ssh
sudo systemctl start ssh
```

2. Configure Port Forwarding in VirtualBox

Open VirtualBox

Select your VM → Settings

Go to Network

Adapter 1:
Attached to: NAT
Click Advanced → Port Forwarding

Add a new rule:

Name	Protocol	Host IP	Host Port	Guest IP	Guest Port
SSH	TCP		2222		22

Click OK to save.

3. Find Your Ubuntu Username

Run:

```bash
whoami
```

Example output:

ubuntuuser

4. Connect Using PuTTY

Open PuTTY and configure:

Host Name: 127.0.0.1
Port: 2222
Connection Type: SSH

Click Open

5. Login

Enter your credentials:

login as: ubuntuuser
password: yourpassword

🔁 Connection Flow
PuTTY (Host Machine)
    ↓
127.0.0.1:2222
    ↓
VirtualBox NAT Port Forwarding
    ↓
Ubuntu VM (port 22)
✅ Test Connection (Optional)

From your host terminal:

ssh username@127.0.0.1 -p 2222
