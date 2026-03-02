#!/bin/bash

# Install act if not already installed
if ! command -v act &> /dev/null; then
    echo "Installing act..."
    curl -s https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
fi

# Run the workflow locally (simulates ubuntu-22.04 by default)
act workflow_dispatch -W .github/workflows/on_release_attach_artifacts.yml
