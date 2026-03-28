---
name: sync
description: Sync the project's documentation and memory to match the current state of the codebase. Run after significant changes — new modules, version bumps, refactors, or completed features. Checks git status/log, build, module versions, then updates README, CLAUDE.md, and Memory as needed.
---

# /sync — Project State Sync

Bring README, CLAUDE.md, and Memory up to date with the actual current state of the codebase.

**Ground truth is the code, not commit messages or docs.**

## Step 1: Gather current state

Run all of the following (each as a separate Bash call — no `&&`):

1. `git log --oneline -15` — recent commits
2. `git status` — uncommitted changes
3. `git tag --list | sort -V` — all published tags and their versions
4. `make build` — verify build passes
5. Read `go.work` — confirm which modules are in the workspace
6. Read each module's `go.mod` to confirm current require versions:
   - `core/go.mod`, `gml3_2_1/go.mod`, `gml3_1_1/go.mod`, `gml2_1_2/go.mod`, `gml/go.mod`, `cmd/gml-parser/go.mod`

## Step 2: Verify implementation against code (ground truth)

**Do not skip this step.** Docs and memory drift from reality; code does not.

1. `grep -rn "TODO\|FIXME" gml3_2_1/ gml3_1_1/ gml2_1_2/ core/ gml/` — find open issues recorded in code
2. Check handler registrations in each parser's `reader.go` (`handlers` map + switch cases) to confirm what elements are actually handled
3. Check `docs/issues/` for any existing issue files (`ls docs/issues/`)
4. For each item marked "未実装" or "残課題" in CLAUDE.md / MEMORY.md — verify whether it is still open by reading the relevant code. A commit message saying "fix X" is not enough; read the code.
5. Check that no MEMORY.md or CLAUDE.md entry references a `docs/issues/` file that no longer exists

## Step 3: Identify what has changed

Compare gathered state against:
- `README.md`
- `.claude/CLAUDE.md`
- Memory files at `/home/youruser/.claude/projects/-workspace-go-gml/memory/`

Look for:
- New or moved modules not reflected in docs
- Version numbers that have been bumped (tags vs go.mod require)
- Stale "未実装" entries that are now implemented (verified in Step 2)
- Stale "実装済み" entries that still have TODO/FIXME in code
- Dangling references to deleted `docs/issues/` files
- Build failures that need to be noted

## Step 4: Update docs

Apply only the changes that are needed. Do not rewrite sections that are still accurate.

### README.md
- Module table: add new modules, remove deleted ones, update paths
- Install instructions: verify `go install` path and version
- Supported geometry table: source of truth for users — update when handlers change

### .claude/CLAUDE.md
- コード構造のツリー: add/remove/rename modules and paths
- **マイルストーン表のみ更新する** (✓/未実装)。完了済み要素の詳細注記は書かない — README が持つ
- **未実装・残課題セクション**: Step 2 の確認結果をもとに更新。コードの TODO/FIXME と一致させる
- リリース手順: keep in sync with `docs/release.md`
- 開発コマンド: verify Makefile targets still match

### Memory (`/home/youruser/.claude/projects/-workspace-go-gml/memory/`)
- Read `MEMORY.md` (index) and relevant detail files
- Update project structure tree in `MEMORY.md`
- **未実装・残課題**: Step 2 の確認結果と一致させる。コードで解決済みなら削除、まだ open なら残す
- `docs/issues/` への参照が切れていたら修正または削除する
- Update `design.md` if module structure or dependency graph changed
- Add a feedback memory if a recurring mistake pattern was noticed this session
- Update or remove stale memories; add new ones only for non-obvious facts

### docs/issues/
- issue ファイルを**削除してよい条件**: コードで解決済みであることを Step 2 で確認した場合のみ
- issue ファイルを削除するときは同時に MEMORY.md / CLAUDE.md の参照も更新する
- 未実装のまま issue ファイルが存在しないなら、新規作成する

## Step 5: Report

Summarize what was checked and what was updated:
- Build status (pass / fail)
- Tags vs go.mod versions (in sync or drift detected)
- TODO/FIXME found in code (list them)
- Files updated and what changed
- Anything that looks inconsistent but was left unchanged (explain why)

## Guidelines

- **Bash commands one at a time** — the hook blocks `&&` and `||`
- **Read before editing** — always Read a file before using Edit or Write on it
- **Minimal changes** — only update what is actually wrong or missing; do not reformat sections that are still accurate
- **Do not commit** — leave committing to the user unless explicitly asked
