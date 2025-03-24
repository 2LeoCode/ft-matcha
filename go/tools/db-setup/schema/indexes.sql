CREATE INDEX IF NOT EXISTS ix_administrators_username_password
  ON administrators (username, password);

CREATE INDEX IF NOT EXISTS ix_gaps_min_max ON gaps (min, max);

CREATE INDEX IF NOT EXISTS ix_interactions_kind ON interactions (kind);
CREATE INDEX IF NOT EXISTS ix_interactions_giver_id ON interactions (giver_id);
CREATE INDEX IF NOT EXISTS ix_interactions_given_id ON interactions (given_id);

CREATE INDEX IF NOT EXISTS ix_messages_author_id ON messages (author_id);
CREATE INDEX IF NOT EXISTS ix_messages_discussion_id
  ON messages (discussion_id);

CREATE INDEX IF NOT EXISTS ix_profile_pictures_user_id
  ON profile_pictures (user_id);

CREATE INDEX IF NOT EXISTS ix_tags_name ON tags (name);

CREATE INDEX IF NOT EXISTS ix_reportx_status ON reports (status);
CREATE INDEX IF NOT EXISTS ix_reports_assigned_administrator_id
  ON reports (assigned_administrator_id);
CREATE INDEX IF NOT EXISTS ix_reports_issuer_id ON reports (issuer_id);
CREATE INDEX IF NOT EXISTS ix_reports_target_id ON reports (target_id);

CREATE INDEX IF NOT EXISTS ix_users_discussions_user_id
  ON users_discussions (user_id);
CREATE INDEX IF NOT EXISTS ix_users_discussions_discussion_id
  ON users_discussions (discussion_id);

CREATE INDEX IF NOT EXISTS ix_positions_user_id
  ON positions (user_id);
CREATE INDEX IF NOT EXISTS ix_positions_latitude_longitude
  ON positions (latitude, longitude);

CREATE INDEX IF NOT EXISTS ix_users_tags_user_id ON users_tags (user_id);
CREATE INDEX IF NOT EXISTS ix_users_tags_tag_id ON users_tags (tag_id);

CREATE INDEX IF NOT EXISTS ix_users_mail_password ON users (mail, password);
CREATE INDEX IF NOT EXISTS ix_users_birth_date ON users (birth_date);
CREATE INDEX IF NOT EXISTS ix_users_last_position ON users (last_position);
CREATE INDEX IF NOT EXISTS ix_users_orientation_sex ON users (orientation, sex);
CREATE INDEX IF NOT EXISTS ix_users_country_city ON users (country, city);
CREATE INDEX IF NOT EXISTS ix_users_fame_points ON users (fame_points);
