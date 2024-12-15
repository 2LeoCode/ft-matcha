PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS User (
  id INTEGER PRIMARY KEY,
  mail TEXT UNIQUE,
  password TEXT,
  firstName TEXT,
  lastName TEXT,
  age INTEGER,
  sex INTEGER,
  orientation INTEGER,
  ageGap BLOB,
  fameRatingGap BLOB,
  country TEXT,
  city TEXT,
  bio TEXT,
  lastPosition INTEGER,
  suspended INTEGER,
  profilePictures BLOB
);

CREATE TABLE IF NOT EXISTS Tag (
  id INTEGER PRIMARY KEY,
  name TEXT
);

CREATE TABLE IF NOT EXISTS Discussion (
  id INTEGER PRIMARY KEY,
  creationTime INTEGER
);

CREATE TABLE IF NOT EXISTS Message (
  authorId INTEGER,
  discussionId INTEGER,
  creationTime INTEGER,
  content TEXT,
  FOREIGN KEY(authorId)
    REFERENCES User(id)
    ON DELETE SET NULL,
  FOREIGN KEY(discussionId)
    REFERENCES Discussion(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Interaction (
  kind TEXT,
  giverId INTEGER,
  givenId INTEGER,
  FOREIGN KEY(giverId)
    REFERENCES User(id)
    ON DELETE SET NULL,
  FOREIGN KEY(givenId)
    REFERENCES User(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Report (
  kind TEXT,
  issuerId INTEGER,
  targetId INTEGER,
  FOREIGN KEY(issuerId)
    REFERENCES User(id)
    ON DELETE SET NULL,
  FOREIGN KEY(targetId)
    REFERENCES User(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS UserTag (
  userId INTEGER,
  tagId INTEGER,
  FOREIGN KEY(userId)
    REFERENCES User(id)
    ON DELETE CASCADE,
  FOREIGN KEY(tagId)
    REFERENCES Tag(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS UserDiscussion (
  userId INTEGER,
  discussionId INTEGER,
  FOREIGN KEY(userId)
    REFERENCES User(id)
    ON DELETE CASCADE,
  FOREIGN KEY(discussionId)
    REFERENCES Discussion(id)
    ON DELETE CASCADE
);

CREATE TRIGGER IF NOT EXISTS delete_unused_tags
  AFTER DELETE ON UserTag
BEGIN
  DELETE FROM Tag WHERE id = OLD.tagId
    AND NOT EXISTS (
      SELECT 1 FROM UserTag WHERE tagId = OLD.tagId
    );
END;

CREATE TRIGGER IF NOT EXISTS delete_empty_discussions
  AFTER DELETE ON UserDiscussion
BEGIN
  DELETE FROM Discussion WHERE id = OLD.discussionId
    AND NOT EXISTS (
      SELECT 1 FROM UserDiscussion WHERE discussionId = OLD.discussionId
    );
END;
