# typed: false
# frozen_string_literal: true

class LetGo < Formula
  desc "A Clojure dialect implemented as a bytecode VM in Go"
  homepage "https://github.com/nooga/let-go"
  license "MIT"
  version "1.6.0"

  on_macos do
    on_intel do
      url "https://github.com/nooga/let-go/releases/download/v1.6.0/let-go_1.6.0_darwin_amd64.tar.gz"
      sha256 "f84a4526433b177e1057de0d7d129f5e346dea337ddf1336565211ebd8a48bda"
    end
    on_arm do
      url "https://github.com/nooga/let-go/releases/download/v1.6.0/let-go_1.6.0_darwin_arm64.tar.gz"
      sha256 "70d6dfde6a14f94f0228c2aea567f3b8aca72c52aceb94b5d3fa6ba364fa81d0"
    end
  end

  on_linux do
    on_intel do
      url "https://github.com/nooga/let-go/releases/download/v1.6.0/let-go_1.6.0_linux_amd64.tar.gz"
      sha256 "61b0b00cdf8a8256d692f73bc9b769b4bc2c666112104b491e34ad8875cd2077"
    end
    on_arm do
      url "https://github.com/nooga/let-go/releases/download/v1.6.0/let-go_1.6.0_linux_arm64.tar.gz"
      sha256 "83375050a923a9034af5f45dec06f2484ada6a49b3d0dd462cad345358b5072f"
    end
  end

  def install
    bin.install "lg"
  end

  test do
    assert_equal "2", shell_output("#{bin}/lg -e '(+ 1 1)'").strip
  end
end
