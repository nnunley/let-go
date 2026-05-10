;; Transducer pipeline — no intermediate collections
(local core (require :io.gitlab.andreyorst.cljlib.core))

(core.transduce
  (core.comp (core.map #(* $ $))
             (core.filter core.even?))
  core.+ 0
  (core.range 100000))
