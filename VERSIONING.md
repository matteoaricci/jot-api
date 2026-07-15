# Versioning

This project uses [Semantic Versioning](https://semver.org/) with automated version bumps.

## How It Works

### Automatic (on merge to main)
When you merge to `main`, GitHub Actions:
1. Analyzes commit messages since the last tag
2. Determines the version bump (major/minor/patch)
3. Creates a new git tag (e.g., `v1.2.3`)
4. Creates a GitHub Release with grouped changelog

**Auto-release is skipped if:**
- Commit contains `[skip release]`, `[no release]`, or `[release skip]`
- Only chore/docs commits (no `feat:` or `fix:`)
- Version tag already exists

### Manual (workflow dispatch)
You can manually trigger a release from GitHub Actions:
1. Go to Actions → Version Bump → Run workflow
2. Choose bump type: `auto`, `major`, `minor`, or `patch`
3. Workflow creates the version immediately

## Commit Message Format

Use [Conventional Commits](https://www.conventionalcommits.org/) to control version bumps:

### Patch Release (0.1.0 → 0.1.1)
Bug fixes and small changes:
```bash
git commit -m "fix: correct pagination bug"
git commit -m "chore: update dependencies"
git commit -m "docs: update README"
```

### Minor Release (0.1.0 → 0.2.0)
New features (backwards compatible):
```bash
git commit -m "feat: add journal sharing endpoint"
git commit -m "feat: add CSV export"
```

### Major Release (0.1.0 → 1.0.0)
Breaking changes:
```bash
git commit -m "feat!: remove deprecated v1 API"
git commit -m "feat: redesign authentication

BREAKING CHANGE: old JWT tokens no longer valid"
```

## Commit Prefixes

- `feat:` - New feature → **MINOR** bump
- `fix:` - Bug fix → **PATCH** bump
- `docs:` - Documentation only → **No release** (unless manual)
- `chore:` - Maintenance → **No release** (unless manual)
- `refactor:` - Code refactoring → **PATCH** bump
- `test:` - Adding tests → **PATCH** bump
- `!` suffix or `BREAKING CHANGE:` → **MAJOR** bump

## Skipping Releases

To prevent automatic version bump, add to commit message:
```bash
git commit -m "chore: update dependencies [skip release]"
git commit -m "docs: fix typo [no release]"
```

Common skip scenarios:
- Documentation updates
- Dependency updates
- CI/CD changes
- Work-in-progress merges

## Manual Versioning

If you need to create a version manually:

```bash
# Create tag
git tag v1.0.0

# Push tag
git push origin v1.0.0
```

## Building with Version

### Local Development
```bash
go run .
# Version shows: 0.1.0-dev
```

### Docker Build
```bash
./build.sh
# Reads git tag and injects into binary
```

### Manual Build with Version
```bash
VERSION=1.2.3
COMMIT=$(git rev-parse --short HEAD)
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

go build -ldflags "\
  -X github.com/matteoaricci/jot-api/version.Version=$VERSION \
  -X github.com/matteoaricci/jot-api/version.GitCommit=$COMMIT \
  -X github.com/matteoaricci/jot-api/version.BuildTime=$BUILD_TIME" \
  -o jot-api
```

## Checking Version

```bash
# API endpoint
curl http://localhost:8080/api/public/version

# Response:
# {
#   "version": "1.2.3",
#   "gitCommit": "a3f12bc",
#   "buildTime": "2026-07-15T14:30:00Z",
#   "goVersion": "go1.23.0"
# }
```

## Current Version

Check the [latest release](https://github.com/matteoaricci/jot-api/releases/latest) or run:

```bash
git describe --tags --abbrev=0
```
