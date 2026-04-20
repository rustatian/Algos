#!/usr/bin/env bash
# Update all git submodules to the latest commit on origin/master,
# then create one parent-repo commit with the moved pointers.
#
# - Ensures .gitmodules declares branch = master for every submodule.
# - Fails fast on any error (no partial commits).
# - Does NOT push.
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"

# Refuse to run with pre-existing staged changes — otherwise our commit
# would silently include unrelated work.
if ! git diff --cached --quiet; then
    echo "error: parent repo has staged changes. Commit or stash them first." >&2
    exit 1
fi

# Read submodule names + paths from .gitmodules into parallel arrays.
names=()
paths=()
while IFS= read -r line; do
    name=${line#submodule.}
    name=${name%.path *}
    path=${line##* }
    names+=("$name")
    paths+=("$path")
done < <(git config --file .gitmodules --get-regexp '^submodule\..*\.path$')

# 1. Normalize .gitmodules: every submodule should track master.
gitmodules_changed=0
for name in "${names[@]}"; do
    current=$(git config --file .gitmodules --get "submodule.${name}.branch" 2>/dev/null || true)
    if [[ "$current" != "master" ]]; then
        git config --file .gitmodules "submodule.${name}.branch" master
        gitmodules_changed=1
        echo "[gitmodules] set branch=master for ${name}"
    fi
done
git submodule sync --recursive >/dev/null

# 2. Update each submodule to origin/master.
moved=0
for i in "${!names[@]}"; do
    name=${names[$i]}
    path=${paths[$i]}

    echo "==> ${name}"
    old=$(git -C "$path" rev-parse HEAD)
    git -C "$path" fetch --quiet origin master
    git -C "$path" checkout --quiet master
    git -C "$path" pull --quiet --ff-only origin master
    new=$(git -C "$path" rev-parse HEAD)

    if [[ "$old" != "$new" ]]; then
        moved=$((moved + 1))
        echo "    ${old:0:8} -> ${new:0:8}"
    else
        echo "    already up to date"
    fi
done

# 3. Commit if anything changed.
if [[ $moved -eq 0 && $gitmodules_changed -eq 0 ]]; then
    echo "Nothing to commit."
    exit 0
fi

git add .gitmodules "${paths[@]}"

if git diff --cached --quiet; then
    echo "Nothing staged; skipping commit."
    exit 0
fi

git commit -m "chore: update submodules"
echo "Committed."
