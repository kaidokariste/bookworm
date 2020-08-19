# Install Apache Spark on Windows
https://phoenixnap.com/kb/install-spark-on-windows-10

## Prerequisites
* A system running Windows 10
* A user account with administrator privileges (required to install software, modify file permissions, and modify system PATH)
* Command Prompt or Powershell
* A tool to extract .tar files, such as 7-Zip

**Step 1: Install Java**
You can check Java version. I have Java 14 
```
C:\Users\kaido>java -version
java version "14.0.2" 2020-07-14
Java(TM) SE Runtime Environment (build 14.0.2+12-46)
Java HotSpot(TM) 64-Bit Server VM (build 14.0.2+12-46, mixed mode, sharing)

C:\Users\kaido>
```
**Step 2: Download Apache Spark**
1. Open a browser and navigate to https://spark.apache.org/downloads.html.

2. Under the Download Apache Spark heading, there are two drop-down menus. Use the current non-preview version.

In our case, in Choose a Spark release drop-down menu select 2.4.5 (Feb 05 2020).
In the second drop-down Choose a package type, leave the selection Pre-built for Apache Hadoop 2.7.
3. Click the spark-2.4.5-bin-hadoop2.7.tgz link.

Apache Spark download page.
4. A page with a list of mirrors loads where you can see different servers to download from. Pick any from the list and save the file to your Downloads folder.

**Step 3: Install Apache Spark**
