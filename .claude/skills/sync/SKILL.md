---
name: sync
description: Sync the project's documentation and memory to match the current state of the codebase. Run after significant changes — new modules, version bumps, refactors, or completed features. Checks git status/log, build, module versions, then updates README, CLAUDE.md, and Memory as needed.
---

# /sync — Project State Sync

Bring README, CLAUDE.md, and Memory up to date with the actual current state of the codebase.

## Step 1: Gather current state

Run all of the following (each as a separate Bash call — no `&&`):

1. `git log --oneline -15` — recent commits
2. `git status` — uncommitted changes
3. `git tag --list | sort -V` — all published tags and their versions
4. `make build` — verify build passes (use `GOTMPDIR=/workspace/go-gml/.tmp`)
5. Read `go.work` — confirm which modules are in the workspace
6. Read each module's `go.mod` to confirm current require versions:
   - `core/go.mod`, `gml3_2_1/go.mod`, `gml3_1_1/go.mod`, `gml2_1_2/go.mod`, `gml/go.mod`, `cmd/gml-parser/go.mod`

## Step 2: Identify what has changed

Compare the gathered state against the current content of:
- `README.md`
- `.claude/CLAUDE.md`
- Memory files at `/home/youruser/.claude/projects/-workspace-go-gml/memory/`

Look for:
- New or moved modules not reflected in docs
- Version numbers that have been bumped (tags vs go.mod require)
- New features or implementation status changes
- Outdated file paths, module paths, or Makefile targets
- Missing `go install` instructions or changed CLI location
- Build failures that need to be noted

## Step 3: Update docs

Apply only the changes that are needed. Do not rewrite sections that are still accurate.

### README.md
- Module table: add new modules, remove deleted ones, update paths
- Install instructions: verify `go install` path and version
- Supported geometry table: reflect implementation status from CLAUDE.md

### .claude/CLAUDE.md
- コード構造のツリー: add/remove/rename modules and paths
- 要素別実装状況: update ✓/△/未実装 based on recent commits
- リリース手順: keep in sync with `docs/release.md`
- 開発コマンド: verify Makefile targets still match

### Memory (`/home/youruser/.claude/projects/-workspace-go-gml/memory/`)
- Read `MEMORY.md` (index) and relevant detail files
- Update project structure tree in `MEMORY.md`
- Update `design.md` if module structure or dependency graph changed
- Add a feedback memory if a recurring mistake pattern was noticed this session
- Do NOT update implementation detail memories (those come from code, not this sync)
- Update or remove stale memories; add new ones only for non-obvious facts

## Step 4: Report

Summarize what was checked and what was updated:
- Build status (pass / fail)
- Tags vs go.mod versions (in sync or drift detected)
- Files updated and what changed
- Anything that looks inconsistent but was left unchanged (explain why)

## Guidelines

- **Bash commands one at a time** — the hook blocks `&&` and `||`
- **Read before editing** — always Read a file before using Edit or Write on it
- **Minimal changes** — only update what is actually wrong or missing; do not reformat sections that are still accurate
- **Do not commit** — leave committing to the user unless explicitly asked
