CREATE TABLE IF NOT EXISTS report (
  reason TEXT NOT NULL,
  status TEXT NOT NULL CHECK (
    status IN ('pending', 'being-resolved', 'resolved')
  ),
  creation_time INTEGER NOT NULL DEFAULT (unixepoch()),
  assigned_administrator_id INTEGER
    REFERENCES administrator (id) ON DELETE SET NULL,
  issuer_id INTEGER REFERENCES user (id) ON DELETE SET NULL,
  target_id INTEGER NOT NULL REFERENCES user (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS i_report_status ON report (status);
CREATE INDEX IF NOT EXISTS i_report_assigned_administrator_id
  ON report (assigned_administrator_id);
CREATE INDEX IF NOT EXISTS i_report_issuer_id ON report (issuer_id);
CREATE INDEX IF NOT EXISTS i_report_target_id ON report (target_id);
