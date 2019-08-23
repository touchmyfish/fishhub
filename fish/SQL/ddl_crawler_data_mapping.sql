-- --------------------------------------------------
--  Table Structure for `fishhub/fish/models.CrawlerDataMapping`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS "crawler_data_mapping" (
                                                    "id" serial NOT NULL PRIMARY KEY,
                                                    "video_id" text NOT NULL DEFAULT '' ,
                                                    "source_id" text NOT NULL DEFAULT '' ,
                                                    "uri" text NOT NULL DEFAULT '' ,
                                                    "source" text NOT NULL DEFAULT '' ,
                                                    "create_time" timestamp with time zone NOT NULL,
                                                    "update_time" timestamp with time zone NOT NULL
);