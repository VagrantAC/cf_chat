{ pkgs ? import <nixpkgs> {}}:

pkgs.mkShell {
  nativeBuildInputs = [ 
    pkgs.go_1_18
    pkgs.golangci-lint
    pkgs.neovim
    pkgs.protobuf3_8
    pkgs.protoc-gen-go
    pkgs.protoc-gen-go-grpc
  ];
}

