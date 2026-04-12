(ns bench.reduce)

(defn -main []
  (reduce + 0 (range 1000000)))
