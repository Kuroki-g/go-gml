import re
from dataclasses import dataclass, field
from typing import Callable

# ---------------------------------------------------------------------------
# DLP – secret patterns
# ---------------------------------------------------------------------------
_DLP_PATTERNS: list[tuple[str, re.Pattern[str]]] = [
    ("anthropic_api_key",   re.compile(r"sk-" r"ant-[a-zA-Z0-9\-_]{20,}")),
    ("openai_api_key",      re.compile(r"sk" r"-[a-zA-Z0-9]{48}")),
    ("github_token",        re.compile(r"gh[pousr]" r"_[a-zA-Z0-9]{36,255}")),
    ("github_pat",          re.compile(r"github_pat" r"_[a-zA-Z0-9_]{82}")),
    ("aws_access_key_id",   re.compile(r"AK" r"IA[0-9A-Z]{16}")),
    ("aws_secret_key",      re.compile(r"(?i)aws.{0,20}" r"secret.{0,20}['\"][0-9a-zA-Z/+]{40}['\"]")),
    ("gcp_api_key",         re.compile(r"AI" r"za[0-9A-Za-z\-_]{35}")),
    ("slack_token",         re.compile(r"xox" r"[baprs]-[0-9A-Za-z\-]+")),
    ("private_key",         re.compile(r"-----" r"BEGIN (?:RSA |EC |OPENSSH )?PRIVATE KEY" r"-----")),
    ("jwt",                 re.compile(r"eyJ" r"[a-zA-Z0-9_-]+\.eyJ[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+")),
]


def _redact(m: re.Match[str]) -> str:
    return m.group()[:6] + "[redacted]"


def _dlp_scan(text: str) -> tuple[str, str] | None:
    """Return (label, redacted_snippet) for the first pattern hit, else None."""
    for label, pattern in _DLP_PATTERNS:
        m = pattern.search(text)
        if m:
            return label, _redact(m)
    return None


# ---------------------------------------------------------------------------
# DLP functions  DlpFn = (url, body) -> (label, snippet) | None
# ---------------------------------------------------------------------------
DlpFn = Callable[[str, str], tuple[str, str] | None]


def default_dlp(url: str, body: str) -> tuple[str, str] | None:
    """Scan URL and request body for known secret patterns."""
    for text in (url, body):
        hit = _dlp_scan(text)
        if hit:
            return hit
    return None


def no_dlp(url: str, body: str) -> tuple[str, str] | None:
    """Disable DLP for this domain."""
    return None


# ---------------------------------------------------------------------------
# Per-domain policy
# ---------------------------------------------------------------------------
@dataclass
class DomainPolicy:
    methods: frozenset[str] = frozenset({"GET", "HEAD"})
    dlp: DlpFn = field(default=default_dlp)


POLICIES: dict[str, DomainPolicy] = {
    # --- Anthropic (POST allowed, DLP enforced) ---
    "anthropic.com":            DomainPolicy(methods=frozenset({"GET", "POST", "HEAD"})),
    "claude.ai":                DomainPolicy(methods=frozenset({"GET", "POST", "HEAD"})),
    "claude.com":               DomainPolicy(methods=frozenset({"GET", "POST", "HEAD"})),
    # --- Dev tools ---
    "github.com":               DomainPolicy(),
    "githubusercontent.com":    DomainPolicy(),
    # --- Python ecosystem ---
    "pypi.org":                 DomainPolicy(),
    "pythonhosted.org":         DomainPolicy(),
    "docs.python.org":          DomainPolicy(),
    "docs.astral.sh":           DomainPolicy(),
    "docs.pydantic.dev":        DomainPolicy(),
    # --- Go ecosystem ---
    "go.dev":                   DomainPolicy(),
    "pkg.go.dev":               DomainPolicy(),
    "proxy.golang.org":         DomainPolicy(),
    "sum.golang.org":           DomainPolicy(),
    # --- Docs / community ---
    "learn.microsoft.com":      DomainPolicy(),
    "modelcontextprotocol.io":  DomainPolicy(),
    "speakerdeck.com":          DomainPolicy(),
    "qiita.com":                DomainPolicy(),
    "zenn.dev":                 DomainPolicy(),
    "testing.googleblog.com":   DomainPolicy(),
    "www.w3.org":               DomainPolicy(),
    # --- Geometry---
    "nlftp.mlit.go.jp":         DomainPolicy(),
    "www.geopackage.org"        DomainPolicy(),
    "gdal.org"                  DomainPolicy(),
    # --- 学術論文 ---
    "jstage.jst.go.jp":         DomainPolicy(),
    "ci.nii.ac.jp":             DomainPolicy(),
}


def get_policy(host: str) -> DomainPolicy | None:
    for domain, policy in POLICIES.items():
        if host == domain or host.endswith("." + domain):
            return policy
    return None
