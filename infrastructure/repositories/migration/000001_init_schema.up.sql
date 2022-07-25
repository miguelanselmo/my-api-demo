CREATE TABLE USERS (
   ID 		INT PRIMARY KEY NOT NULL,
   NAME     TEXT NOT NULL,
   EMAIL    TEXT NOT NULL,
   PASSWORD TEXT NOT NULL,
   CREATEDAT TEXT NOT NULL,
   UPDATEDAT TEXT,
   GROUPID  INT
);

CREATE TABLE GROUPS (
   ID          INT PRIMARY KEY NOT NULL,
   NAME        TEXT NOT NULL,
   CREATEDAT   TEXT NOT NULL,
   UPDATEDAT   TEXT
);

