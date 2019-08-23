--  Table Structure for `fishhub/fish/models.Video`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS "video" (
                                     "id" serial NOT NULL PRIMARY KEY,
                                     "video_id" text NOT NULL DEFAULT '' ,
                                     "title" text NOT NULL DEFAULT '' ,
                                     "series" text NOT NULL DEFAULT '' ,
                                     "area" text NOT NULL DEFAULT '' ,
                                     "status" text NOT NULL DEFAULT '' ,
                                     "captions" text NOT NULL DEFAULT '' ,
                                     "director" text NOT NULL DEFAULT '' ,
                                     "actor" text NOT NULL DEFAULT '' ,
                                     "plot_category" text NOT NULL DEFAULT '' ,
                                     "level_category" text NOT NULL DEFAULT '' ,
                                     "other_category" text NOT NULL DEFAULT '' ,
                                     "description" text NOT NULL DEFAULT '' ,
                                     "published_time" timestamp with time zone NOT NULL,
                                     "create_time" timestamp with time zone NOT NULL,
                                     "update_time" timestamp with time zone NOT NULL
);

