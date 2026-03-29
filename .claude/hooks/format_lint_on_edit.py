#!/usr/bin/env python3
"""PostToolUse hook: 編集した .go ファイルに gofmt / nilaway をかける。"""

import json
import subprocess
import sys


def main() -> None:
    try:
        data = json.load(sys.stdin)
    except json.JSONDecodeError:
        sys.exit(0)

    file_path = data.get("tool_input", {}).get("file_path", "")
    if file_path.endswith(".go"):
        subprocess.run(["gofmt", "-w", file_path], check=False)
        subprocess.run(["nilaway", "./core/...", "./gml/..." ,"./citygml2_0/..." ,"./gml3_2_1/..."], check=False, cwd="/workspace/go-gml")
        subprocess.run(["go", "mod", "tidy"], check=False, cwd="/workspace/go-gml")


if __name__ == "__main__":
    main()
