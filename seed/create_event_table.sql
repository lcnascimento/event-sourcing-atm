CREATE KEYSPACE IF NOT EXISTS accounts
WITH REPLICATION = { 
  'class' : 'SimpleStrategy', 
  'replication_factor' : 1 
};

CREATE TABLE IF NOT EXISTS "events" (
  row_id TEXT,
  aggregate_id TEXT,
  event_time TIMESTAMP,
  payload BLOB,
  PRIMARY KEY(row_id, aggregate_id, event_time)
);

INSERT INTO events(
  row_id,
  aggregate_id,
  event_time,
  payload
)
VALUES (
  '1234:Test',
  '321',
  DATEOF(NOW()),
  bigintAsBlob(3)
);
