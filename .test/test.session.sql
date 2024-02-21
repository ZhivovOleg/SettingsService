--SELECT settings FROM settings s WHERE s.servicename = 'testService';

--SELECT settings->'d' FROM settings WHERE servicename = 'testService';

--SELECT settings FROM settings WHERE serviceName = 'testService';

--SELECT settings->'c'->'ca' FROM settings WHERE servicename = 'testService';

--SELECT * FROM settings;

--DELETE from settings WHERE servicename = 'newService';

--UPDATE settings SET settings = '{"a":"a1"}' WHERE servicename = 'newService';

--UPDATE settings SET settings = jsonb_set(settings, '{b, ba}', '"zzzzzzz"', true) WHERE servicename = 'newService';

--UPDATE settings SET settings = jsonb_set(settings, '{c,ca}', '10', FALSE) WHERE servicename = 'testService';
