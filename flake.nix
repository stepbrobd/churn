{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "flake-utils";
    };
  };

  outputs =
    { self
    , nixpkgs
    , flake-utils
    , gomod2nix
    }: flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs {
        inherit system;
        overlays = [
          gomod2nix.overlays.default
        ];
      };
    in
    {
      packages.default = pkgs.buildGoApplication {
        pname = "churn";
        version = self.shortRev or self.dirtyShortRev;
        src = ./.;
        modules = ./gomod2nix.toml;
      };

      apps.default = flake-utils.lib.mkApp {
        drv = self.packages.${system}.default;
      };

      devShells.default = pkgs.mkShell {
        packages = with pkgs; [
          cobra-cli
          delve
          go
          gopls
          gotools
          go-tools
          gomod2nix.packages.${system}.default
          mysql80
          sqlite
          typst
          typstfmt
        ];
        GOROOT = "${pkgs.go}/share/go";
      };

      formatter = pkgs.writeShellScriptBin "formatter" ''
        set -eoux pipefail
        ${pkgs.nixpkgs-fmt}/bin/nixpkgs-fmt .
        ${pkgs.go}/bin/gofmt -s -w .
        ${pkgs.typstfmt}/bin/typstfmt **/*.typ
      '';
    });
}
