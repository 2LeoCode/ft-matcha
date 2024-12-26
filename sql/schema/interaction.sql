CREATE TABLE IF NOT EXISTS interaction (
  kind TEXT NOT NULL CHECK (kind IN ('kiss', 'punch', 'block')),
  creation_time INTEGER NOT NULL DEFAULT (unixepoch()),
  giver_id INTEGER REFERENCES user (id) ON DELETE SET NULL,
  given_id INTEGER NOT NULL REFERENCES user (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS i_interaction_kind ON interaction (kind);
CREATE INDEX IF NOT EXISTS i_interaction_giver_id ON interaction (giver_id);
CREATE INDEX IF NOT EXISTS i_interaction_given_id ON interaction (given_id);
