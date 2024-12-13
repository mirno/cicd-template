{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go             # Go programming language
    pkgs.gotests        # Generate unit tests for Go
    pkgs.richgo         # Rich data structure rendering for Go
    pkgs.lefthook       # Git hooks manager
    pkgs.go-task           # Taskfile (task runner)

    # Additional Go tools
    pkgs.lefthook
    pkgs.golangci-lint  # Linter for Go code
    pkgs.delve          # Debugger for Go
    pkgs.gotools        # Miscellaneous Go tools

    # CICD
    # pkgs.docker
    pkgs.git
    pkgs.earthly
    pkgs.goreleaser
    pkgs.go-task
    pkgs.gofumpt

    # Linters
    pkgs.golangci-lint
    pkgs.editorconfig-checker
    pkgs.typos
    pkgs.plantuml
    pkgs.yamllint

    pkgs.gotests
    pkgs.gotools
    pkgs.gopls
    pkgs.gosec

    # node
    pkgs.nodejs # Provides Node.js and npm

    # dev
    pkgs.firefox-devedition

  ];
}

