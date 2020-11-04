-- https://docs.datastax.com/en/archived/cassandra/3.0/cassandra/dml/dmlConfigConsistency.html
CONSISTENCY QUORUM;

--https://docs.datastax.com/en/cql-oss/3.3/cql/cql_using/useCreateTable.html
CREATE TABLE korter.communal_fees
(
    id              UUID,
    appartment_code text,
    fees            map<text,double>,
    PRIMARY KEY (id, appartment_code)
);

select * from korter.communal_fees;

-- https://docs.datastax.com/en/cql-oss/3.3/cql/cql_using/useInsertMap.html

INSERT INTO korter.communal_fees(id, appartment_code, fees)
VALUES (uuid(), 'MURAKA3-5', {'küte': 8.43});

-- Add new key:value into map. All clustering keys should be presented.
UPDATE korter.communal_fees
SET fees = fees + {'SoeVesi': 4.65}
WHERE id = 57c339a4-9ad8-41ff-a5e2-30182460551c
  and appartment_code = 'MURAKA3-5';