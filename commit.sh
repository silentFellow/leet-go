#!/bin/bash

# Get the list of new and changed files
new_files=$(git ls-files --others --exclude-standard)
changed_files=$(git diff --name-only)

# Function to commit files by directory
commit_by_directory() {
  files=$1
  message_prefix=$2

  for dir in $(echo "$files" | xargs -n1 dirname | sort -u); do
    files_in_dir=$(echo "$files" | grep "^$dir/")
    if [ -n "$files_in_dir" ]; then
      git add $files_in_dir
      git commit -m "$message_prefix files in $dir"
    fi
  done
}

# Commit new files by directory
commit_by_directory "$new_files" "Add new"

# Commit changed files by directory
commit_by_directory "$changed_files" "Update"
