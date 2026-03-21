---
name: search
description: Deep research skill that spawns parallel subagents to investigate topics thoroughly and write findings to markdown. Use this when the user types /search or asks for deep research, investigation, latest information lookup, or wants multiple topics researched in parallel. Also trigger when the user points to a directory or set of files and wants them researched, enriched, or updated with latest findings. Even if the user just says "調べて" or "deep dive" without the slash command, use this skill. The key signal is: multiple topics or files to research, or wanting more thorough investigation than a single direct answer.
---

# /search — Deep Research with Subagents

Spawn parallel subagents to research topics deeply, then write summarized findings as markdown.

## Step 1: Understand the target

When invoked, first determine:
- **What** needs to be researched (topics, files, directory, question)
- **Where** to write results (append to existing files, or create new ones)
- **How deep** to go (breadth overview vs. deep dive on one thing)

If the target or direction is ambiguous — e.g., the user just typed `/search` with no arguments — use AskUserQuestion to clarify before proceeding. Ask about the target and the desired output location.

If a directory is given, read its README or index file first to understand the structure, then identify which files or topics to research.

## Step 2: Plan the subagents

Break the research into independent units — one subagent per topic or file. Aim for parallelism: if there are 5 files, spawn 5 subagents at once.

For each unit, decide:
- What to research (the topic or existing file content)
- Where to write results (path to existing file, or path to new file)
- What angle to investigate (latest ecosystem status, available libraries, known gaps, recent changes, etc.)

Tell the user your plan before spawning: list the subagents you're about to launch and what each will investigate.

## Step 3: Spawn all subagents in parallel

Launch all subagents in a single turn. Each subagent receives this instruction template:

```
## Research task

Topic: <topic or file subject>
Existing content (if any): <paste existing file content, or "none">
Output file: <absolute path>

## Instructions

1. Read the existing content above (if any) to understand context and what's already covered.
2. Use WebSearch to investigate: current state of the ecosystem, available libraries/tools (especially in Go if relevant), recent developments, known pain points, and anything that would add value beyond what's already written.
3. Synthesize your findings into a concise markdown summary. Focus on new information not already in the existing content.
4. Append your findings to the output file using this format:

---

## 調査メモ (<today's date>)

### <Finding category 1>
<content>

### <Finding category 2>
<content>

(add sections as needed based on what you found)
```

If the output file doesn't exist, create it. Start with a `# <Topic>` heading, then write your findings.

## Step 4: Collect and confirm

Once all subagents complete:
- Report a brief summary of what was written and where
- If any subagent found nothing new or hit an error, mention it
- Ask the user if they want to dive deeper on any specific area

## Guidelines

**Breadth vs. depth**: If given a directory with many files, default to breadth — one subagent per file, each doing a focused search. If given a single topic, go deep — use the subagent to search multiple angles (official docs, GitHub, recent blog posts, community discussions).

**Avoid duplication**: Have each subagent read the existing file content first so they don't repeat what's already there.

**Output quality**: Summaries should be informative but concise. Bullet points and short paragraphs are preferred over walls of text. Include sources (URLs) where relevant.

**Language**: Match the language already used in the existing files. If mixed, default to Japanese for prose and English for technical terms/code.
