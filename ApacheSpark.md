![](img/apachesparklogo.png )

* [Spark SQL Guide](#spark-sql-guide)  
   * [Reading in JSON file](#reading-in-the-json-file) 

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

# Spark SQL guide
* [Handling JSON files](https://spark.apache.org/docs/latest/sql-data-sources-json.html)  
* [Code for following short example](https://github.com/kaidokariste/java-textbook/tree/master/SPARK)  
Let's do different parsing on json file looking like this  

```
{
    "_id": {
        "$oid": "5bfbedba3262d7327c1849b9"
    },
    "address": {
        "city": "Tartu",
        "county": "Tartu maakond",
        "district": "Ränilinn",
        "street": "Kaubanduse",
        "houseNumber": "50",
        "houseType": "Apartment building"
    },
    "paymentDate": {
        "year": "2015",
        "month": "July",
        "season": "Summer"
    },
	"otherFees":[{"telefon": 17.00, "internet": 24.00}],
    "paymentDetails": {
        "centralHeating": 8.59,
        "waterHeating": 14.4,
        "waterConsumption": 16.22,
        "gas": 0.52,
        "garbageCollection": 2.88,
        "overallElectricity": [{"night": 5.0},{"day": 2.99}],
        "renovationFund": 31.48,
        "administrationFee": 10.5,
        "insurance": 2.3,
        "janitor": 5.09,
        "total": 103.03
    }
}
```
## Reading in the JSON File
```
Dataset<Row> df = ss.read().option("multiline","true").json(<path-to-json-file>);
```
`.option("multiline","true")` - necessary only if the JSON looks like above. If the JSOn is in one row (no line breaks), then this is not necessary

Using `df.PrintSchema()` we will get JSON schema as follows
```
root
 |-- _id: struct (nullable = true)
 |    |-- $oid: string (nullable = true)
 |-- address: struct (nullable = true)
 |    |-- city: string (nullable = true)
 |    |-- county: string (nullable = true)
 |    |-- district: string (nullable = true)
 |    |-- houseNumber: string (nullable = true)
 |    |-- houseType: string (nullable = true)
 |    |-- street: string (nullable = true)
 |-- otherFees: array (nullable = true)
 |    |-- element: struct (containsNull = true)
 |    |    |-- internet: double (nullable = true)
 |    |    |-- telefon: double (nullable = true)
 |-- paymentDate: struct (nullable = true)
 |    |-- month: string (nullable = true)
 |    |-- season: string (nullable = true)
 |    |-- year: string (nullable = true)
 |-- paymentDetails: struct (nullable = true)
 |    |-- administrationFee: double (nullable = true)
 |    |-- centralHeating: double (nullable = true)
 |    |-- garbageCollection: double (nullable = true)
 |    |-- gas: double (nullable = true)
 |    |-- insurance: double (nullable = true)
 |    |-- janitor: double (nullable = true)
 |    |-- overallElectricity: array (nullable = true)
 |    |    |-- element: struct (containsNull = true)
 |    |    |    |-- day: double (nullable = true)
 |    |    |    |-- night: double (nullable = true)
 |    |-- renovationFund: double (nullable = true)
 |    |-- total: double (nullable = true)
 |    |-- waterConsumption: double (nullable = true)
 |    |-- waterHeating: double (nullable = true)
```
If the key-value pair is inside JSOn object, the quering is quite straight forward (example _address.city_)
Things will go a bit complicated, when they are inside arrays. But following example will show the logic.
```
        // SQL statements can be run by using the sql methods provided by spark
        Dataset<Row> namesDF =
                ss.sql("SELECT address.city, otherFees[0].internet, paymentDetails.overallElectricity[1].day FROM maksud");
        namesDF.show();
```
In array cases, you have to specify array index where to search and then appropriate key. 




