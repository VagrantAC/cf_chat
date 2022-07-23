{ pkgs ? import <nixpkgs> {}}:

pkgs.mkShell {
  nativeBuildInputs = [ 
    pkgs.hello
    pkgs.go_1_18
    pkgs.golangci-lint
    # pkgs.gcc11
    pkgs.neovim
    pkgs.httpie
  ];
}

