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

## Preparing the phone

Goal is to reach place where the developer mode is enabled and you can turn USB connection there.
For me it was "Information about phone" -> "Software info" > "Build number" aka _Järgu number_ and tap it like 5 times.  
Now go to developers mode and turn on "USB debugging" and connect the phone via USB to computer.

For pictured instruction look chapter "How to delete Samsung apps using Android Debug Bridge" in [How to delete Samsung apps from your phone](https://www.androidpolice.com/how-to-delete-samsung-apps/).

### Some problems I faced in this step

When I looked connected devices and looked if i can see them, then saw
```
PS C:\> adb start-server
* daemon not running; starting now at tcp:5037
* daemon started successfully
PS C:\> adb devices
List of devices attached
33002f65b6cc14d7        unauthorized
```
Device was unauthorized and i couldnt execute the shell command. What helped was disconnecting the phone and pressing "Revoke USB debugging authorization" (Under the USB debugging mode) and then reconnecting with USB.


### References
[How to install and use ADB on Windows](https://www.androidpolice.com/install-adb-windows-mac-linux-guide/)  
[How to delete Samsung apps from your phone](https://www.androidpolice.com/how-to-delete-samsung-apps/)

