{ lib, stdenv, buildGoModule }:

buildGoModule rec {
  pname = "sunlist";
  version = "0.0.1";

  src = ./src;
  vendorSha256 = null;

  # disable tests since these depend on the db :(
  doCheck = false;

  buildPhase = ''
    go build -o sunlist
  '';

  installPhase = ''
    mkdir -p $out/bin
    cp sunlist $out/bin
  '';

  meta = {
    description = "generates a game list for sunshine";
    homepage = "";
    maintainers = [ "christoph.ruckstetter@weinig.com" ];
    platforms = lib.platforms.linux;
  };
}
