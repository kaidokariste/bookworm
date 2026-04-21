![](../img/snowflake.png )

# Introduction 


## Administrative commands
```sql
SHOW SCHEMAS; -- Schemade vaatamine hetke andmebaasis
SHOW SCHEMAS IN DATABASE db_name;
```

## Kasutajate haldus
```sql
SHOW USERS; -- kasutajate vaatamine
DESC USER KAIDOKARISTE; -- Kasutaja parameetrite vaatamine
SHOW GRANTS TO USER "VICTOR.HUGO"; -- kasutaja rollide ja muude õiguste vaatamine. 
SHOW ROLES; -- Kõik sulle antud rollid
SELECT CURRENT_ROLE(); -- Hetke rolli vaatamine
```

## Ajavööndid
```sql
ALTER SESSION SET TIMEZONE = 'UTC';
ALTER SESSION SET TIMEZONE = 'Europe/Tallinn';
``` 

## Auditeerimine
Auditi mõttes viimase logini kontrollimine
```sql
SELECT name,
       login_name,
       last_success_login,
       disabled,
       default_role,
       default_secondary_role,
       default_warehouse
FROM snowflake.account_usage.users
WHERE deleted_on IS NULL
ORDER BY last_success_login DESC;
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

### Warehouside (Compute) manageerimine
``` 
DROP WAREHOUSE IF EXISTS COMPUTE_WH;
```

### Anonüümsed protseduurid

```sql
DECLARE msg STRING;

BEGIN
    msg := 'Hello World';
    RETURN msg;
END;
```

### JSON Protseduurid
JSON Flat

```sql
    WITH flat AS (SELECT record_content:id::NUMBER AS id,
                             record_content:modified_dtime::TIMESTAMP_LTZ as modified_dtime,
                             f.key                     AS key,
                             f.value                   AS value
                      FROM <schema.json_containing_table>,
                           LATERAL FLATTEN(INPUT => record_content) f
                      )
            SELECT *
                             FROM flat
                                 PIVOT (
                                 MAX(value) FOR key IN (ANY)
                                 )
            order by id, modified_dtime asc;
```
