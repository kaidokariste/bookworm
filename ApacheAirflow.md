![](img/airflowlogo.png )
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
[Install/update Python](https://github.com/kaidokariste/python#updating-python-in-ubuntu)
