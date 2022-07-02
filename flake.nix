{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-22.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachSystem [ "x86_64-linux" ] (system:
      let pkgs = nixpkgs.legacyPackages.${system};
      in
      rec {
        defaultPackage = pkgs.callPackage ./default.nix { };
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            # ide
            kakoune
            kak-lsp
            gopls

            # dev
            go

            # helpers
            gotools
            pgformatter
          ];

          shellHook = ''
          '';
        };
      });
}
