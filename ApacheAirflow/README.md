![](../img/airflowlogo.png )
# Installin Apache Airflow in Ubuntu 20.04
## Create new linux user
First thing what would be good to do is add separate Linux user for airflow related things and not run everything under root:
```bash
# Creating linux user for airflow usage:
sudo useradd -m -s /bin/bash airflow_user
sudo passwd airflow_user ( <password>)
#Adding user airflow_user to group sudo
$sudo gpasswd -a airflow_user sudo
```
## Install Postgres database
Lot of tutorials available in google. Maybe will add later also here how it is done in Ubuntu, but some ide can be already read out in
[Postgres server management](https://github.com/kaidokariste/postgresql/blob/master/PostgresServerManagement.md)
If you have managed to install Postgres server and made initdb then create database and database user:  
```sql
psql # login to db as postgres user
CREATE DATABASE airflow_db;
CREATE USER airflow_user WITH PASSWORD '............';
GRANT ALL PRIVILEGES ON DATABASE airflow_db TO airflow_user;
```
## Intsall Python (also when you need update)
I have written it here [How to install or update Python](https://github.com/kaidokariste/python#updating-python-in-ubuntu)

## Installing airflow
The main command is below, but suggestion is to check if there exists corresponding Python version constraints file. Here you should check if there exists Python3.9 constraints for Airflow 2.3.3. Upgrading Airflow, the the command is the same, just versions change. 
```python
# Airflow upgrade
pip3 install "apache-airflow==2.3.3" --constraint "https://raw.githubusercontent.com/apache/airflow/constraints-2.3.3/constraints-3.9.txt"
```

> **airflow command not found**  
> Normal behaviour would be that as airflow_user you would be able to use command $ airflow ... when airflow is installed. As we installed airflow user, then the path > where this command sits is not added automatically under PATH variable. It is sitting under ~/.local/bin/ folder. You can do every time export PATH or add it to the > end of ~/.bashrc as described here [Permanently add a directory to shell path](https://linuxconfig.org/permanently-add-a-directory-to-shell-path)  

Although we should now see airflow folder then python3 install does not create it automatically. We do now airflow database initiation but this time it is still SQLite database. 
```bash
#Python3 does not creat airflow folder automatically. So initate database
airflow db init
```

Now you should see airflow_folder under airflow_user. Go to airflow configuration file and point it to postgres database using user, database and credentials we created eralier.
```
# Setting database to PG running locally in airflow.cfg:
sql_alchemy_conn = postgresql+psycopg2://airflow_user:............@localhost:5432/airflow_db
```
