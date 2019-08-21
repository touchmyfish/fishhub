CREATE TABLE video (
  id serial NOT NULL,
  video_id varchar DEFAULT NULL,
  title varchar DEFAULT NULL,
  series integer DEFAULT NULL,
  area varchar DEFAULT NULL,
  status integer DEFAULT NULL,
  published_time timestamp without time zone,
  create_time timestamp without time zone,
  update_time timestamp without time zone,
  captions integer DEFAULT NULL,
  director varchar DEFAULT NULL,
  actor varchar DEFAULT NULL,
  plot_category integer DEFAULT NULL,
  level_category integer DEFAULT NULL,
  other_category integer DEFAULT NULL,
  description varchar DEFAULT NULL,
  PRIMARY KEY (id)
);