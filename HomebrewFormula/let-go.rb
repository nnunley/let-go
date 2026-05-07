# typed: false
# frozen_string_literal: true

class LetGo < Formula
  desc "A Clojure dialect implemented as a bytecode VM in Go"
  homepage "https://github.com/nooga/let-go"
  license "MIT"
  version "1.7.1"

  on_macos do
    on_intel do
      url "https://github.com/nooga/let-go/releases/download/v1.7.1/let-go_1.7.1_darwin_amd64.tar.gz"
      sha256 "e1879abcf78f62c839f72ed26d11531ce74e4312ee62247a5a24b86da16a3d2f"
    end
    on_arm do
      url "https://github.com/nooga/let-go/releases/download/v1.7.1/let-go_1.7.1_darwin_arm64.tar.gz"
      sha256 "d00999b1656424c58a55790dc58e263f20f73097e82e9e55979fb3359578b6c6"
    end
  end

  on_linux do
    on_intel do
      url "https://github.com/nooga/let-go/releases/download/v1.7.1/let-go_1.7.1_linux_amd64.tar.gz"
      sha256 "a41a4ba8a5699fcc115956e262fd5471d040adbc34f2f840bb30191b51d67fd2"
    end
    on_arm do
      url "https://github.com/nooga/let-go/releases/download/v1.7.1/let-go_1.7.1_linux_arm64.tar.gz"
      sha256 "ce3f796ba5f006baa567687b65bb5886b4faba6da498c9a19ae8ebaedc87d0af"
    end
  end

  def install
    bin.install "lg"
  end

  test do
    assert_equal "2", shell_output("#{bin}/lg -e '(+ 1 1)'").strip
  end
end
