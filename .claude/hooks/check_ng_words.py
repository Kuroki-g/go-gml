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
        "keywords": ["スコープ外", "スコープの問題", "仕様外", "対応しない", "特別扱い", "使い勝手", "動作上は問題なし"],
        "message": """\
NG表現が検出されました: {words}

GML 3.2.1 XSD に準拠することが目的です。難易度に負けて、安易な手段に逃げることは許しません。

回答を修正してください。""",
    },
    {
        "keywords": ["はず", "思われ", "だろう", "かもしれ", "おそらく", "probably", "maybe", "might"],
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
]


def main() -> None:
    try:
        data = json.load(sys.stdin)
    except json.JSONDecodeError:
        sys.exit(0)

    transcript = data.get("transcript", [])
    last_assistant = ""

    for msg in reversed(transcript):
        if msg.get("role") == "assistant":
            content = msg.get("content", "")
            if isinstance(content, list):
                for block in content:
                    if isinstance(block, dict) and block.get("type") == "text":
                        last_assistant += block.get("text", "")
            elif isinstance(content, str):
                last_assistant = content
            break

    for rule in RULES:
        found = [w for w in rule["keywords"] if w in last_assistant]
        if found:
            result = {
                "decision": "block",
                "reason": rule["message"].format(words=", ".join(f'"{w}"' for w in found)),
            }
            print(json.dumps(result, ensure_ascii=False))
            return


if __name__ == "__main__":
    main()
