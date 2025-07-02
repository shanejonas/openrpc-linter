# Release Process

This project uses automated releases triggered by merges to the `main` branch.

## How It Works

1. **Automatic**: Every merge to `main` triggers a new release
2. **Semantic Versioning**: Version bumping based on commit messages
3. **Cross-Platform**: Builds binaries for Linux, Windows, and macOS
4. **GitHub Releases**: Automatically creates GitHub releases with assets

## Version Bumping

The release process uses **Conventional Commits** to automatically determine version bumps:

- **Patch**: `fix:` → `v1.0.0` → `v1.0.1`
- **Minor**: `feat:` → `v1.0.0` → `v1.1.0`  
- **Major**: Any commit with `BREAKING` → `v1.0.0` → `v2.0.0`

## Skipping Releases

To skip a release, include `[skip release]` or `[no release]` in your commit message.

## Examples

```bash
# Patch release
git commit -m "fix: resolve linting issue"

# Minor release  
git commit -m "feat: add new validation rule"

# Major release (breaking change)
git commit -m "feat: BREAKING change to API"
# or
git commit -m "fix: BREAKING fix that changes behavior"

# Skip release
git commit -m "docs: update README [skip release]"
```

## Manual Override

If you need to create a specific version manually:

```bash
git tag v1.2.3
git push origin v1.2.3
```

This will trigger the release workflow with the specified version. 