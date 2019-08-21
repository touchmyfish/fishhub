CREATE TABLE episode (
  id serial NOT NULL ,
  video_id varchar DEFAULT NULL,
  title varchar DEFAULT NULL,
  duration time DEFAULT NULL,
  play_url varchar DEFAULT NULL,
  source integer DEFAULT NULL,
  create_time timestamp without time zone,
  update_time timestamp without time zone,
  description varchar DEFAULT NULL,
  PRIMARY KEY (id)
);