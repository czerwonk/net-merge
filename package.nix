{ pkgs, lib, buildGo122Module }:

buildGo122Module {
  pname = "net-merge";
  version = "0.0.1";

  src = lib.cleanSource ./.;

  vendorHash = pkgs.lib.fileContents ./go.mod.sri;

  CGO_ENABLED = 0;

  meta = with lib; {
    description = "Simple tool for merging CIDR/IP addresses from standard input";
    homepage = "https://github.com/czerwonk/net-merge";
    license = licenses.mit;
  };
}
