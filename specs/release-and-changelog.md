# Release and changelog

This document defines versioning, CHANGELOG format, and the release checklist.

## Versioning

- **Semantic versioning** in spirit; example format: `v0.96.0-teamgram-server`.
  - **Major**: Incompatible API or architecture changes.
  - **Minor**: Backward-compatible new features.
  - **Patch**: Backward-compatible fixes and small improvements.
- The suffix (e.g. `-teamgram-server`) may be used to distinguish from dependencies (e.g. proto).

## CHANGELOG

- **Location**: Root **CHANGELOG.md** (when present).
- **Format**: Prefer [Keep a Changelog](https://keepachangelog.com/) (sections: Added, Changed, Deprecated, Removed, Fixed, Security).
- **Responsibility**: Update CHANGELOG for each release with the changes for that version.

## Release checklist

Before cutting a new release:

1. **Version**: Update `VERSION` in Makefile or version injection if used.
2. **CHANGELOG**: Add an entry for the new version and summarize changes.
3. **Tag**: Create a Git tag (e.g. `v0.96.0`) and push.
4. **Artifacts**: Build and publish binaries or Docker images per project practice.
5. **README**: Update any version or “current release” references if needed.
