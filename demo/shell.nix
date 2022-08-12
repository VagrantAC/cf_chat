{ pkgs ? import <nixpkgs> {}}:

pkgs.mkShell {
  nativeBuildInputs = [
    pkgs.neovim
	  pkgs.nodePackages.yarn
  	pkgs.nodejs-16_x
    pkgs.go_1_18
    pkgs.golangci-lint
    pkgs.neovim
    pkgs.protobuf3_8
    pkgs.protoc-gen-go
    pkgs.protoc-gen-go-grpc
    pkgs.protoc-gen-grpc-web
    pkgs.nodePackages.ts-node
    pkgs.nodePackages.typescript
  ];
}
