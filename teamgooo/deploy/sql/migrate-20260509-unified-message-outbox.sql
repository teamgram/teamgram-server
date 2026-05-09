-- Precondition: run once against schemas where canonical_messages does not yet
-- have the unified outbox attrs/forward-ref columns. Existing migration style
-- does not use ADD COLUMN IF NOT EXISTS.
ALTER TABLE `canonical_messages`
  ADD COLUMN `message_attrs_schema_version` int NOT NULL DEFAULT '0' AFTER `media_ref_payload`,
  ADD COLUMN `message_attrs_payload` blob AFTER `message_attrs_schema_version`,
  ADD COLUMN `forward_ref_schema_version` int NOT NULL DEFAULT '0' AFTER `message_attrs_payload`,
  ADD COLUMN `forward_ref_payload` blob AFTER `forward_ref_schema_version`;
