{ pkgs ? import <nixpkgs> {}}:

pkgs.mkShell {
  nativeBuildInputs = [ 
    pkgs.neovim
	  pkgs.nodePackages.yarn
  	pkgs.nodejs-16_x
    pkgs.nodePackages.webpack
    pkgs.nodePackages.webpack-cli
  ];
}

