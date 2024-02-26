CREATE TABLE settings (
    Id          SERIAL  PRIMARY KEY,
    serviceName varchar UNIQUE NOT NULL,
    settings    jsonb   NOT NULL
);
INSERT INTO settings(serviceName, settings) VALUES ('testService', '{"a":1, "b":"testData", "c":{"ca":1, "cb":"ccc"}, "d":[1,2,3]}')
