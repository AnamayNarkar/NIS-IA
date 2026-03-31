#!/usr/bin/env bash
set -euo pipefail

trivy fs --scanners secret --secret-config trivy-secret.yaml --skip-dirs vendor app/
