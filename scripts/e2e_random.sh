#!/usr/bin/env bash
# E2E smoke test: run inspect + convert on testdata XML files in random order.
# Usage: ./scripts/e2e_random.sh [N]
#   N  max files to test (default: all)

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BIN="$REPO_ROOT/cmd/gml-parser/gml-parser"
TESTDATA="$REPO_ROOT/testdata"

# Build
echo "==> building gml-parser"
make -C "$REPO_ROOT" gml-parser-build

# Collect data files (exclude KS-META metadata files)
mapfile -t ALL_FILES < <(
  find "$TESTDATA" -name "*.xml" -not -name "KS-META-*" | sort
)

TOTAL="${#ALL_FILES[@]}"

# Shuffle using shuf if available, otherwise sort -R
if command -v shuf &>/dev/null; then
  mapfile -t FILES < <(printf '%s\n' "${ALL_FILES[@]}" | shuf)
else
  mapfile -t FILES < <(printf '%s\n' "${ALL_FILES[@]}" | sort -R)
fi

# Optionally limit count
LIMIT="${1:-$TOTAL}"
FILES=("${FILES[@]:0:$LIMIT}")

PASS=0
FAIL=0
FAILURES=()

echo "==> testing ${#FILES[@]} / $TOTAL files"
echo ""

for f in "${FILES[@]}"; do
  rel="${f#$REPO_ROOT/}"
  printf "  inspect  %s ... " "$rel"
  if "$BIN" inspect -i "$f" > /dev/null 2>&1; then
    printf "ok\n"
    PASS=$((PASS+1))
  else
    printf "FAIL\n"
    FAIL=$((FAIL+1))
    FAILURES+=("inspect  $rel")
  fi

  printf "  convert  %s ... " "$rel"
  if "$BIN" convert --swap -i "$f" > /dev/null 2>&1; then
    printf "ok\n"
    PASS=$((PASS+1))
  else
    printf "FAIL\n"
    FAIL=$((FAIL+1))
    FAILURES+=("convert  $rel")
  fi
done

echo ""
echo "==> results: $PASS passed, $FAIL failed"

if [[ $FAIL -gt 0 ]]; then
  echo ""
  echo "FAILURES:"
  for line in "${FAILURES[@]}"; do
    echo "  $line"
  done
  exit 1
fi
