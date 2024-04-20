# Removing vendor apps from Android phone in Windows

As the phone started to lag, I started to seek, if it is possible to remove some preinstalled Samsung apps from my phone

## ADB (Android Debug Bridge) installing in Windows
Firstly I used the automated install, but didn't work for me, so I just followed [How to install and use ADB on Windows](https://www.androidpolice.com/install-adb-windows-mac-linux-guide/) the **Perform a manual setup in Windows** part:

1. Downloaded the github repo as zip
2. Unzipped the folder and moved to C:\
3. Added the C:\nexus-tools-main to uservariables PATH.
4. Type `adb` in command line and if it work, then it prints out possible command list.
```
PS C:\> adb
```
I will list also all the other commands I used during clean up process  
```bash
adb start-server - starts the adb server  
adb kill-server - kills adb server
adb devices - lists all connected devices
adb shell - opens the shell so you can execute commands inside your phone
```

### References
[How to install and use ADB on Windows](https://www.androidpolice.com/install-adb-windows-mac-linux-guide/)  
[How to delete Samsung apps from your phone](https://www.androidpolice.com/how-to-delete-samsung-apps/)

