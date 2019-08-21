CREATE TABLE crawler_data_mapping (
  id serial NOT NULL,
  video_id varchar DEFAULT NULL,
  source_id varchar DEFAULT NULL,
  uri varchar DEFAULT NULL,
  source varchar DEFAULT NULL,
  create_time timestamp without time zone,
  update_time timestamp without time zone,
  PRIMARY KEY (id)
);
