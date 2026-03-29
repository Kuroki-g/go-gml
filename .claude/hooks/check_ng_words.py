#!/usr/bin/env python3
"""
Stop hook: アシスタントの応答に GML 仕様に関する NG 表現が含まれていたらブロックする。

GML 3.2.1 XSD に定義されている要素はすべて「仕様内」「未実装」と表現すること。
「スコープ外」「仕様外」「対応しない」は使用禁止。
"""

import json
import sys

RULES = [
    {
        "keywords": ["スコープ外", "スコープの問題", "仕様外", "対応しない", "特別扱い", "使い勝手", "動作上は問題なし", "実用上は", "実用上問題", "実際の動作", "XSD非準拠", "非準拠ですが", "非準拠だが", "準拠していませんが", "準拠していないが"],
        "message": """\
NG表現が検出されました: {words}

GML 3.2.1 XSD に準拠することが目的です。難易度に負けて、安易な手段に逃げることは許しません。
「実用上は動く」という言い訳は不要です。XSD仕様に準拠した実装を行ってください。

回答を修正してください。""",
    },
    {
        "keywords": ["はず", "思われ", "だろう", "かもしれ", "おそらく", "probably", "maybe", "might", "ほぼ準拠", "概ね準拠", "ほぼ正しい", "非準拠だが"],
        "message": """\
NG表現が検出されました: {words}

推測で実装を行ってはなりません。XSD仕様に厳密に準拠することを求めます。回答を修正してください。""",
    },
    {
        "keywords": ["ついでに", "序に", "一緒に"],
        "message": """\
NG表現が検出されました: {words}

指示を遵守してください。追加で作業を行いたい場合には、ユーザーに提案を行い、明確な指示を受領したのちに実行してください。""",
    },
    {
        "keywords": ["GitHub Issues", "gh", "ghコマンド", "GitHubにIssue"],
        "message": """\
NG表現が検出されました: {words}

このプロジェクトではGithub Issuesを使用しません。@.claude/agents/issue-editor.md を読んでください """,
    },
]


ISSUE_DOCS_PREFIX = "docs/issues/"


def _last_assistant_turn(transcript: list) -> tuple[str, list]:
    """Return (text, tool_use_blocks) from the last assistant message."""
    for msg in reversed(transcript):
        if msg.get("role") != "assistant":
            continue
        content = msg.get("content", "")
        text = ""
        tools = []
        if isinstance(content, str):
            text = content
        elif isinstance(content, list):
            for block in content:
                if not isinstance(block, dict):
                    continue
                if block.get("type") == "text":
                    text += block.get("text", "")
                elif block.get("type") == "tool_use":
                    tools.append(block)
        return text, tools
    return "", []


def _check_issue_path(tools: list) -> dict | None:
    """Block Write/Edit calls that write issue content outside docs/issues/."""
    for block in tools:
        if block.get("name") not in ("Write", "Edit"):
            continue
        file_path = block.get("input", {}).get("file_path", "")
        lower = file_path.lower()
        if "issue" in lower and ISSUE_DOCS_PREFIX not in file_path:
            return {
                "decision": "block",
                "reason": (
                    f"NG: Issue ドキュメントは {ISSUE_DOCS_PREFIX} に作成してください。\n"
                    f"書き込み先: {file_path}\n\n"
                    "@.claude/agents/issue-editor.md を使って docs/issues/<filename>.md に書いてください。"
                ),
            }
    return None


def main() -> None:
    try:
        data = json.load(sys.stdin)
    except json.JSONDecodeError:
        sys.exit(0)

    transcript = data.get("transcript", [])
    last_text, tool_blocks = _last_assistant_turn(transcript)

    for rule in RULES:
        found = [w for w in rule["keywords"] if w in last_text]
        if found:
            result = {
                "decision": "block",
                "reason": rule["message"].format(words=", ".join(f'"{w}"' for w in found)),
            }
            print(json.dumps(result, ensure_ascii=False))
            return

    violation = _check_issue_path(tool_blocks)
    if violation:
        print(json.dumps(violation, ensure_ascii=False))


if __name__ == "__main__":
    main()
