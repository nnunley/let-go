;; A portable consumer that pins its host let-go build with require-letgo
;; (Layer 3). require-letgo is let-go-specific, so BOTH the :refer and the
;; load-time call are guarded behind :lg — a JVM-Clojure reader skips them and
;; never tries to load let-go.semver or resolve the require-letgo symbol.
;; Under let-go on the dev/none test binary the version check warns-and-passes.
(ns portable-requires
  #?(:lg
     (:require
      [let-go.semver :refer [require-letgo]])))

#?(:lg (require-letgo ">=1.9.0"))

(defn ok [] :loaded)
