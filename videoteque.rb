class Videoteque < Formula
  desc "Tool for watching movies"
  homepage "https://github.com/yassernasc/videoteque"
  url "https://github.com/yassernasc/videoteque/archive/refs/tags/2024.01.24.tar.gz"
  sha256 "70aab6491f8a55cfe660c8d07aafd20d060f88cbf29f2805ce403d5a5ec18ccc"
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
