CREATE TABLE IF NOT EXISTS User (
  id INTEGER PRIMARY KEY,
  mail TEXT UNIQUE,
  password TEXT,
  firstName TEXT,
  lastName TEXT,
  age INTEGER,
  country TEXT,
  city TEXT,
  bio TEXT,
  lastPosition INTEGER,
  fameRating INTEGER
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

CREATE TABLE IF NOT EXISTS View (
  viewerId INTEGER,
  viewedId INTEGER,
  FOREIGN KEY(viewerId)
    REFERENCES User(id)
    ON DELETE SET NULL,
  FOREIGN KEY(viewedId)
    REFERENCES User(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Kiss (
  giverId INTEGER,
  givenId INTEGER,
  FOREIGN KEY(giverId)
    REFERENCES User(id)
    ON DELETE SET NULL,
  FOREIGN KEY(givenId)
    REFERENCES User(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Punch (
  giverId INTEGER,
  givenId INTEGER,
  FOREIGN KEY(giverId)
    REFERENCES User(id)
    ON DELETE SET NULL,
  FOREIGN KEY(givenId)
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
