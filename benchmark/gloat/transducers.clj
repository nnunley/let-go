(ns bench.transducers)

(defn -main []
  (transduce
    (comp (map #(* % %))
          (filter even?))
    + 0
    (range 100000)))
