![](../img/snowflake.png )

# Introduction 


## Administrative commands
```sql
SHOW ROLES; -- Kõik sulle antud rollid
SHOW GRANTS TO USER "VICTOR.HUGO"; -- kasutaja rollide ja muude õiguste vaatamine. 
SELECT CURRENT_ROLE(); -- Hetke rolli vaatamine
SHOW SCHEMAS; -- Schemade vaatamine hetke andmebaasis
SHOW SCHEMAS IN DATABASE db_name;   
```

## Schema loomine

```sql
CREATE SCHEMA IF NOT EXISTS bigdata.kaido_kariste WITH MANAGED ACCESS
COMMENT = 'Kaido isiklik arendusschema' ;
```
**WITH MANAGED ACCESS**  
Määrate, kes saab jagada õigusi skeemi sees olevatele objektidele.
* Tavaline skeem: Objekti looja (nt tabeli loonud kasutaja) saab jagada teistele õigusi.
* Managed Access: Ainult skeemi omanik (roll, millel on OWNERSHIP õigus) saab jagada õigusi. See on turvalisuse mõttes eelistatud valik ettevõtetes.

### Kuidas protsess tavaliselt välja näeb?
1. SYSADMIN loob andmebaasi ja skeemi:

```sql
USE ROLE SYSADMIN;
CREATE SCHEMA prod.sales_analytics WITH MANAGED ACCESS;
```

2. SECURITYADMIN (või USERADMIN) loob rollid ja jagab ligipääsu.

```sql
USE ROLE SECURITYADMIN;
GRANT USAGE ON DATABASE prod TO ROLE analyst_role;
GRANT USAGE ON SCHEMA prod.sales_analytics TO ROLE analyst_role;
GRANT SELECT ON ALL TABLES IN SCHEMA prod.sales_analytics TO ROLE analyst_role;
```
