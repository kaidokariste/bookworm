![](img/cassandralogo.png )

* [Connecting to Cassandra from Datagrip](#Connecting-to-Cassandra-from-Datagrip)  
   * [CSQL](#CSQL) 

# Install Apache Cassandra on Windows
https://phoenixnap.com/kb/install-spark-on-windows-10

# Fast links

# Connecting to Cassandra from Datagrip
As our company is mainly using DataGrip for SQL related works, then I started to investigate possibility
to connect Cassandra, running in Datastax Cloud. As it was not trivial and finally took me around 2h, to set up connection, then
thought I would describe it also here.

## Steps to go trough
1. Make sure you have installed Cassandra Driver to DataGrip.  
**File > Data Sources..** , and if you add new Cassandra data source, you see also before connection credentials, wich driver it is using.  

2. Go to Datastax UI, choose your database, click connect, and under Connection Method choose driver. Now you have opportunity to "Download Secure Connect Bundle".
Save it some place, where you can find it later and unzip it. Inside the folder are some files. Important ones are:
- ca - under newly created datasource window, choose SSH/SSL, thick **Use SSL** and map it to *CA file* box
- cert - same thing, map it under *Client certificate file* box.
- config - open with notepad. There you get your connection parameters
- key - map it under *Client key file*
Mode can stay required

3. Connection credentials 
From config.json the mapping to General tab boxes are:
- host -> Host
- cql_port -> Port
- DB Username (not in config, but find in Datastax UI) -> User
- DB Password (you choose it when creating a dtabase) -> Password
- keyspace -> keyspace   

If filled correctly then now you should be able to connect.
You can do ```SELECT dateof(now()) FROM system.local ;``` for testing purpose. Should return local time
