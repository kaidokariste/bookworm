# Install Apache Spark on Windows
https://phoenixnap.com/kb/install-spark-on-windows-10
# Fast links
`java -version` - Check if java is installed  
`python –-version` - Check if python is installed  
http://localhost:4040/ - UI for Spark
`spark-shell` - open Spark Shell

## Prerequisites
* A system running Windows 10
* A user account with administrator privileges (required to install software, modify file permissions, and modify system PATH)
* Command Prompt or Powershell
* A tool to extract .tar files, such as 7-Zip

**Step 1: Install Java**
You can check Java version. I have Java 14  
- [How to install Java](www.example.com)

```
C:\Users\kaido>java -version
java version "14.0.2" 2020-07-14
Java(TM) SE Runtime Environment (build 14.0.2+12-46)
Java HotSpot(TM) 64-Bit Server VM (build 14.0.2+12-46, mixed mode, sharing)

C:\Users\kaido>
```
**Step 2: Download Apache Spark**
1. Open a browser and navigate to https://spark.apache.org/downloads.html.
Currently available was
```
Choose a Spark release:  3.0.0 (Jun 18 2020)
Choose a package type:  Pre-built for Apache Hadoop 3.2 and later
```

Verify the download  
First check your file and then signatures given in spark dowload file.
```
C:\Users\kaido>certutil -hashfile C:\Users\kaido.kariste\Downloads\spark-3.0.0-bin-hadoop3.2.tgz SHA512
SHA512 hash of C:\Users\kaido.kariste\Downloads\spark-3.0.0-bin-hadoop3.2.tgz:
bfe45406c67cc4ae00411ad18cc438f51e7d4b6f14eb61e7bf6b5450897c2e8d3ab020152657c0239f253735c263512ffabf538ac5b9fffa38b8295736a9c387
CertUtil: -hashfile command completed successfully.
```

**Step 3: Install Apache Spark**
- create new folder "Spark" in C:\ and extract files from tgz file into this folder. Although the tutorial didn't say that, but i had to also extract.tar file. Then i got my necessary files

**Step 4: Add winutils.exe File**
- Now, create new folders Hadoop and bin on C
- Navigate to https://github.com/cdarlint/winutils and choose corresponding one.

**Step 5: Environment variables**
1. "Edit environment variables for your account"
2.  User variables:  
```
   SPARK_HOME  C:\Spark\spark-3.0.0-bin-hadoop3.2  
   HADOOP_HOME C:\Hadoop
```
3. In the Path add ```%SPARK_HOME%\bin``` and ```%HADOOP_HOME\bin%```

**Step 6: Launch Spark**
Open cmd and type `spark-shell`

```
C:\Users\kaido.kariste>spark-shell
Welcome to
      ____              __
     / __/__  ___ _____/ /__
    _\ \/ _ \/ _ `/ __/  '_/
   /___/ .__/\_,_/_/ /_/\_\   version 3.0.0
      /_/

Using Scala version 2.12.10 (Java HotSpot(TM) 64-Bit Server VM, Java 13.0.2)
Type in expressions to have them evaluated.
Type :help for more information.

scala>
```
Open a web browser and navigate to http://localhost:4040/. and you should see Spark UI
To exit Spark and close the Scala shell, press ctrl-d in the command-prompt window.

# Test spark
I created a file test-file-for-spark.txt
```
scala> val x =sc.textFile("C:\\Users\\kaido.kariste\\Documents\\test-file-for-spark.txt")
x: org.apache.spark.rdd.RDD[String] = C:\Users\kaido.kariste\Documents\test-file-for-spark.txt MapPartitionsRDD[1] at textFile at <console>:24
``` 
Output shows that RDD is created. Lets print out 11 lines
```
x.take(11).foreach(println)
```
