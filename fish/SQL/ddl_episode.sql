-- --------------------------------------------------
--  Table Structure for `fishhub/fish/models.Episode`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS "episode" (
                                       "id" serial NOT NULL PRIMARY KEY,
                                       "video_id" text NOT NULL DEFAULT '' ,
                                       "title" text NOT NULL DEFAULT '' ,
                                       "play_url" text NOT NULL DEFAULT '' ,
                                       "source" text NOT NULL DEFAULT '' ,
                                       "description" text NOT NULL DEFAULT '' ,
                                       "duration" text NOT NULL DEFAULT '' ,
                                       "create_time" timestamp with time zone NOT NULL,
                                       "update_time" timestamp with time zone NOT NULL
);
