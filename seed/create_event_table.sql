CREATE KEYSPACE IF NOT EXISTS accounts
WITH REPLICATION = { 
  'class' : 'SimpleStrategy', 
  'replication_factor' : 1 
};

USE accounts;

CREATE TABLE IF NOT EXISTS "events" (
  row_id TEXT,
  aggregate_id TEXT,
  event_time TIMESTAMP,
  event_type TEXT,
  payload TEXT,
  PRIMARY KEY(row_id, aggregate_id, event_time)
);

INSERT INTO events(
  row_id,
  aggregate_id,
  event_time,
  event_type,
  payload
)
VALUES (
  '1234:Test',
  '321',
  DATEOF(NOW()),
  'Event1',
  bigintAsBlob(5)
);

SELECT * FROM events WHERE row_id = '1234' AND event_type = 'Event2';
