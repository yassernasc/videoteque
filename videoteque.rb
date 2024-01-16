class Videoteque < Formula
  desc "Tool for watching movies"
  homepage "https://github.com/yassernasc/videoteque"
  url "https://github.com/yassernasc/videoteque/archive/refs/tags/2024.01.16.tar.gz"
  sha256 "83fa7d95dc2254d2d09a6ec64e90771f62ced60dc0fddddca865cd24074aee5b"
  license "GPL-3.0-or-later"

  depends_on "go" => :build
  depends_on "node" => :build

  def install
    system "make"
    bin.install "vt"
  end

  test do
    system bin/"vt", "version"
  end
end
