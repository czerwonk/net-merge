{
  description = "net-merge - simple tool for merging CIDR/IP addresses from standard input";

  outputs = { self, nixpkgs }:
    let
      forAllSystems = nixpkgs.lib.genAttrs [
        "x86_64-linux"
        "aarch64-darwin"
        "x86_64-darwin"
      ];

      pkgsForSystem = system: (import nixpkgs {
        inherit system;
        overlays = [ self.overlays.default ];
      });
    in
    {
      overlays.default = _final: prev:
        let
          inherit (prev) buildGo122Module callPackage lib;
        in
        {
          net-merge = callPackage ./package.nix { inherit buildGo122Module lib; };
        };

      packages = forAllSystems (system: rec {
        inherit (pkgsForSystem system) net-merge;
        default = net-merge;
      });
    };
}
